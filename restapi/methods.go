// Code generated for API Clients. DO NOT EDIT.

package restapi

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"text/template"
)

// Creates a new abuse report which will be reviewed by our system and abuse response team. This API is only available to authorized accounts. Contact abuse@ngrok.com to request access
func (c *Client) AbuseReportsCreate(ctx context.Context, arg *AbuseReportCreate) (*AbuseReport, *http.Response, error) {
	var res AbuseReport
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/abuse_reports")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get the detailed status of abuse report by ID.
func (c *Client) AbuseReportsGet(ctx context.Context, arg *Item) (*AbuseReport, *http.Response, error) {
	var res AbuseReport
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/abuse_reports/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create a new Agent Ingress. The ngrok agent can be configured to connect to ngrok via the new set of addresses on the returned Agent Ingress.
func (c *Client) AgentIngressesCreate(ctx context.Context, arg *AgentIngressCreate) (*AgentIngress, *http.Response, error) {
	var res AgentIngress
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/agent_ingresses")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete an Agent Ingress by ID
func (c *Client) AgentIngressesDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/agent_ingresses/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get the details of an Agent Ingress by ID.
func (c *Client) AgentIngressesGet(ctx context.Context, arg *Item) (*AgentIngress, *http.Response, error) {
	var res AgentIngress
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/agent_ingresses/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all Agent Ingresses owned by this account
func (c *Client) AgentIngressesList(ctx context.Context, arg *Paging) (*AgentIngressList, *http.Response, error) {
	var res AgentIngressList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/agent_ingresses")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update attributes of an Agent Ingress by ID.
func (c *Client) AgentIngressesUpdate(ctx context.Context, arg *AgentIngressUpdate) (*AgentIngress, *http.Response, error) {
	var res AgentIngress
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/agent_ingresses/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create a new API key. The generated API key can be used to authenticate to the ngrok API.
func (c *Client) APIKeysCreate(ctx context.Context, arg *APIKeyCreate) (*APIKey, *http.Response, error) {
	var res APIKey
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/api_keys")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete an API key by ID
func (c *Client) APIKeysDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/api_keys/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get the details of an API key by ID.
func (c *Client) APIKeysGet(ctx context.Context, arg *Item) (*APIKey, *http.Response, error) {
	var res APIKey
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/api_keys/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all API keys owned by this account
func (c *Client) APIKeysList(ctx context.Context, arg *Paging) (*APIKeyList, *http.Response, error) {
	var res APIKeyList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/api_keys")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update attributes of an API key by ID.
func (c *Client) APIKeysUpdate(ctx context.Context, arg *APIKeyUpdate) (*APIKey, *http.Response, error) {
	var res APIKey
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/api_keys/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get an application session by ID.
func (c *Client) ApplicationSessionsGet(ctx context.Context, arg *Item) (*ApplicationSession, *http.Response, error) {
	var res ApplicationSession
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/app/sessions/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete an application session by ID.
func (c *Client) ApplicationSessionsDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/app/sessions/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all application sessions for this account.
func (c *Client) ApplicationSessionsList(ctx context.Context, arg *Paging) (*ApplicationSessionList, *http.Response, error) {
	var res ApplicationSessionList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/app/sessions")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get an application user by ID.
func (c *Client) ApplicationUsersGet(ctx context.Context, arg *Item) (*ApplicationUser, *http.Response, error) {
	var res ApplicationUser
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/app/users/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete an application user by ID.
func (c *Client) ApplicationUsersDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/app/users/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all application users for this account.
func (c *Client) ApplicationUsersList(ctx context.Context, arg *Paging) (*ApplicationUserList, *http.Response, error) {
	var res ApplicationUserList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/app/users")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all online tunnel sessions running on this account.
func (c *Client) TunnelSessionsList(ctx context.Context, arg *Paging) (*TunnelSessionList, *http.Response, error) {
	var res TunnelSessionList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/tunnel_sessions")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get the detailed status of a tunnel session by ID
func (c *Client) TunnelSessionsGet(ctx context.Context, arg *Item) (*TunnelSession, *http.Response, error) {
	var res TunnelSession
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/tunnel_sessions/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Issues a command instructing the ngrok agent to restart. The agent restarts itself by calling exec() on platforms that support it. This operation is notably not supported on Windows. When an agent restarts, it reconnects with a new tunnel session ID.
func (c *Client) TunnelSessionsRestart(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/tunnel_sessions/{{ .ID }}/restart")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Issues a command instructing the ngrok agent that started this tunnel session to exit.
func (c *Client) TunnelSessionsStop(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/tunnel_sessions/{{ .ID }}/stop")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Issues a command instructing the ngrok agent to update itself to the latest version. After this call completes successfully, the ngrok agent will be in the update process. A caller should wait some amount of time to allow the update to complete (at least 10 seconds) before making a call to the Restart endpoint to request that the agent restart itself to start using the new code. This call will never update an ngrok agent to a new major version which could cause breaking compatibility issues. If you wish to update to a new major version, that must be done manually. Still, please be aware that updating your ngrok agent could break your integration. This call will fail in any of the following circumstances: there is no update available the ngrok agent's configuration disabled update checks the agent is currently in process of updating the agent has already successfully updated but has not yet been restarted
func (c *Client) TunnelSessionsUpdate(ctx context.Context, arg *TunnelSessionsUpdate) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/tunnel_sessions/{{ .ID }}/update")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create a new Failover backend
func (c *Client) FailoverBackendsCreate(ctx context.Context, arg *FailoverBackendCreate) (*FailoverBackend, *http.Response, error) {
	var res FailoverBackend
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/failover")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete a Failover backend by ID.
func (c *Client) FailoverBackendsDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/failover/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get detailed information about a Failover backend by ID
func (c *Client) FailoverBackendsGet(ctx context.Context, arg *Item) (*FailoverBackend, *http.Response, error) {
	var res FailoverBackend
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/failover/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all Failover backends on this account
func (c *Client) FailoverBackendsList(ctx context.Context, arg *Paging) (*FailoverBackendList, *http.Response, error) {
	var res FailoverBackendList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/failover")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update Failover backend by ID
func (c *Client) FailoverBackendsUpdate(ctx context.Context, arg *FailoverBackendUpdate) (*FailoverBackend, *http.Response, error) {
	var res FailoverBackend
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/failover/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) HTTPResponseBackendsCreate(ctx context.Context, arg *HTTPResponseBackendCreate) (*HTTPResponseBackend, *http.Response, error) {
	var res HTTPResponseBackend
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/http_response")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) HTTPResponseBackendsDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/http_response/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) HTTPResponseBackendsGet(ctx context.Context, arg *Item) (*HTTPResponseBackend, *http.Response, error) {
	var res HTTPResponseBackend
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/http_response/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) HTTPResponseBackendsList(ctx context.Context, arg *Paging) (*HTTPResponseBackendList, *http.Response, error) {
	var res HTTPResponseBackendList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/http_response")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) HTTPResponseBackendsUpdate(ctx context.Context, arg *HTTPResponseBackendUpdate) (*HTTPResponseBackend, *http.Response, error) {
	var res HTTPResponseBackend
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/http_response/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create a new static backend
func (c *Client) StaticBackendsCreate(ctx context.Context, arg *StaticBackendCreate) (*StaticBackend, *http.Response, error) {
	var res StaticBackend
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/static")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete a static backend by ID.
func (c *Client) StaticBackendsDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/static/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get detailed information about a static backend by ID
func (c *Client) StaticBackendsGet(ctx context.Context, arg *Item) (*StaticBackend, *http.Response, error) {
	var res StaticBackend
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/static/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all static backends on this account
func (c *Client) StaticBackendsList(ctx context.Context, arg *Paging) (*StaticBackendList, *http.Response, error) {
	var res StaticBackendList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/static")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update static backend by ID
func (c *Client) StaticBackendsUpdate(ctx context.Context, arg *StaticBackendUpdate) (*StaticBackend, *http.Response, error) {
	var res StaticBackend
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/static/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create a new TunnelGroup backend
func (c *Client) TunnelGroupBackendsCreate(ctx context.Context, arg *TunnelGroupBackendCreate) (*TunnelGroupBackend, *http.Response, error) {
	var res TunnelGroupBackend
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/tunnel_group")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete a TunnelGroup backend by ID.
func (c *Client) TunnelGroupBackendsDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/tunnel_group/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get detailed information about a TunnelGroup backend by ID
func (c *Client) TunnelGroupBackendsGet(ctx context.Context, arg *Item) (*TunnelGroupBackend, *http.Response, error) {
	var res TunnelGroupBackend
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/tunnel_group/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all TunnelGroup backends on this account
func (c *Client) TunnelGroupBackendsList(ctx context.Context, arg *Paging) (*TunnelGroupBackendList, *http.Response, error) {
	var res TunnelGroupBackendList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/tunnel_group")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update TunnelGroup backend by ID
func (c *Client) TunnelGroupBackendsUpdate(ctx context.Context, arg *TunnelGroupBackendUpdate) (*TunnelGroupBackend, *http.Response, error) {
	var res TunnelGroupBackend
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/tunnel_group/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create a new Weighted backend
func (c *Client) WeightedBackendsCreate(ctx context.Context, arg *WeightedBackendCreate) (*WeightedBackend, *http.Response, error) {
	var res WeightedBackend
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/weighted")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete a Weighted backend by ID.
func (c *Client) WeightedBackendsDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/weighted/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get detailed information about a Weighted backend by ID
func (c *Client) WeightedBackendsGet(ctx context.Context, arg *Item) (*WeightedBackend, *http.Response, error) {
	var res WeightedBackend
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/weighted/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all Weighted backends on this account
func (c *Client) WeightedBackendsList(ctx context.Context, arg *Paging) (*WeightedBackendList, *http.Response, error) {
	var res WeightedBackendList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/weighted")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update Weighted backend by ID
func (c *Client) WeightedBackendsUpdate(ctx context.Context, arg *WeightedBackendUpdate) (*WeightedBackend, *http.Response, error) {
	var res WeightedBackend
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/backends/weighted/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create a new bot user
func (c *Client) BotUsersCreate(ctx context.Context, arg *BotUserCreate) (*BotUser, *http.Response, error) {
	var res BotUser
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/bot_users")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete a bot user by ID
func (c *Client) BotUsersDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/bot_users/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get the details of a Bot User by ID.
func (c *Client) BotUsersGet(ctx context.Context, arg *Item) (*BotUser, *http.Response, error) {
	var res BotUser
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/bot_users/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all bot users in this account.
func (c *Client) BotUsersList(ctx context.Context, arg *Paging) (*BotUserList, *http.Response, error) {
	var res BotUserList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/bot_users")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update attributes of a bot user by ID.
func (c *Client) BotUsersUpdate(ctx context.Context, arg *BotUserUpdate) (*BotUser, *http.Response, error) {
	var res BotUser
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/bot_users/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Upload a new Certificate Authority
func (c *Client) CertificateAuthoritiesCreate(ctx context.Context, arg *CertificateAuthorityCreate) (*CertificateAuthority, *http.Response, error) {
	var res CertificateAuthority
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/certificate_authorities")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete a Certificate Authority
func (c *Client) CertificateAuthoritiesDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/certificate_authorities/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get detailed information about a certficate authority
func (c *Client) CertificateAuthoritiesGet(ctx context.Context, arg *Item) (*CertificateAuthority, *http.Response, error) {
	var res CertificateAuthority
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/certificate_authorities/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all Certificate Authority on this account
func (c *Client) CertificateAuthoritiesList(ctx context.Context, arg *Paging) (*CertificateAuthorityList, *http.Response, error) {
	var res CertificateAuthorityList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/certificate_authorities")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update attributes of a Certificate Authority by ID
func (c *Client) CertificateAuthoritiesUpdate(ctx context.Context, arg *CertificateAuthorityUpdate) (*CertificateAuthority, *http.Response, error) {
	var res CertificateAuthority
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/certificate_authorities/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create a new tunnel authtoken credential. This authtoken credential can be used to start a new tunnel session. The response to this API call is the only time the generated token is available. If you need it for future use, you must save it securely yourself.
func (c *Client) CredentialsCreate(ctx context.Context, arg *CredentialCreate) (*Credential, *http.Response, error) {
	var res Credential
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/credentials")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete a tunnel authtoken credential by ID
func (c *Client) CredentialsDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/credentials/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get detailed information about a tunnel authtoken credential
func (c *Client) CredentialsGet(ctx context.Context, arg *Item) (*Credential, *http.Response, error) {
	var res Credential
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/credentials/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all tunnel authtoken credentials on this account
func (c *Client) CredentialsList(ctx context.Context, arg *Paging) (*CredentialList, *http.Response, error) {
	var res CredentialList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/credentials")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update attributes of an tunnel authtoken credential by ID
func (c *Client) CredentialsUpdate(ctx context.Context, arg *CredentialUpdate) (*Credential, *http.Response, error) {
	var res Credential
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/credentials/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create a new endpoint configuration
func (c *Client) EndpointConfigurationsCreate(ctx context.Context, arg *EndpointConfigurationCreate) (*EndpointConfiguration, *http.Response, error) {
	var res EndpointConfiguration
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete an endpoint configuration. This operation will fail if the endpoint configuration is still referenced by any reserved domain or reserved address.
func (c *Client) EndpointConfigurationsDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Returns detailed information about an endpoint configuration
func (c *Client) EndpointConfigurationsGet(ctx context.Context, arg *Item) (*EndpointConfiguration, *http.Response, error) {
	var res EndpointConfiguration
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Returns a list of all endpoint configurations on this account
func (c *Client) EndpointConfigurationsList(ctx context.Context, arg *Paging) (*EndpointConfigurationList, *http.Response, error) {
	var res EndpointConfigurationList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Updates an endpoint configuration. If a module is not specified in the update, it will not be modified. However, each module configuration that is specified will completely replace the existing value. There is no way to delete an existing module via this API, instead use the delete module API.
func (c *Client) EndpointConfigurationsUpdate(ctx context.Context, arg *EndpointConfigurationUpdate) (*EndpointConfiguration, *http.Response, error) {
	var res EndpointConfiguration
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create an HTTPS Edge Route
func (c *Client) EdgesHTTPSRoutesCreate(ctx context.Context, arg *HTTPSEdgeRouteCreate) (*HTTPSEdgeRoute, *http.Response, error) {
	var res HTTPSEdgeRoute
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get an HTTPS Edge Route by ID
func (c *Client) EdgesHTTPSRoutesGet(ctx context.Context, arg *EdgeRouteItem) (*HTTPSEdgeRoute, *http.Response, error) {
	var res HTTPSEdgeRoute
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Updates an HTTPS Edge Route by ID. If a module is not specified in the update, it will not be modified. However, each module configuration that is specified will completely replace the existing value. There is no way to delete an existing module via this API, instead use the delete module API.
func (c *Client) EdgesHTTPSRoutesUpdate(ctx context.Context, arg *HTTPSEdgeRouteUpdate) (*HTTPSEdgeRoute, *http.Response, error) {
	var res HTTPSEdgeRoute
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete an HTTPS Edge Route by ID
func (c *Client) EdgesHTTPSRoutesDelete(ctx context.Context, arg *EdgeRouteItem) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create an HTTPS Edge
func (c *Client) EdgesHTTPSCreate(ctx context.Context, arg *HTTPSEdgeCreate) (*HTTPSEdge, *http.Response, error) {
	var res HTTPSEdge
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get an HTTPS Edge by ID
func (c *Client) EdgesHTTPSGet(ctx context.Context, arg *Item) (*HTTPSEdge, *http.Response, error) {
	var res HTTPSEdge
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Returns a list of all HTTPS Edges on this account
func (c *Client) EdgesHTTPSList(ctx context.Context, arg *Paging) (*HTTPSEdgeList, *http.Response, error) {
	var res HTTPSEdgeList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Updates an HTTPS Edge by ID. If a module is not specified in the update, it will not be modified. However, each module configuration that is specified will completely replace the existing value. There is no way to delete an existing module via this API, instead use the delete module API.
func (c *Client) EdgesHTTPSUpdate(ctx context.Context, arg *HTTPSEdgeUpdate) (*HTTPSEdge, *http.Response, error) {
	var res HTTPSEdge
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete an HTTPS Edge by ID
func (c *Client) EdgesHTTPSDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) HTTPSEdgeMutualTLSModuleReplace(ctx context.Context, arg *EdgeMutualTLSReplace) (*EndpointMutualTLS, *http.Response, error) {
	var res EndpointMutualTLS
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .ID }}/mutual_tls")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) HTTPSEdgeMutualTLSModuleGet(ctx context.Context, arg *Item) (*EndpointMutualTLS, *http.Response, error) {
	var res EndpointMutualTLS
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .ID }}/mutual_tls")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) HTTPSEdgeMutualTLSModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .ID }}/mutual_tls")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) HTTPSEdgeTLSTerminationModuleReplace(ctx context.Context, arg *EdgeTLSTerminationAtEdgeReplace) (*EndpointTLSTermination, *http.Response, error) {
	var res EndpointTLSTermination
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .ID }}/tls_termination")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) HTTPSEdgeTLSTerminationModuleGet(ctx context.Context, arg *Item) (*EndpointTLSTermination, *http.Response, error) {
	var res EndpointTLSTermination
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .ID }}/tls_termination")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) HTTPSEdgeTLSTerminationModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .ID }}/tls_termination")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteBackendModuleReplace(ctx context.Context, arg *EdgeRouteBackendReplace) (*EndpointBackend, *http.Response, error) {
	var res EndpointBackend
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/backend")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteBackendModuleGet(ctx context.Context, arg *EdgeRouteItem) (*EndpointBackend, *http.Response, error) {
	var res EndpointBackend
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/backend")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteBackendModuleDelete(ctx context.Context, arg *EdgeRouteItem) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/backend")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteIPRestrictionModuleReplace(ctx context.Context, arg *EdgeRouteIPRestrictionReplace) (*EndpointIPPolicy, *http.Response, error) {
	var res EndpointIPPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/ip_restriction")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteIPRestrictionModuleGet(ctx context.Context, arg *EdgeRouteItem) (*EndpointIPPolicy, *http.Response, error) {
	var res EndpointIPPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/ip_restriction")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteIPRestrictionModuleDelete(ctx context.Context, arg *EdgeRouteItem) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/ip_restriction")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteRequestHeadersModuleReplace(ctx context.Context, arg *EdgeRouteRequestHeadersReplace) (*EndpointRequestHeaders, *http.Response, error) {
	var res EndpointRequestHeaders
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/request_headers")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteRequestHeadersModuleGet(ctx context.Context, arg *EdgeRouteItem) (*EndpointRequestHeaders, *http.Response, error) {
	var res EndpointRequestHeaders
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/request_headers")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteRequestHeadersModuleDelete(ctx context.Context, arg *EdgeRouteItem) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/request_headers")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteResponseHeadersModuleReplace(ctx context.Context, arg *EdgeRouteResponseHeadersReplace) (*EndpointResponseHeaders, *http.Response, error) {
	var res EndpointResponseHeaders
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/response_headers")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteResponseHeadersModuleGet(ctx context.Context, arg *EdgeRouteItem) (*EndpointResponseHeaders, *http.Response, error) {
	var res EndpointResponseHeaders
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/response_headers")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteResponseHeadersModuleDelete(ctx context.Context, arg *EdgeRouteItem) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/response_headers")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteCompressionModuleReplace(ctx context.Context, arg *EdgeRouteCompressionReplace) (*EndpointCompression, *http.Response, error) {
	var res EndpointCompression
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/compression")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteCompressionModuleGet(ctx context.Context, arg *EdgeRouteItem) (*EndpointCompression, *http.Response, error) {
	var res EndpointCompression
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/compression")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteCompressionModuleDelete(ctx context.Context, arg *EdgeRouteItem) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/compression")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteCircuitBreakerModuleReplace(ctx context.Context, arg *EdgeRouteCircuitBreakerReplace) (*EndpointCircuitBreaker, *http.Response, error) {
	var res EndpointCircuitBreaker
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/circuit_breaker")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteCircuitBreakerModuleGet(ctx context.Context, arg *EdgeRouteItem) (*EndpointCircuitBreaker, *http.Response, error) {
	var res EndpointCircuitBreaker
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/circuit_breaker")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteCircuitBreakerModuleDelete(ctx context.Context, arg *EdgeRouteItem) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/circuit_breaker")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteWebhookVerificationModuleReplace(ctx context.Context, arg *EdgeRouteWebhookVerificationReplace) (*EndpointWebhookValidation, *http.Response, error) {
	var res EndpointWebhookValidation
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/webhook_verification")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteWebhookVerificationModuleGet(ctx context.Context, arg *EdgeRouteItem) (*EndpointWebhookValidation, *http.Response, error) {
	var res EndpointWebhookValidation
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/webhook_verification")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteWebhookVerificationModuleDelete(ctx context.Context, arg *EdgeRouteItem) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/webhook_verification")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteOAuthModuleReplace(ctx context.Context, arg *EdgeRouteOAuthReplace) (*EndpointOAuth, *http.Response, error) {
	var res EndpointOAuth
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/oauth")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteOAuthModuleGet(ctx context.Context, arg *EdgeRouteItem) (*EndpointOAuth, *http.Response, error) {
	var res EndpointOAuth
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/oauth")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteOAuthModuleDelete(ctx context.Context, arg *EdgeRouteItem) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/oauth")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteSAMLModuleReplace(ctx context.Context, arg *EdgeRouteSAMLReplace) (*EndpointSAML, *http.Response, error) {
	var res EndpointSAML
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/saml")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteSAMLModuleGet(ctx context.Context, arg *EdgeRouteItem) (*EndpointSAML, *http.Response, error) {
	var res EndpointSAML
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/saml")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteSAMLModuleDelete(ctx context.Context, arg *EdgeRouteItem) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/saml")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteOIDCModuleReplace(ctx context.Context, arg *EdgeRouteOIDCReplace) (*EndpointOIDC, *http.Response, error) {
	var res EndpointOIDC
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/oidc")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteOIDCModuleGet(ctx context.Context, arg *EdgeRouteItem) (*EndpointOIDC, *http.Response, error) {
	var res EndpointOIDC
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/oidc")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteOIDCModuleDelete(ctx context.Context, arg *EdgeRouteItem) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/oidc")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteWebsocketTCPConverterModuleReplace(ctx context.Context, arg *EdgeRouteWebsocketTCPConverterReplace) (*EndpointWebsocketTCPConverter, *http.Response, error) {
	var res EndpointWebsocketTCPConverter
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/websocket_tcp_converter")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteWebsocketTCPConverterModuleGet(ctx context.Context, arg *EdgeRouteItem) (*EndpointWebsocketTCPConverter, *http.Response, error) {
	var res EndpointWebsocketTCPConverter
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/websocket_tcp_converter")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteWebsocketTCPConverterModuleDelete(ctx context.Context, arg *EdgeRouteItem) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/websocket_tcp_converter")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteUserAgentFilterModuleReplace(ctx context.Context, arg *EdgeRouteUserAgentFilterReplace) (*EndpointUserAgentFilter, *http.Response, error) {
	var res EndpointUserAgentFilter
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/user_agent_filter")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteUserAgentFilterModuleGet(ctx context.Context, arg *EdgeRouteItem) (*EndpointUserAgentFilter, *http.Response, error) {
	var res EndpointUserAgentFilter
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/user_agent_filter")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteUserAgentFilterModuleDelete(ctx context.Context, arg *EdgeRouteItem) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/user_agent_filter")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRoutePolicyModuleReplace(ctx context.Context, arg *EdgeRoutePolicyReplace) (*EndpointPolicy, *http.Response, error) {
	var res EndpointPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/policy")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRoutePolicyModuleGet(ctx context.Context, arg *EdgeRouteItem) (*EndpointPolicy, *http.Response, error) {
	var res EndpointPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/policy")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRoutePolicyModuleDelete(ctx context.Context, arg *EdgeRouteItem) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/policy")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteTrafficPolicyModuleReplace(ctx context.Context, arg *EdgeRouteTrafficPolicyReplace) (*EndpointTrafficPolicy, *http.Response, error) {
	var res EndpointTrafficPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/traffic_policy")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteTrafficPolicyModuleGet(ctx context.Context, arg *EdgeRouteItem) (*EndpointTrafficPolicy, *http.Response, error) {
	var res EndpointTrafficPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/traffic_policy")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EdgeRouteTrafficPolicyModuleDelete(ctx context.Context, arg *EdgeRouteItem) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/traffic_policy")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.EdgeID = ""
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create a TCP Edge
func (c *Client) EdgesTCPCreate(ctx context.Context, arg *TCPEdgeCreate) (*TCPEdge, *http.Response, error) {
	var res TCPEdge
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tcp")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get a TCP Edge by ID
func (c *Client) EdgesTCPGet(ctx context.Context, arg *Item) (*TCPEdge, *http.Response, error) {
	var res TCPEdge
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tcp/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Returns a list of all TCP Edges on this account
func (c *Client) EdgesTCPList(ctx context.Context, arg *Paging) (*TCPEdgeList, *http.Response, error) {
	var res TCPEdgeList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tcp")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Updates a TCP Edge by ID. If a module is not specified in the update, it will not be modified. However, each module configuration that is specified will completely replace the existing value. There is no way to delete an existing module via this API, instead use the delete module API.
func (c *Client) EdgesTCPUpdate(ctx context.Context, arg *TCPEdgeUpdate) (*TCPEdge, *http.Response, error) {
	var res TCPEdge
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tcp/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete a TCP Edge by ID
func (c *Client) EdgesTCPDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tcp/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TCPEdgeBackendModuleReplace(ctx context.Context, arg *EdgeBackendReplace) (*EndpointBackend, *http.Response, error) {
	var res EndpointBackend
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tcp/{{ .ID }}/backend")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TCPEdgeBackendModuleGet(ctx context.Context, arg *Item) (*EndpointBackend, *http.Response, error) {
	var res EndpointBackend
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tcp/{{ .ID }}/backend")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TCPEdgeBackendModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tcp/{{ .ID }}/backend")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TCPEdgeIPRestrictionModuleReplace(ctx context.Context, arg *EdgeIPRestrictionReplace) (*EndpointIPPolicy, *http.Response, error) {
	var res EndpointIPPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tcp/{{ .ID }}/ip_restriction")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TCPEdgeIPRestrictionModuleGet(ctx context.Context, arg *Item) (*EndpointIPPolicy, *http.Response, error) {
	var res EndpointIPPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tcp/{{ .ID }}/ip_restriction")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TCPEdgeIPRestrictionModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tcp/{{ .ID }}/ip_restriction")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TCPEdgePolicyModuleReplace(ctx context.Context, arg *EdgePolicyReplace) (*EndpointPolicy, *http.Response, error) {
	var res EndpointPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tcp/{{ .ID }}/policy")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TCPEdgePolicyModuleGet(ctx context.Context, arg *Item) (*EndpointPolicy, *http.Response, error) {
	var res EndpointPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tcp/{{ .ID }}/policy")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TCPEdgePolicyModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tcp/{{ .ID }}/policy")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TCPEdgeTrafficPolicyModuleReplace(ctx context.Context, arg *EdgeTrafficPolicyReplace) (*EndpointTrafficPolicy, *http.Response, error) {
	var res EndpointTrafficPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tcp/{{ .ID }}/traffic_policy")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TCPEdgeTrafficPolicyModuleGet(ctx context.Context, arg *Item) (*EndpointTrafficPolicy, *http.Response, error) {
	var res EndpointTrafficPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tcp/{{ .ID }}/traffic_policy")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TCPEdgeTrafficPolicyModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tcp/{{ .ID }}/traffic_policy")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create a TLS Edge
func (c *Client) EdgesTLSCreate(ctx context.Context, arg *TLSEdgeCreate) (*TLSEdge, *http.Response, error) {
	var res TLSEdge
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tls")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get a TLS Edge by ID
func (c *Client) EdgesTLSGet(ctx context.Context, arg *Item) (*TLSEdge, *http.Response, error) {
	var res TLSEdge
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tls/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Returns a list of all TLS Edges on this account
func (c *Client) EdgesTLSList(ctx context.Context, arg *Paging) (*TLSEdgeList, *http.Response, error) {
	var res TLSEdgeList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tls")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Updates a TLS Edge by ID. If a module is not specified in the update, it will not be modified. However, each module configuration that is specified will completely replace the existing value. There is no way to delete an existing module via this API, instead use the delete module API.
func (c *Client) EdgesTLSUpdate(ctx context.Context, arg *TLSEdgeUpdate) (*TLSEdge, *http.Response, error) {
	var res TLSEdge
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tls/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete a TLS Edge by ID
func (c *Client) EdgesTLSDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tls/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TLSEdgeBackendModuleReplace(ctx context.Context, arg *EdgeBackendReplace) (*EndpointBackend, *http.Response, error) {
	var res EndpointBackend
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tls/{{ .ID }}/backend")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TLSEdgeBackendModuleGet(ctx context.Context, arg *Item) (*EndpointBackend, *http.Response, error) {
	var res EndpointBackend
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tls/{{ .ID }}/backend")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TLSEdgeBackendModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tls/{{ .ID }}/backend")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TLSEdgeIPRestrictionModuleReplace(ctx context.Context, arg *EdgeIPRestrictionReplace) (*EndpointIPPolicy, *http.Response, error) {
	var res EndpointIPPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tls/{{ .ID }}/ip_restriction")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TLSEdgeIPRestrictionModuleGet(ctx context.Context, arg *Item) (*EndpointIPPolicy, *http.Response, error) {
	var res EndpointIPPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tls/{{ .ID }}/ip_restriction")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TLSEdgeIPRestrictionModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tls/{{ .ID }}/ip_restriction")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TLSEdgeMutualTLSModuleReplace(ctx context.Context, arg *EdgeMutualTLSReplace) (*EndpointMutualTLS, *http.Response, error) {
	var res EndpointMutualTLS
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tls/{{ .ID }}/mutual_tls")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TLSEdgeMutualTLSModuleGet(ctx context.Context, arg *Item) (*EndpointMutualTLS, *http.Response, error) {
	var res EndpointMutualTLS
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tls/{{ .ID }}/mutual_tls")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TLSEdgeMutualTLSModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tls/{{ .ID }}/mutual_tls")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TLSEdgeTLSTerminationModuleReplace(ctx context.Context, arg *EdgeTLSTerminationReplace) (*EndpointTLSTermination, *http.Response, error) {
	var res EndpointTLSTermination
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tls/{{ .ID }}/tls_termination")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TLSEdgeTLSTerminationModuleGet(ctx context.Context, arg *Item) (*EndpointTLSTermination, *http.Response, error) {
	var res EndpointTLSTermination
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tls/{{ .ID }}/tls_termination")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TLSEdgeTLSTerminationModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tls/{{ .ID }}/tls_termination")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TLSEdgePolicyModuleReplace(ctx context.Context, arg *EdgePolicyReplace) (*EndpointPolicy, *http.Response, error) {
	var res EndpointPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tls/{{ .ID }}/policy")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TLSEdgePolicyModuleGet(ctx context.Context, arg *Item) (*EndpointPolicy, *http.Response, error) {
	var res EndpointPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tls/{{ .ID }}/policy")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TLSEdgePolicyModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tls/{{ .ID }}/policy")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TLSEdgeTrafficPolicyModuleReplace(ctx context.Context, arg *EdgeTrafficPolicyReplace) (*EndpointTrafficPolicy, *http.Response, error) {
	var res EndpointTrafficPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tls/{{ .ID }}/traffic_policy")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TLSEdgeTrafficPolicyModuleGet(ctx context.Context, arg *Item) (*EndpointTrafficPolicy, *http.Response, error) {
	var res EndpointTrafficPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tls/{{ .ID }}/traffic_policy")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) TLSEdgeTrafficPolicyModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/edges/tls/{{ .ID }}/traffic_policy")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create an endpoint, currently available only for cloud endpoints
func (c *Client) EndpointsCreate(ctx context.Context, arg *EndpointCreate) (*Endpoint, *http.Response, error) {
	var res Endpoint
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoints")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all active endpoints on the account
func (c *Client) EndpointsList(ctx context.Context, arg *Paging) (*EndpointList, *http.Response, error) {
	var res EndpointList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoints")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get the status of an endpoint by ID
func (c *Client) EndpointsGet(ctx context.Context, arg *Item) (*Endpoint, *http.Response, error) {
	var res Endpoint
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoints/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update an Endpoint by ID, currently available only for cloud endpoints
func (c *Client) EndpointsUpdate(ctx context.Context, arg *EndpointUpdate) (*Endpoint, *http.Response, error) {
	var res Endpoint
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoints/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete an Endpoint by ID, currently available only for cloud endpoints
func (c *Client) EndpointsDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoints/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create a new Event Destination. It will not apply to anything until it is associated with an Event Subscription.
func (c *Client) EventDestinationsCreate(ctx context.Context, arg *EventDestinationCreate) (*EventDestination, *http.Response, error) {
	var res EventDestination
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/event_destinations")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete an Event Destination. If the Event Destination is still referenced by an Event Subscription.
func (c *Client) EventDestinationsDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/event_destinations/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get detailed information about an Event Destination by ID.
func (c *Client) EventDestinationsGet(ctx context.Context, arg *Item) (*EventDestination, *http.Response, error) {
	var res EventDestination
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/event_destinations/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all Event Destinations on this account.
func (c *Client) EventDestinationsList(ctx context.Context, arg *Paging) (*EventDestinationList, *http.Response, error) {
	var res EventDestinationList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/event_destinations")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update attributes of an Event Destination.
func (c *Client) EventDestinationsUpdate(ctx context.Context, arg *EventDestinationUpdate) (*EventDestination, *http.Response, error) {
	var res EventDestination
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/event_destinations/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Send a test event to an Event Destination
func (c *Client) EventDestinationsSendTestEvent(ctx context.Context, arg *Item) (*SentEvent, *http.Response, error) {
	var res SentEvent
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/event_destinations/{{ .ID }}/send_test_event")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create an Event Subscription.
func (c *Client) EventSubscriptionsCreate(ctx context.Context, arg *EventSubscriptionCreate) (*EventSubscription, *http.Response, error) {
	var res EventSubscription
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/event_subscriptions")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete an Event Subscription.
func (c *Client) EventSubscriptionsDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/event_subscriptions/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get an Event Subscription by ID.
func (c *Client) EventSubscriptionsGet(ctx context.Context, arg *Item) (*EventSubscription, *http.Response, error) {
	var res EventSubscription
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/event_subscriptions/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List this Account's Event Subscriptions.
func (c *Client) EventSubscriptionsList(ctx context.Context, arg *Paging) (*EventSubscriptionList, *http.Response, error) {
	var res EventSubscriptionList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/event_subscriptions")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update an Event Subscription.
func (c *Client) EventSubscriptionsUpdate(ctx context.Context, arg *EventSubscriptionUpdate) (*EventSubscription, *http.Response, error) {
	var res EventSubscription
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/event_subscriptions/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Add an additional type for which this event subscription will trigger
func (c *Client) EventSourcesCreate(ctx context.Context, arg *EventSourceCreate) (*EventSource, *http.Response, error) {
	var res EventSource
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/event_subscriptions/{{ .SubscriptionID }}/sources")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.SubscriptionID = ""

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Remove a type for which this event subscription will trigger
func (c *Client) EventSourcesDelete(ctx context.Context, arg *EventSourceItem) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/event_subscriptions/{{ .SubscriptionID }}/sources/{{ .Type }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.SubscriptionID = ""
	arg.Type = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get the details for a given type that triggers for the given event subscription
func (c *Client) EventSourcesGet(ctx context.Context, arg *EventSourceItem) (*EventSource, *http.Response, error) {
	var res EventSource
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/event_subscriptions/{{ .SubscriptionID }}/sources/{{ .Type }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.SubscriptionID = ""
	arg.Type = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List the types for which this event subscription will trigger
func (c *Client) EventSourcesList(ctx context.Context, arg *EventSourcePaging) (*EventSourceList, *http.Response, error) {
	var res EventSourceList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/event_subscriptions/{{ .SubscriptionID }}/sources")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.SubscriptionID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update the type for which this event subscription will trigger
func (c *Client) EventSourcesUpdate(ctx context.Context, arg *EventSourceUpdate) (*EventSource, *http.Response, error) {
	var res EventSource
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/event_subscriptions/{{ .SubscriptionID }}/sources/{{ .Type }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.SubscriptionID = ""
	arg.Type = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create a new IP policy. It will not apply to any traffic until you associate to a traffic source via an endpoint configuration or IP restriction.
func (c *Client) IPPoliciesCreate(ctx context.Context, arg *IPPolicyCreate) (*IPPolicy, *http.Response, error) {
	var res IPPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ip_policies")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete an IP policy. If the IP policy is referenced by another object for the purposes of traffic restriction it will be treated as if the IP policy remains but has zero rules.
func (c *Client) IPPoliciesDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ip_policies/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get detailed information about an IP policy by ID.
func (c *Client) IPPoliciesGet(ctx context.Context, arg *Item) (*IPPolicy, *http.Response, error) {
	var res IPPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ip_policies/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all IP policies on this account
func (c *Client) IPPoliciesList(ctx context.Context, arg *Paging) (*IPPolicyList, *http.Response, error) {
	var res IPPolicyList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ip_policies")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update attributes of an IP policy by ID
func (c *Client) IPPoliciesUpdate(ctx context.Context, arg *IPPolicyUpdate) (*IPPolicy, *http.Response, error) {
	var res IPPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ip_policies/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create a new IP policy rule attached to an IP Policy.
func (c *Client) IPPolicyRulesCreate(ctx context.Context, arg *IPPolicyRuleCreate) (*IPPolicyRule, *http.Response, error) {
	var res IPPolicyRule
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ip_policy_rules")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete an IP policy rule.
func (c *Client) IPPolicyRulesDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ip_policy_rules/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get detailed information about an IP policy rule by ID.
func (c *Client) IPPolicyRulesGet(ctx context.Context, arg *Item) (*IPPolicyRule, *http.Response, error) {
	var res IPPolicyRule
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ip_policy_rules/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all IP policy rules on this account
func (c *Client) IPPolicyRulesList(ctx context.Context, arg *Paging) (*IPPolicyRuleList, *http.Response, error) {
	var res IPPolicyRuleList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ip_policy_rules")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update attributes of an IP policy rule by ID
func (c *Client) IPPolicyRulesUpdate(ctx context.Context, arg *IPPolicyRuleUpdate) (*IPPolicyRule, *http.Response, error) {
	var res IPPolicyRule
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ip_policy_rules/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create a new IP restriction
func (c *Client) IPRestrictionsCreate(ctx context.Context, arg *IPRestrictionCreate) (*IPRestriction, *http.Response, error) {
	var res IPRestriction
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ip_restrictions")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete an IP restriction
func (c *Client) IPRestrictionsDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ip_restrictions/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get detailed information about an IP restriction
func (c *Client) IPRestrictionsGet(ctx context.Context, arg *Item) (*IPRestriction, *http.Response, error) {
	var res IPRestriction
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ip_restrictions/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all IP restrictions on this account
func (c *Client) IPRestrictionsList(ctx context.Context, arg *Paging) (*IPRestrictionList, *http.Response, error) {
	var res IPRestrictionList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ip_restrictions")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update attributes of an IP restriction by ID
func (c *Client) IPRestrictionsUpdate(ctx context.Context, arg *IPRestrictionUpdate) (*IPRestriction, *http.Response, error) {
	var res IPRestriction
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ip_restrictions/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create a new Kubernetes Operator
func (c *Client) KubernetesOperatorsCreate(ctx context.Context, arg *KubernetesOperatorCreate) (*KubernetesOperator, *http.Response, error) {
	var res KubernetesOperator
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/kubernetes_operators")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update an existing Kubernetes operator by ID.
func (c *Client) KubernetesOperatorsUpdate(ctx context.Context, arg *KubernetesOperatorUpdate) (*KubernetesOperator, *http.Response, error) {
	var res KubernetesOperator
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/kubernetes_operators/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete a Kubernetes Operator
func (c *Client) KubernetesOperatorsDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/kubernetes_operators/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get of a Kubernetes Operator
func (c *Client) KubernetesOperatorsGet(ctx context.Context, arg *Item) (*KubernetesOperator, *http.Response, error) {
	var res KubernetesOperator
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/kubernetes_operators/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all Kubernetes Operators owned by this account
func (c *Client) KubernetesOperatorsList(ctx context.Context, arg *Paging) (*KubernetesOperatorList, *http.Response, error) {
	var res KubernetesOperatorList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/kubernetes_operators")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointBasicAuthModuleReplace(ctx context.Context, arg *EndpointBasicAuthReplace) (*EndpointBasicAuth, *http.Response, error) {
	var res EndpointBasicAuth
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/basic_auth")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointBasicAuthModuleGet(ctx context.Context, arg *Item) (*EndpointBasicAuth, *http.Response, error) {
	var res EndpointBasicAuth
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/basic_auth")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointBasicAuthModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/basic_auth")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointCircuitBreakerModuleReplace(ctx context.Context, arg *EndpointCircuitBreakerReplace) (*EndpointCircuitBreaker, *http.Response, error) {
	var res EndpointCircuitBreaker
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/circuit_breaker")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointCircuitBreakerModuleGet(ctx context.Context, arg *Item) (*EndpointCircuitBreaker, *http.Response, error) {
	var res EndpointCircuitBreaker
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/circuit_breaker")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointCircuitBreakerModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/circuit_breaker")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointCompressionModuleReplace(ctx context.Context, arg *EndpointCompressionReplace) (*EndpointCompression, *http.Response, error) {
	var res EndpointCompression
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/compression")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointCompressionModuleGet(ctx context.Context, arg *Item) (*EndpointCompression, *http.Response, error) {
	var res EndpointCompression
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/compression")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointCompressionModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/compression")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointTLSTerminationModuleReplace(ctx context.Context, arg *EndpointTLSTerminationReplace) (*EndpointTLSTermination, *http.Response, error) {
	var res EndpointTLSTermination
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/tls_termination")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointTLSTerminationModuleGet(ctx context.Context, arg *Item) (*EndpointTLSTermination, *http.Response, error) {
	var res EndpointTLSTermination
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/tls_termination")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointTLSTerminationModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/tls_termination")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointIPPolicyModuleReplace(ctx context.Context, arg *EndpointIPPolicyReplace) (*EndpointIPPolicy, *http.Response, error) {
	var res EndpointIPPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/ip_policy")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointIPPolicyModuleGet(ctx context.Context, arg *Item) (*EndpointIPPolicy, *http.Response, error) {
	var res EndpointIPPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/ip_policy")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointIPPolicyModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/ip_policy")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointMutualTLSModuleReplace(ctx context.Context, arg *EndpointMutualTLSReplace) (*EndpointMutualTLS, *http.Response, error) {
	var res EndpointMutualTLS
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/mutual_tls")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointMutualTLSModuleGet(ctx context.Context, arg *Item) (*EndpointMutualTLS, *http.Response, error) {
	var res EndpointMutualTLS
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/mutual_tls")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointMutualTLSModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/mutual_tls")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointRequestHeadersModuleReplace(ctx context.Context, arg *EndpointRequestHeadersReplace) (*EndpointRequestHeaders, *http.Response, error) {
	var res EndpointRequestHeaders
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/request_headers")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointRequestHeadersModuleGet(ctx context.Context, arg *Item) (*EndpointRequestHeaders, *http.Response, error) {
	var res EndpointRequestHeaders
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/request_headers")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointRequestHeadersModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/request_headers")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointResponseHeadersModuleReplace(ctx context.Context, arg *EndpointResponseHeadersReplace) (*EndpointResponseHeaders, *http.Response, error) {
	var res EndpointResponseHeaders
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/response_headers")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointResponseHeadersModuleGet(ctx context.Context, arg *Item) (*EndpointResponseHeaders, *http.Response, error) {
	var res EndpointResponseHeaders
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/response_headers")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointResponseHeadersModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/response_headers")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointOAuthModuleReplace(ctx context.Context, arg *EndpointOAuthReplace) (*EndpointOAuth, *http.Response, error) {
	var res EndpointOAuth
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/oauth")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointOAuthModuleGet(ctx context.Context, arg *Item) (*EndpointOAuth, *http.Response, error) {
	var res EndpointOAuth
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/oauth")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointOAuthModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/oauth")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointWebhookValidationModuleReplace(ctx context.Context, arg *EndpointWebhookValidationReplace) (*EndpointWebhookValidation, *http.Response, error) {
	var res EndpointWebhookValidation
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/webhook_validation")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointWebhookValidationModuleGet(ctx context.Context, arg *Item) (*EndpointWebhookValidation, *http.Response, error) {
	var res EndpointWebhookValidation
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/webhook_validation")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointWebhookValidationModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/webhook_validation")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointSAMLModuleReplace(ctx context.Context, arg *EndpointSAMLReplace) (*EndpointSAML, *http.Response, error) {
	var res EndpointSAML
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/saml")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointSAMLModuleGet(ctx context.Context, arg *Item) (*EndpointSAML, *http.Response, error) {
	var res EndpointSAML
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/saml")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointSAMLModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/saml")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointOIDCModuleReplace(ctx context.Context, arg *EndpointOIDCReplace) (*EndpointOIDC, *http.Response, error) {
	var res EndpointOIDC
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/oidc")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointOIDCModuleGet(ctx context.Context, arg *Item) (*EndpointOIDC, *http.Response, error) {
	var res EndpointOIDC
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/oidc")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointOIDCModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/oidc")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointBackendModuleReplace(ctx context.Context, arg *EndpointBackendReplace) (*EndpointBackend, *http.Response, error) {
	var res EndpointBackend
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/backend")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Put(ctx, uri, arg.Module, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointBackendModuleGet(ctx context.Context, arg *Item) (*EndpointBackend, *http.Response, error) {
	var res EndpointBackend
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/backend")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) EndpointBackendModuleDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/endpoint_configurations/{{ .ID }}/backend")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create a new reserved address.
func (c *Client) ReservedAddrsCreate(ctx context.Context, arg *ReservedAddrCreate) (*ReservedAddr, *http.Response, error) {
	var res ReservedAddr
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/reserved_addrs")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete a reserved address.
func (c *Client) ReservedAddrsDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/reserved_addrs/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get the details of a reserved address.
func (c *Client) ReservedAddrsGet(ctx context.Context, arg *Item) (*ReservedAddr, *http.Response, error) {
	var res ReservedAddr
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/reserved_addrs/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all reserved addresses on this account.
func (c *Client) ReservedAddrsList(ctx context.Context, arg *Paging) (*ReservedAddrList, *http.Response, error) {
	var res ReservedAddrList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/reserved_addrs")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update the attributes of a reserved address.
func (c *Client) ReservedAddrsUpdate(ctx context.Context, arg *ReservedAddrUpdate) (*ReservedAddr, *http.Response, error) {
	var res ReservedAddr
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/reserved_addrs/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Detach the endpoint configuration attached to a reserved address.
func (c *Client) ReservedAddrsDeleteEndpointConfig(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/reserved_addrs/{{ .ID }}/endpoint_configuration")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create a new reserved domain.
func (c *Client) ReservedDomainsCreate(ctx context.Context, arg *ReservedDomainCreate) (*ReservedDomain, *http.Response, error) {
	var res ReservedDomain
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/reserved_domains")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete a reserved domain.
func (c *Client) ReservedDomainsDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/reserved_domains/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get the details of a reserved domain.
func (c *Client) ReservedDomainsGet(ctx context.Context, arg *Item) (*ReservedDomain, *http.Response, error) {
	var res ReservedDomain
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/reserved_domains/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all reserved domains on this account.
func (c *Client) ReservedDomainsList(ctx context.Context, arg *Paging) (*ReservedDomainList, *http.Response, error) {
	var res ReservedDomainList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/reserved_domains")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update the attributes of a reserved domain.
func (c *Client) ReservedDomainsUpdate(ctx context.Context, arg *ReservedDomainUpdate) (*ReservedDomain, *http.Response, error) {
	var res ReservedDomain
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/reserved_domains/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Detach the certificate management policy attached to a reserved domain.
func (c *Client) ReservedDomainsDeleteCertificateManagementPolicy(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/reserved_domains/{{ .ID }}/certificate_management_policy")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Detach the certificate attached to a reserved domain.
func (c *Client) ReservedDomainsDeleteCertificate(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/reserved_domains/{{ .ID }}/certificate")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Detach the http endpoint configuration attached to a reserved domain.
func (c *Client) ReservedDomainsDeleteHTTPEndpointConfig(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/reserved_domains/{{ .ID }}/http_endpoint_configuration")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Detach the https endpoint configuration attached to a reserved domain.
func (c *Client) ReservedDomainsDeleteHTTPSEndpointConfig(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/reserved_domains/{{ .ID }}/https_endpoint_configuration")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

func (c *Client) RootGet(ctx context.Context, arg *Empty) (*RootResponse, *http.Response, error) {
	var res RootResponse
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create a new SSH Certificate Authority
func (c *Client) SSHCertificateAuthoritiesCreate(ctx context.Context, arg *SSHCertificateAuthorityCreate) (*SSHCertificateAuthority, *http.Response, error) {
	var res SSHCertificateAuthority
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ssh_certificate_authorities")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete an SSH Certificate Authority
func (c *Client) SSHCertificateAuthoritiesDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ssh_certificate_authorities/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get detailed information about an SSH Certficate Authority
func (c *Client) SSHCertificateAuthoritiesGet(ctx context.Context, arg *Item) (*SSHCertificateAuthority, *http.Response, error) {
	var res SSHCertificateAuthority
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ssh_certificate_authorities/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all SSH Certificate Authorities on this account
func (c *Client) SSHCertificateAuthoritiesList(ctx context.Context, arg *Paging) (*SSHCertificateAuthorityList, *http.Response, error) {
	var res SSHCertificateAuthorityList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ssh_certificate_authorities")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update an SSH Certificate Authority
func (c *Client) SSHCertificateAuthoritiesUpdate(ctx context.Context, arg *SSHCertificateAuthorityUpdate) (*SSHCertificateAuthority, *http.Response, error) {
	var res SSHCertificateAuthority
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ssh_certificate_authorities/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create a new ssh_credential from an uploaded public SSH key. This ssh credential can be used to start new tunnels via ngrok's SSH gateway.
func (c *Client) SSHCredentialsCreate(ctx context.Context, arg *SSHCredentialCreate) (*SSHCredential, *http.Response, error) {
	var res SSHCredential
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ssh_credentials")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete an ssh_credential by ID
func (c *Client) SSHCredentialsDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ssh_credentials/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get detailed information about an ssh_credential
func (c *Client) SSHCredentialsGet(ctx context.Context, arg *Item) (*SSHCredential, *http.Response, error) {
	var res SSHCredential
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ssh_credentials/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all ssh credentials on this account
func (c *Client) SSHCredentialsList(ctx context.Context, arg *Paging) (*SSHCredentialList, *http.Response, error) {
	var res SSHCredentialList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ssh_credentials")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update attributes of an ssh_credential by ID
func (c *Client) SSHCredentialsUpdate(ctx context.Context, arg *SSHCredentialUpdate) (*SSHCredential, *http.Response, error) {
	var res SSHCredential
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ssh_credentials/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create a new SSH Host Certificate
func (c *Client) SSHHostCertificatesCreate(ctx context.Context, arg *SSHHostCertificateCreate) (*SSHHostCertificate, *http.Response, error) {
	var res SSHHostCertificate
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ssh_host_certificates")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete an SSH Host Certificate
func (c *Client) SSHHostCertificatesDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ssh_host_certificates/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get detailed information about an SSH Host Certficate
func (c *Client) SSHHostCertificatesGet(ctx context.Context, arg *Item) (*SSHHostCertificate, *http.Response, error) {
	var res SSHHostCertificate
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ssh_host_certificates/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all SSH Host Certificates issued on this account
func (c *Client) SSHHostCertificatesList(ctx context.Context, arg *Paging) (*SSHHostCertificateList, *http.Response, error) {
	var res SSHHostCertificateList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ssh_host_certificates")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update an SSH Host Certificate
func (c *Client) SSHHostCertificatesUpdate(ctx context.Context, arg *SSHHostCertificateUpdate) (*SSHHostCertificate, *http.Response, error) {
	var res SSHHostCertificate
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ssh_host_certificates/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Create a new SSH User Certificate
func (c *Client) SSHUserCertificatesCreate(ctx context.Context, arg *SSHUserCertificateCreate) (*SSHUserCertificate, *http.Response, error) {
	var res SSHUserCertificate
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ssh_user_certificates")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete an SSH User Certificate
func (c *Client) SSHUserCertificatesDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ssh_user_certificates/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get detailed information about an SSH User Certficate
func (c *Client) SSHUserCertificatesGet(ctx context.Context, arg *Item) (*SSHUserCertificate, *http.Response, error) {
	var res SSHUserCertificate
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ssh_user_certificates/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all SSH User Certificates issued on this account
func (c *Client) SSHUserCertificatesList(ctx context.Context, arg *Paging) (*SSHUserCertificateList, *http.Response, error) {
	var res SSHUserCertificateList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ssh_user_certificates")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update an SSH User Certificate
func (c *Client) SSHUserCertificatesUpdate(ctx context.Context, arg *SSHUserCertificateUpdate) (*SSHUserCertificate, *http.Response, error) {
	var res SSHUserCertificate
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/ssh_user_certificates/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Upload a new TLS certificate
func (c *Client) TLSCertificatesCreate(ctx context.Context, arg *TLSCertificateCreate) (*TLSCertificate, *http.Response, error) {
	var res TLSCertificate
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/tls_certificates")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()

	resp, err := c.Post(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Delete a TLS certificate
func (c *Client) TLSCertificatesDelete(ctx context.Context, arg *Item) (*Empty, *http.Response, error) {
	var res Empty
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/tls_certificates/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Delete(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get detailed information about a TLS certificate
func (c *Client) TLSCertificatesGet(ctx context.Context, arg *Item) (*TLSCertificate, *http.Response, error) {
	var res TLSCertificate
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/tls_certificates/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all TLS certificates on this account
func (c *Client) TLSCertificatesList(ctx context.Context, arg *Paging) (*TLSCertificateList, *http.Response, error) {
	var res TLSCertificateList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/tls_certificates")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Update attributes of a TLS Certificate by ID
func (c *Client) TLSCertificatesUpdate(ctx context.Context, arg *TLSCertificateUpdate) (*TLSCertificate, *http.Response, error) {
	var res TLSCertificate
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/tls_certificates/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Patch(ctx, uri, arg, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// List all online tunnels currently running on the account.
func (c *Client) TunnelsList(ctx context.Context, arg *Paging) (*TunnelList, *http.Response, error) {
	var res TunnelList
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/tunnels")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	pathUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	if arg.BeforeID != nil {
		params.Add("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		params.Add("limit", *arg.Limit)
	}
	pathUrl.RawQuery = params.Encode()
	uri = pathUrl.String()

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}

// Get the status of a tunnel by ID
func (c *Client) TunnelsGet(ctx context.Context, arg *Item) (*Tunnel, *http.Response, error) {
	var res Tunnel
	var path bytes.Buffer
	if err := template.Must(template.New("").Parse("/tunnels/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	uri := path.String()
	arg.ID = ""

	resp, err := c.Get(ctx, uri, &res)
	if errors.Is(err, io.EOF) && resp != nil && resp.StatusCode == 204 {
		err = nil
	}
	return &res, resp, err
}
