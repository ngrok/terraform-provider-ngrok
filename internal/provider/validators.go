package provider

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = jsonSyntaxValidator{}

type jsonSyntaxValidator struct{}

func (v jsonSyntaxValidator) Description(_ context.Context) string {
	return "value must be valid JSON"
}

func (v jsonSyntaxValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v jsonSyntaxValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()
	if !json.Valid([]byte(value)) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid JSON",
			"The value must be valid JSON syntax.",
		)
	}
}

// JSONSyntax returns a validator that checks if a string is valid JSON.
func JSONSyntax() validator.String {
	return jsonSyntaxValidator{}
}
