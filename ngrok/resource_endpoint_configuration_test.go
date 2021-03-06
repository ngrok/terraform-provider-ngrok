// Code generated by apic. DO NOT EDIT.

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
	resourceEndpointConfigurations_createConfig = `resource "ngrok_endpoint_configuration" "example" {
  description = "app servers"
  request_headers {
    add = {
      x-frontend = "ngrok"
    }
    remove = [ "cache-control" ]
  }
  type = "https"
}`
	resourceEndpointConfigurations_updateConfig = `resource "ngrok_endpoint_configuration" "example" {
  ip_policy {
  }
}`
)

func init() {
	resource.AddTestSweepers("endpoint_configurations", &resource.Sweeper{
		Name: "endpoint_configurations",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.EndpointConfigurationsList(ctx, nil)
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.EndpointConfigurations {
				// Assume items with empty Description and Metadata are system defined (i.e. API Keys)
				if item.Description != "" && item.Metadata != "" {
					_, _, err := conn.EndpointConfigurationsDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}

			return nil
		},
	})
}

func TestAccResourceEndpointConfigurations(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyEndpointConfigurations,
		Steps: []resource.TestStep{
			{
				Config: resourceEndpointConfigurations_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateEndpointConfigurations,
				// ),
			},
		},
	})
}

func testAccCheckDestroyEndpointConfigurations(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateEndpointConfigurations(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
