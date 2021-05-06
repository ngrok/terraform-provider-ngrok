---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ngrok_reserved_addr Resource - terraform-provider-ngrok"
subcategory: ""
description: |-
  
---

# ngrok_reserved_addr (Resource)



## Example Usage

```terraform
resource "ngrok_reserved_addr" "example" {
  description = "SSH for device #001"
  region = "us"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **addr** (String) hostname:port of the reserved address that was assigned at creation time
- **created_at** (String) timestamp when the reserved address was created, RFC 3339 format
- **description** (String) human-readable description of what this reserved address will be used for
- **endpoint_configuration** (Block Set) object reference to the endpoint configuration that will be applied to traffic to this address (see [below for nested schema](#nestedblock--endpoint_configuration))
- **endpoint_configuration_id** (String) ID of an endpoint configuration of type tcp that will be used to handle inbound traffic to this address
- **id** (String) The ID of this resource.
- **metadata** (String) arbitrary user-defined machine-readable data of this reserved address. Optional, max 4096 bytes.
- **ngrok_id** (String) unique reserved address resource identifier
- **region** (String) reserve the address in this geographic ngrok datacenter. Optional, default is us. (au, eu, ap, us, jp, in, sa)
- **uri** (String) URI of the reserved address API resource

<a id="nestedblock--endpoint_configuration"></a>
### Nested Schema for `endpoint_configuration`

Optional:

- **ngrok_id** (String) a resource identifier
- **uri** (String) a uri for locating a resource

