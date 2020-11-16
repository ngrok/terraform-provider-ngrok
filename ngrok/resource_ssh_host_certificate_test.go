// Code generated by apic. DO NOT EDIT.

package ngrok

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var (
	resourceSSHHostCertificates_createConfig = `resource "ngrok_ssh_host_certificate" "example" {
  description = "personal server"
  principals = [ "inconshreveable.com", "10.2.42.9" ]
  public_key = "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBI3oSgxrOEJ+tIJ/n6VYtxQIFvynqlOHpfOAJ4x4OfmMYDkbf8dr6RAuUSf+ZC2HMCujta7EjZ9t+6v08Ue+Cgk= inconshreveable.com"
  ssh_certificate_authority_id = "sshca_1kOTHco3uORrsjO2vdJK5FckNed"
  valid_until = "2021-02-14T23:49:11Z"
}`
	resourceSSHHostCertificates_updateConfig = `resource "ngrok_ssh_host_certificate" "example" {
  metadata = "{\"region\": \"us-west-2\"}"
}`
)

func TestAccResourceSSHHostCertificates(t *testing.T) {
	t.Skip("Test skipped. See: https://github.com/ngrok-private/ngrok/issues/4719")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroySSHHostCertificates,
		Steps: []resource.TestStep{
			{
				Config: resourceSSHHostCertificates_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateSSHHostCertificates,
				// ),
			},
		},
	})
}

func testAccCheckDestroySSHHostCertificates(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateSSHHostCertificates(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
