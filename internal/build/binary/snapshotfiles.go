// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package binary

import (
	"context"
	"strings"

	"github.com/google/go-containerregistry/pkg/v1/empty"
	"github.com/google/go-containerregistry/pkg/v1/mutate"
	"namespacelabs.dev/foundation/internal/artifacts/oci"
	"namespacelabs.dev/foundation/internal/build"
	"namespacelabs.dev/foundation/internal/compute"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/fnfs/memfs"
	"namespacelabs.dev/foundation/std/pkggraph"
	"namespacelabs.dev/foundation/std/tasks"
)

type snapshotFiles struct {
	rel   string
	files []string
}

func (m snapshotFiles) BuildImage(ctx context.Context, env pkggraph.SealedContext, conf build.Configuration) (compute.Computable[oci.Image], error) {
	w := conf.Workspace().VersionedFS(m.rel, true)

	return compute.Map(tasks.Action("binary.snapshot-files"),
		compute.Inputs().Computable("fsys", w),
		compute.Output{},
		func(ctx context.Context, r compute.Resolved) (oci.Image, error) {
			y := compute.MustGetDepValue(r, w, "fsys").FS()

			var files []string
			for _, file := range m.files {
				if strings.HasPrefix(file, "/") {
					return nil, fnerrors.BadInputError("absolute paths not supported (saw %q)", file)
				} else if strings.HasPrefix(file, "../") {
					return nil, fnerrors.BadInputError("relative paths that leave %q not supported (saw %q)", m.rel, file)
				} else if strings.HasPrefix(file, "./") {
					files = append(files, strings.TrimPrefix(file, "./"))
				} else {
					files = append(files, file)
				}
			}

			result, err := memfs.Snapshot(y, memfs.SnapshotOpts{IncludeFiles: files, RequireIncludeFiles: true})
			if err != nil {
				return nil, err
			}

			layer, err := oci.LayerFromFS(ctx, result)
			if err != nil {
				return nil, err
			}

			return mutate.AppendLayers(empty.Image, layer)
		}), nil
}

func (m snapshotFiles) PlatformIndependent() bool { return true }
