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
	resourceCredentials_createConfig = `resource "ngrok_credential" "example" {
  description = "development cred for alan@example.com"
}`
	resourceCredentials_updateConfig = `resource "ngrok_credential" "example" {
  description = "device alpha-2"
  metadata = "{\"device_id\": \"d5111ba7-0cc5-4ba3-8398-e6c79e4e89c2\"}"
}`
)

func init() {
	resource.AddTestSweepers("credentials", &resource.Sweeper{
		Name: "credentials",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.CredentialsList(ctx, &restapi.Paging{})
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.Credentials {
				// Assume items with empty Description or Metadata are system defined
				// (i.e. API Keys) so do not sweep them for cleanup.
				// However, not all items have Description and Metadata fields, so need to reflect.
				iv := reflect.ValueOf(item)
				dv := iv.FieldByName("Description")
				mv := iv.FieldByName("Metadata")
				shouldKeep := (dv.IsValid() && dv.IsZero()) || (mv.IsValid() && mv.IsZero())
				if !shouldKeep {
					_, _, err := conn.CredentialsDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}

			return nil
		},
	})
}

func TestAccResourceCredentials(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyCredentials,
		Steps: []resource.TestStep{
			{
				Config: resourceCredentials_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateCredentials,
				// ),
			},
		},
	})
}

func testAccCheckDestroyCredentials(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateCredentials(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
