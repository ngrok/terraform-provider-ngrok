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
	resourceIPPolicyRules_createConfig = `resource "ngrok_ip_policy_rule" "example" {
  action = "allow"
  cidr = "212.3.14.0/24"
  description = "nyc office"
  ip_policy_id = "ipp_26rOydjEUNZSLTi8bYXBg278qUT"
}`
	resourceIPPolicyRules_updateConfig = `resource "ngrok_ip_policy_rule" "example" {
  cidr = "212.3.15.0/24"
}`
)

func init() {
	resource.AddTestSweepers("ip_policy_rules", &resource.Sweeper{
		Name: "ip_policy_rules",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.IPPolicyRulesList(ctx, &restapi.Paging{})
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.IPPolicyRules {
				// Assume items with empty Description or Metadata are system defined
				// (i.e. API Keys) so do not sweep them for cleanup.
				// However, not all items have Description and Metadata fields, so need to reflect.
				iv := reflect.ValueOf(item)
				dv := iv.FieldByName("Description")
				mv := iv.FieldByName("Metadata")
				shouldKeep := (dv.IsValid() && dv.IsZero()) || (mv.IsValid() && mv.IsZero())
				if !shouldKeep {
					_, _, err := conn.IPPolicyRulesDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}

			return nil
		},
	})
}

func TestAccResourceIPPolicyRules(t *testing.T) {
	t.Skip("Test skipped.")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyIPPolicyRules,
		Steps: []resource.TestStep{
			{
				Config: resourceIPPolicyRules_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateIPPolicyRules,
				// ),
			},
		},
	})
}

func testAccCheckDestroyIPPolicyRules(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateIPPolicyRules(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
