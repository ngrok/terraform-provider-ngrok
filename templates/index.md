<!-- Code generated for API Clients. DO NOT EDIT. -->

# Ngrok Provider

The ngrok provider is used to configure your [ngrok](https://ngrok.com/) infrastructure.
See the [Getting Started](https://ngrok.com/docs#getting-started) page for an introduction to using ngrok.
Detailed [API documentation](https://ngrok.com/docs/ngrok-link) is also available.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the ngrok provider
provider "ngrok" {
  api_key = "my-api-key"
}

# Create a new reserved domain
resource "ngrok_reserved_domain" "my_domain" {
  name   = "my-domain.example.com"
  region = "us"
  certificate_management_policy {
    authority        = "letsencrypt"
    private_key_type = "ecdsa"
  }
}
```

## Argument Reference

The following arguments are supported:

* `api_key` - (Required) The ngrok API Key. Use the [ngrok dashboard](https://dashboard.ngrok.com/api/keys) to locate keys. It can be sourced from the `NGROK_API_KEY` environment variable.
* `api_base_url` - (Optional) The API URL used to talk with ngrok. The default is `https://api.ngrok.com`.  It can be sourced from the `NGROK_API_BASE_URL` environment variable.
