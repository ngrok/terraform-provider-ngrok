# log_destinations Resource

## Example Usage

Define the Log Destination resource `ngrok_log_destination.example`:

```
resource "ngrok_log_destination" "example" {
  description = "kinesis dev stream"
  format = "json"
  metadata = "{\"environment\":\"dev\"}"
  target {
    kinesis {
      auth {
        creds {
          aws_access_key_id = "AKIAIOSFODNN7EXAMPLE"
          aws_secret_access_key = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
        }
      }
      stream_arn = "arn:aws:kinesis:us-east-2:123456789012:stream/mystream2"
    }
  }
}
```

## Argument Reference

* `description` - (Optional) Human-readable description of the Log Destination. Optional, max 255 bytes.
* `format` - (Optional) The output format you would like to serialize your logs into before they post to their target. Currently the only accepted value is <code>JSON</code>.
* `metadata` - (Optional) Arbitrary user-defined machine-readable data of this Log Destination. Optional, max 4096 bytes.
* `target` - (Optional) An object that encapsulates where and how to send your logs to their ultimate destination. A log destination must contain exactly one of the following objects, leaving the rest null: <code>kinesis</code>, <code>firehose</code>, <code>cloudwatch</code>, or <code>S3</code>.

## Attribute Reference

* `created_at` - Timestamp when the Log Destination was created, RFC 3339 format.
* `description` - Human-readable description of the Log Destination. Optional, max 255 bytes.
* `format` - The output format you would like to serialize your logs into before they post to their target. Currently the only accepted value is <code>JSON</code>.
* `metadata` - Arbitrary user-defined machine-readable data of this Log Destination. Optional, max 4096 bytes.
* `ngrok_id` - Unique identifier for this Log Destination.
* `target` - An object that encapsulates where and how to send your logs to their ultimate destination. A log destination must contain exactly one of the following objects, leaving the rest null: <code>kinesis</code>, <code>firehose</code>, <code>cloudwatch</code>, or <code>S3</code>.
* `uri` - URI of the Log Destination API resource.

