resource "ngrok_ssh_credential" "example" {
  public_key  = file("${path.module}/id_ed25519.pub")
  description = "SSH credential for the deploy bot"
  metadata    = jsonencode({ environment = "production" })
  acl         = ["bind:app.example.com"]
}
