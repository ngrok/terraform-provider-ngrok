# static_backends Resource

## Example Usage

Define the Static Backend resource `ngrok_static_backend.example`:

```
resource "ngrok_static_backend" "example" {
}
```

## Argument Reference

* `address` - (Optional) the address to forward to
* `description` - (Optional) human-readable description of this backend. Optional
* `enabled` - (Optional) if tls is checked
* `metadata` - (Optional) arbitrary user-defined machine-readable data of this backend. Optional
* `tls` - (Optional) tls configuration to use

## Attribute Reference

* `address` - the address to forward to
* `created_at` - timestamp when the backend was created, RFC 3339 format
* `description` - human-readable description of this backend. Optional
* `metadata` - arbitrary user-defined machine-readable data of this backend. Optional
* `ngrok_id` - unique identifier for this static backend
* `tls` - tls configuration to use

