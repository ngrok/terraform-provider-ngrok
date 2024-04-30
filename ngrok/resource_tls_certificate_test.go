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
	resourceTLSCertificates_createConfig = `resource "ngrok_tls_certificate" "example" {
  certificate_pem = "-----BEGIN CERTIFICATE-----\nMIIDDTCCAfWgAwIBAgIUBUunDdA4gjgtEbZA8w9Ljhvl3bEwDQYJKoZIhvcNAQEL\nBQAwFjEUMBIGA1UEAwwLZXhhbXBsZS5jb20wHhcNMjAwMzI0MTgxODE5WhcNMjAw\nNDIzMTgxODE5WjAWMRQwEgYDVQQDDAtleGFtcGxlLmNvbTCCASIwDQYJKoZIhvcN\nAQEBBQADggEPADCCAQoCggEBAPKVkkKYNl3d9cqrz4tIFlwsohED5W4y1dcBixy4\nGANFFnw43nc2wPyKwYXumJqJIFrcW/NkUZL07bd+dou6mT6Gh/zbaTW91IkREPXL\n7b3KfVu4XkFosVXpWs0U6o4GrZ81CLiKBWI+H03x/ij5OSiJ1l71pqLeTJLOydAR\nAl8kpp7axeHU4UbDrAZkW5SnuZTjIKwVg0UNsBg1yNfUOu1Uah3BYaqPgQitC0Yg\nLW+NUGu/T91bkD7tLsVInkQXeQGdXBAqOycfJ7wj8OlIpyuXjTnGFA0izVmbQw5f\nrQnZ0geGyhLamvz9Gcd7mIlD/+/AEN9Lht82tAOzKG98/O8CAwEAAaNTMFEwHQYD\nVR0OBBYEFKv6RsvEC6T+zCtJZwB0FCR1sEkhMB8GA1UdIwQYMBaAFKv6RsvEC6T+\nzCtJZwB0FCR1sEkhMA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEB\nAC5fBrouinespo5+9AipjhY/HOKTg+OCnppFnSnqeU1eXZZJ0oakdHTpTNxtbQP9\ntOJTA2f3KWvmpNDMohEQXZz8wHDkdbrIXJKVp6zs1pEp+0BIjA4y9mSywa5xuyk0\noGeChRgGqp2JujDyPCb7LEaKKQEEdMqy73QG+jEAh14+wKixlAf1nATBdeCUvssK\n2x1uZMyqjJFB5y/5EdnWQzD4WJkrsCkxsZHVMN1d+dqf2sf3dTRV8fzsFGOG17NS\n6u2n9iGcFdBA82XN8yeLIWhy1t3GWutG1sdxENbFRRXea+iUqzDsmRtkaBma2GLQ\nd6JTpFbsCtwDjP23UEi7SZo=\n-----END CERTIFICATE-----"
  private_key_pem = "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDylZJCmDZd3fXK\nq8+LSBZcLKIRA+VuMtXXAYscuBgDRRZ8ON53NsD8isGF7piaiSBa3FvzZFGS9O23\nfnaLupk+hof822k1vdSJERD1y+29yn1buF5BaLFV6VrNFOqOBq2fNQi4igViPh9N\n8f4o+TkoidZe9aai3kySzsnQEQJfJKae2sXh1OFGw6wGZFuUp7mU4yCsFYNFDbAY\nNcjX1DrtVGodwWGqj4EIrQtGIC1vjVBrv0/dW5A+7S7FSJ5EF3kBnVwQKjsnHye8\nI/DpSKcrl405xhQNIs1Zm0MOX60J2dIHhsoS2pr8/RnHe5iJQ//vwBDfS4bfNrQD\nsyhvfPzvAgMBAAECggEBALLv7YE98exvi5zB+0fMFuJK8gkHDLequ93q/4hhqyTO\nU3WyJTdepiAi4fk/NEXZnIopPZJdj2aNUMQnfp43OE7MwYac+hBwRFQOyKnmkSmM\nMcf0SWKKLTUn+piIMzQsbOmhHxuwg6QiGslOFaJ3o9fpRL2rCg3dWDJ6Ypcd1NgE\nK0uy7gg+DwIpU6MeG6lA+HbxbGi+yd2x88Gjn9dGr7FZK34RUDooH60BCX9P8N9X\nT+n10MzzX7ZQOsLfe8FKc1/X8AybI5SYm1GMyfKD4QBt6JG4HKAjPHzBzcIpfN3d\n7BM11Imkrz7LcbUG+F23NVsi6n5IIGT1WqwCRIH2PpECgYEA/SJ5Ra4d0hUS5RYB\nzABquM3sp7JsKxCn7O5PqNLB4TgH9dXtWFhaFVB6juMGyHbvktVH0j4lps/Te0rk\nVU2zU1XxvCTFhtcCYUtNk0cRw6LH8feKiorXHdDRB33t0c47QSD/6AGOjBtxqD7B\n3ZxyR3P+7RdQopLLRFN+FHAnmzsCgYEA9VSGZDFSK+fbg4CgwkWdzuHrAXaUEv0U\novqqWd/yXB9wauEvRHnOrSgW6hFZQiatJOXx0KnalJQzohz/SLGO0MqGtwQbYWVT\nWiJgjUbNeiPEHBeUA6U55lVQr26kQSUWdXEtRbDz+hqV1K+6tTEMzaSPmJiHNgki\nlNMO2gqGQd0CgYBJ268qx5zn2UJEGWG41j5NYbg1TfgFsLxugzI2/heX0TNxZVP1\nPQI7ydmYq2ElSJ6qZxSnoX5255i7FqT8xskV/bOkw83mhAGrxb8Cw+/I90wDq8h+\nl/ggOPdkijfDybq8TBae6SVgd/l3r6f9M1KcypmNMApVBSPN8daNvBOyVQKBgQDo\nsj2utyFrx8Xsm4rf+kxOuPbBMooM4MQ8OmpuSP6G5sMofWLqHmcs0sO5TK9PEYRV\nZU3ST+ml2FSJRdvWRaRi4laZLWoTHZrL+aN/HVM0sMwIoUyhkIy0ruOTIuzlZZpB\n1xHL8qXX6nOHgw8jYdz1CUuyv6owVMXaR77kjer+eQKBgByYZlR/eNTzlot0SdFl\nIbgQ9bV7VLIo+vKzOXE3trfzRJMgUosLTp+5wdSVSW/VBdYZ7Ir3n0bbpY/dGinI\nVShxPbChhCZnhvG2lEEiekI44m5jHSA6hhtRdt/CrhL65Rw2SE5lMEe8htg1UGus\nwzLHWHBl72FjbjdhvEgrq60W\n-----END PRIVATE KEY-----"
}`
	resourceTLSCertificates_updateConfig = `resource "ngrok_tls_certificate" "example" {
  metadata = "{\"example\": true}"
}`
)

func init() {
	resource.AddTestSweepers("tls_certificates", &resource.Sweeper{
		Name: "tls_certificates",
		F: func(region string) error {
			ctx := context.Background()
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*restapi.Client)

			list, _, err := conn.TLSCertificatesList(ctx, &restapi.Paging{})
			if err != nil {
				return fmt.Errorf("Error getting list of items: %s", err)
			}

			for _, item := range list.TLSCertificates {
				// Assume items with empty Description or Metadata are system defined
				// (i.e. API Keys) so do not sweep them for cleanup.
				// However, not all items have Description and Metadata fields, so need to reflect.
				iv := reflect.ValueOf(item)
				dv := iv.FieldByName("Description")
				mv := iv.FieldByName("Metadata")
				shouldKeep := (dv.IsValid() && dv.IsZero()) || (mv.IsValid() && mv.IsZero())
				if !shouldKeep {
					_, _, err := conn.TLSCertificatesDelete(ctx, &restapi.Item{ID: item.ID})

					if err != nil {
						log.Printf("Error destroying id %s during sweep: %s", item.ID, err)
					}
				}
			}

			return nil
		},
	})
}

func TestAccResourceTLSCertificates(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroyTLSCertificates,
		Steps: []resource.TestStep{
			{
				Config: resourceTLSCertificates_createConfig,
				// Check: resource.ComposeAggregateTestCheckFunc(
				// 	testAccCheckCreateTLSCertificates,
				// ),
			},
		},
	})
}

func testAccCheckDestroyTLSCertificates(s *terraform.State) (err error) {
	return err
}

func testAccCheckCreateTLSCertificates(s *terraform.State) (err error) {
	fmt.Sprintf("state=%#v", s)
	return err
}
