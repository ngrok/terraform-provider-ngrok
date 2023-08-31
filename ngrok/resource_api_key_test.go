package ngrok

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

var (
	resourceAPIKeys_createConfig = `resource "ngrok_api_key" "example" {
  description = "ad-hoc dev testing"
  metadata = "{\"environment\":\"dev\"}"
}`
	resourceAPIKeys_updateConfig = `resource "ngrok_api_key" "example" {
  metadata = "{\"environment\":\"dev\", \"owner_id\": 123}"
}`
)

func init() {
	resource.AddTestSweepers("api_keys", &resource.Sweeper{
		Name: "api_keys",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.APIKeysList(ctx, &restapi.Paging{})
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.Keys {
				// Assume items with empty Description and Metadata are system defined (i.e. API Keys)
				if item.Description != "" && item.Metadata != "" {
					_, _, err := conn.APIKeysDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}

			return nil
		},
	})
}

func TestAccResourceAPIKeys(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyAPIKeys,
		Steps: []resource.TestStep{
			{
				Config: resourceAPIKeys_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateAPIKeys,
				// ),
			},
		},
	})
}

func testAccCheckDestroyAPIKeys(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateAPIKeys(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
