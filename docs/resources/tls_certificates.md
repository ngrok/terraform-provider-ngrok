# tls_certificates Resource

## Example Usage

Define the TLS Certificate resource `ngrok_tls_certificate.example`:

```
resource "ngrok_tls_certificate" "example" {
  certificate_pem = "-----BEGIN CERTIFICATE-----\nMIIDDTCCAfWgAwIBAgIUBUunDdA4gjgtEbZA8w9Ljhvl3bEwDQYJKoZIhvcNAQEL\nBQAwFjEUMBIGA1UEAwwLZXhhbXBsZS5jb20wHhcNMjAwMzI0MTgxODE5WhcNMjAw\nNDIzMTgxODE5WjAWMRQwEgYDVQQDDAtleGFtcGxlLmNvbTCCASIwDQYJKoZIhvcN\nAQEBBQADggEPADCCAQoCggEBAPKVkkKYNl3d9cqrz4tIFlwsohED5W4y1dcBixy4\nGANFFnw43nc2wPyKwYXumJqJIFrcW/NkUZL07bd+dou6mT6Gh/zbaTW91IkREPXL\n7b3KfVu4XkFosVXpWs0U6o4GrZ81CLiKBWI+H03x/ij5OSiJ1l71pqLeTJLOydAR\nAl8kpp7axeHU4UbDrAZkW5SnuZTjIKwVg0UNsBg1yNfUOu1Uah3BYaqPgQitC0Yg\nLW+NUGu/T91bkD7tLsVInkQXeQGdXBAqOycfJ7wj8OlIpyuXjTnGFA0izVmbQw5f\nrQnZ0geGyhLamvz9Gcd7mIlD/+/AEN9Lht82tAOzKG98/O8CAwEAAaNTMFEwHQYD\nVR0OBBYEFKv6RsvEC6T+zCtJZwB0FCR1sEkhMB8GA1UdIwQYMBaAFKv6RsvEC6T+\nzCtJZwB0FCR1sEkhMA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEB\nAC5fBrouinespo5+9AipjhY/HOKTg+OCnppFnSnqeU1eXZZJ0oakdHTpTNxtbQP9\ntOJTA2f3KWvmpNDMohEQXZz8wHDkdbrIXJKVp6zs1pEp+0BIjA4y9mSywa5xuyk0\noGeChRgGqp2JujDyPCb7LEaKKQEEdMqy73QG+jEAh14+wKixlAf1nATBdeCUvssK\n2x1uZMyqjJFB5y/5EdnWQzD4WJkrsCkxsZHVMN1d+dqf2sf3dTRV8fzsFGOG17NS\n6u2n9iGcFdBA82XN8yeLIWhy1t3GWutG1sdxENbFRRXea+iUqzDsmRtkaBma2GLQ\nd6JTpFbsCtwDjP23UEi7SZo=\n-----END CERTIFICATE-----"
  private_key_pem = "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDylZJCmDZd3fXK\nq8+LSBZcLKIRA+VuMtXXAYscuBgDRRZ8ON53NsD8isGF7piaiSBa3FvzZFGS9O23\nfnaLupk+hof822k1vdSJERD1y+29yn1buF5BaLFV6VrNFOqOBq2fNQi4igViPh9N\n8f4o+TkoidZe9aai3kySzsnQEQJfJKae2sXh1OFGw6wGZFuUp7mU4yCsFYNFDbAY\nNcjX1DrtVGodwWGqj4EIrQtGIC1vjVBrv0/dW5A+7S7FSJ5EF3kBnVwQKjsnHye8\nI/DpSKcrl405xhQNIs1Zm0MOX60J2dIHhsoS2pr8/RnHe5iJQ//vwBDfS4bfNrQD\nsyhvfPzvAgMBAAECggEBALLv7YE98exvi5zB+0fMFuJK8gkHDLequ93q/4hhqyTO\nU3WyJTdepiAi4fk/NEXZnIopPZJdj2aNUMQnfp43OE7MwYac+hBwRFQOyKnmkSmM\nMcf0SWKKLTUn+piIMzQsbOmhHxuwg6QiGslOFaJ3o9fpRL2rCg3dWDJ6Ypcd1NgE\nK0uy7gg+DwIpU6MeG6lA+HbxbGi+yd2x88Gjn9dGr7FZK34RUDooH60BCX9P8N9X\nT+n10MzzX7ZQOsLfe8FKc1/X8AybI5SYm1GMyfKD4QBt6JG4HKAjPHzBzcIpfN3d\n7BM11Imkrz7LcbUG+F23NVsi6n5IIGT1WqwCRIH2PpECgYEA/SJ5Ra4d0hUS5RYB\nzABquM3sp7JsKxCn7O5PqNLB4TgH9dXtWFhaFVB6juMGyHbvktVH0j4lps/Te0rk\nVU2zU1XxvCTFhtcCYUtNk0cRw6LH8feKiorXHdDRB33t0c47QSD/6AGOjBtxqD7B\n3ZxyR3P+7RdQopLLRFN+FHAnmzsCgYEA9VSGZDFSK+fbg4CgwkWdzuHrAXaUEv0U\novqqWd/yXB9wauEvRHnOrSgW6hFZQiatJOXx0KnalJQzohz/SLGO0MqGtwQbYWVT\nWiJgjUbNeiPEHBeUA6U55lVQr26kQSUWdXEtRbDz+hqV1K+6tTEMzaSPmJiHNgki\nlNMO2gqGQd0CgYBJ268qx5zn2UJEGWG41j5NYbg1TfgFsLxugzI2/heX0TNxZVP1\nPQI7ydmYq2ElSJ6qZxSnoX5255i7FqT8xskV/bOkw83mhAGrxb8Cw+/I90wDq8h+\nl/ggOPdkijfDybq8TBae6SVgd/l3r6f9M1KcypmNMApVBSPN8daNvBOyVQKBgQDo\nsj2utyFrx8Xsm4rf+kxOuPbBMooM4MQ8OmpuSP6G5sMofWLqHmcs0sO5TK9PEYRV\nZU3ST+ml2FSJRdvWRaRi4laZLWoTHZrL+aN/HVM0sMwIoUyhkIy0ruOTIuzlZZpB\n1xHL8qXX6nOHgw8jYdz1CUuyv6owVMXaR77kjer+eQKBgByYZlR/eNTzlot0SdFl\nIbgQ9bV7VLIo+vKzOXE3trfzRJMgUosLTp+5wdSVSW/VBdYZ7Ir3n0bbpY/dGinI\nVShxPbChhCZnhvG2lEEiekI44m5jHSA6hhtRdt/CrhL65Rw2SE5lMEe8htg1UGus\nwzLHWHBl72FjbjdhvEgrq60W\n-----END PRIVATE KEY-----"
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

