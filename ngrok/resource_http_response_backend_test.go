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
	resourceHTTPResponseBackends_createConfig = `resource "ngrok_http_response_backend" "example" {
  body = "I'm a teapot"
  description = "acme http response"
  headers = {
    Content-Type = "text/plain"
  }
  metadata = "{\"environment\": \"staging\"}"
  status_code = 418
}`
	resourceHTTPResponseBackends_updateConfig = `resource "ngrok_http_response_backend" "example" {
  metadata = "{\"environment\": \"production\"}"
}`
)

func init() {
	resource.AddTestSweepers("http_response_backends", &resource.Sweeper{
		Name: "http_response_backends",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.HTTPResponseBackendsList(ctx, &restapi.Paging{})
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.Backends {
				// Assume items with empty Description or Metadata are system defined
				// (i.e. API Keys) so do not sweep them for cleanup.
				// However, not all items have Description and Metadata fields, so need to reflect.
				iv := reflect.ValueOf(item)
				dv := iv.FieldByName("Description")
				mv := iv.FieldByName("Metadata")
				shouldKeep := (dv.IsValid() && dv.IsZero()) || (mv.IsValid() && mv.IsZero())
				if !shouldKeep {
					_, _, err := conn.HTTPResponseBackendsDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}

			return nil
		},
	})
}

func TestAccResourceHTTPResponseBackends(t *testing.T) {
	t.Skip("Headers need to be diff-suppressed / canonicalized, but https://github.com/hashicorp/terraform-plugin-sdk/issues/477#issuecomment-646351613")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyHTTPResponseBackends,
		Steps: []resource.TestStep{
			{
				Config: resourceHTTPResponseBackends_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateHTTPResponseBackends,
				// ),
			},
		},
	})
}

func testAccCheckDestroyHTTPResponseBackends(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateHTTPResponseBackends(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
