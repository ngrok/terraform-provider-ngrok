# Lookup by ID
data "ngrok_reserved_domain" "by_id" {
  id = "rd_2example"
}

# Lookup by domain name
data "ngrok_reserved_domain" "by_name" {
  domain = "app.example.com"
}
