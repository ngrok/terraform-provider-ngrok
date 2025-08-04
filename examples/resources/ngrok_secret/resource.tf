# Code generated for API Clients. DO NOT EDIT.


resource "ngrok_secret" "example" {
  description = "Database password for prod postgres instance"
  metadata = "env=prod,service=postgres"
  name = "db-password"
  value = "supersecret123"
  vault_id = "vault_2y0YkHvDtItsU4xNJpBPGx8EW2K"
}

