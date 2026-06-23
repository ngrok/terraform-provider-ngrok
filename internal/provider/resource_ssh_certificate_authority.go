package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/ssh_certificate_authorities"
	"github.com/ngrok/terraform-provider-ngrok/internal/resource_ssh_certificate_authority"
)

var (
	_ resource.Resource                = &sshCertificateAuthorityResource{}
	_ resource.ResourceWithImportState = &sshCertificateAuthorityResource{}
)

type sshCertificateAuthorityResourceModel struct {
	ID             types.String `tfsdk:"id"`
	URI            types.String `tfsdk:"uri"`
	CreatedAt      types.String `tfsdk:"created_at"`
	Description    types.String `tfsdk:"description"`
	Metadata       types.String `tfsdk:"metadata"`
	PublicKey      types.String `tfsdk:"public_key"`
	KeyType        types.String `tfsdk:"key_type"`
	PrivateKeyType types.String `tfsdk:"private_key_type"`
	EllipticCurve  types.String `tfsdk:"elliptic_curve"`
	KeySize        types.Int64  `tfsdk:"key_size"`
}

type sshCertificateAuthorityResource struct {
	client *ssh_certificate_authorities.Client
}

func NewSSHCertificateAuthorityResource() resource.Resource {
	return &sshCertificateAuthorityResource{}
}

func (r *sshCertificateAuthorityResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssh_certificate_authority"
}

func (r *sshCertificateAuthorityResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	s := resource_ssh_certificate_authority.SshCertificateAuthorityResourceSchema(ctx)
	attrs := s.Attributes

	addStringPlanModifiers(attrs, "id", useStateForUnknownString())
	addStringPlanModifiers(attrs, "uri", useStateForUnknownString())
	addStringPlanModifiers(attrs, "created_at", useStateForUnknownString())
	addStringPlanModifiers(attrs, "description", useStateForUnknownString())
	addStringPlanModifiers(attrs, "metadata", useStateForUnknownString())
	addStringPlanModifiers(attrs, "public_key", useStateForUnknownString())
	addStringPlanModifiers(attrs, "key_type", useStateForUnknownString())
	addStringPlanModifiers(attrs, "private_key_type", requiresReplaceString())
	addStringPlanModifiers(attrs, "elliptic_curve", useStateForUnknownString(), requiresReplaceString())
	addInt64PlanModifiers(attrs, "key_size", useStateForUnknownInt64(), requiresReplaceInt64())

	resp.Schema = s
}

func (r *sshCertificateAuthorityResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = ssh_certificate_authorities.NewClient(clientConfig)
}

func (r *sshCertificateAuthorityResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan sshCertificateAuthorityResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq := &ngrok.SSHCertificateAuthorityCreate{
		Description:    plan.Description.ValueString(),
		Metadata:       plan.Metadata.ValueString(),
		PrivateKeyType: plan.PrivateKeyType.ValueString(),
		EllipticCurve:  plan.EllipticCurve.ValueString(),
		KeySize:        plan.KeySize.ValueInt64(),
	}

	ca, err := r.client.Create(ctx, createReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating SSH certificate authority", err.Error())
		return
	}

	flattenSSHCertificateAuthority(ca, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *sshCertificateAuthorityResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state sshCertificateAuthorityResourceModel
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
		resp.Diagnostics.AddError("Error reading SSH certificate authority", err.Error())
		return
	}

	flattenSSHCertificateAuthority(ca, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *sshCertificateAuthorityResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan sshCertificateAuthorityResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state sshCertificateAuthorityResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq := &ngrok.SSHCertificateAuthorityUpdate{
		ID:          state.ID.ValueString(),
		Description: stringPtrFromFramework(plan.Description),
		Metadata:    stringPtrFromFramework(plan.Metadata),
	}

	ca, err := r.client.Update(ctx, updateReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating SSH certificate authority", err.Error())
		return
	}

	flattenSSHCertificateAuthority(ca, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *sshCertificateAuthorityResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state sshCertificateAuthorityResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			return
		}
		resp.Diagnostics.AddError("Error deleting SSH certificate authority", err.Error())
	}
}

func (r *sshCertificateAuthorityResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func flattenSSHCertificateAuthority(ca *ngrok.SSHCertificateAuthority, model *sshCertificateAuthorityResourceModel) {
	model.ID = types.StringValue(ca.ID)
	model.URI = types.StringValue(ca.URI)
	model.CreatedAt = types.StringValue(ca.CreatedAt)
	model.Description = types.StringValue(ca.Description)
	model.Metadata = types.StringValue(ca.Metadata)
	model.PublicKey = types.StringValue(ca.PublicKey)
	model.KeyType = types.StringValue(ca.KeyType)
	if model.EllipticCurve.IsUnknown() {
		model.EllipticCurve = types.StringValue("")
	}
	if model.KeySize.IsUnknown() {
		model.KeySize = types.Int64Value(0)
	}
}
