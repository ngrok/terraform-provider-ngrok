
resource "ngrok_failover_backend" "example" {
  backends = [ "bkdhr_26rOyncxuCZ0JdIjYiEDGlsh1lO" ]
  description = "acme failover"
  metadata = "{\"environment\": \"staging\"}"
}

