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

* `basic_auth` - (Optional) basic auth module configuration or <code>null</code>
* `circuit_breaker` - (Optional) circuit breaker module configuration or <code>null</code>
* `compression` - (Optional) compression module configuration or <code>null</code>
* `description` - (Optional) human-readable description of what this endpoint configuration will be do when applied or what traffic it will be applied to. Optional, max 255 bytes
* `ip_policy` - (Optional) ip policy module configuration or <code>null</code>
* `logging` - (Optional) logging module configuration or <code>null</code>
* `metadata` - (Optional) arbitrary user-defined machine-readable data of this endpoint configuration. Optional, max 4096 bytes.
* `mutual_tls` - (Optional) mutual TLS module configuration or <code>null</code>
* `oauth` - (Optional) oauth module configuration or <code>null</code>
* `oidc` - (Optional) oidc module configuration or <code>null</code>
* `request_headers` - (Optional) request headers module configuration or <code>null</code>
* `response_headers` - (Optional) response headers module configuration or <code>null</code>
* `saml` - (Optional) saml module configuration or <code>null</code>
* `tls_termination` - (Optional) TLS termination module configuration or <code>null</code>
* `type` - (Optional) they type of traffic this endpoint configuration can be applied to. one of: <code>http</code>, <code>https</code>, <code>tcp</code>
* `webhook_validation` - (Optional) webhook validation module configuration or <code>null</code>

## Attribute Reference

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

