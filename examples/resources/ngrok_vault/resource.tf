resource "ngrok_vault" "example" {
  name        = "production-secrets"
  description = "Vault holding production secrets"
  metadata    = jsonencode({ environment = "production" })
}
