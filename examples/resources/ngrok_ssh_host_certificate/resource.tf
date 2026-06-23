resource "ngrok_ssh_host_certificate" "example" {
  ssh_certificate_authority_id = ngrok_ssh_certificate_authority.example.id
  public_key                   = file("${path.module}/ssh_host_ed25519_key.pub")
  principals                   = ["host.example.com"]
  description                  = "Host certificate for the production bastion"
  metadata                     = jsonencode({ environment = "production" })
}
