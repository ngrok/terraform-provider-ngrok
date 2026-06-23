package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/vaults"
	"github.com/ngrok/terraform-provider-ngrok/internal/resource_vault"
)

var (
	_ resource.Resource                = &vaultResource{}
	_ resource.ResourceWithImportState = &vaultResource{}
	_ resource.ResourceWithModifyPlan  = &vaultResource{}
)

type vaultResourceModel struct {
	ID            types.String `tfsdk:"id"`
	URI           types.String `tfsdk:"uri"`
	CreatedAt     types.String `tfsdk:"created_at"`
	UpdatedAt     types.String `tfsdk:"updated_at"`
	Name          types.String `tfsdk:"name"`
	Description   types.String `tfsdk:"description"`
	Metadata      types.String `tfsdk:"metadata"`
	CreatedBy     types.String `tfsdk:"created_by"`
	LastUpdatedBy types.String `tfsdk:"last_updated_by"`
}

type vaultResource struct {
	client *vaults.Client
}

func NewVaultResource() resource.Resource {
	return &vaultResource{}
}

func (r *vaultResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vault"
}

func (r *vaultResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	s := resource_vault.VaultResourceSchema(ctx)
	attrs := s.Attributes

	// The generated schema has name as Optional+Computed; override to Required.
	if a, ok := attrs["name"]; ok {
		sa := a.(schema.StringAttribute)
		sa.Required = true
		sa.Optional = false
		sa.Computed = false
		attrs["name"] = sa
	}

	addStringPlanModifiers(attrs, "id", useStateForUnknownString())
	addStringPlanModifiers(attrs, "description", useStateForUnknownString())
	addStringPlanModifiers(attrs, "metadata", useStateForUnknownString())
	addStringPlanModifiers(attrs, "uri", useStateForUnknownString())
	addStringPlanModifiers(attrs, "created_at", useStateForUnknownString())
	addStringPlanModifiers(attrs, "updated_at", useStateForUnknownString())
	addStringPlanModifiers(attrs, "created_by", useStateForUnknownString())
	addStringPlanModifiers(attrs, "last_updated_by", useStateForUnknownString())

	resp.Schema = s
}

func (r *vaultResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = vaults.NewClient(clientConfig)
}

func (r *vaultResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan vaultResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq := &ngrok.VaultCreate{
		Name:        plan.Name.ValueString(),
		Description: plan.Description.ValueString(),
		Metadata:    plan.Metadata.ValueString(),
	}

	vault, err := r.client.Create(ctx, createReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating vault", err.Error())
		return
	}

	flattenVault(vault, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *vaultResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state vaultResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	vault, err := r.client.Get(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading vault", err.Error())
		return
	}

	flattenVault(vault, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *vaultResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan vaultResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state vaultResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq := &ngrok.VaultUpdate{
		ID:          state.ID.ValueString(),
		Description: stringPtrFromFramework(plan.Description),
		Metadata:    stringPtrFromFramework(plan.Metadata),
	}
	if !plan.Name.Equal(state.Name) {
		updateReq.Name = stringPtrFromFramework(plan.Name)
	}

	vault, err := r.client.Update(ctx, updateReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating vault", err.Error())
		return
	}

	flattenVault(vault, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *vaultResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state vaultResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			return
		}
		resp.Diagnostics.AddError("Error deleting vault", err.Error())
	}
}

func (r *vaultResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *vaultResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Skip on create or destroy
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return
	}

	var plan vaultResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state vaultResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If any user-configurable field changed, mark updated_at and last_updated_by
	// as unknown so the provider can return the new value from the API.
	hasChanges := !plan.Name.Equal(state.Name) ||
		!plan.Description.Equal(state.Description) ||
		!plan.Metadata.Equal(state.Metadata)

	if hasChanges {
		resp.Plan.SetAttribute(ctx, path.Root("updated_at"), types.StringUnknown())
		resp.Plan.SetAttribute(ctx, path.Root("last_updated_by"), types.StringUnknown())
	}
}

func flattenVault(vault *ngrok.Vault, model *vaultResourceModel) {
	model.ID = types.StringValue(vault.ID)
	model.URI = types.StringValue(vault.URI)
	model.CreatedAt = types.StringValue(vault.CreatedAt)
	model.UpdatedAt = types.StringValue(vault.UpdatedAt)
	model.Name = types.StringValue(vault.Name)
	model.Description = types.StringValue(vault.Description)
	model.Metadata = types.StringValue(vault.Metadata)
	model.CreatedBy = types.StringValue(vault.CreatedBy)
	model.LastUpdatedBy = types.StringValue(vault.LastUpdatedBy)
}
