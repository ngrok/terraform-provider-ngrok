resource "ngrok_tls_certificate" "example" {
  certificate_pem = file("${path.module}/cert.pem")
  private_key_pem = file("${path.module}/key.pem")
  description     = "TLS certificate for app.example.com"
  metadata        = jsonencode({ environment = "production" })
}
