package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/api_keys"
	"github.com/ngrok/terraform-provider-ngrok/internal/resource_api_key"
)

var (
	_ resource.Resource                = &apiKeyResource{}
	_ resource.ResourceWithImportState = &apiKeyResource{}
)

type apiKeyResourceModel struct {
	ID          types.String `tfsdk:"id"`
	URI         types.String `tfsdk:"uri"`
	Description types.String `tfsdk:"description"`
	Metadata    types.String `tfsdk:"metadata"`
	CreatedAt   types.String `tfsdk:"created_at"`
	Token       types.String `tfsdk:"token"`
	OwnerID     types.String `tfsdk:"owner_id"`
}

type apiKeyResource struct {
	client *api_keys.Client
}

func NewAPIKeyResource() resource.Resource {
	return &apiKeyResource{}
}

func (r *apiKeyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_key"
}

func (r *apiKeyResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	s := resource_api_key.ApiKeyResourceSchema(ctx)
	attrs := s.Attributes

	addStringPlanModifiers(attrs, "id", useStateForUnknownString())
	addStringPlanModifiers(attrs, "uri", useStateForUnknownString())
	addStringPlanModifiers(attrs, "description", useStateForUnknownString())
	addStringPlanModifiers(attrs, "metadata", useStateForUnknownString())
	addStringPlanModifiers(attrs, "created_at", useStateForUnknownString())
	addStringPlanModifiers(attrs, "token", useStateForUnknownString())
	markSensitive(attrs, "token")
	addStringPlanModifiers(attrs, "owner_id", requiresReplaceString(), useStateForUnknownString())

	resp.Schema = s
}

func (r *apiKeyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = api_keys.NewClient(clientConfig)
}

func (r *apiKeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan apiKeyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq := &ngrok.APIKeyCreate{
		Description: plan.Description.ValueString(),
		Metadata:    plan.Metadata.ValueString(),
		OwnerID:     stringPtrFromFramework(plan.OwnerID),
	}

	apiKey, err := r.client.Create(ctx, createReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating API key", err.Error())
		return
	}

	flattenAPIKey(apiKey, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *apiKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state apiKeyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiKey, err := r.client.Get(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading API key", err.Error())
		return
	}

	// Preserve the token from state since it's only available on create
	token := state.Token
	flattenAPIKey(apiKey, &state)
	state.Token = token
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *apiKeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan apiKeyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state apiKeyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq := &ngrok.APIKeyUpdate{
		ID:          state.ID.ValueString(),
		Description: stringPtrFromFramework(plan.Description),
		Metadata:    stringPtrFromFramework(plan.Metadata),
	}

	apiKey, err := r.client.Update(ctx, updateReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating API key", err.Error())
		return
	}

	// Preserve the token from state since it's only available on create
	token := state.Token
	flattenAPIKey(apiKey, &plan)
	plan.Token = token
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *apiKeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state apiKeyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			return
		}
		resp.Diagnostics.AddError("Error deleting API key", err.Error())
	}
}

func (r *apiKeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func flattenAPIKey(apiKey *ngrok.APIKey, model *apiKeyResourceModel) {
	model.ID = types.StringValue(apiKey.ID)
	model.URI = types.StringValue(apiKey.URI)
	model.Description = types.StringValue(apiKey.Description)
	model.Metadata = types.StringValue(apiKey.Metadata)
	model.CreatedAt = types.StringValue(apiKey.CreatedAt)

	if apiKey.Token != nil {
		model.Token = types.StringValue(*apiKey.Token)
	}

	if apiKey.OwnerID != nil {
		model.OwnerID = types.StringValue(*apiKey.OwnerID)
	} else {
		model.OwnerID = types.StringNull()
	}
}
