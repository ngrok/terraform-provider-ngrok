package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/ssh_credentials"
	"github.com/ngrok/terraform-provider-ngrok/internal/resource_ssh_credential"
)

var (
	_ resource.Resource                = &sshCredentialResource{}
	_ resource.ResourceWithImportState = &sshCredentialResource{}
)

type sshCredentialResourceModel struct {
	ID          types.String `tfsdk:"id"`
	URI         types.String `tfsdk:"uri"`
	CreatedAt   types.String `tfsdk:"created_at"`
	Description types.String `tfsdk:"description"`
	Metadata    types.String `tfsdk:"metadata"`
	PublicKey   types.String `tfsdk:"public_key"`
	ACL         types.List   `tfsdk:"acl"`
	OwnerID     types.String `tfsdk:"owner_id"`
}

type sshCredentialResource struct {
	client *ssh_credentials.Client
}

func NewSSHCredentialResource() resource.Resource {
	return &sshCredentialResource{}
}

func (r *sshCredentialResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssh_credential"
}

func (r *sshCredentialResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	s := resource_ssh_credential.SshCredentialResourceSchema(ctx)
	attrs := s.Attributes

	// Plan modifiers
	addStringPlanModifiers(attrs, "id", useStateForUnknownString())
	addStringPlanModifiers(attrs, "uri", useStateForUnknownString())
	addStringPlanModifiers(attrs, "created_at", useStateForUnknownString())
	addStringPlanModifiers(attrs, "description", useStateForUnknownString())
	addStringPlanModifiers(attrs, "metadata", useStateForUnknownString())
	addStringPlanModifiers(attrs, "public_key", requiresReplaceString())
	addStringPlanModifiers(attrs, "owner_id", requiresReplaceString(), useStateForUnknownString())

	resp.Schema = s
}

func (r *sshCredentialResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = ssh_credentials.NewClient(clientConfig)
}

func (r *sshCredentialResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan sshCredentialResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var acl []string
	plan.ACL.ElementsAs(ctx, &acl, false)

	createReq := &ngrok.SSHCredentialCreate{
		Description: plan.Description.ValueString(),
		Metadata:    plan.Metadata.ValueString(),
		ACL:         acl,
		PublicKey:   plan.PublicKey.ValueString(),
		OwnerID:     stringPtrFromFramework(plan.OwnerID),
	}

	cred, err := r.client.Create(ctx, createReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating SSH credential", err.Error())
		return
	}

	flattenSSHCredential(ctx, cred, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *sshCredentialResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state sshCredentialResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	cred, err := r.client.Get(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading SSH credential", err.Error())
		return
	}

	flattenSSHCredential(ctx, cred, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *sshCredentialResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan sshCredentialResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state sshCredentialResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var acl []string
	plan.ACL.ElementsAs(ctx, &acl, false)

	updateReq := &ngrok.SSHCredentialUpdate{
		ID:          state.ID.ValueString(),
		Description: stringPtrFromFramework(plan.Description),
		Metadata:    stringPtrFromFramework(plan.Metadata),
		ACL:         acl,
	}

	cred, err := r.client.Update(ctx, updateReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating SSH credential", err.Error())
		return
	}

	flattenSSHCredential(ctx, cred, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *sshCredentialResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state sshCredentialResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			return
		}
		resp.Diagnostics.AddError("Error deleting SSH credential", err.Error())
	}
}

func (r *sshCredentialResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func flattenSSHCredential(ctx context.Context, cred *ngrok.SSHCredential, model *sshCredentialResourceModel) {
	model.ID = types.StringValue(cred.ID)
	model.URI = types.StringValue(cred.URI)
	model.CreatedAt = types.StringValue(cred.CreatedAt)
	model.Description = types.StringValue(cred.Description)
	model.Metadata = types.StringValue(cred.Metadata)
	model.PublicKey = types.StringValue(cred.PublicKey)
	model.ACL, _ = types.ListValueFrom(ctx, types.StringType, cred.ACL)

	if cred.OwnerID != nil {
		model.OwnerID = types.StringValue(*cred.OwnerID)
	} else {
		model.OwnerID = types.StringNull()
	}
}
