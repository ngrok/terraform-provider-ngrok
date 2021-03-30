# tls_certificates Resource

## Example Usage

Define the TLS Certificate resource `ngrok_tls_certificate.example`:

```
resource "ngrok_tls_certificate" "example" {
}
```

## Argument Reference

* `certificate_pem` - (Required) chain of PEM-encoded certificates, leaf first. See <a href="#tls-certificates-pem">Certificate Bundles</a>.
* `description` - (Optional) human-readable description of this TLS certificate. optional, max 255 bytes.
* `metadata` - (Optional) arbitrary user-defined machine-readable data of this TLS certificate. optional, max 4096 bytes.
* `private_key_pem` - (Optional) private key for the TLS certificate, PEM-encoded. See <a href="#tls-certificates-key">Private Keys</a>.

## Attribute Reference

* `certificate_pem` - chain of PEM-encoded certificates, leaf first. See <a href="#tls-certificates-pem">Certificate Bundles</a>.
* `created_at` - timestamp when the TLS certificate was created, RFC 3339 format
* `description` - human-readable description of this TLS certificate. optional, max 255 bytes.
* `extended_key_usages` - extended set of actions the private key of this TLS certificate can be used for
* `issued_at` - timestamp (in RFC 3339 format) when this TLS certificate was issued automatically, or null if this certificate was user-uploaded
* `issuer_common_name` - issuer common name from the leaf of this TLS certificate
* `key_usages` - set of actions the private key of this TLS certificate can be used for
* `metadata` - arbitrary user-defined machine-readable data of this TLS certificate. optional, max 4096 bytes.
* `ngrok_id` - unique identifier for this TLS certificate
* `not_after` - timestamp when this TLS certificate becomes invalid, RFC 3339 format
* `not_before` - timestamp when this TLS certificate becomes valid, RFC 3339 format
* `private_key_pem` - private key for the TLS certificate, PEM-encoded. See <a href="#tls-certificates-key">Private Keys</a>.
* `private_key_type` - type of the private key of this TLS certificate. One of rsa, ecdsa, or ed25519.
* `serial_number` - serial number of the leaf of this TLS certificate
* `subject_alternative_names` - subject alternative names (SANs) from the leaf of this TLS certificate
* `subject_common_name` - subject common name from the leaf of this TLS certificate
* `subject_country` - subject country from the leaf of this TLS certificate
* `subject_locality` - subject locality from the leaf of this TLS certificate
* `subject_organization` - subject organization from the leaf of this TLS certificate
* `subject_organizational_unit` - subject organizational unit from the leaf of this TLS certificate
* `subject_province` - subject province from the leaf of this TLS certificate
* `uri` - URI of the TLS certificate API resource

