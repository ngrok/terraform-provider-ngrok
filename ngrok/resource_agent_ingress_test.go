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
	resourceAgentIngresses_createConfig = `resource "ngrok_agent_ingress" "example" {
  description = "acme devices"
  domain = "connect.acme.com"
}`
	resourceAgentIngresses_updateConfig = `resource "ngrok_agent_ingress" "example" {
  description = "ACME Co. Device Ingress"
  metadata = "{\"device_sku\": \"824JS4RZ1F8X\"}"
}`
)

func init() {
	resource.AddTestSweepers("agent_ingresses", &resource.Sweeper{
		Name: "agent_ingresses",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.AgentIngressesList(ctx, &restapi.Paging{})
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.Ingresses {
				// Assume items with empty Description or Metadata are system defined
				// (i.e. API Keys) so do not sweep them for cleanup.
				// However, not all items have Description and Metadata fields, so need to reflect.
				iv := reflect.ValueOf(item)
				dv := iv.FieldByName("Description")
				mv := iv.FieldByName("Metadata")
				shouldKeep := (dv.IsValid() && dv.IsZero()) || (mv.IsValid() && mv.IsZero())
				if !shouldKeep {
					_, _, err := conn.AgentIngressesDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}

			return nil
		},
	})
}

func TestAccResourceAgentIngresses(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyAgentIngresses,
		Steps: []resource.TestStep{
			{
				Config: resourceAgentIngresses_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateAgentIngresses,
				// ),
			},
		},
	})
}

func testAccCheckDestroyAgentIngresses(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateAgentIngresses(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
