// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package service

import (
	"google.golang.org/protobuf/proto"
	"namespacelabs.dev/foundation/internal/protos"
	"namespacelabs.dev/foundation/providers/aws"
	"namespacelabs.dev/foundation/runtime/kubernetes/client"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/std/planning"
)

func makeEnv(plan *schema.DeployPlan, awsConf *aws.Conf) planning.Context {
	messages := []proto.Message{&client.HostEnv{Incluster: true}}

	if awsConf != nil {
		messages = append(messages, awsConf)
	}

	cfg := planning.MakeConfigurationWith(plan.Environment.Name, planning.ConfigurationSlice{Configuration: protos.WrapAnysOrDie(messages...)})

	return planning.MakeUnverifiedContext(cfg, planning.MakeWorkspace(plan.Workspace, nil), plan.Environment, plan.Workspace.ModuleName)
}