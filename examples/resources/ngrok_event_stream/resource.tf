
resource "ngrok_event_stream" "example" {
  description = "low sampling, basic HTTP logs"
  destination_ids = [ "ed_1ro7aG1J2tGT6neX0PHJLTuzQ9E" ]
  event_type = "http_request_complete"
  fields = [ "http.request.method", "http.response.status_code", "conn.client_ip" ]
  metadata = "{\"environment\": \"staging\"}"
  sampling_rate = 0.1
}

