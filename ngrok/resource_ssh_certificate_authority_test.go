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
	resourceSSHCertificateAuthorities_createConfig = `resource "ngrok_ssh_certificate_authority" "example" {
  description = "Staging Environment Hosts"
  private_key_type = "ed25519"
}`
	resourceSSHCertificateAuthorities_updateConfig = `resource "ngrok_ssh_certificate_authority" "example" {
  metadata = "{\"region\": \"us-east-1\"}"
}`
)

func init() {
	resource.AddTestSweepers("ssh_certificate_authorities", &resource.Sweeper{
		Name: "ssh_certificate_authorities",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.SSHCertificateAuthoritiesList(ctx, &restapi.Paging{})
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.SSHCertificateAuthorities {
				// Assume items with empty Description or Metadata are system defined
				// (i.e. API Keys) so do not sweep them for cleanup.
				// However, not all items have Description and Metadata fields, so need to reflect.
				iv := reflect.ValueOf(item)
				dv := iv.FieldByName("Description")
				mv := iv.FieldByName("Metadata")
				shouldKeep := (dv.IsValid() && dv.IsZero()) || (mv.IsValid() && mv.IsZero())
				if !shouldKeep {
					_, _, err := conn.SSHCertificateAuthoritiesDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}

			return nil
		},
	})
}

func TestAccResourceSSHCertificateAuthorities(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroySSHCertificateAuthorities,
		Steps: []resource.TestStep{
			{
				Config: resourceSSHCertificateAuthorities_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateSSHCertificateAuthorities,
				// ),
			},
		},
	})
}

func testAccCheckDestroySSHCertificateAuthorities(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateSSHCertificateAuthorities(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
