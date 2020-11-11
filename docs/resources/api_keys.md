# api_keys Resource

## Example Usage

Define the API Key resource `ngrok_api_key.example`:

```
resource "ngrok_api_key" "example" {
  description = "ad-hoc dev testing"
  metadata = "{\"environment\":\"dev\"}"
}
```

## Argument Reference

* `description` - (Optional) human-readable description of what uses the API key to authenticate. optional, max 255 bytes.
* `metadata` - (Optional) arbitrary user-defined data of this API key. optional, max 4096 bytes

## Attribute Reference

* `created_at` - timestamp when the api key was created, RFC 3339 format
* `description` - human-readable description of what uses the API key to authenticate. optional, max 255 bytes.
* `metadata` - arbitrary user-defined data of this API key. optional, max 4096 bytes
* `ngrok_id` - unique API key resource identifier
* `token` - the bearer token that can be placed into the Authorization header to authenticate request to the ngrok API. <strong>This value is only available one time, on the API response from key creation. Otherwise it is null.</strong>
* `uri` - URI to the API resource of this API key

