# ssh_certificate_authorities Resource

## Example Usage

Define the SSH Certificate Authority resource `ngrok_ssh_certificate_authority.example`:

```
resource "ngrok_ssh_certificate_authority" "example" {
  description = "Staging Environment Hosts"
  private_key_type = "ed25519"
}
```

## Argument Reference

* `description` - (Optional) human-readable description of this SSH Certificate Authority. optional, max 255 bytes.
* `elliptic_curve` - (Optional) the type of elliptic curve to use when creating an ECDSA key
* `key_size` - (Optional) the key size to use when creating an RSA key. one of <code>2048</code> or <code>4096</code>
* `metadata` - (Optional) arbitrary user-defined machine-readable data of this SSH Certificate Authority. optional, max 4096 bytes.
* `private_key_type` - (Optional) the type of private key to generate. one of <code>rsa</code>, <code>ecdsa</code>, <code>ed25519</code>

## Attribute Reference

* `created_at` - timestamp when the SSH Certificate Authority API resource was created, RFC 3339 format
* `description` - human-readable description of this SSH Certificate Authority. optional, max 255 bytes.
* `elliptic_curve` - the type of elliptic curve to use when creating an ECDSA key
* `key_size` - the key size to use when creating an RSA key. one of <code>2048</code> or <code>4096</code>
* `key_type` - the type of private key for this SSH Certificate Authority
* `metadata` - arbitrary user-defined machine-readable data of this SSH Certificate Authority. optional, max 4096 bytes.
* `ngrok_id` - unique identifier for this SSH Certificate Authority
* `private_key_type` - the type of private key to generate. one of <code>rsa</code>, <code>ecdsa</code>, <code>ed25519</code>
* `public_key` - raw public key for this SSH Certificate Authority
* `uri` - URI of the SSH Certificate Authority API resource

