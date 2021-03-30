# certificate_authorities Resource

## Example Usage

Define the Certificate Authority resource `ngrok_certificate_authority.example`:

```
resource "ngrok_certificate_authority" "example" {
}
```

## Argument Reference

* `ca_pem` - (Optional) raw PEM of the Certificate Authority
* `description` - (Optional) human-readable description of this Certificate Authority. optional, max 255 bytes.
* `metadata` - (Optional) arbitrary user-defined machine-readable data of this Certificate Authority. optional, max 4096 bytes.

## Attribute Reference

* `ca_pem` - raw PEM of the Certificate Authority
* `created_at` - timestamp when the Certificate Authority was created, RFC 3339 format
* `description` - human-readable description of this Certificate Authority. optional, max 255 bytes.
* `extended_key_usages` - extended set of actions the private key of this Certificate Authority can be used for
* `key_usages` - set of actions the private key of this Certificate Authority can be used for
* `metadata` - arbitrary user-defined machine-readable data of this Certificate Authority. optional, max 4096 bytes.
* `ngrok_id` - unique identifier for this Certificate Authority
* `not_after` - timestamp when this Certificate Authority becomes invalid, RFC 3339 format
* `not_before` - timestamp when this Certificate Authority becomes valid, RFC 3339 format
* `subject_common_name` - subject common name of the Certificate Authority
* `uri` - URI of the Certificate Authority API resource

