// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package posh

import (
	"fmt"
	"os"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewCredentials(email string, password string) *Credentials {
	return &Credentials{
		email,
		password,
	}
}

func NewEnvCredentials() (*Credentials, error) {

	email, found := os.LookupEnv(EmailEnvironmentVariable)
	if !found {
		return nil, fmt.Errorf("no %s found in environment", EmailEnvironmentVariable)
	}

	password, found := os.LookupEnv(PasswordEnvironmentVariable)
	if !found {
		return nil, fmt.Errorf("no %s found in environment", PasswordEnvironmentVariable)
	}

	return &Credentials{
		email,
		password,
	}, nil
}
