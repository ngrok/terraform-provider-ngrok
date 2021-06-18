
resource "ngrok_event_subscription" "example" {
  description = "low sampling, basic HTTP logs"
  destination_ids = [ "ed_1ro7aylyqQ1LLMWNWrOISvlfveQ" ]
  metadata = "{\"environment\": \"staging\"}"
  sources [ {
    type = "http_request_complete"
  } ]
}

