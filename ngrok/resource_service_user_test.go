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
	resourceServiceUsers_createConfig = `resource "ngrok_service_user" "example" {
  name = "new service user from API"
}`
	resourceServiceUsers_updateConfig = `resource "ngrok_service_user" "example" {
  active = false
  name = "inactive service user from API"
}`
)

func init() {
	resource.AddTestSweepers("service_users", &resource.Sweeper{
		Name: "service_users",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.ServiceUsersList(ctx, &restapi.Paging{})
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.ServiceUsers {
				// Assume items with empty Description or Metadata are system defined
				// (i.e. API Keys) so do not sweep them for cleanup.
				// However, not all items have Description and Metadata fields, so need to reflect.
				iv := reflect.ValueOf(item)
				dv := iv.FieldByName("Description")
				mv := iv.FieldByName("Metadata")
				shouldKeep := (dv.IsValid() && dv.IsZero()) || (mv.IsValid() && mv.IsZero())
				if !shouldKeep {
					_, _, err := conn.ServiceUsersDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}

			return nil
		},
	})
}

func TestAccResourceServiceUsers(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyServiceUsers,
		Steps: []resource.TestStep{
			{
				Config: resourceServiceUsers_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateServiceUsers,
				// ),
			},
		},
	})
}

func testAccCheckDestroyServiceUsers(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateServiceUsers(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
