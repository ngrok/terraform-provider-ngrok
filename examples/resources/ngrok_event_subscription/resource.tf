resource "ngrok_event_subscription" "example" {
  description     = "Subscription for HTTP request events"
  metadata        = jsonencode({ environment = "production" })
  destination_ids = [ngrok_event_destination.example.id]

  sources = [
    {
      type = "http_request_complete.v0"
    }
  ]
}
