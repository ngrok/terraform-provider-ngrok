package ngrok

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

var (
	resourceSSHUserCertificates_createConfig = `resource "ngrok_ssh_user_certificate" "example" {
  description = "temporary access to staging machine"
  principals = [ "ec2-user", "root" ]
  public_key = "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBK58lFzmWlDimDtBz78wVT4oauA8PjY0CiXTCEIsBNC6UwOJvZ0jdSaYNhDaa7dRV84DfBb/gKzqlXC7cVMZjl0= alan@work-laptop"
  ssh_certificate_authority_id = "sshca_26rOyirnW8khUZJ8xjNfPu3GPdi"
  valid_until = "2022-06-22T22:21:34-05:00"
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

			list, _, err := conn.SSHUserCertificatesList(ctx, &restapi.Paging{})
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.SSHUserCertificates {
				// Assume items with empty Description or Metadata are system defined
				// (i.e. API Keys) so do not sweep them for cleanup.
				// However, not all items have Description and Metadata fields, so need to reflect.
				iv := reflect.ValueOf(item)
				dv := iv.FieldByName("Description")
				mv := iv.FieldByName("Metadata")
				shouldKeep := (dv.IsValid() && dv.IsZero()) || (mv.IsValid() && mv.IsZero())
				if !shouldKeep {
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
	t.Skip("Test skipped.")
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
