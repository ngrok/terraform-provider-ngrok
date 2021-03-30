# event_streams Resource

## Example Usage

Define the Event Stream resource `ngrok_event_stream.example`:

```
resource "ngrok_event_stream" "example" {
}
```

## Argument Reference

* `description` - (Optional) Human-readable description of the Event Stream. Optional, max 255 bytes.
* `destination_ids` - (Optional) A list of Event Destination IDs which should be used for this Event Stream. Event Streams are required to have at least one Event Destination.
* `event_type` - (Optional) The protocol that determines which events will be collected. Supported values are <code>tcp_connection_closed</code> and <code>http_request_complete</code>.
* `fields` - (Optional) A list of protocol-specific fields you want to collect on each event.
* `metadata` - (Optional) Arbitrary user-defined machine-readable data of this Event Stream. Optional, max 4096 bytes.
* `sampling_rate` - (Optional) The percentage of all events you would like to capture. Valid values range from 0.01, representing 1% of all events to 1.00, representing 100% of all events.

## Attribute Reference

* `created_at` - Timestamp when the Event Stream was created, RFC 3339 format.
* `description` - Human-readable description of the Event Stream. Optional, max 255 bytes.
* `destination_ids` - A list of Event Destination IDs which should be used for this Event Stream. Event Streams are required to have at least one Event Destination.
* `event_type` - The protocol that determines which events will be collected. Supported values are <code>tcp_connection_closed</code> and <code>http_request_complete</code>.
* `fields` - A list of protocol-specific fields you want to collect on each event.
* `metadata` - Arbitrary user-defined machine-readable data of this Event Stream. Optional, max 4096 bytes.
* `ngrok_id` - Unique identifier for this Event Stream.
* `sampling_rate` - The percentage of all events you would like to capture. Valid values range from 0.01, representing 1% of all events to 1.00, representing 100% of all events.
* `uri` - URI of the Event Stream API resource.

