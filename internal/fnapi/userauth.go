// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package fnapi

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"namespacelabs.dev/foundation/internal/clerk"
	"namespacelabs.dev/foundation/internal/workspace/dirs"
)

var ErrRelogin = errors.New("not logged in, please run `ns login`")

const userAuthJson = "auth.json"

type UserAuth struct {
	Username       string `json:"username,omitempty"`
	Org            string `json:"org,omitempty"` // The organization this user is acting as. Only really relevant for robot accounts which authenticate against a repository.
	InternalOpaque []byte `json:"opaque,omitempty"`

	Clerk *clerk.State `json:"clerk,omitempty"`
}

func (user UserAuth) GenerateToken(ctx context.Context) (string, error) {
	if user.Clerk != nil {
		jwt, err := clerk.JWT(ctx, user.Clerk)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("jwt:%s", jwt), nil
	}

	return base64.RawStdEncoding.EncodeToString(user.InternalOpaque), nil
}

func LoginAsRobotAndStore(ctx context.Context, repository, accessToken string) (string, error) {
	userAuth, err := RobotLogin(ctx, repository, accessToken)
	if err != nil {
		return "", err
	}

	return StoreUser(ctx, userAuth)
}

func StoreUser(ctx context.Context, userAuth *UserAuth) (string, error) {
	userAuthData, err := json.Marshal(userAuth)
	if err != nil {
		return "", err
	}

	return userAuth.Username, StoreMarshalledUser(ctx, userAuthData)
}

func StoreMarshalledUser(ctx context.Context, userAuthData []byte) error {
	configDir, err := dirs.Ensure(dirs.Config())
	if err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(configDir, userAuthJson), userAuthData, 0600); err != nil {
		return err
	}

	return nil
}

func LoadUser() (*UserAuth, error) {
	dir, err := dirs.Config()
	if err != nil {
		return nil, err
	}

	p := filepath.Join(dir, userAuthJson)
	data, err := os.ReadFile(p)
	if err != nil {
		if os.IsNotExist(err) {
			// XXX use fnerrors
			return nil, ErrRelogin
		}

		return nil, err
	}

	userAuth := &UserAuth{}
	if err := json.Unmarshal(data, userAuth); err != nil {
		return nil, err
	}

	return userAuth, nil
}
