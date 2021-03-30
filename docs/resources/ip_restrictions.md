# ip_restrictions Resource

## Example Usage

Define the IP Restriction resource `ngrok_ip_restriction.example`:

```
resource "ngrok_ip_restriction" "example" {
  ip_policy_ids = [ "ipp_1qTwJu99fJ7XzbLhyiVELgIdjE1" ]
  type = "dashboard"
}
```

## Argument Reference

* `description` - (Optional) human-readable description of this IP restriction. optional, max 255 bytes.
* `enforced` - (Optional) true if the IP restriction will be enforce. if false, only warnings will be issued
* `ip_policy_ids` - (Optional) the set of IP policy identifiers that are used to enforce the restriction
* `metadata` - (Optional) arbitrary user-defined machine-readable data of this IP restriction. optional, max 4096 bytes.
* `type` - (Optional) the type of IP restriction. this defines what traffic will be restricted with the attached policies. four values are currently supported: <code>dashboard</code>, <code>api</code>, <code>agent</code>, and <code>endpoints</code>

## Attribute Reference

* `created_at` - timestamp when the IP restriction was created, RFC 3339 format
* `description` - human-readable description of this IP restriction. optional, max 255 bytes.
* `enforced` - true if the IP restriction will be enforce. if false, only warnings will be issued
* `ip_policies` - the set of IP policies that are used to enforce the restriction
* `ip_policy_ids` - the set of IP policy identifiers that are used to enforce the restriction
* `metadata` - arbitrary user-defined machine-readable data of this IP restriction. optional, max 4096 bytes.
* `ngrok_id` - unique identifier for this IP restriction
* `type` - the type of IP restriction. this defines what traffic will be restricted with the attached policies. four values are currently supported: <code>dashboard</code>, <code>api</code>, <code>agent</code>, and <code>endpoints</code>
* `uri` - URI of the IP restriction API resource

