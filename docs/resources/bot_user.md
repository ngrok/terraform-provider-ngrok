---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ngrok_bot_user Resource - terraform-provider-ngrok"
subcategory: ""
description: |-
  
---

# ngrok_bot_user (Resource)



## Example Usage

```terraform
# Code generated for API Clients. DO NOT EDIT.


resource "ngrok_bot_user" "example" {
  name = "new bot user from API"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **active** (Boolean) whether or not the bot is active
- **name** (String) human-readable name used to identify the bot

### Read-Only

- **id** (String) unique API key resource identifier

