
resource "ngrok_ip_whitelist_entry" "example" {
  description = "outbound proxy servers"
  ip_net = "10.1.1.0/24"
}

