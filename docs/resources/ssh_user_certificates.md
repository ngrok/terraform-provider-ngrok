# ssh_user_certificates Resource

## Example Usage

Define the SSH User Certificate resource `ngrok_ssh_user_certificate.example`:

```
resource "ngrok_ssh_user_certificate" "example" {
  description = "temporary access to staging machine"
  principals = [ "ec2-user", "root" ]
  public_key = "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBK58lFzmWlDimDtBz78wVT4oauA8PjY0CiXTCEIsBNC6UwOJvZ0jdSaYNhDaa7dRV84DfBb/gKzqlXC7cVMZjl0= alan@work-laptop"
  ssh_certificate_authority_id = "sshca_1knYod4EQ6mRQwwD5zZzFgHndHk"
  valid_until = "2021-02-23T20:59:58Z"
}
```

## Argument Reference

* `critical_options` - (Optional) A map of critical options included in the certificate. Only two critical options are currently defined by OpenSSH: <code>force-command</code> and <code>source-address</code>. See <a href="https://github.com/openssh/openssh-portable/blob/master/PROTOCOL.certkeys">the OpenSSH certificate protocol spec</a> for additional details.
* `description` - (Optional) human-readable description of this SSH User Certificate. optional, max 255 bytes.
* `extensions` - (Optional) A map of extensions included in the certificate. Extensions are additional metadata that can be interpreted by the SSH server for any purpose. These can be used to permit or deny the ability to open a terminal, do port forwarding, x11 forwarding, and more. If unspecified, the certificate will include limited permissions with the following extension map: <code>{"permit-pty": "", "permit-user-rc": ""}</code> OpenSSH understands a number of predefined extensions. See <a href="https://github.com/openssh/openssh-portable/blob/master/PROTOCOL.certkeys">the OpenSSH certificate protocol spec</a> for additional details.
* `metadata` - (Optional) arbitrary user-defined machine-readable data of this SSH User Certificate. optional, max 4096 bytes.
* `principals` - (Optional) the list of principals included in the ssh user certificate. This is the list of usernames that the certificate holder may sign in as on a machine authorizinig the signing certificate authority. Dangerously, if no principals are specified, this certificate may be used to log in as any user.
* `public_key` - (Optional) a public key in OpenSSH Authorized Keys format that this certificate signs
* `ssh_certificate_authority_id` - (Optional) the ssh certificate authority that is used to sign this ssh user certificate
* `valid_after` - (Optional) The time when the user certificate becomes valid, in RFC 3339 format. Defaults to the current time if unspecified.
* `valid_until` - (Optional) The time when this host certificate becomes invalid, in RFC 3339 format. If unspecified, a default value of 24 hours will be used. The OpenSSH certificates RFC calls this <code>valid_before</code>.

## Attribute Reference

* `certificate` - the signed SSH certificate in OpenSSH Authorized Keys Format. this value should be placed in a <code>-cert.pub</code> certificate file on disk that should be referenced in your <code>sshd_config</code> configuration file with a <code>HostCertificate</code> directive
* `created_at` - timestamp when the SSH User Certificate API resource was created, RFC 3339 format
* `critical_options` - A map of critical options included in the certificate. Only two critical options are currently defined by OpenSSH: <code>force-command</code> and <code>source-address</code>. See <a href="https://github.com/openssh/openssh-portable/blob/master/PROTOCOL.certkeys">the OpenSSH certificate protocol spec</a> for additional details.
* `description` - human-readable description of this SSH User Certificate. optional, max 255 bytes.
* `extensions` - A map of extensions included in the certificate. Extensions are additional metadata that can be interpreted by the SSH server for any purpose. These can be used to permit or deny the ability to open a terminal, do port forwarding, x11 forwarding, and more. If unspecified, the certificate will include limited permissions with the following extension map: <code>{"permit-pty": "", "permit-user-rc": ""}</code> OpenSSH understands a number of predefined extensions. See <a href="https://github.com/openssh/openssh-portable/blob/master/PROTOCOL.certkeys">the OpenSSH certificate protocol spec</a> for additional details.
* `key_type` - the key type of the <code>public_key</code>, one of <code>rsa</code>, <code>ecdsa</code> or <code>ed25519</code>
* `metadata` - arbitrary user-defined machine-readable data of this SSH User Certificate. optional, max 4096 bytes.
* `ngrok_id` - unique identifier for this SSH User Certificate
* `principals` - the list of principals included in the ssh user certificate. This is the list of usernames that the certificate holder may sign in as on a machine authorizinig the signing certificate authority. Dangerously, if no principals are specified, this certificate may be used to log in as any user.
* `public_key` - a public key in OpenSSH Authorized Keys format that this certificate signs
* `ssh_certificate_authority_id` - the ssh certificate authority that is used to sign this ssh user certificate
* `uri` - URI of the SSH User Certificate API resource
* `valid_after` - the time when the ssh host certificate becomes valid, in RFC 3339 format.
* `valid_until` - the time after which the ssh host certificate becomes invalid, in RFC 3339 format. the OpenSSH certificates RFC calls this <code>valid_before</code>.

