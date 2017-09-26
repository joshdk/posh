// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package posh

import (
	"fmt"
)

type Client struct {
	baseURL string
	session Session
}

func NewClient(config Config) (*Client, error) {
	client := Client{
		baseURL: BaseURL,
	}

	if config.BaseURL != "" {
		client.baseURL = config.BaseURL
	}

	if config.Credentials != nil {
		request := LoginRequest{
			UserHandle: config.Credentials.Email,
			Password:   config.Credentials.Password,
		}

		response, err := client.Login(&request)
		if err != nil {
			return nil, err
		}

		client.session = Session{
			Token:  response.AccessToken,
			UserID: response.User.ID,
		}

		return &client, nil
	}

	return nil, fmt.Errorf("incomplete client config")
}

func (client *Client) Session() Session {
	return client.session
}
