// +build github

package transform

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DiffSuppressWhitespace(k, old, new string, _ *schema.ResourceData) bool {
	return strings.TrimSpace(old) == strings.TrimSpace(new)
}

func DiffSuppressCase(k string, old string, new string, _ *schema.ResourceData) bool {
	return old == new || strings.EqualFold(old, new)
}
