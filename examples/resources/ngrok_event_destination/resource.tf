# Code generated for API Clients. DO NOT EDIT.


resource "ngrok_event_destination" "example" {
  description = "kinesis dev stream"
  format = "json"
  metadata = "{\"environment\":\"dev\"}"
  target {
    kinesis {
      auth {
        role {
          role_arn = "arn:aws:iam::123456789012:role/example"
        }
      }
      stream_arn = "arn:ngrok-local:kinesis:us-east-2:123456789012:stream/mystream2"
    }
  }
}

