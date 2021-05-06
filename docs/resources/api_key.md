---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ngrok_api_key Resource - terraform-provider-ngrok"
subcategory: ""
description: |-
  
---

# ngrok_api_key (Resource)



## Example Usage

```terraform
resource "ngrok_api_key" "example" {
  description = "ad-hoc dev testing"
  metadata = "{\"environment\":\"dev\"}"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **created_at** (String) timestamp when the api key was created, RFC 3339 format
- **description** (String) human-readable description of what uses the API key to authenticate. optional, max 255 bytes.
- **id** (String) The ID of this resource.
- **metadata** (String) arbitrary user-defined data of this API key. optional, max 4096 bytes
- **ngrok_id** (String) unique API key resource identifier
- **token** (String, Sensitive) the bearer token that can be placed into the Authorization header to authenticate request to the ngrok API. This value is only available one time, on the API response from key creation. Otherwise it is null.
- **uri** (String) URI to the API resource of this API key

