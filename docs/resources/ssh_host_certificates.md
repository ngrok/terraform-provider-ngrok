# ssh_host_certificates Resource

## Example Usage

Define the SSH Host Certificate resource `ngrok_ssh_host_certificate.example`:

```
resource "ngrok_ssh_host_certificate" "example" {
  description = "personal server"
  principals = [ "inconshreveable.com", "10.2.42.9" ]
  public_key = "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBI3oSgxrOEJ+tIJ/n6VYtxQIFvynqlOHpfOAJ4x4OfmMYDkbf8dr6RAuUSf+ZC2HMCujta7EjZ9t+6v08Ue+Cgk= inconshreveable.com"
  ssh_certificate_authority_id = "sshca_1kOTHco3uORrsjO2vdJK5FckNed"
  valid_until = "2021-02-14T23:49:11Z"
}
```

## Argument Reference

* `description` - (Optional) human-readable description of this SSH Host Certificate. optional, max 255 bytes.
* `metadata` - (Optional) arbitrary user-defined machine-readable data of this SSH Host Certificate. optional, max 4096 bytes.
* `principals` - (Optional) the list of principals included in the ssh host certificate. This is the list of hostnames and/or IP addresses that are authorized to serve SSH traffic with this certificate. Dangerously, if no principals are specified, this certificate is considered valid for all hosts.
* `public_key` - (Optional) a public key in OpenSSH Authorized Keys format that this certificate signs
* `ssh_certificate_authority_id` - (Optional) the ssh certificate authority that is used to sign this ssh host certificate
* `valid_after` - (Optional) The time when the host certificate becomes valid, in RFC 3339 format. Defaults to the current time if unspecified.
* `valid_until` - (Optional) The time when this host certificate becomes invalid, in RFC 3339 format. If unspecified, a default value of one year in the future will be used. The OpenSSH certificates RFC calls this <code>valid_before</code>.

## Attribute Reference

* `certificate` - the signed SSH certificate in OpenSSH Authorized Keys format. this value should be placed in a <code>-cert.pub</code> certificate file on disk that should be referenced in your <code>sshd_config</code> configuration file with a <code>HostCertificate</code> directive
* `created_at` - timestamp when the SSH Host Certificate API resource was created, RFC 3339 format
* `description` - human-readable description of this SSH Host Certificate. optional, max 255 bytes.
* `key_type` - the key type of the <code>public_key</code>, one of <code>rsa</code>, <code>ecdsa</code> or <code>ed25519</code>
* `metadata` - arbitrary user-defined machine-readable data of this SSH Host Certificate. optional, max 4096 bytes.
* `ngrok_id` - unique identifier for this SSH Host Certificate
* `principals` - the list of principals included in the ssh host certificate. This is the list of hostnames and/or IP addresses that are authorized to serve SSH traffic with this certificate. Dangerously, if no principals are specified, this certificate is considered valid for all hosts.
* `public_key` - a public key in OpenSSH Authorized Keys format that this certificate signs
* `ssh_certificate_authority_id` - the ssh certificate authority that is used to sign this ssh host certificate
* `uri` - URI of the SSH Host Certificate API resource
* `valid_after` - the time when the ssh host certificate becomes valid, in RFC 3339 format.
* `valid_until` - the time after which the ssh host certificate becomes invalid, in RFC 3339 format. the OpenSSH certificates RFC calls this <code>valid_before</code>.

