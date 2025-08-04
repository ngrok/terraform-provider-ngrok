# Code generated for API Clients. DO NOT EDIT.


resource "ngrok_vault" "example" {
  description = "Vault containing production environment secrets"
  metadata = "env=prod,team=devops"
  name = "prod-secrets"
}

