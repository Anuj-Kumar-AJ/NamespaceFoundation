// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package devsession

import (
	"context"
	"fmt"

	"github.com/gorilla/mux"
	"namespacelabs.dev/foundation/internal/artifacts/oci"
	"namespacelabs.dev/foundation/internal/compute"
	"namespacelabs.dev/foundation/languages/web"
	"namespacelabs.dev/foundation/schema"
)

const (
	WebPackage schema.PackageName = "namespacelabs.dev/foundation/internal/webui/devui"

	baseRepository = "us-docker.pkg.dev/foundation-344819/prebuilts"
	prebuilt       = "sha256:6c920558b6b92b241bb7a96f9af77b242a9a0fc6d56fe376ba67163358a1f457"
)

func PrebuiltWebUI(ctx context.Context) (*mux.Router, error) {
	image := oci.ImageP(fmt.Sprintf("%s/%s@%s", baseRepository, WebPackage, prebuilt), nil, oci.ResolveOpts{PublicImage: true})

	return compute.GetValue(ctx, web.ServeFS(image, true))
}