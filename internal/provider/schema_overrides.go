package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
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

// addMapPlanModifiers adds map plan modifiers to the named attribute.
func addMapPlanModifiers(attrs map[string]schema.Attribute, name string, modifiers ...planmodifier.Map) {
	if a, ok := attrs[name]; ok {
		sa := a.(schema.MapAttribute)
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

// Convenience aliases for common plan modifiers
var (
	useStateForUnknownString = stringplanmodifier.UseStateForUnknown
	requiresReplaceString    = stringplanmodifier.RequiresReplace
	useStateForUnknownList   = listplanmodifier.UseStateForUnknown
	requiresReplaceList      = listplanmodifier.RequiresReplace
	useStateForUnknownInt64  = int64planmodifier.UseStateForUnknown
	requiresReplaceInt64     = int64planmodifier.RequiresReplace
	requiresReplaceMap       = mapplanmodifier.RequiresReplace
	useStateForUnknownObject = objectplanmodifier.UseStateForUnknown
)
