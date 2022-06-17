// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package fnnet

import (
	"context"
	"fmt"
	"net"
)

func ListenPort(ctx context.Context, localAddr string, localPort, containerPort int) (net.Listener, error) {
	var cfg net.ListenConfig

	if localPort == 0 {
		// First we try to listen on a local port that matches the container port.
		lst, err := cfg.Listen(ctx, "tcp", fmt.Sprintf("%s:%d", localAddr, containerPort))
		if err == nil {
			return lst, nil
		}

		// Any failures fallback to the open any port path.
	}

	return cfg.Listen(ctx, "tcp", fmt.Sprintf("%s:%d", localAddr, localPort))
}
