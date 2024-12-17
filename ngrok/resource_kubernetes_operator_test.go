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
	resourceKubernetesOperators_createConfig = `resource "ngrok_kubernetes_operator" "example" {
  deployment {
    name = "ngrok-operator"
    namespace = "ngrok-operator"
    version = "0.12.2"
  }
  description = "Created by ngrok-operator"
  enabled_features = [ "Ingress", "Bindings" ]
  metadata = "{\"namespace.uid\":\"9663c1aa-10e4-4933-8576-398a49a5caf6\",\"owned-by\":\"ngrok-operator\"}"
  region = "global"
}`
	resourceKubernetesOperators_updateConfig = `resource "ngrok_kubernetes_operator" "example" {
  deployment {
    cluster_name = ""
    name = "ngrok-operator"
    namespace = "ngrok-operator"
    version = "0.12.2"
  }
  description = "Created by ngrok-operator"
  enabled_features = [ "Ingress", "Bindings" ]
  metadata = "{\"namespace.uid\":\"9663c1aa-10e4-4933-8576-398a49a5caf6\",\"owned-by\":\"ngrok-operator\"}"
  region = "global"
}`
)

func init() {
	resource.AddTestSweepers("kubernetes_operators", &resource.Sweeper{
		Name: "kubernetes_operators",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.KubernetesOperatorsList(ctx, &restapi.Paging{})
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.Operators {
				// Assume items with empty Description or Metadata are system defined
				// (i.e. API Keys) so do not sweep them for cleanup.
				// However, not all items have Description and Metadata fields, so need to reflect.
				iv := reflect.ValueOf(item)
				dv := iv.FieldByName("Description")
				mv := iv.FieldByName("Metadata")
				shouldKeep := (dv.IsValid() && dv.IsZero()) || (mv.IsValid() && mv.IsZero())
				if !shouldKeep {
					_, _, err := conn.KubernetesOperatorsDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}

			return nil
		},
	})
}

func TestAccResourceKubernetesOperators(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyKubernetesOperators,
		Steps: []resource.TestStep{
			{
				Config: resourceKubernetesOperators_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateKubernetesOperators,
				// ),
			},
		},
	})
}

func testAccCheckDestroyKubernetesOperators(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateKubernetesOperators(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
