package util

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

// AgentName ..
var AgentName = "Go-Client/v0.0.1"

// Client type
type Client struct {
	endpointBase *url.URL     // default APIEndpointBase
	httpClient   *http.Client // default http.DefaultClient
	reqHooks     []func(req *http.Request)
}

// ClientOption type
type ClientOption func(*Client) error

// NewClient returns a new client instance.
func NewClient(options ...ClientOption) (*Client, error) {
	c := &Client{}
	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}
	if c.httpClient == nil {
		c.httpClient = http.DefaultClient
	}
	return c, nil
}

// WithHTTPClient function
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) error {
		c.httpClient = client
		return nil
	}
}

// WithReqHook function
func WithReqHook(hook func(req *http.Request)) ClientOption {
	return func(c *Client) error {
		c.reqHooks = append(c.reqHooks, hook)
		return nil
	}
}

// WithEndpointBase function
func WithEndpointBase(endpointBase string) ClientOption {
	return func(c *Client) error {
		u, err := url.ParseRequestURI(endpointBase)
		if err != nil {
			return err
		}
		c.endpointBase = u
		return nil
	}
}
func (c *Client) url(base *url.URL, endpoint string, query url.Values) string {
	u := *base
	u.Path = path.Join(u.Path, endpoint)
	if query != nil {
		u.RawQuery = query.Encode()
	}
	return u.String()
}

// Do ..
func (c *Client) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	// req.Header.Set("Authorization", "Bearer "+c.channelToken)
	req.Header.Set("User-Agent", AgentName)
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	for _, hook := range c.reqHooks {
		hook(req)
	}
	return c.httpClient.Do(req)
}

// Get ..
func (c *Client) Get(ctx context.Context, endpoint string, query url.Values) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, c.url(c.endpointBase, endpoint, query), nil)
	if err != nil {
		return nil, err
	}
	// query.Add("_ts", strconv.FormatInt(time.Now().Unix(), 10))
	// req.URL.RawQuery = query.Encode()
	return c.Do(ctx, req)
}

// Post ..
func (c *Client) Post(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, c.url(c.endpointBase, endpoint, nil), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return c.Do(ctx, req)
}

// Postform ..
func (c *Client) Postform(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", c.url(c.endpointBase, endpoint, nil), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return c.Do(ctx, req)
}

// Put ..
func (c *Client) put(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPut, c.url(c.endpointBase, endpoint, nil), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return c.Do(ctx, req)
}

// Delete ..
func (c *Client) Delete(ctx context.Context, endpoint string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, c.url(c.endpointBase, endpoint, nil), nil)
	if err != nil {
		return nil, err
	}
	return c.Do(ctx, req)
}

func closeResponse(res *http.Response) error {
	defer res.Body.Close()
	_, err := io.Copy(ioutil.Discard, res.Body)
	return err
}
