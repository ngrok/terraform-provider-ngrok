// Code generated by apic. DO NOT EDIT.

package ngrok

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var (
	resourceSSHUserCertificates_createConfig = `resource "ngrok_ssh_user_certificate" "example" {
  description = "temporary access to staging machine"
  principals = [ "ec2-user", "root" ]
  public_key = "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBK58lFzmWlDimDtBz78wVT4oauA8PjY0CiXTCEIsBNC6UwOJvZ0jdSaYNhDaa7dRV84DfBb/gKzqlXC7cVMZjl0= alan@work-laptop"
  ssh_certificate_authority_id = "sshca_1kONCO9Ja0gvqHAZG9sTIgPZPdd"
  valid_until = "2021-02-14T22:59:09Z"
}`
	resourceSSHUserCertificates_updateConfig = `resource "ngrok_ssh_user_certificate" "example" {
  description = "temporary access to staging machine for alan"
  metadata = "{\"user_email\": \"alan@example.com\"}"
}`
)

func TestAccResourceSSHUserCertificates(t *testing.T) {
	t.Skip("Test skipped. See: https://github.com/ngrok-private/ngrok/issues/4720")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroySSHUserCertificates,
		Steps: []resource.TestStep{
			{
				Config: resourceSSHUserCertificates_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateSSHUserCertificates,
				// ),
			},
		},
	})
}

func testAccCheckDestroySSHUserCertificates(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateSSHUserCertificates(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
