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
	resourceBotUsers_createConfig = `resource "ngrok_bot_user" "example" {
  name = "new bot user from API"
}`
	resourceBotUsers_updateConfig = `resource "ngrok_bot_user" "example" {
  active = false
  name = "inactive bot user from API"
}`
)

func init() {
	resource.AddTestSweepers("bot_users", &resource.Sweeper{
		Name: "bot_users",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.BotUsersList(ctx, &restapi.Paging{})
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.BotUsers {
				// Assume items with empty Description or Metadata are system defined
				// (i.e. API Keys) so do not sweep them for cleanup.
				// However, not all items have Description and Metadata fields, so need to reflect.
				iv := reflect.ValueOf(item)
				dv := iv.FieldByName("Description")
				mv := iv.FieldByName("Metadata")
				shouldKeep := (dv.IsValid() && dv.IsZero()) || (mv.IsValid() && mv.IsZero())
				if !shouldKeep {
					_, _, err := conn.BotUsersDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}

			return nil
		},
	})
}

func TestAccResourceBotUsers(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyBotUsers,
		Steps: []resource.TestStep{
			{
				Config: resourceBotUsers_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateBotUsers,
				// ),
			},
		},
	})
}

func testAccCheckDestroyBotUsers(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateBotUsers(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
