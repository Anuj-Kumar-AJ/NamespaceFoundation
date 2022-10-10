// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package runs

import (
	"encoding/json"
	"fmt"
	"os"

	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/storedrun"
)

func MakeUrl(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fnerrors.BadInputError("%s: failed to read: %w", path, err)
	}

	runId := &storedrun.StoredRunID{}
	if err := json.Unmarshal(content, runId); err != nil {
		return "", fnerrors.InternalError("%s: unable to parse stored run: %w", path, err)
	}

	return fmt.Sprintf("https://dashboard.namespace.so/run/%s", runId.RunId), nil
}
