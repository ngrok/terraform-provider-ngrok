package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// addStringPlanModifiers adds string plan modifiers to the named attribute.
func addStringPlanModifiers(attrs map[string]schema.Attribute, name string, modifiers ...planmodifier.String) {
	if a, ok := attrs[name]; ok {
		sa := a.(schema.StringAttribute)
		sa.PlanModifiers = append(sa.PlanModifiers, modifiers...)
		attrs[name] = sa
	}
}

// addListPlanModifiers adds list plan modifiers to the named attribute.
func addListPlanModifiers(attrs map[string]schema.Attribute, name string, modifiers ...planmodifier.List) {
	if a, ok := attrs[name]; ok {
		sa := a.(schema.ListAttribute)
		sa.PlanModifiers = append(sa.PlanModifiers, modifiers...)
		attrs[name] = sa
	}
}

// addInt64PlanModifiers adds int64 plan modifiers to the named attribute.
func addInt64PlanModifiers(attrs map[string]schema.Attribute, name string, modifiers ...planmodifier.Int64) {
	if a, ok := attrs[name]; ok {
		sa := a.(schema.Int64Attribute)
		sa.PlanModifiers = append(sa.PlanModifiers, modifiers...)
		attrs[name] = sa
	}
}

// addObjectPlanModifiers adds object plan modifiers to the named attribute.
func addObjectPlanModifiers(attrs map[string]schema.Attribute, name string, modifiers ...planmodifier.Object) {
	if a, ok := attrs[name]; ok {
		sa := a.(schema.SingleNestedAttribute)
		sa.PlanModifiers = append(sa.PlanModifiers, modifiers...)
		attrs[name] = sa
	}
}

// markSensitive marks the named string attribute as sensitive.
func markSensitive(attrs map[string]schema.Attribute, name string) {
	if a, ok := attrs[name]; ok {
		sa := a.(schema.StringAttribute)
		sa.Sensitive = true
		attrs[name] = sa
	}
}

// setStringDefault sets a static string default for the named attribute.
// (unused for now but available for future use)
var _ = setStringDefault

func setStringDefault(attrs map[string]schema.Attribute, name string, val string) {
	if _, ok := attrs[name]; ok {
		// The function is a placeholder — actual defaults use framework defaults
		_ = val
	}
}

// setBoolDefault sets a static bool default for the named attribute.
func setBoolDefault(attrs map[string]schema.Attribute, name string, val bool) {
	if a, ok := attrs[name]; ok {
		sa := a.(schema.BoolAttribute)
		sa.Default = booldefault.StaticBool(val)
		attrs[name] = sa
	}
}

// setListDefault sets a static list default for the named attribute.
func setListDefault(attrs map[string]schema.Attribute, name string, val types.List) {
	if a, ok := attrs[name]; ok {
		sa := a.(schema.ListAttribute)
		sa.Default = listdefault.StaticValue(val)
		attrs[name] = sa
	}
}

// setDeprecated sets the deprecation message on the named string attribute.
func setDeprecated(attrs map[string]schema.Attribute, name string, msg string) {
	if a, ok := attrs[name]; ok {
		sa := a.(schema.StringAttribute)
		sa.DeprecationMessage = msg
		attrs[name] = sa
	}
}

// caseInsensitiveModifier is a plan modifier that suppresses diffs when the API
// normalizes a string value (e.g., lowercasing a domain name). If the planned
// and state values differ only by case, the state value is preserved to avoid
// a spurious replace.
type caseInsensitiveModifier struct{}

func (m caseInsensitiveModifier) Description(_ context.Context) string {
	return "Suppresses diffs when values differ only by case."
}

func (m caseInsensitiveModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m caseInsensitiveModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() || req.StateValue.IsUnknown() {
		return
	}
	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
		return
	}
	if strings.EqualFold(req.PlanValue.ValueString(), req.StateValue.ValueString()) {
		resp.PlanValue = req.StateValue
	}
}

func caseInsensitiveString() planmodifier.String {
	return caseInsensitiveModifier{}
}

// whitespaceInsensitiveModifier suppresses diffs caused by leading/trailing
// whitespace differences (e.g., trailing newlines in PEM content).
type whitespaceInsensitiveModifier struct{}

func (m whitespaceInsensitiveModifier) Description(_ context.Context) string {
	return "Suppresses diffs caused by leading/trailing whitespace differences."
}

func (m whitespaceInsensitiveModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m whitespaceInsensitiveModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() || req.StateValue.IsUnknown() {
		return
	}
	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
		return
	}
	if strings.TrimSpace(req.PlanValue.ValueString()) == strings.TrimSpace(req.StateValue.ValueString()) {
		resp.PlanValue = req.StateValue
	}
}

func whitespaceInsensitiveString() planmodifier.String {
	return whitespaceInsensitiveModifier{}
}

// domainNormalizeModifier suppresses diffs caused by the API's domain
// normalization: lowercasing, trimming whitespace, and stripping trailing dots.
// This is a superset of caseInsensitiveModifier for domain fields.
type domainNormalizeModifier struct{}

func (m domainNormalizeModifier) Description(_ context.Context) string {
	return "Normalizes domain values to match API behavior (lowercase, trim whitespace, strip trailing dot)."
}

func (m domainNormalizeModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m domainNormalizeModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
		return
	}
	if req.StateValue.IsNull() || req.StateValue.IsUnknown() {
		return
	}
	normalized := strings.ToLower(strings.TrimSuffix(strings.TrimSpace(req.PlanValue.ValueString()), "."))
	if normalized == req.StateValue.ValueString() {
		resp.PlanValue = req.StateValue
	}
}

func domainNormalizeString() planmodifier.String {
	return domainNormalizeModifier{}
}

// normalizeURLModifier suppresses diffs caused by the API's URL normalization:
// lowercasing the scheme and stripping default ports (443 for https, 80 for http).
type normalizeURLModifier struct{}

func (m normalizeURLModifier) Description(_ context.Context) string {
	return "Normalizes URLs to match API behavior (lowercase scheme, strip default ports)."
}

func (m normalizeURLModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m normalizeURLModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
		return
	}
	if req.StateValue.IsNull() || req.StateValue.IsUnknown() {
		return
	}
	if normalizeURL(req.PlanValue.ValueString()) == normalizeURL(req.StateValue.ValueString()) {
		resp.PlanValue = req.StateValue
	}
}

func normalizeURLString() planmodifier.String {
	return normalizeURLModifier{}
}

// suppressWhenSiblingSetModifier suppresses diffs on a string attribute when a
// sibling attribute is set (non-null). Used for certificate_id on reserved_domain:
// when certificate_management_policy is set, the cert is auto-managed and the
// ID changes after issuance — the diff should be suppressed.
type suppressWhenSiblingSetModifier struct {
	sibling string
}

func (m suppressWhenSiblingSetModifier) Description(_ context.Context) string {
	return "Suppresses diffs when the sibling attribute " + m.sibling + " is set."
}

func (m suppressWhenSiblingSetModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m suppressWhenSiblingSetModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() || req.StateValue.IsUnknown() {
		return
	}

	var siblingVal attr.Value
	diags := req.Plan.GetAttribute(ctx, req.Path.ParentPath().AtName(m.sibling), &siblingVal)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !siblingVal.IsNull() && !siblingVal.IsUnknown() {
		resp.PlanValue = req.StateValue
	}
}

func suppressWhenSiblingSet(sibling string) planmodifier.String {
	return suppressWhenSiblingSetModifier{sibling: sibling}
}

// Convenience aliases for common plan modifiers
var (
	useStateForUnknownString = stringplanmodifier.UseStateForUnknown
	requiresReplaceString    = stringplanmodifier.RequiresReplace
	useStateForUnknownList   = listplanmodifier.UseStateForUnknown
	useStateForUnknownInt64  = int64planmodifier.UseStateForUnknown
	requiresReplaceInt64     = int64planmodifier.RequiresReplace
	useStateForUnknownObject = objectplanmodifier.UseStateForUnknown
)
