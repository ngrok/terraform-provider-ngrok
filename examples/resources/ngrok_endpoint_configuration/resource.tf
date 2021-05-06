
resource "ngrok_endpoint_configuration" "example" {
  description = "app servers"
  request_headers {
    add = {
      X-Frontend = "ngrok"
    }
    remove = [ "Cache-Control" ]
  }
  type = "https"
}

