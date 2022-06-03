// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package secrets

import (
	"context"
	"fmt"
	"unicode/utf8"

	"github.com/kr/text"
	"github.com/spf13/cobra"
	"namespacelabs.dev/foundation/internal/cli/fncobra"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/secrets"
)

func newRevealCmd() *cobra.Command {
	var secretKey, specificEnv string

	cmd := &cobra.Command{
		Use:   "reveal",
		Short: "Reveals the specified secret value.",
		Args:  cobra.MaximumNArgs(1),
		RunE: fncobra.RunE(func(ctx context.Context, args []string) error {
			_, bundle, err := loadBundleFromArgs(ctx, args, nil)
			if err != nil {
				return err
			}

			key, err := parseKey(secretKey)
			if err != nil {
				return err
			}

			key.EnvironmentName = specificEnv

			results, err := bundle.LookupValues(ctx, key)
			if err != nil {
				return err
			}

			out := console.Stdout(ctx)

			if len(results) == 1 && utf8.Valid(results[0].Value) {
				fmt.Fprintf(out, "%s\n", results[0].Value)
				return nil
			}

			for k, result := range results {
				if k > 0 {
					fmt.Fprintln(out)
				}

				secrets.DescribeKey(out, result.Key)

				if utf8.Valid(result.Value) {
					fmt.Fprintf(out, "\n\n  %s\n", result.Value)
				} else {
					fmt.Fprintf(out, " raw value:\n\n")
					if err := secrets.OutputBase64(text.NewIndentWriter(out, []byte("  ")), result.Value); err != nil {
						return err
					}
				}
			}

			return nil
		}),
	}

	cmd.Flags().StringVar(&secretKey, "secret", "", "The secret key, in {package_name}:{name} format.")
	cmd.Flags().StringVar(&specificEnv, "env", "", "If set, matches specified secret with the named environment (e.g. dev, or prod).")
	_ = cmd.MarkFlagRequired("secret")

	return cmd
}