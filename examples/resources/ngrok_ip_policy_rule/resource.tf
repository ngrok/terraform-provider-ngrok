
resource "ngrok_ip_policy_rule" "example" {
  action = "allow"
  cidr = "212.3.14.0/24"
  description = "nyc office"
  ip_policy_id = "ipp_26rOydjEUNZSLTi8bYXBg278qUT"
}

