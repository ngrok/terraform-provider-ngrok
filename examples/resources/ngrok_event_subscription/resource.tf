
resource "ngrok_event_subscription" "example" {
  description = "ip policy creations"
  destination_ids = [ "ed_25auH2H0JNlDGXUH01Z3sZdgFFM" ]
  metadata = "{\"environment\": \"staging\"}"
  sources [ {
    type = "ip_policy_created.v0"
  } ]
}

