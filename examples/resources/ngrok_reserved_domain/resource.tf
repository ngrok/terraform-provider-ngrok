
resource "ngrok_reserved_domain" "example" {
  certificate_id = "cert_1rV51OQetZPK9V6vTWUVy3Onjir"
  name = "myapp.mydomain.com"
  region = "us"
}

