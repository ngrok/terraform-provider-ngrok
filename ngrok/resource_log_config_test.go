// Code generated by apic. DO NOT EDIT.

package ngrok

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var (
	resourceLogConfigs_createConfig = `resource "ngrok_log_config" "example" {
  description = "low sampling, basic HTTP logs"
  destination_ids = [ "ld_1kFGlQM7h9b1nW5VX28F5veBM2j" ]
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
