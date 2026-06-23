resource "ngrok_ssh_certificate_authority" "example" {
  description      = "SSH certificate authority for production hosts"
  metadata         = jsonencode({ environment = "production" })
  private_key_type = "ecdsa"
  elliptic_curve   = "p256"
}
