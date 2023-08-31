# Code generated for API Clients. DO NOT EDIT.


resource "ngrok_http_response_backend" "example" {
  body = "I'm a teapot"
  description = "acme http response"
  headers = {
    Content-Type = "text/plain"
  }
  metadata = "{\"environment\": \"staging\"}"
  status_code = 418
}

