// Code generated by apic. DO NOT EDIT.

package ngrok

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

var (
	resourceSSHUserCertificates_createConfig = `resource "ngrok_ssh_user_certificate" "example" {
  description = "temporary access to staging machine"
  principals = [ "ec2-user", "root" ]
  public_key = "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBK58lFzmWlDimDtBz78wVT4oauA8PjY0CiXTCEIsBNC6UwOJvZ0jdSaYNhDaa7dRV84DfBb/gKzqlXC7cVMZjl0= alan@work-laptop"
  ssh_certificate_authority_id = "sshca_1knYod4EQ6mRQwwD5zZzFgHndHk"
  valid_until = "2021-02-23T20:59:58Z"
}`
	resourceSSHUserCertificates_updateConfig = `resource "ngrok_ssh_user_certificate" "example" {
  description = "temporary access to staging machine for alan"
  metadata = "{\"user_email\": \"alan@example.com\"}"
}`
)

func init() {
	resource.AddTestSweepers("ssh_user_certificates", &resource.Sweeper{
		Name: "ssh_user_certificates",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.SSHUserCertificatesList(ctx, nil)
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}
			for _, item := range list.SSHUserCertificates {
				// Assume items with empty Description and Metadata are system defined (i.e. API Keys)
				if item.Description != "" && item.Metadata != "" {
					_, _, err := conn.SSHUserCertificatesDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}
			return nil
		},
	})
}

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
