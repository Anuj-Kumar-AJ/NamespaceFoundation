// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"namespacelabs.dev/foundation/framework/resources"
	"namespacelabs.dev/foundation/framework/resources/provider"
	redisclass "namespacelabs.dev/foundation/library/database/redis"
)

const providerPkg = "namespacelabs.dev/foundation/library/oss/redis"

func main() {
	ctx, p := provider.MustPrepare[*redisclass.DatabaseIntent]()

	endpoint, err := resources.LookupServerEndpoint(p.Resources, fmt.Sprintf("%s:redisServer", providerPkg), "redis")
	if err != nil {
		log.Fatalf("failed to get redis server endpoint: %v", err)
	}

	instance := &redisclass.DatabaseInstance{
		Database: p.Intent.Database,
		Url:      endpoint,
		Password: "", // TODO model password as a generated secret
	}

	client := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     instance.Url,
		Password: instance.Password,
		DB:       int(instance.Database),
	})

	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("redis database never became ready: %v", err)
	}

	p.EmitResult(instance)
}
