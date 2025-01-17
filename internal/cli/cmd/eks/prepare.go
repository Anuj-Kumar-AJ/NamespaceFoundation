// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package eks

import (
	"context"

	"github.com/spf13/cobra"
	"namespacelabs.dev/foundation/internal/cli/fncobra"
	"namespacelabs.dev/foundation/internal/prepare"
	"namespacelabs.dev/foundation/internal/runtime/kubernetes"
	"namespacelabs.dev/foundation/std/cfg"
)

func newPrepareCmd() *cobra.Command {
	cmd := fncobra.CmdWithEnv(&cobra.Command{
		Use:   "prepare-ingress",
		Short: "Runs ingress preparation.",
		Args:  cobra.NoArgs,
	}, func(ctx context.Context, env cfg.Context, args []string) error {
		rt, err := kubernetes.ConnectToCluster(ctx, env.Configuration())
		if err != nil {
			return err
		}

		return prepare.PrepareIngressInKube(ctx, env, rt)
	})

	return cmd
}
