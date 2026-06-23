package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/secrets"
	"github.com/ngrok/terraform-provider-ngrok/internal/resource_secret"
)

var (
	_ resource.Resource                = &secretResource{}
	_ resource.ResourceWithImportState = &secretResource{}
	_ resource.ResourceWithModifyPlan  = &secretResource{}
)

type secretResourceModel struct {
	ID              types.String `tfsdk:"id"`
	URI             types.String `tfsdk:"uri"`
	CreatedAt       types.String `tfsdk:"created_at"`
	UpdatedAt       types.String `tfsdk:"updated_at"`
	Name            types.String `tfsdk:"name"`
	Value           types.String `tfsdk:"value"`
	Description     types.String `tfsdk:"description"`
	Metadata        types.String `tfsdk:"metadata"`
	VaultID         types.String `tfsdk:"vault_id"`
	VaultName       types.String `tfsdk:"vault_name"`
	CreatedByID     types.String `tfsdk:"created_by_id"`
	LastUpdatedByID types.String `tfsdk:"last_updated_by_id"`
}

type secretResource struct {
	client *secrets.Client
}

func NewSecretResource() resource.Resource {
	return &secretResource{}
}

func (r *secretResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_secret"
}

func (r *secretResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	s := resource_secret.SecretResourceSchema(ctx)
	attrs := s.Attributes

	// Remove Ref-type nested objects that don't match hand-written model
	delete(attrs, "vault")
	delete(attrs, "created_by")
	delete(attrs, "last_updated_by")

	// Make name Required (generated has it as Optional+Computed)
	if a, ok := attrs["name"]; ok {
		sa := a.(schema.StringAttribute)
		sa.Required = true
		sa.Optional = false
		sa.Computed = false
		attrs["name"] = sa
	}

	// Make value Required + Sensitive (generated has it as Optional+Computed)
	if a, ok := attrs["value"]; ok {
		sa := a.(schema.StringAttribute)
		sa.Required = true
		sa.Optional = false
		sa.Computed = false
		attrs["value"] = sa
	}
	addStringPlanModifiers(attrs, "value", useStateForUnknownString())
	markSensitive(attrs, "value")

	// Make vault_id Required with RequiresReplace (generated has it as Optional+Computed)
	if a, ok := attrs["vault_id"]; ok {
		sa := a.(schema.StringAttribute)
		sa.Required = true
		sa.Optional = false
		sa.Computed = false
		attrs["vault_id"] = sa
	}
	addStringPlanModifiers(attrs, "vault_id", requiresReplaceString())

	// Add hand-written fields not in generated schema
	attrs["created_by_id"] = schema.StringAttribute{
		Description: "The ID of the user or bot that created the secret.",
		Computed:    true,
	}
	attrs["last_updated_by_id"] = schema.StringAttribute{
		Description: "The ID of the user or bot that last updated the secret.",
		Computed:    true,
	}

	addStringPlanModifiers(attrs, "id", useStateForUnknownString())
	addStringPlanModifiers(attrs, "uri", useStateForUnknownString())
	addStringPlanModifiers(attrs, "created_at", useStateForUnknownString())
	addStringPlanModifiers(attrs, "updated_at", useStateForUnknownString())
	addStringPlanModifiers(attrs, "description", useStateForUnknownString())
	addStringPlanModifiers(attrs, "metadata", useStateForUnknownString())
	addStringPlanModifiers(attrs, "created_by_id", useStateForUnknownString())
	addStringPlanModifiers(attrs, "last_updated_by_id", useStateForUnknownString())

	resp.Schema = s
}

func (r *secretResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = secrets.NewClient(clientConfig)
}

func (r *secretResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan secretResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq := &ngrok.SecretCreate{
		Name:        plan.Name.ValueString(),
		Value:       plan.Value.ValueString(),
		Description: plan.Description.ValueString(),
		Metadata:    plan.Metadata.ValueString(),
		VaultID:     plan.VaultID.ValueString(),
	}

	secret, err := r.client.Create(ctx, createReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating secret", err.Error())
		return
	}

	flattenSecret(secret, &plan)
	// Preserve the value from the plan since the API does not return it
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *secretResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state secretResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	secret, err := r.client.Get(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading secret", err.Error())
		return
	}

	// Preserve the value from state since the API does not return it
	preservedValue := state.Value
	flattenSecret(secret, &state)
	state.Value = preservedValue
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *secretResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan secretResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state secretResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq := &ngrok.SecretUpdate{
		ID:          state.ID.ValueString(),
		Description: stringPtrFromFramework(plan.Description),
		Metadata:    stringPtrFromFramework(plan.Metadata),
	}
	if !plan.Name.Equal(state.Name) {
		updateReq.Name = stringPtrFromFramework(plan.Name)
	}
	if !plan.Value.Equal(state.Value) {
		updateReq.Value = stringPtrFromFramework(plan.Value)
	}

	secret, err := r.client.Update(ctx, updateReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating secret", err.Error())
		return
	}

	flattenSecret(secret, &plan)
	// Preserve the value from the plan since the API does not return it
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *secretResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state secretResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			return
		}
		resp.Diagnostics.AddError("Error deleting secret", err.Error())
	}
}

func (r *secretResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *secretResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Skip on create or destroy
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return
	}

	var plan secretResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state secretResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If any user-configurable field changed, mark updated_at and
	// last_updated_by_id as unknown so the provider returns the new
	// values from the API.
	hasChanges := !plan.Name.Equal(state.Name) ||
		!plan.Value.Equal(state.Value) ||
		!plan.Description.Equal(state.Description) ||
		!plan.Metadata.Equal(state.Metadata)

	if hasChanges {
		resp.Plan.SetAttribute(ctx, path.Root("updated_at"), types.StringUnknown())
		resp.Plan.SetAttribute(ctx, path.Root("last_updated_by_id"), types.StringUnknown())
	}
}

func flattenSecret(secret *ngrok.Secret, model *secretResourceModel) {
	model.ID = types.StringValue(secret.ID)
	model.URI = types.StringValue(secret.URI)
	model.CreatedAt = types.StringValue(secret.CreatedAt)
	model.UpdatedAt = types.StringValue(secret.UpdatedAt)
	model.Name = types.StringValue(secret.Name)
	model.Description = types.StringValue(secret.Description)
	model.Metadata = types.StringValue(secret.Metadata)
	model.VaultID = types.StringValue(secret.Vault.ID)
	model.VaultName = types.StringValue(secret.VaultName)
	model.CreatedByID = types.StringValue(secret.CreatedBy.ID)
	model.LastUpdatedByID = types.StringValue(secret.LastUpdatedBy.ID)
}
