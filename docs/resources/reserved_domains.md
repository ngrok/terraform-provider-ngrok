# reserved_domains Resource

## Example Usage

Define the Reserved Domain resource `ngrok_reserved_domain.example`:

```
resource "ngrok_reserved_domain" "example" {
  certificate_id = "cert_1qWtAbtP9DNiQlh1FfbplVFAFOV"
  name = "myapp.mydomain.com"
  region = "us"
}
```

## Argument Reference

* `authority` - (Optional) certificate authority to request certificates from. The only supported value is letsencrypt.
* `certificate_id` - (Optional) ID of a user-uploaded TLS certificate to use for connections to targeting this domain. Optional, mutually exclusive with `certificate_management_policy`.
* `certificate_management_policy` - (Optional) configuration for automatic management of TLS certificates for this domain, or null if automatic management is disabled. Optional, mutually exclusive with `certificate_id`.
* `description` - (Optional) human-readable description of what this reserved domain will be used for
* `http_endpoint_configuration_id` - (Optional) ID of an endpoint configuration of type http that will be used to handle inbound http traffic to this domain
* `https_endpoint_configuration_id` - (Optional) ID of an endpoint configuration of type https that will be used to handle inbound https traffic to this domain
* `metadata` - (Optional) arbitrary user-defined machine-readable data of this reserved domain. Optional, max 4096 bytes.
* `name` - (Required) the domain name to reserve. It may be a full domain name like app.example.com. If the name does not contain a '.' it will reserve that subdomain on ngrok.io.
* `private_key_type` - (Optional) type of private key to use when requesting certificates. Defaults to rsa, can be either rsa or ecdsa.
* `region` - (Optional) reserve the domain in this geographic ngrok datacenter. Optional, default is us. (au, eu, ap, us, jp, in, sa)

## Attribute Reference

* `certificate` - object referencing the TLS certificate used for connections to this domain. This can be either a user-uploaded certificate, the most recently issued automatic one, or null otherwise.
* `certificate_id` - ID of a user-uploaded TLS certificate to use for connections to targeting this domain. Optional, mutually exclusive with `certificate_management_policy`.
* `certificate_management_policy` - configuration for automatic management of TLS certificates for this domain, or null if automatic management is disabled
* `certificate_management_status` - status of the automatic certificate management for this domain, or null if automatic management is disabled
* `cname_target` - DNS CNAME target for a custom hostname, or null if the reserved domain is a subdomain of *.ngrok.io
* `created_at` - timestamp when the reserved domain was created, RFC 3339 format
* `description` - human-readable description of what this reserved domain will be used for
* `domain` - hostname of the reserved domain
* `http_endpoint_configuration` - object referencing the endpoint configuration applied to http traffic on this domain
* `http_endpoint_configuration_id` - ID of an endpoint configuration of type http that will be used to handle inbound http traffic to this domain
* `https_endpoint_configuration` - object referencing the endpoint configuration applied to https traffic on this domain
* `https_endpoint_configuration_id` - ID of an endpoint configuration of type https that will be used to handle inbound https traffic to this domain
* `metadata` - arbitrary user-defined machine-readable data of this reserved domain. Optional, max 4096 bytes.
* `name` - the domain name to reserve. It may be a full domain name like app.example.com. If the name does not contain a '.' it will reserve that subdomain on ngrok.io.
* `ngrok_id` - unique reserved domain resource identifier
* `region` - reserve the domain in this geographic ngrok datacenter. Optional, default is us. (au, eu, ap, us, jp, in, sa)
* `uri` - URI of the reserved domain API resource

