resource "ngrok_event_destination" "example" {
  description = "Send events to Amazon Kinesis Firehose"
  metadata    = jsonencode({ environment = "production" })
  format      = "json"

  target = {
    firehose = {
      delivery_stream_arn = "arn:aws:firehose:us-east-1:123456789012:deliverystream/ngrok-events"
      auth = {
        role = {
          role_arn = "arn:aws:iam::123456789012:role/ngrok-firehose"
        }
      }
    }
  }
}
