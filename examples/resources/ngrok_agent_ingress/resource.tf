resource "ngrok_agent_ingress" "example" {
  domain      = "agents.example.com"
  description = "Agent ingress for production agents"
  metadata    = jsonencode({ environment = "production" })
}
