resource "ngrok_ip_restriction" "example" {
  type          = "dashboard"
  ip_policy_ids = [ngrok_ip_policy.example.id]
  description   = "Restrict dashboard access to office networks"
  metadata      = jsonencode({ team = "security" })
  enforced      = true
}
