# Code generated for API Clients. DO NOT EDIT.


resource "ngrok_kubernetes_operator" "example" {
  deployment {
    name = "ngrok-operator"
    namespace = "ngrok-operator"
    version = "0.12.2"
  }
  description = "Created by ngrok-operator"
  enabled_features = [ "Ingress", "Bindings" ]
  metadata = "{\"namespace.uid\":\"9663c1aa-10e4-4933-8576-398a49a5caf6\",\"owned-by\":\"ngrok-operator\"}"
  region = "global"
}

