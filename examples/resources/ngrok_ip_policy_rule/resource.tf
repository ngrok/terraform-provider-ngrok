
resource "ngrok_ip_policy_rule" "example" {
  cidr = "212.3.14.0/24"
  description = "nyc office"
  ip_policy_id = "ipp_25auGv9R7vPmi6NKs5Cxcyzc2Cm"
}

