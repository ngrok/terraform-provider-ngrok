
resource "ngrok_ssh_host_certificate" "example" {
  description = "personal server"
  principals = [ "inconshreveable.com", "10.2.42.9" ]
  public_key = "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBI3oSgxrOEJ+tIJ/n6VYtxQIFvynqlOHpfOAJ4x4OfmMYDkbf8dr6RAuUSf+ZC2HMCujta7EjZ9t+6v08Ue+Cgk= inconshreveable.com"
  ssh_certificate_authority_id = "sshca_1rV5GiTaiBQg8AtJiyiVeNKzYiR"
  valid_until = "2021-07-20T23:38:42Z"
}

