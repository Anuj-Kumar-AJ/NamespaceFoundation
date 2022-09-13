// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package cuefrontendopaque

import (
	"context"

	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/frontend/fncue"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/std/pkggraph"
	"namespacelabs.dev/foundation/workspace"
)

type cueTest struct {
	Servers []string `json:"serversUnderTest"`
}

func parseTests(ctx context.Context, pl workspace.EarlyPackageLoader, loc pkggraph.Location, v *fncue.CueV) ([]*schema.Test, error) {
	it, err := v.Val.Fields()
	if err != nil {
		return nil, err
	}

	out := []*schema.Test{}

	for it.Next() {
		parsedTest, err := parseTest(ctx, pl, loc, it.Label(), (&fncue.CueV{Val: it.Value()}))
		if err != nil {
			return nil, err
		}

		out = append(out, parsedTest)
	}

	return out, nil
}

func parseTest(ctx context.Context, pl workspace.EarlyPackageLoader, loc pkggraph.Location, name string, v *fncue.CueV) (*schema.Test, error) {
	var bits cueTest
	if err := v.Val.Decode(&bits); err != nil {
		return nil, err
	}

	out := &schema.Test{
		Name:             name,
		ServersUnderTest: bits.Servers,
	}

	if build := v.LookupPath("build"); build.Exists() {
		cueBuild, err := parseCueBuild(ctx, name, loc, build)
		if err != nil {
			return nil, err
		}

		out.Driver = cueBuild.inlineBinary
	} else {
		return nil, fnerrors.UserError(loc, "missing build definition")
	}

	return out, nil
}