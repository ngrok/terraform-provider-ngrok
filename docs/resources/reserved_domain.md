---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ngrok_reserved_domain Resource - terraform-provider-ngrok"
subcategory: ""
description: |-
  
---

# ngrok_reserved_domain (Resource)



## Example Usage

```terraform
resource "ngrok_reserved_domain" "example" {
  certificate_id = "cert_1rV51OQetZPK9V6vTWUVy3Onjir"
  name = "myapp.mydomain.com"
  region = "us"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) the domain name to reserve. It may be a full domain name like app.example.com. If the name does not contain a '.' it will reserve that subdomain on ngrok.io.

### Optional

- **certificate** (Block Set) object referencing the TLS certificate used for connections to this domain. This can be either a user-uploaded certificate, the most recently issued automatic one, or null otherwise. (see [below for nested schema](#nestedblock--certificate))
- **certificate_id** (String) ID of a user-uploaded TLS certificate to use for connections to targeting this domain. Optional, mutually exclusive with `certificate_management_policy`.
- **certificate_management_policy** (Block Set) configuration for automatic management of TLS certificates for this domain, or null if automatic management is disabled (see [below for nested schema](#nestedblock--certificate_management_policy))
- **certificate_management_status** (Block Set) status of the automatic certificate management for this domain, or null if automatic management is disabled (see [below for nested schema](#nestedblock--certificate_management_status))
- **cname_target** (String) DNS CNAME target for a custom hostname, or null if the reserved domain is a subdomain of *.ngrok.io
- **created_at** (String) timestamp when the reserved domain was created, RFC 3339 format
- **description** (String) human-readable description of what this reserved domain will be used for
- **http_endpoint_configuration** (Block Set) object referencing the endpoint configuration applied to http traffic on this domain (see [below for nested schema](#nestedblock--http_endpoint_configuration))
- **http_endpoint_configuration_id** (String) ID of an endpoint configuration of type http that will be used to handle inbound http traffic to this domain
- **https_endpoint_configuration** (Block Set) object referencing the endpoint configuration applied to https traffic on this domain (see [below for nested schema](#nestedblock--https_endpoint_configuration))
- **https_endpoint_configuration_id** (String) ID of an endpoint configuration of type https that will be used to handle inbound https traffic to this domain
- **id** (String) The ID of this resource.
- **metadata** (String) arbitrary user-defined machine-readable data of this reserved domain. Optional, max 4096 bytes.
- **ngrok_id** (String) unique reserved domain resource identifier
- **region** (String) reserve the domain in this geographic ngrok datacenter. Optional, default is us. (au, eu, ap, us, jp, in, sa)
- **uri** (String) URI of the reserved domain API resource

### Read-Only

- **domain** (String) hostname of the reserved domain

<a id="nestedblock--certificate"></a>
### Nested Schema for `certificate`

Optional:

- **ngrok_id** (String) a resource identifier
- **uri** (String) a uri for locating a resource


<a id="nestedblock--certificate_management_policy"></a>
### Nested Schema for `certificate_management_policy`

Optional:

- **authority** (String) certificate authority to request certificates from. The only supported value is letsencrypt.
- **private_key_type** (String) type of private key to use when requesting certificates. Defaults to rsa, can be either rsa or ecdsa.


<a id="nestedblock--certificate_management_status"></a>
### Nested Schema for `certificate_management_status`

Optional:

- **provisioning_job** (Block Set) status of the certificate provisioning job, or null if the certificiate isn't being provisioned or renewed (see [below for nested schema](#nestedblock--certificate_management_status--provisioning_job))
- **renews_at** (String) timestamp when the next renewal will be requested, RFC 3339 format

<a id="nestedblock--certificate_management_status--provisioning_job"></a>
### Nested Schema for `certificate_management_status.provisioning_job`

Optional:

- **error_code** (String) if present, an error code indicating why provisioning is failing. It may be either a temporary condition (INTERNAL_ERROR), or a permanent one the user must correct (DNS_ERROR).
- **msg** (String) a message describing the current status or error
- **ns_targets** (Block List) if present, indicates the dns nameservers that the user must configure to complete the provisioning process of a wildcard certificate (see [below for nested schema](#nestedblock--certificate_management_status--provisioning_job--ns_targets))
- **retries_at** (String) timestamp when the provisioning job will be retried
- **started_at** (String) timestamp when the provisioning job started, RFC 3339 format

<a id="nestedblock--certificate_management_status--provisioning_job--ns_targets"></a>
### Nested Schema for `certificate_management_status.provisioning_job.ns_targets`

Optional:

- **nameservers** (List of String) the nameservers the user must add
- **zone** (String) the zone that the nameservers need to be applied to




<a id="nestedblock--http_endpoint_configuration"></a>
### Nested Schema for `http_endpoint_configuration`

Optional:

- **ngrok_id** (String) a resource identifier
- **uri** (String) a uri for locating a resource


<a id="nestedblock--https_endpoint_configuration"></a>
### Nested Schema for `https_endpoint_configuration`

Optional:

- **ngrok_id** (String) a resource identifier
- **uri** (String) a uri for locating a resource

