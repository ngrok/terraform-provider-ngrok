package ngrok

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var (
	testAccProviders map[string]*schema.Provider
	testAccProvider  *schema.Provider
)

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"ngrok": testAccProvider,
	}
}

func testAccPreCheck(t *testing.T) {
	for _, variable := range []string{"NGROK_API_KEY", "NGROK_API_BASE_URL"} {
		if v := os.Getenv(variable); v == "" {
			t.Fatalf("%s environment variable must be set for acceptance tests", variable)
		}
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatal(err)
	}
}

func TestProvider_impl(t *testing.T) {
	var (
		_ *schema.Provider = Provider()
	)
}

// helper functions

func instanceState(s *terraform.State, name string) (is *terraform.InstanceState, err error) {
	ms := s.RootModule()
	rs, ok := ms.Resources[name]
	if !ok {
		err = fmt.Errorf("Resource not found: %s in %s", name, ms.Path)
	} else if is = rs.Primary; is == nil {
		err = fmt.Errorf("Resource has no primary instance: %s in %s", name, ms.Path)
	}
	return is, err
}
