# endpoint_configurations Resource
## Endpoint Configuration management

An [Endpoint Configuration](https://ngrok.com.lan/docs/ngrok-link#api-endpoint-configurations) describes
a ngrok network endpoint instance.

_Endpoints are your gateway to ngrok features!_
## Example Usage

Define the Endpoint Configuration resource `ngrok_endpoint_configuration.example`:

```
resource "ngrok_endpoint_configuration" "example" {
  description = "app servers"
  request_headers {
    add = {
      X-Frontend = "ngrok"
    }
    remove = [ "Cache-Control" ]
  }
  type = "https"
}
```

## Argument Reference

* `add` - (Optional) a map of header key to header value that will be injected into the HTTP Response returned to the HTTP client
* `allow_idp_initiated` - (Optional) If true, the IdP may initiate a login directly (e.g. the user does not need to visit the endpoint first and then be redirected). The IdP should set the <code>RelayState</code> parameter to the target URL of the resource they want the user to be redirected to after the SAML login assertion has been processed.
* `allow_options` - (Optional) true or false indicating whether to allow OPTIONS requests through without authentication which is necessary for CORS. default is <code>false</code>
* `auth_check_interval` - (Optional) Integer number of seconds after which ngrok guarantees it will refresh user state from the identity provider and recheck whether the user is still authorized to access the endpoint. This is the preferred tunable to use to enforce a minimum amount of time after which a revoked user will no longer be able to access the resource.
* `auth_provider_id` - (Optional) determines how the basic auth credentials are validated. Currently only the value <code>agent</code> is supported which means that credentials will be validated against the username and password specified by the ngrok agent's <code>-auth</code> flag, if any.
* `authorized_groups` - (Optional) If present, only users who are a member of one of the listed groups may access the target endpoint.
* `backend` - (Optional) backend module configuration or <code>null</code>
* `backend_id` - (Optional) backend to be used to back this endpoint
* `basic_auth` - (Optional) basic auth module configuration or <code>null</code>
* `certificate_authority_ids` - (Optional) list of certificate authorities that will be used to validate the TLS client certificate presnted by the initiatiator of the TLS connection
* `circuit_breaker` - (Optional) circuit breaker module configuration or <code>null</code>
* `client_id` - (Optional) The OIDC app's client ID and OIDC audience.
* `client_secret` - (Optional) The OIDC app's client secret.
* `compression` - (Optional) compression module configuration or <code>null</code>
* `cookie_prefix` - (Optional) the prefix of the session cookie that ngrok sets on the http client to cache authentication. default is 'ngrok.'
* `description` - (Optional) human-readable description of what this endpoint configuration will be do when applied or what traffic it will be applied to. Optional, max 255 bytes
* `email_addresses` - (Optional) a list of email addresses of users authenticated by identity provider who are allowed access to the endpoint
* `email_domains` - (Optional) a list of email domains of users authenticated by identity provider who are allowed access to the endpoint
* `enabled` - (Optional) <code>true</code> if the module will be applied to traffic, <code>false</code> to disable. default <code>true</code> if unspecified
* `error_threshold_percentage` - (Optional) Error threshold percentage should be between 0 - 1.0, not 0-100.0
* `event_stream_ids` - (Optional) list of all EventStreams that will be used to configure and export this endpoint's logs
* `facebook` - (Optional) configuration for using facebook as the identity provider
* `force_authn` - (Optional) If true, indicates that whenever we redirect a user to the IdP for authentication that the IdP must prompt the user for authentication credentials even if the user already has a valid session with the IdP.
* `github` - (Optional) configuration for using github as the identity provider
* `google` - (Optional) configuration for using google as the identity provider
* `idp_metadata` - (Optional) The full XML IdP EntityDescriptor in bytes. This parameter is mutually exclusive with <code>idp_metadata_url</code>. It is recommended to use that parameter instead if the IdP exposes a metadata URL.
* `idp_metadata_url` - (Optional) The IdP's metadata URL which returns the XML IdP EntityDescriptor. The IdP's metadata URL specifies how to connect to the IdP as well as its public key which is then used to validate the signature on incoming SAML assertions to the ACS endpoint.
* `inactivity_timeout` - (Optional) Integer number of seconds of inactivity after which if the user has not accessed the endpoint, their session will time out and they will be forced to reauthenticate.
* `ip_policy` - (Optional) ip policy module configuration or <code>null</code>
* `ip_policy_ids` - (Optional) list of all IP policies that will be used to check if a source IP is allowed access to the endpoint
* `issuer` - (Optional) URL of the OIDC "OpenID provider". This is the base URL used for discovery.
* `logging` - (Optional) logging module configuration or <code>null</code>
* `maximum_duration` - (Optional) Integer number of seconds of the maximum duration of an authenticated session. After this period is exceeded, a user must reauthenticate.
* `metadata` - (Optional) arbitrary user-defined machine-readable data of this endpoint configuration. Optional, max 4096 bytes.
* `microsoft` - (Optional) configuration for using microsoft as the identity provider
* `min_version` - (Optional) The minimum TLS version used for termination and advertised to the client during the TLS handshake. if unspecified, ngrok will choose an industry-safe default. This value must be null if <code>terminate_at</code> is set to <code>upstream</code>.
* `mutual_tls` - (Optional) mutual TLS module configuration or <code>null</code>
* `num_buckets` - (Optional) Integer number of buckets into which metrics are retained. Max 128.
* `oauth` - (Optional) oauth module configuration or <code>null</code>
* `oidc` - (Optional) oidc module configuration or <code>null</code>
* `options_passthrough` - (Optional) Do not enforce authentication on HTTP OPTIONS requests. necessary if you are supporting CORS.
* `organizations` - (Optional) a list of github org identifiers. users who are members of any of the listed organizations will be allowed access. identifiers should be the organization's 'slug'
* `provider` - (Optional) an object which defines the identity provider to use for authentication and configuration for who may access the endpoint
* `realm` - (Optional) an arbitrary string to be specified in as the 'realm' value in the <code>WWW-Authenticate</code> header. default is <code>ngrok</code>
* `remove` - (Optional) a list of header names that will be removed from the HTTP Response returned to the HTTP client
* `request_headers` - (Optional) request headers module configuration or <code>null</code>
* `response_headers` - (Optional) response headers module configuration or <code>null</code>
* `rolling_window` - (Optional) Integer number of seconds in the statistical rolling window that metrics are retained for.
* `saml` - (Optional) saml module configuration or <code>null</code>
* `scopes` - (Optional) The set of scopes to request from the OIDC identity provider.
* `secret` - (Optional) a string secret used to validate requests from the given provider. All providers except AWS SNS require a secret
* `teams` - (Optional) a list of github teams identifiers. users will be allowed access to the endpoint if they are a member of any of these teams. identifiers should be in the 'slug' format qualified with the org name, e.g. <code>org-name/team-name</code>
* `terminate_at` - (Optional) <code>edge</code> if the ngrok edge should terminate TLS traffic, <code>upstream</code> if TLS traffic should be passed through to the upstream ngrok agent / application server for termination. if <code>upstream</code> is chosen, most other modules will be disallowed because they rely on the ngrok edge being able to access the underlying traffic.
* `tls_termination` - (Optional) TLS termination module configuration or <code>null</code>
* `tripped_duration` - (Optional) Integer number of seconds after which the circuit is tripped to wait before re-evaluating upstream health
* `type` - (Optional) they type of traffic this endpoint configuration can be applied to. one of: <code>http</code>, <code>https</code>, <code>tcp</code>
* `volume_threshold` - (Optional) Integer number of requests in a rolling window that will trip the circuit. Helpful if traffic volume is low.
* `webhook_validation` - (Optional) webhook validation module configuration or <code>null</code>

## Attribute Reference

* `backend` - backend module configuration or <code>null</code>
* `basic_auth` - basic auth module configuration or <code>null</code>
* `circuit_breaker` - circuit breaker module configuration or <code>null</code>
* `compression` - compression module configuration or <code>null</code>
* `created_at` - timestamp when the endpoint configuration was created, RFC 3339 format
* `description` - human-readable description of what this endpoint configuration will be do when applied or what traffic it will be applied to. Optional, max 255 bytes
* `ip_policy` - ip policy module configuration or <code>null</code>
* `logging` - logging module configuration or <code>null</code>
* `metadata` - arbitrary user-defined machine-readable data of this endpoint configuration. Optional, max 4096 bytes.
* `mutual_tls` - mutual TLS module configuration or <code>null</code>
* `ngrok_id` - unique identifier of this endpoint configuration
* `oauth` - oauth module configuration or <code>null</code>
* `oidc` - oidc module configuration or <code>null</code>
* `request_headers` - request headers module configuration or <code>null</code>
* `response_headers` - response headers module configuration or <code>null</code>
* `saml` - saml module configuration or <code>null</code>
* `tls_termination` - TLS termination module configuration or <code>null</code>
* `type` - they type of traffic this endpoint configuration can be applied to. one of: <code>http</code>, <code>https</code>, <code>tcp</code>
* `uri` - URI of the endpoint configuration API resource
* `webhook_validation` - webhook validation module configuration or <code>null</code>

