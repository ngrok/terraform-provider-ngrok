# reserved_addrs Resource

## Example Usage

Define the Reserved Address resource `ngrok_reserved_addr.example`:

```
resource "ngrok_reserved_addr" "example" {
}
```

## Argument Reference

* `description` - (Optional) human-readable description of what this reserved address will be used for
* `endpoint_configuration_id` - (Optional) ID of an endpoint configuration of type tcp that will be used to handle inbound traffic to this address
* `metadata` - (Optional) arbitrary user-defined machine-readable data of this reserved address. Optional, max 4096 bytes.
* `region` - (Optional) reserve the address in this geographic ngrok datacenter. Optional, default is us. (au, eu, ap, us, jp, in, sa)

## Attribute Reference

* `addr` - hostname:port of the reserved address that was assigned at creation time
* `created_at` - timestamp when the reserved address was created, RFC 3339 format
* `description` - human-readable description of what this reserved address will be used for
* `endpoint_configuration` - object reference to the endpoint configuration that will be applied to traffic to this address
* `endpoint_configuration_id` - ID of an endpoint configuration of type tcp that will be used to handle inbound traffic to this address
* `metadata` - arbitrary user-defined machine-readable data of this reserved address. Optional, max 4096 bytes.
* `ngrok_id` - unique reserved address resource identifier
* `region` - reserve the address in this geographic ngrok datacenter. Optional, default is us. (au, eu, ap, us, jp, in, sa)
* `uri` - URI of the reserved address API resource

