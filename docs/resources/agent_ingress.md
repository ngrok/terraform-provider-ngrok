---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ngrok_agent_ingress Resource - terraform-provider-ngrok"
subcategory: ""
description: |-
  
---

# ngrok_agent_ingress (Resource)



## Example Usage

```terraform
# Code generated for API Clients. DO NOT EDIT.


resource "ngrok_agent_ingress" "example" {
  description = "acme devices"
  domain = "connect.acme.com"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **domain** (String) the domain that you own to be used as the base domain name to generate regional agent ingress domains.

### Optional

- **certificate_management_policy** (Block Set) configuration for automatic management of TLS certificates for this domain, or null if automatic management is disabled (see [below for nested schema](#nestedblock--certificate_management_policy))
- **description** (String) human-readable description of the use of this Agent Ingress. optional, max 255 bytes.
- **metadata** (String) arbitrary user-defined machine-readable data of this Agent Ingress. optional, max 4096 bytes

### Read-Only

- **created_at** (String) timestamp when the Agent Ingress was created, RFC 3339 format
- **id** (String) unique Agent Ingress resource identifier
- **ns_targets** (List of String) a list of target values to use as the values of NS records for the domain property these values will delegate control over the domain to ngrok
- **region_domains** (List of String) a list of regional agent ingress domains that are subdomains of the value of domain this value may increase over time as ngrok adds more regions
- **uri** (String) URI to the API resource of this Agent ingress

<a id="nestedblock--certificate_management_policy"></a>
### Nested Schema for `certificate_management_policy`

Optional:

- **authority** (String) certificate authority to request certificates from. The only supported value is letsencrypt.
- **private_key_type** (String) type of private key to use when requesting certificates. Defaults to rsa, can be either rsa or ecdsa.


