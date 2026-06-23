package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/ssh_host_certificates"
	"github.com/ngrok/terraform-provider-ngrok/internal/resource_ssh_host_certificate"
)

var (
	_ resource.Resource                = &sshHostCertificateResource{}
	_ resource.ResourceWithImportState = &sshHostCertificateResource{}
)

type sshHostCertificateResourceModel struct {
	ID                        types.String   `tfsdk:"id"`
	URI                       types.String   `tfsdk:"uri"`
	CreatedAt                 types.String   `tfsdk:"created_at"`
	Description               types.String   `tfsdk:"description"`
	Metadata                  types.String   `tfsdk:"metadata"`
	PublicKey                 types.String   `tfsdk:"public_key"`
	KeyType                   types.String   `tfsdk:"key_type"`
	SSHCertificateAuthorityID types.String   `tfsdk:"ssh_certificate_authority_id"`
	Principals                []types.String `tfsdk:"principals"`
	ValidAfter                types.String   `tfsdk:"valid_after"`
	ValidUntil                types.String   `tfsdk:"valid_until"`
	Certificate               types.String   `tfsdk:"certificate"`
}

type sshHostCertificateResource struct {
	client *ssh_host_certificates.Client
}

func NewSSHHostCertificateResource() resource.Resource {
	return &sshHostCertificateResource{}
}

func (r *sshHostCertificateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssh_host_certificate"
}

func (r *sshHostCertificateResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	s := resource_ssh_host_certificate.SshHostCertificateResourceSchema(ctx)
	attrs := s.Attributes

	addStringPlanModifiers(attrs, "id", useStateForUnknownString())
	addStringPlanModifiers(attrs, "uri", useStateForUnknownString())
	addStringPlanModifiers(attrs, "created_at", useStateForUnknownString())
	addStringPlanModifiers(attrs, "description", useStateForUnknownString())
	addStringPlanModifiers(attrs, "metadata", useStateForUnknownString())
	addStringPlanModifiers(attrs, "public_key", requiresReplaceString())
	addStringPlanModifiers(attrs, "key_type", useStateForUnknownString())
	addStringPlanModifiers(attrs, "ssh_certificate_authority_id", requiresReplaceString())
	addListPlanModifiers(attrs, "principals", requiresReplaceList())
	addStringPlanModifiers(attrs, "valid_after", requiresReplaceString(), useStateForUnknownString())
	addStringPlanModifiers(attrs, "valid_until", requiresReplaceString(), useStateForUnknownString())
	addStringPlanModifiers(attrs, "certificate", useStateForUnknownString())

	resp.Schema = s
}

func (r *sshHostCertificateResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = ssh_host_certificates.NewClient(clientConfig)
}

func (r *sshHostCertificateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan sshHostCertificateResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq := &ngrok.SSHHostCertificateCreate{
		SSHCertificateAuthorityID: plan.SSHCertificateAuthorityID.ValueString(),
		PublicKey:                 plan.PublicKey.ValueString(),
		Principals:                expandStringList(plan.Principals),
		ValidAfter:                plan.ValidAfter.ValueString(),
		ValidUntil:                plan.ValidUntil.ValueString(),
		Description:               plan.Description.ValueString(),
		Metadata:                  plan.Metadata.ValueString(),
	}

	cert, err := r.client.Create(ctx, createReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating SSH host certificate", err.Error())
		return
	}

	flattenSSHHostCertificate(cert, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *sshHostCertificateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state sshHostCertificateResourceModel
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
		resp.Diagnostics.AddError("Error reading SSH host certificate", err.Error())
		return
	}

	flattenSSHHostCertificate(cert, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *sshHostCertificateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan sshHostCertificateResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state sshHostCertificateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq := &ngrok.SSHHostCertificateUpdate{
		ID:          state.ID.ValueString(),
		Description: stringPtrFromFramework(plan.Description),
		Metadata:    stringPtrFromFramework(plan.Metadata),
	}

	cert, err := r.client.Update(ctx, updateReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating SSH host certificate", err.Error())
		return
	}

	flattenSSHHostCertificate(cert, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *sshHostCertificateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state sshHostCertificateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			return
		}
		resp.Diagnostics.AddError("Error deleting SSH host certificate", err.Error())
	}
}

func (r *sshHostCertificateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func flattenSSHHostCertificate(cert *ngrok.SSHHostCertificate, model *sshHostCertificateResourceModel) {
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
	model.ValidAfter = types.StringValue(cert.ValidAfter)
	model.ValidUntil = types.StringValue(cert.ValidUntil)
	model.Certificate = types.StringValue(cert.Certificate)
}
