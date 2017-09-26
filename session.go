// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package posh

type Session struct {
	UserID string `json:"user-id"`
	Token  string `json:"token"`
}
