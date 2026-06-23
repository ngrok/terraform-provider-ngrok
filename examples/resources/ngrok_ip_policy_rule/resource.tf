resource "ngrok_ip_policy_rule" "example" {
  ip_policy_id = ngrok_ip_policy.example.id
  action       = "allow"
  cidr         = "203.0.113.0/24"
  description  = "Allow the office network"
  metadata     = jsonencode({ team = "security" })
}
