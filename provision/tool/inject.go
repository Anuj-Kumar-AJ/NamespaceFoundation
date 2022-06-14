// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package tool

import (
	"context"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"namespacelabs.dev/foundation/schema"
)

var (
	registrations = map[string]func(context.Context, *schema.Environment, *schema.Stack_Entry) (*anypb.Any, error){}
)

func RegisterInjection[V proto.Message](name string, provider func(context.Context, *schema.Environment, *schema.Stack_Entry) (V, error)) {
	registrations[name] = func(ctx context.Context, env *schema.Environment, srv *schema.Stack_Entry) (*anypb.Any, error) {
		msg, err := provider(ctx, env, srv)
		if err != nil {
			return nil, err
		}
		return anypb.New(msg)
	}
}
