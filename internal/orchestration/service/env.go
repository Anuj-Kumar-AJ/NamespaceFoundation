// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package service

import (
	"namespacelabs.dev/foundation/internal/protos"
	"namespacelabs.dev/foundation/runtime/kubernetes/client"
	"namespacelabs.dev/foundation/schema"
)

type env struct {
	workspace *schema.Workspace
	devHost   *schema.DevHost
	env       *schema.Environment
}

func makeEnv(plan *schema.DeployPlan) *env {
	return &env{
		workspace: plan.Workspace,
		env:       plan.Environment,
		devHost: &schema.DevHost{
			// TODO add more config bits here when needed.
			Configure: []*schema.DevHost_ConfigureEnvironment{{
				Purpose: plan.Environment.Purpose,
				Runtime: "kubernetes",
				Configuration: protos.WrapAnysOrDie(
					&client.HostEnv{Incluster: true}),
			}},
		},
	}
}

func (e env) ErrorLocation() string        { return e.workspace.ModuleName }
func (e env) Workspace() *schema.Workspace { return e.workspace }
func (e env) DevHost() *schema.DevHost     { return e.devHost }
func (e env) Proto() *schema.Environment   { return e.env }