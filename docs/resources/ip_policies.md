# ip_policies Resource

## Example Usage

Define the IP Policy resource `ngrok_ip_policy.example`:

```
resource "ngrok_ip_policy" "example" {
}
```

## Argument Reference

* `action` - (Required) the IP policy action. Supported values are <code>allow</code> or <code>deny</code>
* `description` - (Optional) human-readable description of the source IPs of this IP policy. optional, max 255 bytes.
* `metadata` - (Optional) arbitrary user-defined machine-readable data of this IP policy. optional, max 4096 bytes.

## Attribute Reference

* `action` - the IP policy action. Supported values are <code>allow</code> or <code>deny</code>
* `created_at` - timestamp when the IP policy was created, RFC 3339 format
* `description` - human-readable description of the source IPs of this IP policy. optional, max 255 bytes.
* `metadata` - arbitrary user-defined machine-readable data of this IP policy. optional, max 4096 bytes.
* `ngrok_id` - unique identifier for this IP policy
* `uri` - URI of the IP Policy API resource

