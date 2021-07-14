package transform

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DiffSuppressWhitespace(_, old, new string, _ *schema.ResourceData) bool {
	return strings.TrimSpace(old) == strings.TrimSpace(new)
}

func DiffSuppressCase(_, old, new string, _ *schema.ResourceData) bool {
	return old == new || strings.EqualFold(old, new)
}

func DiffSuppressAutoCertId(key, _, _ string, d *schema.ResourceData) bool {
	if key == "certificate_id" {
		// When a certificate management policy exists certificates are automatically managed.
		// The ID will change _after_ the resource is initially created and a third-party issues a certificate.
		// To avoid Terraform considering the result an update the diff is suppressed.
		_, ok := d.GetOk("certificate_management_policy")
		return ok
	}

	return false
}
