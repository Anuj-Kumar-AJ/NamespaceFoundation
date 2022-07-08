// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package source

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"namespacelabs.dev/foundation/internal/cli/fncobra"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/sdk/deno"
	"namespacelabs.dev/foundation/provision"
	"namespacelabs.dev/foundation/runtime/rtypes"
)

func newDenoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deno",
		Short: "Run Namespace-configured deno.",
	}

	return fncobra.CmdWithEnv(cmd, func(ctx context.Context, env provision.Env, args []string) error {
		d, err := deno.EnsureSDK(ctx)
		if err != nil {
			return err
		}

		if err := d.CacheImports(ctx, env.Root().Abs()); err != nil {
			return err
		}

		x := console.EnterInputMode(ctx)
		defer x()

		return d.Run(ctx, env.Root().Abs(), rtypes.IO{Stdin: os.Stdin, Stdout: os.Stdout, Stderr: os.Stderr}, "repl", "--cached-only")
	})
}