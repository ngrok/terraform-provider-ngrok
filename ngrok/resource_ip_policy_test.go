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
	resourceIPPolicies_createConfig = `resource "ngrok_ip_policy" "example" {
  description = "API Outbound Gateway"
}`
	resourceIPPolicies_updateConfig = `resource "ngrok_ip_policy" "example" {
  metadata = "metadata={\"pod-id\": \"b3d9c464-4f48-4783-a741-d7d7d5db310f\"}"
}`
)

func init() {
	resource.AddTestSweepers("ip_policies", &resource.Sweeper{
		Name: "ip_policies",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.IPPoliciesList(ctx, &restapi.Paging{})
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.IPPolicies {
				// Assume items with empty Description or Metadata are system defined
				// (i.e. API Keys) so do not sweep them for cleanup.
				// However, not all items have Description and Metadata fields, so need to reflect.
				iv := reflect.ValueOf(item)
				dv := iv.FieldByName("Description")
				mv := iv.FieldByName("Metadata")
				shouldKeep := (dv.IsValid() && dv.IsZero()) || (mv.IsValid() && mv.IsZero())
				if !shouldKeep {
					_, _, err := conn.IPPoliciesDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}

			return nil
		},
	})
}

func TestAccResourceIPPolicies(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyIPPolicies,
		Steps: []resource.TestStep{
			{
				Config: resourceIPPolicies_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateIPPolicies,
				// ),
			},
		},
	})
}

func testAccCheckDestroyIPPolicies(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateIPPolicies(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
