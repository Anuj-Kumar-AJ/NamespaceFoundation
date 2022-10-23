// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package client

import "namespacelabs.dev/foundation/std/cfg"

func NewLocalHostEnv(contextName string, env cfg.Context) *HostEnv {
	hostEnv := &HostEnv{
		Kubeconfig: "~/.kube/config",
		Context:    contextName,
	}

	return hostEnv
}