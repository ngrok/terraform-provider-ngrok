resource "ngrok_ssh_user_certificate" "example" {
  ssh_certificate_authority_id = ngrok_ssh_certificate_authority.example.id
  public_key                   = file("${path.module}/id_ed25519.pub")
  principals                   = ["deploy"]
  description                  = "User certificate for the deploy account"
  metadata                     = jsonencode({ environment = "production" })
}
