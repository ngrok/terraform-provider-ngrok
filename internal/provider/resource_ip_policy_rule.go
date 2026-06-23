package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/ip_policy_rules"
	"github.com/ngrok/terraform-provider-ngrok/internal/resource_ip_policy_rule"
)

var (
	_ resource.Resource                = &ipPolicyRuleResource{}
	_ resource.ResourceWithImportState = &ipPolicyRuleResource{}
)

type ipPolicyRuleResourceModel struct {
	ID          types.String `tfsdk:"id"`
	URI         types.String `tfsdk:"uri"`
	CreatedAt   types.String `tfsdk:"created_at"`
	Description types.String `tfsdk:"description"`
	Metadata    types.String `tfsdk:"metadata"`
	CIDR        types.String `tfsdk:"cidr"`
	IPPolicyID  types.String `tfsdk:"ip_policy_id"`
	Action      types.String `tfsdk:"action"`
}

type ipPolicyRuleResource struct {
	client *ip_policy_rules.Client
}

func NewIPPolicyRuleResource() resource.Resource {
	return &ipPolicyRuleResource{}
}

func (r *ipPolicyRuleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_policy_rule"
}

func (r *ipPolicyRuleResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	s := resource_ip_policy_rule.IpPolicyRuleResourceSchema(ctx)
	attrs := s.Attributes

	// Delete Ref nested object (hand-written uses flat ip_policy_id instead)
	delete(attrs, "ip_policy")

	// Plan modifiers
	addStringPlanModifiers(attrs, "id", useStateForUnknownString())
	addStringPlanModifiers(attrs, "uri", useStateForUnknownString())
	addStringPlanModifiers(attrs, "created_at", useStateForUnknownString())
	addStringPlanModifiers(attrs, "description", useStateForUnknownString())
	addStringPlanModifiers(attrs, "metadata", useStateForUnknownString())
	addStringPlanModifiers(attrs, "ip_policy_id", requiresReplaceString())
	addStringPlanModifiers(attrs, "action", requiresReplaceString())

	resp.Schema = s
}

func (r *ipPolicyRuleResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clientConfig, ok := req.ProviderData.(*ngrok.ClientConfig)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *ngrok.ClientConfig, got: %T.", req.ProviderData),
		)
		return
	}
	r.client = ip_policy_rules.NewClient(clientConfig)
}

func (r *ipPolicyRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ipPolicyRuleResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	rule, err := r.client.Create(ctx, &ngrok.IPPolicyRuleCreate{
		Description: plan.Description.ValueString(),
		Metadata:    plan.Metadata.ValueString(),
		CIDR:        plan.CIDR.ValueString(),
		IPPolicyID:  plan.IPPolicyID.ValueString(),
		Action:      stringPtrFromFramework(plan.Action),
	})
	if err != nil {
		resp.Diagnostics.AddError("Error creating IP policy rule", err.Error())
		return
	}

	flattenIPPolicyRule(rule, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ipPolicyRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ipPolicyRuleResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	rule, err := r.client.Get(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading IP policy rule", err.Error())
		return
	}

	flattenIPPolicyRule(rule, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ipPolicyRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ipPolicyRuleResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state ipPolicyRuleResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	rule, err := r.client.Update(ctx, &ngrok.IPPolicyRuleUpdate{
		ID:          state.ID.ValueString(),
		Description: stringPtrFromFramework(plan.Description),
		Metadata:    stringPtrFromFramework(plan.Metadata),
		CIDR:        stringPtrFromFramework(plan.CIDR),
	})
	if err != nil {
		resp.Diagnostics.AddError("Error updating IP policy rule", err.Error())
		return
	}

	flattenIPPolicyRule(rule, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ipPolicyRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ipPolicyRuleResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			return
		}
		resp.Diagnostics.AddError("Error deleting IP policy rule", err.Error())
	}
}

func (r *ipPolicyRuleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func flattenIPPolicyRule(rule *ngrok.IPPolicyRule, model *ipPolicyRuleResourceModel) {
	model.ID = types.StringValue(rule.ID)
	model.URI = types.StringValue(rule.URI)
	model.CreatedAt = types.StringValue(rule.CreatedAt)
	model.Description = types.StringValue(rule.Description)
	model.Metadata = types.StringValue(rule.Metadata)
	model.CIDR = types.StringValue(rule.CIDR)
	model.IPPolicyID = types.StringValue(rule.IPPolicy.ID)
	model.Action = types.StringValue(rule.Action)
}
