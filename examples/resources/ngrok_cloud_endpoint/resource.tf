resource "ngrok_cloud_endpoint" "example" {
  url         = "https://app.example.com"
  description = "Production cloud endpoint"
  metadata    = jsonencode({ environment = "production" })

  traffic_policy = jsonencode({
    on_http_request = [
      {
        actions = [
          {
            type = "custom-response"
            config = {
              status_code = 200
              body        = "Hello from ngrok"
            }
          }
        ]
      }
    ]
  })
}
