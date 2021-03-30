package transform

import (
	"fmt"
	"strings"
)

func ValidateLowercaseKeys(i interface{}, key string) (s []string, es []error) {
	m, ok := i.(map[string]interface{})
	if !ok {
		es = append(es, fmt.Errorf("ValidateLowercaseKeys can only be used on map values. Want map[string]interface{}, got %T", i))
	}
	for k := range m {
		if strings.ToLower(k) != k {
			es = append(es, fmt.Errorf("%s: unsupported mixed or upper-case key %q", key, k))
		}
	}
	return
}
