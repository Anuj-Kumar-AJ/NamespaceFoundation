// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package opaque

import (
	"context"

	"namespacelabs.dev/foundation/universe/db/postgres"
)

func ProvideDatabase(ctx context.Context, db *Database, single ExtensionDeps, deps DatabaseDeps) (*postgres.DB, error) {

	return single.Wire.ProvideDatabase(ctx, &postgres.Database{
		Name:     db.Name,
		HostedAt: db.HostedAt,
		Credentials: &postgres.Database_Credentials{
			User:     &postgres.Database_Credentials_Secret{Value: deps.Creds.Username},
			Password: &postgres.Database_Credentials_Secret{Value: deps.Creds.Password},
		},
	})
}
