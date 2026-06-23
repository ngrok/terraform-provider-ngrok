resource "ngrok_reserved_domain" "example" {
  domain      = "app.example.com"
  description = "Production application domain"
  metadata    = jsonencode({ team = "platform" })

  certificate_management_policy = {
    authority        = "letsencrypt"
    private_key_type = "ecdsa"
  }
}
