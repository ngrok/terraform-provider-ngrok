# Code generated for API Clients. DO NOT EDIT.


resource "ngrok_event_subscription" "example" {
  description = "ip policy creations"
  destination_ids = [ "ed_26rOygIJTeAVyFkkw0z9hqMSv0p" ]
  metadata = "{\"environment\": \"staging\"}"
  sources [ {
    type = "ip_policy_created.v0"
  } ]
}

