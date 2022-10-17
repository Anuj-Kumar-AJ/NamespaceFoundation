// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package parsing

import (
	"context"

	"golang.org/x/exp/slices"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/std/pkggraph"
)

const StartupTestBinary = "namespacelabs.dev/foundation/std/startup/testdriver"

func transformTest(loc pkggraph.Location, server *schema.Server, test *schema.Test) error {
	if test.PackageName != "" {
		return fnerrors.UserError(loc, "package_name can not be set")
	}

	if test.Name == "" {
		return fnerrors.UserError(loc, "test name must be set")
	}

	if test.Driver == nil {
		return fnerrors.UserError(loc, "driver must be set")
	}

	if test.Driver.Name != "" && test.Driver.Name != test.Name {
		return fnerrors.UserError(loc, "driver.name must be unset or be the same as the test name")
	} else {
		test.Driver.Name = test.Name
	}

	if err := transformBinary(loc, test.Driver); err != nil {
		return err
	}

	if server != nil && !slices.Contains(test.ServersUnderTest, server.PackageName) {
		test.ServersUnderTest = append(test.ServersUnderTest, server.PackageName)
	}

	if len(test.ServersUnderTest) == 0 {
		return fnerrors.UserError(loc, "need at least one server under test")
	}

	test.PackageName = loc.PackageName.String()

	return nil
}

func createServerStartupTest(ctx context.Context, pl pkggraph.PackageLoader, pkgName schema.PackageName) (*schema.Test, error) {
	startupTest, err := pl.LoadByName(ctx, StartupTestBinary)
	if err != nil {
		return nil, err
	}

	if len(startupTest.Binaries) != 1 {
		return nil, fnerrors.InternalError("expected %q to be a single binary", StartupTestBinary)
	}

	return &schema.Test{
		PackageName: pkgName.String(),
		Name:        "startup-test",
		ServersUnderTest: []string{
			pkgName.String(),
		},
		Driver: startupTest.Binaries[0],
	}, nil
}