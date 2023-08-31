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
	resourceTunnelGroupBackends_createConfig = `resource "ngrok_tunnel_group_backend" "example" {
  description = "acme tunnel group"
  labels = {
    baz = "qux"
    foo = "bar"
  }
  metadata = "{\"environment\": \"staging\"}"
}`
	resourceTunnelGroupBackends_updateConfig = `resource "ngrok_tunnel_group_backend" "example" {
  metadata = "{\"environment\": \"production\"}"
}`
)

func init() {
	resource.AddTestSweepers("tunnel_group_backends", &resource.Sweeper{
		Name: "tunnel_group_backends",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.TunnelGroupBackendsList(ctx, &restapi.Paging{})
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.Backends {
				// Assume items with empty Description and Metadata are system defined (i.e. API Keys)
				if item.Description != "" && item.Metadata != "" {
					_, _, err := conn.TunnelGroupBackendsDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}

			return nil
		},
	})
}

func TestAccResourceTunnelGroupBackends(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyTunnelGroupBackends,
		Steps: []resource.TestStep{
			{
				Config: resourceTunnelGroupBackends_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateTunnelGroupBackends,
				// ),
			},
		},
	})
}

func testAccCheckDestroyTunnelGroupBackends(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateTunnelGroupBackends(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
