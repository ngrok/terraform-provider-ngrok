---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ngrok_ip_policy_rule Resource - terraform-provider-ngrok"
subcategory: ""
description: |-
  IP Policy Rules are the IPv4 or IPv6 CIDRs entries that
   make up an IP Policy.
---

# ngrok_ip_policy_rule (Resource)

IP Policy Rules are the IPv4 or IPv6 CIDRs entries that
 make up an IP Policy.

## Example Usage

```terraform
# Code generated for API Clients. DO NOT EDIT.


resource "ngrok_ip_policy_rule" "example" {
  action = "allow"
  cidr = "212.3.14.0/24"
  description = "nyc office"
  ip_policy_id = "ipp_26rOydjEUNZSLTi8bYXBg278qUT"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **action** (String) the action to apply to the policy rule, either `allow` or `deny`
- **cidr** (String) an IP or IP range specified in CIDR notation. IPv4 and IPv6 are both supported.
- **ip_policy_id** (String) ID of the IP policy this IP policy rule will be attached to

### Optional

- **description** (String) human-readable description of the source IPs of this IP rule. optional, max 255 bytes.
- **metadata** (String) arbitrary user-defined machine-readable data of this IP policy rule. optional, max 4096 bytes.

### Read-Only

- **id** (String) unique identifier for this IP policy rule


