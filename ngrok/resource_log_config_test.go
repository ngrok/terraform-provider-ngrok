// Code generated by apic. DO NOT EDIT.

package ngrok

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

var (
	resourceLogConfigs_createConfig = `resource "ngrok_log_config" "example" {
  description = "low sampling, basic HTTP logs"
  destination_ids = [ "ld_1knYnW1PywGEbP89U7CJ8PdTdcO" ]
  event_type = "http_request_complete"
  fields = [ "http.request.method", "http.response.status_code", "conn.client_ip" ]
  metadata = "{\"environment\": \"staging\"}"
  sampling_rate = 0.1
}`
	resourceLogConfigs_updateConfig = `resource "ngrok_log_config" "example" {
  description = "medium sampling, basic HTTP logs"
  sampling_rate = 0.3
}`
)

func init() {
	resource.AddTestSweepers("log_configs", &resource.Sweeper{
		Name: "log_configs",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.LogConfigsList(ctx, nil)
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}
			for _, item := range list.LogConfigs {
				// Assume items with empty Description and Metadata are system defined (i.e. API Keys)
				if item.Description != "" && item.Metadata != "" {
					_, _, err := conn.LogConfigsDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}
			return nil
		},
	})
}

func TestAccResourceLogConfigs(t *testing.T) {
	t.Skip("Test skipped. See: https://github.com/ngrok-private/ngrok/issues/4717")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyLogConfigs,
		Steps: []resource.TestStep{
			{
				Config: resourceLogConfigs_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateLogConfigs,
				// ),
			},
		},
	})
}

func testAccCheckDestroyLogConfigs(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateLogConfigs(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
