
resource "ngrok_tunnel_group_backend" "example" {
  description = "acme tunnel group"
  labels = {
    baz = "qux"
    foo = "bar"
  }
  metadata = "{\"environment\": \"staging\"}"
}

