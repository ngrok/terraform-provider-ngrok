// Code generated by apic. DO NOT EDIT.

package ngrok

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var (
	resourceIPRestrictions_createConfig = `resource "ngrok_ip_restriction" "example" {
  ip_policy_ids = [ "ipp_1k9iKpPOlwEaNRix43oB5WLs7FG" ]
  type = "dashboard"
}`
	resourceIPRestrictions_updateConfig = `resource "ngrok_ip_restriction" "example" {
  ip_policy_ids = [ "ipp_1k9iKpPOlwEaNRix43oB5WLs7FG", "ipp_1k9iKstRLrifshU5FJIEw8zwsGu" ]
}`
)

func TestAccResourceIPRestrictions(t *testing.T) {
	t.Skip("Test skipped. See: https://github.com/ngrok-private/ngrok/issues/4716")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyIPRestrictions,
		Steps: []resource.TestStep{
			{
				Config: resourceIPRestrictions_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateIPRestrictions,
				// ),
			},
		},
	})
}

func testAccCheckDestroyIPRestrictions(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateIPRestrictions(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
