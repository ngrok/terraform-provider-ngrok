package ngrok

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ngrok/terraform-provider-ngrok/restapi"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

// sharedClientForRegion returns a common provider client configured for the specified region
func sharedClientForRegion(region string) (interface{}, error) {
	cfg := restapi.ClientConfig{
		APIKey:  os.Getenv("NGROK_API_KEY"),
		BaseURL: os.Getenv("NGROK_API_BASE_URL"),
	}

	client, err := restapi.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("error creating client: %s", err)
	}

	return client, nil
}
