# ssh_credentials Resource

## Example Usage

Define the SSH Credential resource `ngrok_ssh_credential.example`:

```
resource "ngrok_ssh_credential" "example" {
}
```

## Argument Reference

* `acl` - (Optional) optional list of ACL rules. If unspecified, the credential will have no restrictions. The only allowed ACL rule at this time is the <code>bind</code> rule. The <code>bind</code> rule allows the caller to restrict what domains and addresses the token is allowed to bind. For example, to allow the token to open a tunnel on example.ngrok.io your ACL would include the rule <code>bind:example.ngrok.io</code>. Bind rules may specify a leading wildcard to match multiple domains with a common suffix. For example, you may specify a rule of <code>bind:*.example.com</code> which will allow <code>x.example.com</code>, <code>y.example.com</code>, <code>*.example.com</code>, etc. A rule of <code>'*'</code> is equivalent to no acl at all and will explicitly permit all actions.
* `description` - (Optional) human-readable description of who or what will use the ssh credential to authenticate. Optional, max 255 bytes.
* `metadata` - (Optional) arbitrary user-defined machine-readable data of this ssh credential. Optional, max 4096 bytes.
* `public_key` - (Optional) the PEM-encoded public key of the SSH keypair that will be used to authenticate

## Attribute Reference

* `acl` - optional list of ACL rules. If unspecified, the credential will have no restrictions. The only allowed ACL rule at this time is the <code>bind</code> rule. The <code>bind</code> rule allows the caller to restrict what domains and addresses the token is allowed to bind. For example, to allow the token to open a tunnel on example.ngrok.io your ACL would include the rule <code>bind:example.ngrok.io</code>. Bind rules may specify a leading wildcard to match multiple domains with a common suffix. For example, you may specify a rule of <code>bind:*.example.com</code> which will allow <code>x.example.com</code>, <code>y.example.com</code>, <code>*.example.com</code>, etc. A rule of <code>'*'</code> is equivalent to no acl at all and will explicitly permit all actions.
* `created_at` - timestamp when the ssh credential was created, RFC 3339 format
* `description` - human-readable description of who or what will use the ssh credential to authenticate. Optional, max 255 bytes.
* `metadata` - arbitrary user-defined machine-readable data of this ssh credential. Optional, max 4096 bytes.
* `ngrok_id` - unique ssh credential resource identifier
* `public_key` - the PEM-encoded public key of the SSH keypair that will be used to authenticate
* `uri` - URI of the ssh credential API resource

