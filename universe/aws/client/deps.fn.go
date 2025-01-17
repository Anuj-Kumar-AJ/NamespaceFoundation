// This file was automatically generated by Namespace.
// DO NOT EDIT. To update, re-run `ns generate`.

package client

import (
	"context"
	"namespacelabs.dev/foundation/std/go/core"
	"namespacelabs.dev/foundation/std/monitoring/tracing"
	"namespacelabs.dev/foundation/std/secrets"
)

// Dependencies that are instantiated once for the lifetime of the extension.
type ExtensionDeps struct {
	Credentials   *secrets.Value
	OpenTelemetry tracing.DeferredTracerProvider
}

type _checkProvideClientFactory func(context.Context, *ClientFactoryArgs, ExtensionDeps) (ClientFactory, error)

var _ _checkProvideClientFactory = ProvideClientFactory

var (
	Package__hva50k = &core.Package{
		PackageName: "namespacelabs.dev/foundation/universe/aws/client",
	}

	Provider__hva50k = core.Provider{
		Package:     Package__hva50k,
		Instantiate: makeDeps__hva50k,
	}
)

func makeDeps__hva50k(ctx context.Context, di core.Dependencies) (_ interface{}, err error) {
	var deps ExtensionDeps

	// name: "aws_credentials_file"
	// optional: true
	if deps.Credentials, err = secrets.ProvideSecret(ctx, core.MustUnwrapProto("ChRhd3NfY3JlZGVudGlhbHNfZmlsZSgB", &secrets.Secret{}).(*secrets.Secret)); err != nil {
		return nil, err
	}

	if err := di.Instantiate(ctx, tracing.Provider__70o2mm, func(ctx context.Context, v interface{}) (err error) {
		if deps.OpenTelemetry, err = tracing.ProvideTracerProvider(ctx, nil, v.(tracing.ExtensionDeps)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return deps, nil
}
