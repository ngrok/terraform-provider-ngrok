// Code generated by apic. DO NOT EDIT.

package ngrok

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var (
	resourceReservedAddrs_createConfig = `resource "ngrok_reserved_addr" "example" {
  description = "SSH for device #001"
  region = "us"
}`
	resourceReservedAddrs_updateConfig = `resource "ngrok_reserved_addr" "example" {
  endpoint_configuration_id = "ec_1kOTodlMNbHwofMx4ZaEQbZlMDo"
  metadata = "{\"proto\": \"ssh\"}"
}`
)

func TestAccResourceReservedAddrs(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyReservedAddrs,
		Steps: []resource.TestStep{
			{
				Config: resourceReservedAddrs_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateReservedAddrs,
				// ),
			},
		},
	})
}

func testAccCheckDestroyReservedAddrs(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateReservedAddrs(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
