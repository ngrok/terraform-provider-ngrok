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
	resourceSecrets_createConfig = `resource "ngrok_secret" "example" {
  description = "Database password for prod postgres instance"
  metadata = "env=prod,service=postgres"
  name = "db-password"
  value = "supersecret123"
  vault_id = "vault_2y0YkHvDtItsU4xNJpBPGx8EW2K"
}`
	resourceSecrets_updateConfig = `resource "ngrok_secret" "example" {
  description = "Database password for prod postgres instance 2025"
  metadata = "env=prod,service=postgres,rotation=2025"
  name = "db-password-2025"
  value = "rotated-secret-2025"
}`
)

func init() {
	resource.AddTestSweepers("secrets", &resource.Sweeper{
		Name: "secrets",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.SecretsList(ctx, &restapi.FilteredPaging{})
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.Secrets {
				// Assume items with empty Description or Metadata are system defined
				// (i.e. API Keys) so do not sweep them for cleanup.
				// However, not all items have Description and Metadata fields, so need to reflect.
				iv := reflect.ValueOf(item)
				dv := iv.FieldByName("Description")
				mv := iv.FieldByName("Metadata")
				shouldKeep := (dv.IsValid() && dv.IsZero()) || (mv.IsValid() && mv.IsZero())
				if !shouldKeep {
					_, _, err := conn.SecretsDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}
			return nil
		},
	})
}

func TestAccResourceSecrets(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroySecrets,
		Steps: []resource.TestStep{
			{
				Config: resourceSecrets_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateSecrets,
				// ),
			},
		},
	})
}

func testAccCheckDestroySecrets(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateSecrets(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
