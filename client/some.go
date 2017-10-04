// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// unnamed API: some Resource Client
//
// Command:
// $ goagen
// --design=github.com/kovetskiy/goatest/design
// --out=$(GOPATH)/src/github.com/kovetskiy/goatest
// --version=v1.3.0

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// GetSomePath computes a request path to the get action of some.
func GetSomePath() string {

	return fmt.Sprintf("/some")
}

// Get some value
func (c *Client) GetSome(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetSomeRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetSomeRequest create the request corresponding to the get action endpoint of the some resource.
func (c *Client) NewGetSomeRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
