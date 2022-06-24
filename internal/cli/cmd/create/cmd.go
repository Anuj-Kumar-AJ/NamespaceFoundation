// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package create

import (
	"github.com/spf13/cobra"
)

func NewCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Short:   "Creates a new extension, service, server or workspace (also from a starter template).",
		Aliases: []string{"c"},
	}

	cmd.AddCommand(newExtensionCmd())
	cmd.AddCommand(newServerCmd())
	cmd.AddCommand(newServiceCmd())
	cmd.AddCommand(newWorkspaceCmd())
	cmd.AddCommand(newStarterCmd())

	return cmd
}