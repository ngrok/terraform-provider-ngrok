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
	resourceCertificateAuthorities_createConfig = `resource "ngrok_certificate_authority" "example" {
  ca_pem = "-----BEGIN CERTIFICATE-----\nMIIEETCCAvmgAwIBAgIUU3N6lNzPqar4400cLQMcVHFl+mEwDQYJKoZIhvcNAQEL\nBQAwgZcxCzAJBgNVBAYTAkFVMQwwCgYDVQQIDANOU1cxDzANBgNVBAcMBlN5ZG5l\neTEZMBcGA1UECgwQRHJvcGJlYXIgUHR5IEx0ZDEkMCIGA1UEAwwbSW50cmFuZXQg\nU2VydmljZXMgQXV0aG9yaXR5MSgwJgYJKoZIhvcNAQkBFhlzZWN1cml0eUBkcm9w\nYmVhci5leGFtcGxlMB4XDTIwMDUwMTE2Mjc1OVoXDTIxMDUwMTE2Mjc1OVowgZcx\nCzAJBgNVBAYTAkFVMQwwCgYDVQQIDANOU1cxDzANBgNVBAcMBlN5ZG5leTEZMBcG\nA1UECgwQRHJvcGJlYXIgUHR5IEx0ZDEkMCIGA1UEAwwbSW50cmFuZXQgU2Vydmlj\nZXMgQXV0aG9yaXR5MSgwJgYJKoZIhvcNAQkBFhlzZWN1cml0eUBkcm9wYmVhci5l\neGFtcGxlMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7y/EAN0yZkA0\nnRpMBfomnnS8KMWHb90kvGfhkCDR8WCQz5mX7eDEYDthRQrEgp63qtJ7IoCM5f0A\nUD6J2m/mZecP7SfA8OuTAZ7UyRixpZh0zJQSgj24Sh1LQuYci0DNXrei+R1qBvd+\npmpZwkKygNrbZYe3oY1PZ3jEYPSAQzIObDF7LhdhLLrcfWa9BHOGMLnALNMY558b\nvoijTCEmRrSavdvrAS9LDRipEXT8EQOWZZT9VbPtgSBalvStdoupAptmPIWjXftf\nWi1kry+P0xVFZG9iZwUeAT6fSJ+gJD8M1UXWaQbocYrctESP0sZEFM3rzdWqrZb7\n3cH3K5OCvwIDAQABo1MwUTAdBgNVHQ4EFgQUsZdchgUimRHLiPRWw51+DGBmlfMw\nHwYDVR0jBBgwFoAUsZdchgUimRHLiPRWw51+DGBmlfMwDwYDVR0TAQH/BAUwAwEB\n/zANBgkqhkiG9w0BAQsFAAOCAQEANk25tt8sSfn6Qu1bbhWRbjKgS5z+j9LqyCna\nv3fbSchMthaQR7w0vL69ayroeYdqDZkRMmHjuYKY4NyqyXkkaqVO63wEicCo55d9\npIKuPzc/7xwdRephosjGTQ4QaQ4OnrdpJZieI92m9ODexgsab84AYmwNpbGOI/tK\nnPsQr8x1RfLs2gbBwQ4MYVM3tQQbX0o+yve5nz/NCOq4vdG+eKON5u6VYMkOOg9F\nVyNY1iISQkpNk/AF6Vi9BGuDb5Hg0phEl1Q0ntCO7ZHAUHjy0ucqXZiXoXdXZcs3\n3zKKLUKva59EDBZ5TUucvXh8VemBtNc6hd1mX4Tq7lAreG9pjQ==\n-----END CERTIFICATE-----"
  description = "Internal Coprorates Services Authority"
  metadata = "{\"internal_id\": \"7d2caeee-cdc3-4b26-b2c2-b280b8287552\"}"
}`
	resourceCertificateAuthorities_updateConfig = `resource "ngrok_certificate_authority" "example" {
  description = "Internal Corporate Services Authority (Legacy)"
}`
)

func init() {
	resource.AddTestSweepers("certificate_authorities", &resource.Sweeper{
		Name: "certificate_authorities",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.CertificateAuthoritiesList(ctx, &restapi.Paging{})
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.CertificateAuthorities {
				// Assume items with empty Description or Metadata are system defined
				// (i.e. API Keys) so do not sweep them for cleanup.
				// However, not all items have Description and Metadata fields, so need to reflect.
				iv := reflect.ValueOf(item)
				dv := iv.FieldByName("Description")
				mv := iv.FieldByName("Metadata")
				shouldKeep := (dv.IsValid() && dv.IsZero()) || (mv.IsValid() && mv.IsZero())
				if !shouldKeep {
					_, _, err := conn.CertificateAuthoritiesDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}

			return nil
		},
	})
}

func TestAccResourceCertificateAuthorities(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyCertificateAuthorities,
		Steps: []resource.TestStep{
			{
				Config: resourceCertificateAuthorities_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateCertificateAuthorities,
				// ),
			},
		},
	})
}

func testAccCheckDestroyCertificateAuthorities(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateCertificateAuthorities(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
