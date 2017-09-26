// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package posh

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type LoginRequest struct {
	UserHandle string `json:"user_handle"`
	Password   string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	CreatedAt   int    `json:"created_at"`
	ExpiresAt   string `json:"expires_at"`
	User        User   `json:"user"`
}

func (client *Client) Login(request *LoginRequest) (*LoginResponse, error) {

	resp, err := http.PostForm(client.baseURL+"/auth/users/access_token",
		url.Values{
			"user_handle": {request.UserHandle},
			"password":    {request.Password},
		},
	)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := LoginResponse{}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
