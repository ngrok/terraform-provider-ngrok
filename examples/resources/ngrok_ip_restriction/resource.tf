
resource "ngrok_ip_restriction" "example" {
  ip_policy_ids = [ "ipp_26rOyhglKmVz5ABMOwZwPFBFXBZ" ]
  type = "dashboard"
}

