package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
)

// stringPtrFromFramework converts a Framework types.String to a *string for
// optional API fields. Returns nil for null/unknown (field omitted from API
// request), *"" for empty, and *"value" for set values.
func stringPtrFromFramework(v types.String) *string {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	s := v.ValueString()
	return &s
}

// stringFromPtr converts a *string to a string, returning "" for nil.
func stringFromPtr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// boolPtrFromFramework converts a Framework types.Bool to a *bool for
// optional API fields. Returns nil for null/unknown.
func boolPtrFromFramework(v types.Bool) *bool {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	b := v.ValueBool()
	return &b
}

// flattenRef extracts the ID from an ngrok.Ref, returning "" if nil.
func flattenRef(ref *ngrok.Ref) string {
	if ref == nil {
		return ""
	}
	return ref.ID
}

// flattenStringList converts a slice of strings to a types.List of StringType.
func flattenStringList(vals []string) []types.String {
	if vals == nil {
		return nil
	}
	result := make([]types.String, len(vals))
	for i, v := range vals {
		result[i] = types.StringValue(v)
	}
	return result
}

// expandStringList converts a slice of types.String to a slice of strings.
func expandStringList(vals []types.String) []string {
	if vals == nil {
		return nil
	}
	result := make([]string, len(vals))
	for i, v := range vals {
		result[i] = v.ValueString()
	}
	return result
}

// flattenRefList converts a slice of ngrok.Ref to a slice of types.String (IDs).
func flattenRefList(refs []ngrok.Ref) []types.String {
	if refs == nil {
		return nil
	}
	result := make([]types.String, len(refs))
	for i, r := range refs {
		result[i] = types.StringValue(r.ID)
	}
	return result
}

// stringListEqual compares two []types.String slices for equality.
func stringListEqual(a, b []types.String) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !a[i].Equal(b[i]) {
			return false
		}
	}
	return true
}

// preserveSensitive returns the prior value for a sensitive field when the API
// returns it redacted (empty or nil). This prevents inconsistent-result errors
// for fields like api_key, client_secret, aws_secret_access_key.
func preserveSensitive(apiValue, priorValue types.String) types.String {
	if !priorValue.IsNull() && !priorValue.IsUnknown() && apiValue.ValueString() == "" {
		return priorValue
	}
	return apiValue
}

// stringValueFromPtr converts a *string to types.String, returning null for nil
// pointers. Use this for Optional fields to preserve null vs empty distinction.
func stringValueFromPtr(s *string) types.String {
	if s == nil {
		return types.StringNull()
	}
	return types.StringValue(*s)
}

// isNotFound checks if an error from the ngrok API is a 404 Not Found.
func isNotFound(err error) bool {
	return ngrok.IsNotFound(err)
}
