// Code generated by apic. DO NOT EDIT.

package ngrok

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var (
	resourceReservedDomains_createConfig = `resource "ngrok_reserved_domain" "example" {
  certificate_id = "cert_1kFGd3vIF2bugl1YonNmAoIlTxh"
  name = "myapp.mydomain.com"
  region = "us"
}`
	resourceReservedDomains_updateConfig = `resource "ngrok_reserved_domain" "example" {
  certificate_management_policy {
    authority = "letsencrypt"
  }
  description = "point-of-sale new york #302"
  http_endpoint_configuration_id = "ec_1kFGdFkvB2GVkfmhpWnj7LEx3l2"
  https_endpoint_configuration_id = "ec_1kFGdMWSuQ3OfmTLMaZgdIUIYLr"
  metadata = "{env: \"staging\", \"connector_id\":\"64698fcc-5f5c-4b63-910e-8669d04bd943\"}"
}`
)

func TestAccResourceReservedDomains(t *testing.T) {
	t.Skip("Test skipped. See: https://github.com/ngrok-private/ngrok/issues/4718")
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
