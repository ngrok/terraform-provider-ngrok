resource "ngrok_secret" "example" {
  name        = "datadog-api-key"
  value       = "s3cr3t-value"
  vault_id    = ngrok_vault.example.id
  description = "API key consumed by the traffic policy"
  metadata    = jsonencode({ environment = "production" })
}
