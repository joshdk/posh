// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package posh

import (
	"fmt"
	"os"
)

type Session struct {
	UserID string `json:"user-id"`
	Token  string `json:"token"`
}

func NewSession(userID string, token string) *Session {
	return &Session{
		userID,
		token,
	}
}

func NewEnvSession() (*Session, error) {

	userID, found := os.LookupEnv(UserIDEnvironmentVariable)
	if !found {
		return nil, fmt.Errorf("no %s found in environment", UserIDEnvironmentVariable)
	}

	token, found := os.LookupEnv(TokenEnvironmentVariable)
	if !found {
		return nil, fmt.Errorf("no %s found in environment", TokenEnvironmentVariable)
	}

	return &Session{
		userID,
		token,
	}, nil
}
