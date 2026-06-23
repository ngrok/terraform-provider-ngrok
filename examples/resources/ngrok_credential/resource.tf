resource "ngrok_credential" "example" {
  description = "Tunnel credential for the production agent"
  metadata    = jsonencode({ environment = "production" })
  acl         = ["bind:app.example.com"]
}
