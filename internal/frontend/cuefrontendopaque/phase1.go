// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package cuefrontendopaque

import (
	"context"

	"namespacelabs.dev/foundation/internal/frontend"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/std/planning"
)

type phase1plan struct {
	startupPlan   *schema.StartupPlan
	declaredStack []schema.PackageName
}

func (p1 phase1plan) EvalProvision(ctx context.Context, env planning.Context, inputs frontend.ProvisionInputs) (frontend.ProvisionPlan, error) {
	var pdata frontend.ProvisionPlan

	pdata.Startup = phase2plan{startupPlan: p1.startupPlan}

	pdata.DeclaredStack = p1.declaredStack

	return pdata, nil
}
