// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package posh

type User struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Username      string `json:"username"`
	DisplayHandle string `json:"display_handle"`
	FullName      string `json:"full_name"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Gender        string `json:"gender"`
}
