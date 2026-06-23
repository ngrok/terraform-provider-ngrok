package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/credentials"
	"github.com/ngrok/terraform-provider-ngrok/internal/resource_credential"
)

var (
	_ resource.Resource                = &credentialResource{}
	_ resource.ResourceWithImportState = &credentialResource{}
)

type credentialResourceModel struct {
	ID          types.String `tfsdk:"id"`
	URI         types.String `tfsdk:"uri"`
	CreatedAt   types.String `tfsdk:"created_at"`
	Description types.String `tfsdk:"description"`
	Metadata    types.String `tfsdk:"metadata"`
	Token       types.String `tfsdk:"token"`
	ACL         types.List   `tfsdk:"acl"`
	OwnerID     types.String `tfsdk:"owner_id"`
}

type credentialResource struct {
	client *credentials.Client
}

func NewCredentialResource() resource.Resource {
	return &credentialResource{}
}

func (r *credentialResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_credential"
}

func (r *credentialResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	s := resource_credential.CredentialResourceSchema(ctx)
	attrs := s.Attributes

	addStringPlanModifiers(attrs, "id", useStateForUnknownString())
	addStringPlanModifiers(attrs, "uri", useStateForUnknownString())
	addStringPlanModifiers(attrs, "created_at", useStateForUnknownString())
	addStringPlanModifiers(attrs, "description", useStateForUnknownString())
	addStringPlanModifiers(attrs, "metadata", useStateForUnknownString())
	addStringPlanModifiers(attrs, "token", useStateForUnknownString())
	markSensitive(attrs, "token")
	addStringPlanModifiers(attrs, "owner_id", requiresReplaceString(), useStateForUnknownString())

	resp.Schema = s
}

func (r *credentialResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = credentials.NewClient(clientConfig)
}

func (r *credentialResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan credentialResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var acl []string
	plan.ACL.ElementsAs(ctx, &acl, false)

	createReq := &ngrok.CredentialCreate{
		Description: plan.Description.ValueString(),
		Metadata:    plan.Metadata.ValueString(),
		ACL:         acl,
		OwnerID:     stringPtrFromFramework(plan.OwnerID),
	}

	cred, err := r.client.Create(ctx, createReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating credential", err.Error())
		return
	}

	flattenCredential(ctx, cred, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *credentialResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state credentialResourceModel
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
		resp.Diagnostics.AddError("Error reading credential", err.Error())
		return
	}

	// Preserve the token from state since it's only available on create
	token := state.Token
	flattenCredential(ctx, cred, &state)
	state.Token = token
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *credentialResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan credentialResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state credentialResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var acl []string
	plan.ACL.ElementsAs(ctx, &acl, false)

	updateReq := &ngrok.CredentialUpdate{
		ID:          state.ID.ValueString(),
		Description: stringPtrFromFramework(plan.Description),
		Metadata:    stringPtrFromFramework(plan.Metadata),
		ACL:         acl,
	}

	cred, err := r.client.Update(ctx, updateReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating credential", err.Error())
		return
	}

	// Preserve the token from state since it's only available on create
	token := state.Token
	flattenCredential(ctx, cred, &plan)
	plan.Token = token
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *credentialResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state credentialResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			return
		}
		resp.Diagnostics.AddError("Error deleting credential", err.Error())
	}
}

func (r *credentialResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func flattenCredential(ctx context.Context, cred *ngrok.Credential, model *credentialResourceModel) {
	model.ID = types.StringValue(cred.ID)
	model.URI = types.StringValue(cred.URI)
	model.CreatedAt = types.StringValue(cred.CreatedAt)
	model.Description = types.StringValue(cred.Description)
	model.Metadata = types.StringValue(cred.Metadata)
	model.ACL, _ = types.ListValueFrom(ctx, types.StringType, cred.ACL)

	if cred.Token != nil {
		model.Token = types.StringValue(*cred.Token)
	}

	if cred.OwnerID != nil {
		model.OwnerID = types.StringValue(*cred.OwnerID)
	} else {
		model.OwnerID = types.StringNull()
	}
}
