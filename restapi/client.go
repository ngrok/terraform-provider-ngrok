// Code generated for API Clients. DO NOT EDIT.

package restapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	ctypeJson = "application/json"
	ctypeForm = "application/x-www-form-urlencoded"
)

const (
	LatestVersion  = "2"
	defaultBaseURL = "https://api.ngrok.com"
)

type Client struct {
	http          *http.Client
	baseURL       *url.URL
	version       string
	apiKey        string
	origin        string
	forwarded_for string
	debug         *Debug
}

type ClientConfig struct {
	HTTPClient *http.Client
	BaseURL    string
	Version    string
	APIKey     string
	Debug      Debug
}

func NewClient(cfg ClientConfig) (*Client, error) {
	u, err := url.Parse(cfg.BaseURL)
	if err != nil {
		return nil, err
	}
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}
	if cfg.Version == "" {
		cfg.Version = LatestVersion
	}
	if cfg.BaseURL == "" {
		cfg.BaseURL = defaultBaseURL
	}
	c := &Client{
		http:    cfg.HTTPClient,
		baseURL: u,
		apiKey:  cfg.APIKey,
		version: cfg.Version,
		debug:   &cfg.Debug,
	}
	return c, nil
}

func (c *Client) SetOrigin(origin string) {
	c.origin = origin
}

func (c *Client) SetForwardedFor(forwarded_for string) {
	c.forwarded_for = forwarded_for
}

func (c *Client) SetVersion(version int) {
	if version >= 0 {
		c.version = fmt.Sprintf("%d", version)
	} else {
		c.version = ""
	}
}

func (c *Client) Get(ctx context.Context, path string, resp interface{}) (*http.Response, error) {
	r, err := c.MakeRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}
	return c.ReadResponse(r, resp)
}

func (c *Client) Put(ctx context.Context, path string, body, resp interface{}) (*http.Response, error) {
	r, err := c.MakeRequest(ctx, "PUT", path, body)
	if err != nil {
		return nil, err
	}
	return c.ReadResponse(r, resp)
}

func (c *Client) Post(ctx context.Context, path string, body, resp interface{}) (*http.Response, error) {
	r, err := c.MakeRequest(ctx, "POST", path, body)
	if err != nil {
		return nil, err
	}
	return c.ReadResponse(r, resp)
}

func (c *Client) Patch(ctx context.Context, path string, body, resp interface{}) (*http.Response, error) {
	r, err := c.MakeRequest(ctx, "PATCH", path, body)
	if err != nil {
		return nil, err
	}
	return c.ReadResponse(r, resp)
}

func (c *Client) Delete(ctx context.Context, path string, resp interface{}) (*http.Response, error) {
	r, err := c.MakeRequest(ctx, "DELETE", path, nil)
	if err != nil {
		return nil, err
	}
	return c.ReadResponse(r, resp)
}

func (c *Client) Options(ctx context.Context, method, path string) (*http.Response, error) {
	r, err := c.MakeRequest(ctx, "OPTIONS", path, nil)
	if err != nil {
		return nil, err
	}
	if c.origin != "" {
		r.Header.Set("Access-Control-Request-Method", method)
		r.Header.Set("Access-Control-Request-Headers", "Authorization, Content-Type, Origin, Ngrok-Version")
	}
	return c.ReadResponse(r, nil)
}

func (c *Client) Do(ctx context.Context, method, path string, data interface{}) (*http.Response, error) {
	r, err := c.MakeRequest(ctx, method, path, data)
	if err != nil {
		return nil, err
	}
	return c.http.Do(r)
}

func (c *Client) MakeRequest(ctx context.Context, method, path string, data interface{}) (*http.Request, error) {
	body, ctype, err := encodeBody(data)
	if err != nil {
		return nil, err
	}

	ctx, body = c.debug.makeRequest(ctx, body)

	r, err := http.NewRequestWithContext(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	r.URL = c.baseURL.ResolveReference(r.URL)
	r.Header.Set("Authorization", "Bearer "+c.apiKey)
	r.Header.Set("User-Agent", "ngrok-api-client/0")

	if ctype != "" {
		r.Header.Set("Content-Type", ctype)
	}

	if c.origin != "" {
		r.Header.Set("Origin", c.origin)
	}

	if c.forwarded_for != "" {
		r.Header.Set("X-Forwarded-For", c.forwarded_for)
	}

	if c.version != "" {
		r.Header.Set("Ngrok-Version", c.version)
	}

	return r, nil
}

func (c *Client) ReadResponse(req *http.Request, out interface{}) (*http.Response, error) {
	if c.debug.DryRun {
		return c.debug.dryRunResponse(req)
	}

	r, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	c.debug.printResponse(r)

	if r.StatusCode >= http.StatusBadRequest {
		return r, errorFromResponse(r)
	}
	if out != nil {
		err = json.NewDecoder(r.Body).Decode(out)
	}
	if err == nil {
		err = r.Body.Close()
	} else {
		r.Body.Close()
	}

	return r, err
}

func encodeBody(v interface{}) (io.Reader, string, error) {
	switch x := v.(type) {
	case nil:
		// no body, nothing to do
		return nil, "", nil
	case json.RawMessage:
		return bytes.NewReader(x), ctypeJson, nil
	case io.Reader:
		// a reader, unknown content-type
		return x, "", nil
	case url.Values:
		// form-encode the body
		return strings.NewReader(x.Encode()), ctypeForm, nil
	default:
		// json-encode the body
		b, err := json.Marshal(x)
		if err != nil {
			return nil, "", err
		}
		return bytes.NewReader(b), ctypeJson, nil
	}
}

func (e *Error) Error() string {
	msg := fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Msg)
	if e.ErrorCode != "" {
		msg += fmt.Sprintf(" [%s]", e.ErrorCode)
	}
	if e.operationID() != "" {
		msg += fmt.Sprintf("\n\nOperation ID: %s", e.operationID())
	}
	return msg
}

func (e *Error) operationID() string {
	return e.Details["operation_id"]
}

func errorFromResponse(res *http.Response) error {
	defer res.Body.Close()

	var v Error

	b, err := io.ReadAll(res.Body)
	if err != nil {
		v.Msg = "incomplete error response"
	} else if err = json.Unmarshal(b, &v); err != nil {
		v.Msg = "invalid error response"
	}
	if err != nil {
		v.StatusCode = int32(res.StatusCode)
		v.Details = map[string]string{
			"internal_msg": err.Error(),
			"invalid_body": string(b),
			"operation_id": res.Header.Get("Ngrok-Operation-Id"),
		}
	}

	return &v
}
