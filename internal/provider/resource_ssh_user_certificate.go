package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/ssh_user_certificates"
	"github.com/ngrok/terraform-provider-ngrok/internal/resource_ssh_user_certificate"
)

var (
	_ resource.Resource                = &sshUserCertificateResource{}
	_ resource.ResourceWithImportState = &sshUserCertificateResource{}
)

type sshUserCertificateResourceModel struct {
	ID                        types.String   `tfsdk:"id"`
	URI                       types.String   `tfsdk:"uri"`
	CreatedAt                 types.String   `tfsdk:"created_at"`
	Description               types.String   `tfsdk:"description"`
	Metadata                  types.String   `tfsdk:"metadata"`
	PublicKey                 types.String   `tfsdk:"public_key"`
	KeyType                   types.String   `tfsdk:"key_type"`
	SSHCertificateAuthorityID types.String   `tfsdk:"ssh_certificate_authority_id"`
	Principals                []types.String `tfsdk:"principals"`
	CriticalOptions           types.Map      `tfsdk:"critical_options"`
	Extensions                types.Map      `tfsdk:"extensions"`
	ValidAfter                types.String   `tfsdk:"valid_after"`
	ValidUntil                types.String   `tfsdk:"valid_until"`
	Certificate               types.String   `tfsdk:"certificate"`
}

type sshUserCertificateResource struct {
	client *ssh_user_certificates.Client
}

func NewSSHUserCertificateResource() resource.Resource {
	return &sshUserCertificateResource{}
}

func (r *sshUserCertificateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssh_user_certificate"
}

func (r *sshUserCertificateResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	s := resource_ssh_user_certificate.SshUserCertificateResourceSchema(ctx)
	attrs := s.Attributes

	// Replace generated CustomType SingleNestedAttributes with standard MapAttributes
	attrs["critical_options"] = schema.MapAttribute{
		Description: "A map of critical options included in the certificate.",
		Optional:    true,
		ElementType: types.StringType,
	}
	attrs["extensions"] = schema.MapAttribute{
		Description: "A map of extensions included in the certificate.",
		Optional:    true,
		ElementType: types.StringType,
	}

	addStringPlanModifiers(attrs, "id", useStateForUnknownString())
	addStringPlanModifiers(attrs, "uri", useStateForUnknownString())
	addStringPlanModifiers(attrs, "created_at", useStateForUnknownString())
	addStringPlanModifiers(attrs, "description", useStateForUnknownString())
	addStringPlanModifiers(attrs, "metadata", useStateForUnknownString())
	addStringPlanModifiers(attrs, "public_key", requiresReplaceString())
	addStringPlanModifiers(attrs, "key_type", useStateForUnknownString())
	addStringPlanModifiers(attrs, "ssh_certificate_authority_id", requiresReplaceString())
	addListPlanModifiers(attrs, "principals", requiresReplaceList())
	addMapPlanModifiers(attrs, "critical_options", requiresReplaceMap())
	addMapPlanModifiers(attrs, "extensions", requiresReplaceMap())
	addStringPlanModifiers(attrs, "valid_after", requiresReplaceString(), useStateForUnknownString())
	addStringPlanModifiers(attrs, "valid_until", requiresReplaceString(), useStateForUnknownString())
	addStringPlanModifiers(attrs, "certificate", useStateForUnknownString())

	resp.Schema = s
}

func (r *sshUserCertificateResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = ssh_user_certificates.NewClient(clientConfig)
}

func (r *sshUserCertificateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan sshUserCertificateResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq := &ngrok.SSHUserCertificateCreate{
		SSHCertificateAuthorityID: plan.SSHCertificateAuthorityID.ValueString(),
		PublicKey:                 plan.PublicKey.ValueString(),
		Principals:                expandStringList(plan.Principals),
		CriticalOptions:           expandStringMap(ctx, plan.CriticalOptions),
		Extensions:                expandStringMap(ctx, plan.Extensions),
		ValidAfter:                plan.ValidAfter.ValueString(),
		ValidUntil:                plan.ValidUntil.ValueString(),
		Description:               plan.Description.ValueString(),
		Metadata:                  plan.Metadata.ValueString(),
	}

	cert, err := r.client.Create(ctx, createReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating SSH user certificate", err.Error())
		return
	}

	flattenSSHUserCertificate(ctx, cert, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *sshUserCertificateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state sshUserCertificateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	cert, err := r.client.Get(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading SSH user certificate", err.Error())
		return
	}

	flattenSSHUserCertificate(ctx, cert, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *sshUserCertificateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan sshUserCertificateResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state sshUserCertificateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq := &ngrok.SSHUserCertificateUpdate{
		ID:          state.ID.ValueString(),
		Description: stringPtrFromFramework(plan.Description),
		Metadata:    stringPtrFromFramework(plan.Metadata),
	}

	cert, err := r.client.Update(ctx, updateReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating SSH user certificate", err.Error())
		return
	}

	flattenSSHUserCertificate(ctx, cert, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *sshUserCertificateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state sshUserCertificateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			return
		}
		resp.Diagnostics.AddError("Error deleting SSH user certificate", err.Error())
	}
}

func (r *sshUserCertificateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandStringMap(ctx context.Context, m types.Map) map[string]string {
	if m.IsNull() || m.IsUnknown() {
		return nil
	}
	result := make(map[string]string, len(m.Elements()))
	for k, v := range m.Elements() {
		result[k] = v.(types.String).ValueString()
	}
	return result
}

func flattenStringMap(ctx context.Context, m map[string]string) types.Map {
	if m == nil {
		return types.MapNull(types.StringType)
	}
	elements := make(map[string]types.String, len(m))
	for k, v := range m {
		elements[k] = types.StringValue(v)
	}
	result, _ := types.MapValueFrom(ctx, types.StringType, elements)
	return result
}

func flattenSSHUserCertificate(ctx context.Context, cert *ngrok.SSHUserCertificate, model *sshUserCertificateResourceModel) {
	model.ID = types.StringValue(cert.ID)
	model.URI = types.StringValue(cert.URI)
	model.CreatedAt = types.StringValue(cert.CreatedAt)
	model.Description = types.StringValue(cert.Description)
	model.Metadata = types.StringValue(cert.Metadata)
	model.PublicKey = types.StringValue(cert.PublicKey)
	model.KeyType = types.StringValue(cert.KeyType)
	model.SSHCertificateAuthorityID = types.StringValue(cert.SSHCertificateAuthorityID)
	if model.Principals != nil {
		model.Principals = flattenStringList(cert.Principals)
	} else if len(cert.Principals) > 0 {
		model.Principals = flattenStringList(cert.Principals)
	}
	model.CriticalOptions = flattenStringMap(ctx, cert.CriticalOptions)
	model.Extensions = flattenStringMap(ctx, cert.Extensions)
	model.ValidAfter = types.StringValue(cert.ValidAfter)
	model.ValidUntil = types.StringValue(cert.ValidUntil)
	model.Certificate = types.StringValue(cert.Certificate)
}
