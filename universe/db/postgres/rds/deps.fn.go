// This file was automatically generated by Namespace.
// DO NOT EDIT. To update, re-run `ns generate`.

package rds

import (
	"context"
	fncore "namespacelabs.dev/foundation/std/core"
	"namespacelabs.dev/foundation/std/core/types"
	"namespacelabs.dev/foundation/std/go/core"
	"namespacelabs.dev/foundation/universe/aws/client"
	"namespacelabs.dev/foundation/universe/db/postgres"
	"namespacelabs.dev/foundation/universe/db/postgres/internal/base"
	"namespacelabs.dev/foundation/universe/db/postgres/internal/gencreds"
)

// Dependencies that are instantiated once for the lifetime of the extension.
type ExtensionDeps struct {
	ClientFactory client.ClientFactory
	Creds         *gencreds.Creds
	ServerInfo    *types.ServerInfo
	Wire          base.WireDatabase
}

type _checkProvideDatabase func(context.Context, *Database, ExtensionDeps) (*postgres.DB, error)

var _ _checkProvideDatabase = ProvideDatabase

var (
	Package__4j13h1 = &core.Package{
		PackageName: "namespacelabs.dev/foundation/universe/db/postgres/rds",
	}

	Provider__4j13h1 = core.Provider{
		Package:     Package__4j13h1,
		Instantiate: makeDeps__4j13h1,
	}
)

func makeDeps__4j13h1(ctx context.Context, di core.Dependencies) (_ interface{}, err error) {
	var deps ExtensionDeps

	if err := di.Instantiate(ctx, client.Provider__hva50k, func(ctx context.Context, v interface{}) (err error) {
		if deps.ClientFactory, err = client.ProvideClientFactory(ctx, nil, v.(client.ExtensionDeps)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	if err := di.Instantiate(ctx, gencreds.Provider__ih9ch5, func(ctx context.Context, v interface{}) (err error) {
		if deps.Creds, err = gencreds.ProvideCreds(ctx, nil, v.(gencreds.ExtensionDeps)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	if deps.ServerInfo, err = fncore.ProvideServerInfo(ctx, nil); err != nil {
		return nil, err
	}

	if err := di.Instantiate(ctx, base.Provider__26debk, func(ctx context.Context, v interface{}) (err error) {
		if deps.Wire, err = base.ProvideWireDatabase(ctx, nil, v.(base.ExtensionDeps)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return deps, nil
}