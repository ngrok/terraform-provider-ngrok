package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceReservedDomain_basic(t *testing.T) {
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	domainName := fmt.Sprintf("%s.ngrok.io", rName)
	resourceName := "ngrok_reserved_domain.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccReservedDomainConfig(domainName, "test domain", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "domain", domainName),
					resource.TestCheckResourceAttr(resourceName, "description", "test domain"),
					resource.TestCheckResourceAttrSet(resourceName, "uri"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
				),
			},
			// ImportState
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update description and metadata
			{
				Config: testAccReservedDomainConfig(domainName, "updated description", `{"env":"staging"}`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "updated description"),
					resource.TestCheckResourceAttr(resourceName, "metadata", `{"env":"staging"}`),
				),
			},
		},
	})
}

func testAccReservedDomainConfig(domain, description, metadata string) string {
	metadataAttr := ""
	if metadata != "" {
		metadataAttr = fmt.Sprintf(`  metadata    = %q`, metadata)
	}
	return fmt.Sprintf(`
resource "ngrok_reserved_domain" "test" {
  domain      = %q
  description = %q
%s
}
`, domain, description, metadataAttr)
}
