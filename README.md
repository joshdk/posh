[![GoDoc](https://godoc.org/github.com/joshdk/posh?status.svg)](https://godoc.org/github.com/joshdk/posh)
[![Go Report Card](https://goreportcard.com/badge/github.com/joshdk/posh)](https://goreportcard.com/report/github.com/joshdk/posh)
[![CircleCI](https://circleci.com/gh/joshdk/posh.svg?&style=shield)](https://circleci.com/gh/joshdk/posh/tree/master)

# Posh

Poshmark API client for the Go programming language

## Installing

You can fetch this library by running the following

    go get -u github.com/joshdk/posh

## Example

You can construct a simple client with the following

```go
creds := posh.Credentials{
	Email:    "me@example.com",
	Password: "Pa$sw0rd",
}

config := posh.Config{
	Credentials: &creds,
}

client, err := posh.NewClient(config)
if err != nil {
	panic(err.Error())
}

fmt.Println(client.Session())
```

## License

This library is distributed under the [MIT License](https://opensource.org/licenses/MIT), see LICENSE.txt for more information.