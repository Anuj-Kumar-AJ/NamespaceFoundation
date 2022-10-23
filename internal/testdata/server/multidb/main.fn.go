// This file was automatically generated by Namespace.
// DO NOT EDIT. To update, re-run `ns generate`.
// This file was automatically generated.
package main

import (
	"context"
	"flag"

	"namespacelabs.dev/foundation/std/go/core"
	"namespacelabs.dev/foundation/std/go/server"
)

func main() {
	flag.Parse()

	resources := core.PrepareEnv("namespacelabs.dev/foundation/internal/testdata/server/multidb")
	defer resources.Close(context.Background())

	ctx := core.WithResources(context.Background(), resources)

	depgraph := core.NewDependencyGraph()
	RegisterInitializers(depgraph)
	if err := depgraph.RunInitializers(ctx); err != nil {
		core.Log.Fatal(err)
	}

	server.InitializationDone()

	server.Listen(ctx, func(srv server.Server) {
		if errs := WireServices(ctx, srv, depgraph); len(errs) > 0 {
			for _, err := range errs {
				core.Log.Println(err)
			}
			core.Log.Fatalf("%d services failed to initialize.", len(errs))
		}
	})
}