terraform {
  required_providers {
    ngrok = {
      source  = "ngrok/ngrok"
      version = "~> 1.0"
    }
  }
}

# Configure via NGROK_API_KEY environment variable (recommended)
provider "ngrok" {}

# Or configure explicitly
# provider "ngrok" {
#   api_key = "your-api-key"
# }
