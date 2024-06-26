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
	resourceEventDestinations_createConfig = `resource "ngrok_event_destination" "example" {
  description = "kinesis dev stream"
  format = "json"
  metadata = "{\"environment\":\"dev\"}"
  target {
    kinesis {
      auth {
        role {
          role_arn = "arn:aws:iam::123456789012:role/example"
        }
      }
      stream_arn = "arn:ngrok-local:kinesis:us-east-2:123456789012:stream/mystream2"
    }
  }
}`
	resourceEventDestinations_updateConfig = `resource "ngrok_event_destination" "example" {
  description = "kinesis dev stream 1 of 3"
  metadata = "{\"environment\":\"dev\", \"stream\":1}"
}`
)

func init() {
	resource.AddTestSweepers("event_destinations", &resource.Sweeper{
		Name: "event_destinations",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.EventDestinationsList(ctx, &restapi.Paging{})
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.EventDestinations {
				// Assume items with empty Description or Metadata are system defined
				// (i.e. API Keys) so do not sweep them for cleanup.
				// However, not all items have Description and Metadata fields, so need to reflect.
				iv := reflect.ValueOf(item)
				dv := iv.FieldByName("Description")
				mv := iv.FieldByName("Metadata")
				shouldKeep := (dv.IsValid() && dv.IsZero()) || (mv.IsValid() && mv.IsZero())
				if !shouldKeep {
					_, _, err := conn.EventDestinationsDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}

			return nil
		},
	})
}

func TestAccResourceEventDestinations(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyEventDestinations,
		Steps: []resource.TestStep{
			{
				Config: resourceEventDestinations_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateEventDestinations,
				// ),
			},
		},
	})
}

func testAccCheckDestroyEventDestinations(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateEventDestinations(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
