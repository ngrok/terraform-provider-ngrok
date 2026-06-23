resource "ngrok_ip_policy" "example" {
  description = "Allowlist for office networks"
  metadata    = jsonencode({ team = "security" })
}
