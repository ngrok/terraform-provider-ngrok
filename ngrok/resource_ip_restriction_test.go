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
	resourceIPRestrictions_createConfig = `resource "ngrok_ip_restriction" "example" {
  ip_policy_ids = [ "ipp_26rOyhglKmVz5ABMOwZwPFBFXBZ" ]
  type = "dashboard"
}`
	resourceIPRestrictions_updateConfig = `resource "ngrok_ip_restriction" "example" {
  ip_policy_ids = [ "ipp_26rOyhglKmVz5ABMOwZwPFBFXBZ", "ipp_26rOyh3YNBLliukcO0rQFJJksSp" ]
}`
)

func init() {
	resource.AddTestSweepers("ip_restrictions", &resource.Sweeper{
		Name: "ip_restrictions",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.IPRestrictionsList(ctx, &restapi.Paging{})
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.IPRestrictions {
				// Assume items with empty Description and Metadata are system defined (i.e. API Keys)
				if item.Description != "" && item.Metadata != "" {
					_, _, err := conn.IPRestrictionsDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}

			return nil
		},
	})
}

func TestAccResourceIPRestrictions(t *testing.T) {
	t.Skip("Test skipped.")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyIPRestrictions,
		Steps: []resource.TestStep{
			{
				Config: resourceIPRestrictions_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateIPRestrictions,
				// ),
			},
		},
	})
}

func testAccCheckDestroyIPRestrictions(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateIPRestrictions(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
