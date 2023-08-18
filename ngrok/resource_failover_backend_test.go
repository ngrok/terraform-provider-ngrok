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
	resourceFailoverBackends_createConfig = `resource "ngrok_failover_backend" "example" {
  backends = [ "bkdhr_26rOyncxuCZ0JdIjYiEDGlsh1lO" ]
  description = "acme failover"
  metadata = "{\"environment\": \"staging\"}"
}`
	resourceFailoverBackends_updateConfig = `resource "ngrok_failover_backend" "example" {
  metadata = "{\"environment\": \"production\"}"
}`
)

func init() {
	resource.AddTestSweepers("failover_backends", &resource.Sweeper{
		Name: "failover_backends",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.FailoverBackendsList(ctx, &restapi.Paging{})
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.Backends {
				// Assume items with empty Description and Metadata are system defined (i.e. API Keys)
				if item.Description != "" && item.Metadata != "" {
					_, _, err := conn.FailoverBackendsDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}

			return nil
		},
	})
}

func TestAccResourceFailoverBackends(t *testing.T) {
	t.Skip("References other ngrok IDs by string, not terraform reference; test gen needs to be updated for this")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyFailoverBackends,
		Steps: []resource.TestStep{
			{
				Config: resourceFailoverBackends_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateFailoverBackends,
				// ),
			},
		},
	})
}

func testAccCheckDestroyFailoverBackends(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateFailoverBackends(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
