# event_destinations Resource

## Example Usage

Define the Event Destination resource `ngrok_event_destination.example`:

```
resource "ngrok_event_destination" "example" {
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
      stream_arn = "arn:ngrok-local:kinesis:us-east-2:123456789012:stream/mystream2"
    }
  }
}
```

## Argument Reference

* `auth` - (Optional) Configuration for how to authenticate into your AWS account. Exactly one of <code>role</code> or <code>creds</code> should be configured.
* `aws_access_key_id` - (Optional) The ID portion of an AWS access key.
* `aws_secret_access_key` - (Optional) The secret portion of an AWS access key.
* `callback_url` - (Optional) URL to send events to.
* `cloudwatch_logs` - (Optional) Configuration used to send events to Amazon CloudWatch Logs.
* `creds` - (Optional) Credentials to your AWS account if you prefer ngrok to sign in with long-term access keys.
* `debug` - (Optional) Configuration used for internal debugging.
* `delivery_stream_arn` - (Optional) An Amazon Resource Name specifying the Firehose delivery stream to deposit events into.
* `description` - (Optional) Human-readable description of the Event Destination. Optional, max 255 bytes.
* `firehose` - (Optional) Configuration used to send events to Amazon Kinesis Data Firehose.
* `format` - (Optional) The output format you would like to serialize events into when sending to their target. Currently the only accepted value is <code>JSON</code>.
* `kinesis` - (Optional) Configuration used to send events to Amazon Kinesis.
* `log` - (Optional) Whether or not to output to publisher service logs.
* `log_group_arn` - (Optional) An Amazon Resource Name specifying the CloudWatch Logs group to deposit events into.
* `metadata` - (Optional) Arbitrary user-defined machine-readable data of this Event Destination. Optional, max 4096 bytes.
* `role` - (Optional) A role for ngrok to assume on your behalf to deposit events into your AWS account.
* `role_arn` - (Optional) An ARN that specifies the role that ngrok should use to deliver to the configured target.
* `stream_arn` - (Optional) An Amazon Resource Name specifying the Kinesis stream to deposit events into.
* `target` - (Optional) An object that encapsulates where and how to send your events. An event destination must contain exactly one of the following objects, leaving the rest null: <code>kinesis</code>, <code>firehose</code>, <code>cloudwatch_logs</code>, or <code>s3</code>.
* `verify_with_test_event` - (Optional) 

## Attribute Reference

* `created_at` - Timestamp when the Event Destination was created, RFC 3339 format.
* `description` - Human-readable description of the Event Destination. Optional, max 255 bytes.
* `format` - The output format you would like to serialize events into when sending to their target. Currently the only accepted value is <code>JSON</code>.
* `metadata` - Arbitrary user-defined machine-readable data of this Event Destination. Optional, max 4096 bytes.
* `ngrok_id` - Unique identifier for this Event Destination.
* `target` - An object that encapsulates where and how to send your events. An event destination must contain exactly one of the following objects, leaving the rest null: <code>kinesis</code>, <code>firehose</code>, <code>cloudwatch_logs</code>, or <code>s3</code>.
* `uri` - URI of the Event Destination API resource.
* `verify_with_test_event` - 

