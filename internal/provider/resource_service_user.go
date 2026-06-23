package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/bot_users"
	"github.com/ngrok/terraform-provider-ngrok/internal/resource_service_user"
)

var (
	_ resource.Resource                = &serviceUserResource{}
	_ resource.ResourceWithImportState = &serviceUserResource{}
)

type serviceUserResourceModel struct {
	ID        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Active    types.Bool   `tfsdk:"active"`
	URI       types.String `tfsdk:"uri"`
	CreatedAt types.String `tfsdk:"created_at"`
}

type serviceUserResource struct {
	client *bot_users.Client
}

func NewServiceUserResource() resource.Resource {
	return &serviceUserResource{}
}

func (r *serviceUserResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_user"
}

func (r *serviceUserResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	s := resource_service_user.ServiceUserResourceSchema(ctx)
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
	addStringPlanModifiers(attrs, "uri", useStateForUnknownString())
	addStringPlanModifiers(attrs, "created_at", useStateForUnknownString())
	setBoolDefault(attrs, "active", true)

	resp.Schema = s
}

func (r *serviceUserResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = bot_users.NewClient(clientConfig)
}

func (r *serviceUserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan serviceUserResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq := &ngrok.BotUserCreate{
		Name:   plan.Name.ValueString(),
		Active: boolPtrFromFramework(plan.Active),
	}

	user, err := r.client.Create(ctx, createReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating service user", err.Error())
		return
	}

	flattenServiceUser(user, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *serviceUserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state serviceUserResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	user, err := r.client.Get(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading service user", err.Error())
		return
	}

	flattenServiceUser(user, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *serviceUserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan serviceUserResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state serviceUserResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq := &ngrok.BotUserUpdate{
		ID: state.ID.ValueString(),
	}
	// Only send fields that changed — nil means "don't update this field"
	if !plan.Name.Equal(state.Name) {
		updateReq.Name = stringPtrFromFramework(plan.Name)
	}
	if !plan.Active.Equal(state.Active) {
		updateReq.Active = boolPtrFromFramework(plan.Active)
	}

	user, err := r.client.Update(ctx, updateReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating service user", err.Error())
		return
	}

	flattenServiceUser(user, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *serviceUserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state serviceUserResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			return
		}
		resp.Diagnostics.AddError("Error deleting service user", err.Error())
	}
}

func (r *serviceUserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func flattenServiceUser(user *ngrok.BotUser, model *serviceUserResourceModel) {
	model.ID = types.StringValue(user.ID)
	model.Name = types.StringValue(user.Name)
	model.Active = types.BoolValue(user.Active)
	model.URI = types.StringValue(user.URI)
	model.CreatedAt = types.StringValue(user.CreatedAt)
}
