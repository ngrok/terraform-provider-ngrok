
resource "ngrok_ssh_user_certificate" "example" {
  description = "temporary access to staging machine"
  principals = [ "ec2-user", "root" ]
  public_key = "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBK58lFzmWlDimDtBz78wVT4oauA8PjY0CiXTCEIsBNC6UwOJvZ0jdSaYNhDaa7dRV84DfBb/gKzqlXC7cVMZjl0= alan@work-laptop"
  ssh_certificate_authority_id = "sshca_25auH5JtiUPW9eMiXYzujvcpkGW"
  valid_until = "2022-05-26T08:23:47Z"
}

