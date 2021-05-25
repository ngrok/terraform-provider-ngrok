
resource "ngrok_endpoint_configuration" "example" {
  description = "app servers"
  request_headers {
    add = {
      x-frontend = "ngrok"
    }
    remove = [ "cache-control" ]
  }
  type = "https"
}

