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
	resourceReservedDomains_createConfig = `resource "ngrok_reserved_domain" "example" {
  certificate_id = "cert_26rOxyrxCJlOc0frz7MK0HQjRvd"
  name = "myapp.mydomain.com"
  region = "us"
}`
	resourceReservedDomains_updateConfig = `resource "ngrok_reserved_domain" "example" {
  certificate_management_policy {
    authority = "letsencrypt"
  }
  description = "point-of-sale new york #302"
  http_endpoint_configuration_id = "ec_26rOy1P7jzFxlJtlpw0ZnCXGUUs"
  https_endpoint_configuration_id = "ec_26rOy0zF5qSTR8fMz97lg00i1Mm"
  metadata = "{env: \"staging\", \"connector_id\":\"64698fcc-5f5c-4b63-910e-8669d04bd943\"}"
}`
)

func init() {
	resource.AddTestSweepers("reserved_domains", &resource.Sweeper{
		Name: "reserved_domains",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.ReservedDomainsList(ctx, &restapi.Paging{})
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.ReservedDomains {
				// Assume items with empty Description or Metadata are system defined
				// (i.e. API Keys) so do not sweep them for cleanup.
				// However, not all items have Description and Metadata fields, so need to reflect.
				iv := reflect.ValueOf(item)
				dv := iv.FieldByName("Description")
				mv := iv.FieldByName("Metadata")
				shouldKeep := (dv.IsValid() && dv.IsZero()) || (mv.IsValid() && mv.IsZero())
				if !shouldKeep {
					_, _, err := conn.ReservedDomainsDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}

			return nil
		},
	})
}

func TestAccResourceReservedDomains(t *testing.T) {
	t.Skip("Test skipped.")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyReservedDomains,
		Steps: []resource.TestStep{
			{
				Config: resourceReservedDomains_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateReservedDomains,
				// ),
			},
		},
	})
}

func testAccCheckDestroyReservedDomains(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateReservedDomains(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
