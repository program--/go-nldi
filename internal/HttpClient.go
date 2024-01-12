package internal

import (
	"context"
	"io"
	"net/http"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type nldiHttpClient struct {
	client HttpClient
	ctx    context.Context
}

func (c *nldiHttpClient) WithClient(client HttpClient) *nldiHttpClient {
	c.client = client
	return c
}

func (c *nldiHttpClient) WithContext(ctx context.Context) *nldiHttpClient {
	c.ctx = ctx
	return c
}

func (c *nldiHttpClient) Request(method string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(c.ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *nldiHttpClient) Head(url string) (*http.Response, error) {
	return c.Request("HEAD", url, nil)
}

func (c *nldiHttpClient) Get(url string) (*http.Response, error) {
	return c.Request("GET", url, nil)
}

func (c *nldiHttpClient) Options(url string) (*http.Response, error) {
	return c.Request("OPTIONS", url, nil)
}

func (c *nldiHttpClient) Post(url string, body io.Reader) (*http.Response, error) {
	return c.Request("POST", url, body)
}

func (c *nldiHttpClient) Put(url string, body io.Reader) (*http.Response, error) {
	return c.Request("PUT", url, body)
}

func (c *nldiHttpClient) Delete(url string, body io.Reader) (*http.Response, error) {
	return c.Request("DELETE", url, body)
}

var NLDIHttpClient = new(nldiHttpClient).
	WithClient(http.DefaultClient).
	WithContext(context.Background())
