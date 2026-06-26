package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/certificate_authorities"
	"github.com/ngrok/terraform-provider-ngrok/internal/resource_certificate_authority"
)

var (
	_ resource.Resource                = &certificateAuthorityResource{}
	_ resource.ResourceWithImportState = &certificateAuthorityResource{}
)

type certificateAuthorityResourceModel struct {
	ID                types.String `tfsdk:"id"`
	URI               types.String `tfsdk:"uri"`
	CreatedAt         types.String `tfsdk:"created_at"`
	Description       types.String `tfsdk:"description"`
	Metadata          types.String `tfsdk:"metadata"`
	CAPEM             types.String `tfsdk:"ca_pem"`
	SubjectCommonName types.String `tfsdk:"subject_common_name"`
	NotBefore         types.String `tfsdk:"not_before"`
	NotAfter          types.String `tfsdk:"not_after"`
	KeyUsages         types.List   `tfsdk:"key_usages"`
	ExtendedKeyUsages types.List   `tfsdk:"extended_key_usages"`
}

type certificateAuthorityResource struct {
	client *certificate_authorities.Client
}

func NewCertificateAuthorityResource() resource.Resource {
	return &certificateAuthorityResource{}
}

func (r *certificateAuthorityResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_certificate_authority"
}

func (r *certificateAuthorityResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	s := resource_certificate_authority.CertificateAuthorityResourceSchema(ctx)
	attrs := s.Attributes

	addStringPlanModifiers(attrs, "id", useStateForUnknownString())
	addStringPlanModifiers(attrs, "uri", useStateForUnknownString())
	addStringPlanModifiers(attrs, "created_at", useStateForUnknownString())
	addStringPlanModifiers(attrs, "description", useStateForUnknownString())
	addStringPlanModifiers(attrs, "metadata", useStateForUnknownString())
	addStringPlanModifiers(attrs, "ca_pem", whitespaceInsensitiveString(), requiresReplaceString())
	addStringPlanModifiers(attrs, "subject_common_name", useStateForUnknownString())
	addStringPlanModifiers(attrs, "not_before", useStateForUnknownString())
	addStringPlanModifiers(attrs, "not_after", useStateForUnknownString())
	addListPlanModifiers(attrs, "key_usages", useStateForUnknownList())
	addListPlanModifiers(attrs, "extended_key_usages", useStateForUnknownList())

	resp.Schema = s
}

func (r *certificateAuthorityResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = certificate_authorities.NewClient(clientConfig)
}

func (r *certificateAuthorityResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan certificateAuthorityResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq := &ngrok.CertificateAuthorityCreate{
		Description: plan.Description.ValueString(),
		Metadata:    plan.Metadata.ValueString(),
		CAPEM:       plan.CAPEM.ValueString(),
	}

	ca, err := r.client.Create(ctx, createReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating certificate authority", err.Error())
		return
	}

	flattenCertificateAuthority(ctx, ca, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *certificateAuthorityResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state certificateAuthorityResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ca, err := r.client.Get(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading certificate authority", err.Error())
		return
	}

	flattenCertificateAuthority(ctx, ca, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *certificateAuthorityResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan certificateAuthorityResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state certificateAuthorityResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq := &ngrok.CertificateAuthorityUpdate{
		ID:          state.ID.ValueString(),
		Description: stringPtrFromFramework(plan.Description),
		Metadata:    stringPtrFromFramework(plan.Metadata),
	}

	ca, err := r.client.Update(ctx, updateReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating certificate authority", err.Error())
		return
	}

	flattenCertificateAuthority(ctx, ca, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *certificateAuthorityResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state certificateAuthorityResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			return
		}
		resp.Diagnostics.AddError("Error deleting certificate authority", err.Error())
	}
}

func (r *certificateAuthorityResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func flattenCertificateAuthority(ctx context.Context, ca *ngrok.CertificateAuthority, model *certificateAuthorityResourceModel, diags *diag.Diagnostics) {
	model.ID = types.StringValue(ca.ID)
	model.URI = types.StringValue(ca.URI)
	model.CreatedAt = types.StringValue(ca.CreatedAt)
	model.Description = types.StringValue(ca.Description)
	model.Metadata = types.StringValue(ca.Metadata)
	model.CAPEM = types.StringValue(ca.CAPEM)
	model.SubjectCommonName = types.StringValue(ca.SubjectCommonName)
	model.NotBefore = types.StringValue(ca.NotBefore)
	model.NotAfter = types.StringValue(ca.NotAfter)
	keyUsages, d := types.ListValueFrom(ctx, types.StringType, ca.KeyUsages)
	diags.Append(d...)
	model.KeyUsages = keyUsages
	extKeyUsages, d := types.ListValueFrom(ctx, types.StringType, ca.ExtendedKeyUsages)
	diags.Append(d...)
	model.ExtendedKeyUsages = extKeyUsages
}
