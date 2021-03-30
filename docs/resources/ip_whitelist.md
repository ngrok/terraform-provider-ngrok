# ip_whitelist Resource

## Example Usage

Define the IP Whitelist Entry resource `ngrok_ip_whitelist_entry.example`:

```
resource "ngrok_ip_whitelist_entry" "example" {
}
```

## Argument Reference

* `description` - (Optional) human-readable description of the source IPs for this IP whitelist entry. optional, max 255 bytes.
* `ip_net` - (Optional) an IP address or IP network range in CIDR notation (e.g. 10.1.1.1 or 10.1.0.0/16) of addresses that will be whitelisted to communicate with your tunnel endpoints
* `metadata` - (Optional) arbitrary user-defined machine-readable data of this IP whitelist entry. optional, max 4096 bytes.

## Attribute Reference

* `created_at` - timestamp when the IP whitelist entry was created, RFC 3339 format
* `description` - human-readable description of the source IPs for this IP whitelist entry. optional, max 255 bytes.
* `ip_net` - an IP address or IP network range in CIDR notation (e.g. 10.1.1.1 or 10.1.0.0/16) of addresses that will be whitelisted to communicate with your tunnel endpoints
* `metadata` - arbitrary user-defined machine-readable data of this IP whitelist entry. optional, max 4096 bytes.
* `ngrok_id` - unique identifier for this IP whitelist entry
* `uri` - URI of the IP whitelist entry API resource

