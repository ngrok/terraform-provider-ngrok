resource "ngrok_certificate_authority" "example" {
  ca_pem      = file("${path.module}/ca.pem")
  description = "Corporate root certificate authority"
  metadata    = jsonencode({ team = "platform" })
}
