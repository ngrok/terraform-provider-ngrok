
resource "ngrok_ip_restriction" "example" {
  ip_policy_ids = [ "ipp_25auGwa4eEWUeCOBfCZkwtwqFey" ]
  type = "dashboard"
}

