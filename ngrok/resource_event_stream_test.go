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
	resourceEventStreams_createConfig = `resource "ngrok_event_stream" "example" {
  description = "low sampling, basic HTTP logs"
  destination_ids = [ "ed_1ro7aG1J2tGT6neX0PHJLTuzQ9E" ]
  event_type = "http_request_complete"
  fields = [ "http.request.method", "http.response.status_code", "conn.client_ip" ]
  metadata = "{\"environment\": \"staging\"}"
  sampling_rate = 0.1
}`
	resourceEventStreams_updateConfig = `resource "ngrok_event_stream" "example" {
  description = "medium sampling, basic HTTP logs"
  sampling_rate = 0.3
}`
)

func init() {
	resource.AddTestSweepers("event_streams", &resource.Sweeper{
		Name: "event_streams",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.EventStreamsList(ctx, nil)
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}
			for _, item := range list.EventStreams {
				// Assume items with empty Description and Metadata are system defined (i.e. API Keys)
				if item.Description != "" && item.Metadata != "" {
					_, _, err := conn.EventStreamsDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}
			return nil
		},
	})
}

func TestAccResourceEventStreams(t *testing.T) {
	t.Skip("Test skipped.")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyEventStreams,
		Steps: []resource.TestStep{
			{
				Config: resourceEventStreams_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateEventStreams,
				// ),
			},
		},
	})
}

func testAccCheckDestroyEventStreams(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateEventStreams(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
