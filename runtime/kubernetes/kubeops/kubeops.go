// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package kubeops

import "namespacelabs.dev/foundation/runtime/kubernetes/networking/ingress"

func Register() {
	registerApply()
	registerCreate()
	RegisterCreateSecret()
	registerDelete()
	registerDeleteList()

	ingress.RegisterGraphHandlers()
}