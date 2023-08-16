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
	resourceReservedAddrs_createConfig = `resource "ngrok_reserved_addr" "example" {
  description = "SSH for device #001"
  region = "us"
}`
	resourceReservedAddrs_updateConfig = `resource "ngrok_reserved_addr" "example" {
  endpoint_configuration_id = "ec_26rOxvXFvasY5kQMOlxsY37ZINe"
  metadata = "{\"proto\": \"ssh\"}"
}`
)

func init() {
	resource.AddTestSweepers("reserved_addrs", &resource.Sweeper{
		Name: "reserved_addrs",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.ReservedAddrsList(ctx, &restapi.Paging{})
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.ReservedAddrs {
				// Assume items with empty Description and Metadata are system defined (i.e. API Keys)
				if item.Description != "" && item.Metadata != "" {
					_, _, err := conn.ReservedAddrsDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}

			return nil
		},
	})
}

func TestAccResourceReservedAddrs(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyReservedAddrs,
		Steps: []resource.TestStep{
			{
				Config: resourceReservedAddrs_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateReservedAddrs,
				// ),
			},
		},
	})
}

func testAccCheckDestroyReservedAddrs(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateReservedAddrs(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
