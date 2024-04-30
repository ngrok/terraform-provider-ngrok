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
	resourceSSHHostCertificates_createConfig = `resource "ngrok_ssh_host_certificate" "example" {
  description = "personal server"
  principals = [ "inconshreveable.com", "10.2.42.9" ]
  public_key = "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBI3oSgxrOEJ+tIJ/n6VYtxQIFvynqlOHpfOAJ4x4OfmMYDkbf8dr6RAuUSf+ZC2HMCujta7EjZ9t+6v08Ue+Cgk= inconshreveable.com"
  ssh_certificate_authority_id = "sshca_26rOyuA7GzMmCmvfOui9TPWNxLa"
  valid_until = "2022-06-22T22:21:35-05:00"
}`
	resourceSSHHostCertificates_updateConfig = `resource "ngrok_ssh_host_certificate" "example" {
  metadata = "{\"region\": \"us-west-2\"}"
}`
)

func init() {
	resource.AddTestSweepers("ssh_host_certificates", &resource.Sweeper{
		Name: "ssh_host_certificates",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.SSHHostCertificatesList(ctx, &restapi.Paging{})
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.SSHHostCertificates {
				// Assume items with empty Description or Metadata are system defined
				// (i.e. API Keys) so do not sweep them for cleanup.
				// However, not all items have Description and Metadata fields, so need to reflect.
				iv := reflect.ValueOf(item)
				dv := iv.FieldByName("Description")
				mv := iv.FieldByName("Metadata")
				shouldKeep := (dv.IsValid() && dv.IsZero()) || (mv.IsValid() && mv.IsZero())
				if !shouldKeep {
					_, _, err := conn.SSHHostCertificatesDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}

			return nil
		},
	})
}

func TestAccResourceSSHHostCertificates(t *testing.T) {
	t.Skip("Test skipped.")
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
