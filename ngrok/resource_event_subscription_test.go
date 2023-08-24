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
	resourceEventSubscriptions_createConfig = `resource "ngrok_event_subscription" "example" {
  description = "ip policy creations"
  destination_ids = [ "ed_26rOygIJTeAVyFkkw0z9hqMSv0p" ]
  metadata = "{\"environment\": \"staging\"}"
  sources [ {
    type = "ip_policy_created.v0"
  } ]
}`
	resourceEventSubscriptions_updateConfig = `resource "ngrok_event_subscription" "example" {
  description = "IP Policy Creations"
}`
)

func init() {
	resource.AddTestSweepers("event_subscriptions", &resource.Sweeper{
		Name: "event_subscriptions",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.EventSubscriptionsList(ctx, &restapi.Paging{})
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.EventSubscriptions {
				// Assume items with empty Description and Metadata are system defined (i.e. API Keys)
				if item.Description != "" && item.Metadata != "" {
					_, _, err := conn.EventSubscriptionsDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}

			return nil
		},
	})
}

func TestAccResourceEventSubscriptions(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyEventSubscriptions,
		Steps: []resource.TestStep{
			{
				Config: resourceEventSubscriptions_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateEventSubscriptions,
				// ),
			},
		},
	})
}

func testAccCheckDestroyEventSubscriptions(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateEventSubscriptions(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
