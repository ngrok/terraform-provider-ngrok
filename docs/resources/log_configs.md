# log_configs Resource

## Example Usage

Define the Log Config resource `ngrok_log_config.example`:

```
resource "ngrok_log_config" "example" {
  description = "low sampling, basic HTTP logs"
  destination_ids = [ "ld_1kD5meVhc5hKZyEFuaDr59DRIT7" ]
  event_type = "http_request_complete"
  fields = [ "http.request.method", "http.response.status_code", "conn.client_ip" ]
  metadata = "{\"environment\": \"staging\"}"
  sampling_rate = 0.1
}
```

## Argument Reference

* `description` - (Optional) Human-readable description of the Log Config. Optional, max 255 bytes.
* `destination_ids` - (Optional) A list of Log Destination ids which should be applied to this Log Config. Log Configs are required to have at least one Log Destination.
* `event_type` - (Optional) The protocol that determines which events can be logged. Supported values are <code>tcp_connection_closed</code> and <code>http_request_complete</code>.
* `fields` - (Optional) A list of protocol-specific fields you want to collect on each logging event.
* `metadata` - (Optional) Arbitrary user-defined machine-readable data of this Log Config. Optional, max 4096 bytes.
* `sampling_rate` - (Optional) The percentage of all events you would like to log. Valid values range from 0.01, representing 1% of all events to 1.00, representing 100% of all events.

## Attribute Reference

* `created_at` - Timestamp when the Log Config was created, RFC 3339 format.
* `description` - Human-readable description of the Log Config. Optional, max 255 bytes.
* `destination_ids` - A list of Log Destination ids which should be applied to this Log Config. Log Configs are required to have at least one Log Destination.
* `event_type` - The protocol that determines which events can be logged. Supported values are <code>tcp_connection_closed</code> and <code>http_request_complete</code>.
* `fields` - A list of protocol-specific fields you want to collect on each logging event.
* `metadata` - Arbitrary user-defined machine-readable data of this Log Config. Optional, max 4096 bytes.
* `ngrok_id` - Unique identifier for this Log Config.
* `sampling_rate` - The percentage of all events you would like to log. Valid values range from 0.01, representing 1% of all events to 1.00, representing 100% of all events.
* `uri` - URI of the Log Config API resource.

