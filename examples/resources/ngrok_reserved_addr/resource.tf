resource "ngrok_reserved_addr" "example" {
  description = "Reserved TCP address for the database tunnel"
  metadata    = jsonencode({ environment = "production" })
  region      = "us"
}
