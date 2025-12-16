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
	resourceVaults_createConfig = `resource "ngrok_vault" "example" {
  description = "Vault containing production environment secrets"
  metadata = "env=prod,team=devops"
  name = "prod-secrets"
}`
	resourceVaults_updateConfig = `resource "ngrok_vault" "example" {
  description = "Updated vault for production team"
  metadata = "env=prod,team=secops"
  name = "prod-secrets-renamed"
}`
)

func init() {
	resource.AddTestSweepers("vaults", &resource.Sweeper{
		Name: "vaults",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.VaultsList(ctx, &restapi.FilteredPaging{})
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.Vaults {
				// Assume items with empty Description or Metadata are system defined
				// (i.e. API Keys) so do not sweep them for cleanup.
				// However, not all items have Description and Metadata fields, so need to reflect.
				iv := reflect.ValueOf(item)
				dv := iv.FieldByName("Description")
				mv := iv.FieldByName("Metadata")
				shouldKeep := (dv.IsValid() && dv.IsZero()) || (mv.IsValid() && mv.IsZero())
				if !shouldKeep {
					_, _, err := conn.VaultsDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}
			return nil
		},
	})
}

func TestAccResourceVaults(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyVaults,
		Steps: []resource.TestStep{
			{
				Config: resourceVaults_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateVaults,
				// ),
			},
		},
	})
}

func testAccCheckDestroyVaults(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateVaults(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
