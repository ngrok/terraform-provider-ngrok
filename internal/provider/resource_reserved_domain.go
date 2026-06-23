package provider

import (
	"context"
	"fmt"

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
	"github.com/ngrok/ngrok-api-go/v9/reserved_domains"
	"github.com/ngrok/terraform-provider-ngrok/internal/resource_reserved_domain"
)

var (
	_ resource.Resource                = &reservedDomainResource{}
	_ resource.ResourceWithImportState = &reservedDomainResource{}
)

type reservedDomainResourceModel struct {
	ID                          types.String `tfsdk:"id"`
	Domain                      types.String `tfsdk:"domain"`
	Region                      types.String `tfsdk:"region"`
	Description                 types.String `tfsdk:"description"`
	Metadata                    types.String `tfsdk:"metadata"`
	CertificateID               types.String `tfsdk:"certificate_id"`
	CertificateManagementPolicy types.Object `tfsdk:"certificate_management_policy"`
	CNAMETarget                 types.String `tfsdk:"cname_target"`
	ACMEChallengeCNAMETarget    types.String `tfsdk:"acme_challenge_cname_target"`
	ResolvesTo                  types.List   `tfsdk:"resolves_to"`
	URI                         types.String `tfsdk:"uri"`
	CreatedAt                   types.String `tfsdk:"created_at"`
}

type reservedDomainResource struct {
	client *reserved_domains.Client
}

func NewReservedDomainResource() resource.Resource {
	return &reservedDomainResource{}
}

func (r *reservedDomainResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_reserved_domain"
}

func (r *reservedDomainResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	s := resource_reserved_domain.ReservedDomainResourceSchema(ctx)
	attrs := s.Attributes

	// Remove Ref nested objects not in hand-written model
	delete(attrs, "certificate")
	delete(attrs, "certificate_management_status")

	// Replace generated CustomType resolves_to with flat list of strings
	attrs["resolves_to"] = schema.ListAttribute{
		Description: "A list of ngrok point-of-presence shortcodes (or \"global\") that the domain resolves to.",
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
	}
	addListPlanModifiers(attrs, "resolves_to", useStateForUnknownList())

	// Replace generated CustomType certificate_management_policy with standard SingleNestedAttribute
	attrs["certificate_management_policy"] = schema.SingleNestedAttribute{
		Description: "Configuration for automatic management of TLS certificates for this domain, or null if automatic management is disabled. Mutually exclusive with certificate_id.",
		Optional:    true,
		Computed:    true,
		Attributes: map[string]schema.Attribute{
			"authority": schema.StringAttribute{
				Description: "Certificate authority to request certificates from. The only supported value is letsencrypt.",
				Optional:    true,
				Computed:    true,
			},
			"private_key_type": schema.StringAttribute{
				Description: "Type of private key to use when requesting certificates. Defaults to ecdsa, can be either rsa or ecdsa.",
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
	addStringPlanModifiers(attrs, "region", useStateForUnknownString())
	setDeprecated(attrs, "region", "This field is deprecated and will be removed in a future version.")
	addStringPlanModifiers(attrs, "description", useStateForUnknownString())
	addStringPlanModifiers(attrs, "metadata", useStateForUnknownString())
	addStringPlanModifiers(attrs, "certificate_id", useStateForUnknownString())
	addStringPlanModifiers(attrs, "cname_target", useStateForUnknownString())
	addStringPlanModifiers(attrs, "acme_challenge_cname_target", useStateForUnknownString())
	addStringPlanModifiers(attrs, "uri", useStateForUnknownString())
	addStringPlanModifiers(attrs, "created_at", useStateForUnknownString())

	resp.Schema = s
}

func (r *reservedDomainResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = reserved_domains.NewClient(clientConfig)
}

func (r *reservedDomainResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan reservedDomainResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq := &ngrok.ReservedDomainCreate{
		Domain:      plan.Domain.ValueString(),
		Description: plan.Description.ValueString(),
		Metadata:    plan.Metadata.ValueString(),
	}
	if !plan.Region.IsNull() && !plan.Region.IsUnknown() {
		createReq.Region = plan.Region.ValueString()
	}
	createReq.CertificateID = stringPtrFromFramework(plan.CertificateID)
	createReq.CertificateManagementPolicy = expandCertPolicy(ctx, plan.CertificateManagementPolicy, &resp.Diagnostics)
	createReq.ResolvesTo = expandResolvesTo(ctx, plan.ResolvesTo)
	if resp.Diagnostics.HasError() {
		return
	}

	domain, err := r.client.Create(ctx, createReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating reserved domain", err.Error())
		return
	}

	flattenReservedDomain(ctx, domain, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *reservedDomainResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state reservedDomainResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	domain, err := r.client.Get(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading reserved domain", err.Error())
		return
	}

	flattenReservedDomain(ctx, domain, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *reservedDomainResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan reservedDomainResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state reservedDomainResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq := &ngrok.ReservedDomainUpdate{
		ID:          state.ID.ValueString(),
		Description: stringPtrFromFramework(plan.Description),
		Metadata:    stringPtrFromFramework(plan.Metadata),
	}
	updateReq.CertificateID = stringPtrFromFramework(plan.CertificateID)
	updateReq.CertificateManagementPolicy = expandCertPolicy(ctx, plan.CertificateManagementPolicy, &resp.Diagnostics)
	updateReq.ResolvesTo = expandResolvesTo(ctx, plan.ResolvesTo)
	if resp.Diagnostics.HasError() {
		return
	}

	domain, err := r.client.Update(ctx, updateReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating reserved domain", err.Error())
		return
	}

	flattenReservedDomain(ctx, domain, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *reservedDomainResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state reservedDomainResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			return
		}
		resp.Diagnostics.AddError("Error deleting reserved domain", err.Error())
	}
}

func (r *reservedDomainResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func flattenReservedDomain(ctx context.Context, domain *ngrok.ReservedDomain, model *reservedDomainResourceModel, diags *diag.Diagnostics) {
	model.ID = types.StringValue(domain.ID)
	model.Domain = types.StringValue(domain.Domain)
	model.Region = types.StringValue(domain.Region)
	model.Description = types.StringValue(domain.Description)
	model.Metadata = types.StringValue(domain.Metadata)
	model.URI = types.StringValue(domain.URI)
	model.CreatedAt = types.StringValue(domain.CreatedAt)

	if domain.CNAMETarget != nil {
		model.CNAMETarget = types.StringValue(*domain.CNAMETarget)
	} else {
		model.CNAMETarget = types.StringNull()
	}

	if domain.ACMEChallengeCNAMETarget != nil {
		model.ACMEChallengeCNAMETarget = types.StringValue(*domain.ACMEChallengeCNAMETarget)
	} else {
		model.ACMEChallengeCNAMETarget = types.StringNull()
	}

	model.CertificateID = types.StringValue(flattenRef(domain.Certificate))

	model.ResolvesTo = flattenResolvesToOrdered(ctx, domain.ResolvesTo, model.ResolvesTo, diags)

	model.CertificateManagementPolicy = flattenCertPolicy(ctx, domain.CertificateManagementPolicy, diags)
}

func flattenCertPolicy(ctx context.Context, policy *ngrok.ReservedDomainCertPolicy, diags *diag.Diagnostics) types.Object {
	if policy == nil {
		return types.ObjectNull(certPolicyAttrTypes())
	}

	obj, d := types.ObjectValueFrom(ctx, certPolicyAttrTypes(), &certPolicyModel{
		Authority:      types.StringValue(policy.Authority),
		PrivateKeyType: types.StringValue(policy.PrivateKeyType),
	})
	diags.Append(d...)
	return obj
}

func expandCertPolicy(ctx context.Context, obj types.Object, diags *diag.Diagnostics) *ngrok.ReservedDomainCertPolicy {
	if obj.IsNull() || obj.IsUnknown() {
		return nil
	}

	var model certPolicyModel
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return &ngrok.ReservedDomainCertPolicy{
		Authority:      model.Authority.ValueString(),
		PrivateKeyType: model.PrivateKeyType.ValueString(),
	}
}

type certPolicyModel struct {
	Authority      types.String `tfsdk:"authority"`
	PrivateKeyType types.String `tfsdk:"private_key_type"`
}

func certPolicyAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"authority":        types.StringType,
		"private_key_type": types.StringType,
	}
}

func expandResolvesTo(ctx context.Context, list types.List) []ngrok.ReservedDomainResolvesToEntry {
	if list.IsNull() || list.IsUnknown() {
		return nil
	}
	var vals []string
	list.ElementsAs(ctx, &vals, false)
	entries := make([]ngrok.ReservedDomainResolvesToEntry, len(vals))
	for i, v := range vals {
		entries[i] = ngrok.ReservedDomainResolvesToEntry{Value: v}
	}
	return entries
}

func flattenResolvesToOrdered(ctx context.Context, entries []ngrok.ReservedDomainResolvesToEntry, prior types.List, diags *diag.Diagnostics) types.List {
	apiList := flattenResolvesTo(ctx, entries, diags)
	if diags.HasError() || prior.IsNull() || prior.IsUnknown() {
		return apiList
	}
	// Reorder API response to match prior (plan/state) order to avoid spurious diffs.
	var priorVals []string
	prior.ElementsAs(ctx, &priorVals, false)
	var apiVals []string
	apiList.ElementsAs(ctx, &apiVals, false)

	apiSet := make(map[string]bool, len(apiVals))
	for _, v := range apiVals {
		apiSet[v] = true
	}
	var result []string
	for _, v := range priorVals {
		if apiSet[v] {
			result = append(result, v)
			delete(apiSet, v)
		}
	}
	for _, v := range apiVals {
		if apiSet[v] {
			result = append(result, v)
		}
	}
	list, d := types.ListValueFrom(ctx, types.StringType, result)
	diags.Append(d...)
	return list
}

func flattenResolvesTo(ctx context.Context, entries []ngrok.ReservedDomainResolvesToEntry, diags *diag.Diagnostics) types.List {
	if entries == nil {
		return types.ListNull(types.StringType)
	}
	vals := make([]string, len(entries))
	for i, e := range entries {
		vals[i] = e.Value
	}
	list, d := types.ListValueFrom(ctx, types.StringType, vals)
	diags.Append(d...)
	return list
}
