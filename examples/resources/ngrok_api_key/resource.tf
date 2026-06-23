resource "ngrok_api_key" "example" {
  description = "ngrok API key for CI automation"
  metadata    = jsonencode({ environment = "production" })
}
