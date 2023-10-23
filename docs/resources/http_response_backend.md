---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ngrok_http_response_backend Resource - terraform-provider-ngrok"
subcategory: ""
description: |-
  
---

# ngrok_http_response_backend (Resource)



## Example Usage

```terraform
# Code generated for API Clients. DO NOT EDIT.


resource "ngrok_http_response_backend" "example" {
  body = "I'm a teapot"
  description = "acme http response"
  headers = {
    Content-Type = "text/plain"
  }
  metadata = "{\"environment\": \"staging\"}"
  status_code = 418
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **body** (String) body to return as fixed content
- **description** (String) human-readable description of this backend. Optional
- **headers** (Map of String) headers to return
- **metadata** (String) arbitrary user-defined machine-readable data of this backend. Optional
- **status_code** (Number) status code to return

### Read-Only

- **created_at** (String) timestamp when the backend was created, RFC 3339 format
- **id** (String) The ID of this resource.
- **uri** (String) URI of the HTTPResponseBackend API resource

