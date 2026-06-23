package provider

import (
	"context"
	"fmt"
	"sort"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/agent_ingresses"
	"github.com/ngrok/terraform-provider-ngrok/internal/resource_agent_ingress"
)

var (
	_ resource.Resource                = &agentIngressResource{}
	_ resource.ResourceWithImportState = &agentIngressResource{}
)

type agentIngressResourceModel struct {
	ID                          types.String `tfsdk:"id"`
	Domain                      types.String `tfsdk:"domain"`
	Description                 types.String `tfsdk:"description"`
	Metadata                    types.String `tfsdk:"metadata"`
	NSTargets                   types.List   `tfsdk:"ns_targets"`
	RegionDomains               types.List   `tfsdk:"region_domains"`
	CertificateManagementPolicy types.Object `tfsdk:"certificate_management_policy"`
	URI                         types.String `tfsdk:"uri"`
	CreatedAt                   types.String `tfsdk:"created_at"`
}

type agentIngressResource struct {
	client *agent_ingresses.Client
}

func NewAgentIngressResource() resource.Resource {
	return &agentIngressResource{}
}

func (r *agentIngressResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_agent_ingress"
}

func (r *agentIngressResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	s := resource_agent_ingress.AgentIngressResourceSchema(ctx)
	attrs := s.Attributes

	// Remove fields not in hand-written model
	delete(attrs, "certificate_management_status")

	// Replace generated CustomType certificate_management_policy with standard SingleNestedAttribute
	attrs["certificate_management_policy"] = schema.SingleNestedAttribute{
		Description: "Configuration for automatic management of TLS certificates for this domain.",
		Optional:    true,
		Computed:    true,
		Attributes: map[string]schema.Attribute{
			"authority": schema.StringAttribute{
				Description: "Certificate authority to request certificates from.",
				Optional:    true,
				Computed:    true,
			},
			"private_key_type": schema.StringAttribute{
				Description: "Type of private key to use when requesting certificates.",
				Optional:    true,
				Computed:    true,
			},
		},
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
	}

	addStringPlanModifiers(attrs, "id", useStateForUnknownString())
	addStringPlanModifiers(attrs, "domain", requiresReplaceString())
	addStringPlanModifiers(attrs, "description", useStateForUnknownString())
	addStringPlanModifiers(attrs, "metadata", useStateForUnknownString())
	addListPlanModifiers(attrs, "ns_targets", useStateForUnknownList())
	addListPlanModifiers(attrs, "region_domains", useStateForUnknownList())
	addStringPlanModifiers(attrs, "uri", useStateForUnknownString())
	addStringPlanModifiers(attrs, "created_at", useStateForUnknownString())

	resp.Schema = s
}

func (r *agentIngressResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = agent_ingresses.NewClient(clientConfig)
}

func (r *agentIngressResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan agentIngressResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq := &ngrok.AgentIngressCreate{
		Domain:      plan.Domain.ValueString(),
		Description: plan.Description.ValueString(),
		Metadata:    plan.Metadata.ValueString(),
	}
	createReq.CertificateManagementPolicy = expandAgentIngressCertPolicy(ctx, plan.CertificateManagementPolicy, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	ingress, err := r.client.Create(ctx, createReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating agent ingress", err.Error())
		return
	}

	flattenAgentIngress(ctx, ingress, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *agentIngressResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state agentIngressResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ingress, err := r.client.Get(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading agent ingress", err.Error())
		return
	}

	flattenAgentIngress(ctx, ingress, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *agentIngressResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan agentIngressResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state agentIngressResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq := &ngrok.AgentIngressUpdate{
		ID:          state.ID.ValueString(),
		Description: stringPtrFromFramework(plan.Description),
		Metadata:    stringPtrFromFramework(plan.Metadata),
	}
	updateReq.CertificateManagementPolicy = expandAgentIngressCertPolicy(ctx, plan.CertificateManagementPolicy, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	ingress, err := r.client.Update(ctx, updateReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating agent ingress", err.Error())
		return
	}

	flattenAgentIngress(ctx, ingress, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *agentIngressResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state agentIngressResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			return
		}
		resp.Diagnostics.AddError("Error deleting agent ingress", err.Error())
	}
}

func (r *agentIngressResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func flattenAgentIngress(ctx context.Context, ingress *ngrok.AgentIngress, model *agentIngressResourceModel, diags *diag.Diagnostics) {
	model.ID = types.StringValue(ingress.ID)
	model.Domain = types.StringValue(ingress.Domain)
	model.Description = types.StringValue(ingress.Description)
	model.Metadata = types.StringValue(ingress.Metadata)
	model.URI = types.StringValue(ingress.URI)
	model.CreatedAt = types.StringValue(ingress.CreatedAt)

	// Sort lists to ensure deterministic ordering (API may return in any order)
	sortedNS := make([]string, len(ingress.NSTargets))
	copy(sortedNS, ingress.NSTargets)
	sort.Strings(sortedNS)
	nsTargets, d := types.ListValueFrom(ctx, types.StringType, sortedNS)
	diags.Append(d...)
	model.NSTargets = nsTargets

	sortedRegions := make([]string, len(ingress.RegionDomains))
	copy(sortedRegions, ingress.RegionDomains)
	sort.Strings(sortedRegions)
	regionDomains, d := types.ListValueFrom(ctx, types.StringType, sortedRegions)
	diags.Append(d...)
	model.RegionDomains = regionDomains

	// Only populate cert policy if user configured it or it was previously in state
	if !model.CertificateManagementPolicy.IsNull() && !model.CertificateManagementPolicy.IsUnknown() {
		model.CertificateManagementPolicy = flattenAgentIngressCertPolicy(ctx, ingress.CertificateManagementPolicy, diags)
	} else if model.CertificateManagementPolicy.IsUnknown() {
		model.CertificateManagementPolicy = types.ObjectNull(agentIngressCertPolicyAttrTypes())
	}
}

func flattenAgentIngressCertPolicy(ctx context.Context, policy *ngrok.AgentIngressCertPolicy, diags *diag.Diagnostics) types.Object {
	if policy == nil {
		return types.ObjectNull(agentIngressCertPolicyAttrTypes())
	}

	obj, d := types.ObjectValueFrom(ctx, agentIngressCertPolicyAttrTypes(), &agentIngressCertPolicyModel{
		Authority:      types.StringValue(policy.Authority),
		PrivateKeyType: types.StringValue(policy.PrivateKeyType),
	})
	diags.Append(d...)
	return obj
}

func expandAgentIngressCertPolicy(ctx context.Context, obj types.Object, diags *diag.Diagnostics) *ngrok.AgentIngressCertPolicy {
	if obj.IsNull() || obj.IsUnknown() {
		return nil
	}

	var model agentIngressCertPolicyModel
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return &ngrok.AgentIngressCertPolicy{
		Authority:      model.Authority.ValueString(),
		PrivateKeyType: model.PrivateKeyType.ValueString(),
	}
}

type agentIngressCertPolicyModel struct {
	Authority      types.String `tfsdk:"authority"`
	PrivateKeyType types.String `tfsdk:"private_key_type"`
}

func agentIngressCertPolicyAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"authority":        types.StringType,
		"private_key_type": types.StringType,
	}
}
