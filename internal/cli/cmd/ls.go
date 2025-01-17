// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"namespacelabs.dev/foundation/internal/cli/fncobra"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/parsing"
	"namespacelabs.dev/foundation/internal/parsing/module"
	"namespacelabs.dev/foundation/std/cfg"
)

func NewLsCmd() *cobra.Command {
	var (
		env cfg.Context
	)

	return fncobra.Cmd(
		&cobra.Command{
			Use:     "ls",
			Short:   "List all known packages in the current workspace.",
			Args:    cobra.NoArgs,
			Aliases: []string{"list"},
		}).
		With(fncobra.HardcodeEnv(&env, "dev")).
		Do(func(ctx context.Context) error {
			root, err := module.FindRoot(ctx, ".")
			if err != nil {
				return err
			}

			list, err := parsing.ListSchemas(ctx, env, root)
			if err != nil {
				return err
			}

			stdout := console.Stdout(ctx)
			for _, s := range list.Locations {
				fmt.Fprintln(stdout, s.Rel())
			}

			return nil
		})
}
