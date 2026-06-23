package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/endpoints"
	"github.com/ngrok/terraform-provider-ngrok/internal/resource_cloud_endpoint"
)

var (
	_ resource.Resource                = &cloudEndpointResource{}
	_ resource.ResourceWithImportState = &cloudEndpointResource{}
	_ resource.ResourceWithModifyPlan  = &cloudEndpointResource{}
)

type cloudEndpointResourceModel struct {
	ID             types.String   `tfsdk:"id"`
	URL            types.String   `tfsdk:"url"`
	Type           types.String   `tfsdk:"type"`
	TrafficPolicy  types.String   `tfsdk:"traffic_policy"`
	Description    types.String   `tfsdk:"description"`
	Metadata       types.String   `tfsdk:"metadata"`
	Bindings       []types.String `tfsdk:"bindings"`
	PoolingEnabled types.Bool     `tfsdk:"pooling_enabled"`
	DomainID       types.String   `tfsdk:"domain_id"`
	Region         types.String   `tfsdk:"region"`
	URI            types.String   `tfsdk:"uri"`
	CreatedAt      types.String   `tfsdk:"created_at"`
	UpdatedAt      types.String   `tfsdk:"updated_at"`
}

type cloudEndpointResource struct {
	client *endpoints.Client
}

func NewCloudEndpointResource() resource.Resource {
	return &cloudEndpointResource{}
}

func (r *cloudEndpointResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloud_endpoint"
}

func (r *cloudEndpointResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	s := resource_cloud_endpoint.CloudEndpointResourceSchema(ctx)
	attrs := s.Attributes

	// Delete Ref nested object fields not used by the hand-written model
	delete(attrs, "domain")
	delete(attrs, "edge")
	delete(attrs, "tunnel")
	delete(attrs, "tunnel_session")
	delete(attrs, "tcp_addr")
	delete(attrs, "principal")

	// Delete other non-TF fields
	delete(attrs, "host")
	delete(attrs, "hostport")
	delete(attrs, "port")
	delete(attrs, "public_url")
	delete(attrs, "scheme")
	delete(attrs, "upstream_protocol")
	delete(attrs, "upstream_url")
	delete(attrs, "name")
	delete(attrs, "proto")

	// Add domain_id (flat string instead of nested Ref)
	attrs["domain_id"] = schema.StringAttribute{
		Description: "ID of the domain reserved for this endpoint.",
		Computed:    true,
	}

	// Override type to be Computed only (generated has it as Required)
	attrs["type"] = schema.StringAttribute{
		Description:   "The type of endpoint. Always \"cloud\" for cloud endpoints.",
		Computed:      true,
		PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
	}

	// Add JSON syntax validator to traffic_policy
	if a, ok := attrs["traffic_policy"]; ok {
		sa := a.(schema.StringAttribute)
		sa.Validators = []validator.String{JSONSyntax()}
		attrs["traffic_policy"] = sa
	}

	// Plan modifiers
	addStringPlanModifiers(attrs, "id", useStateForUnknownString())
	addStringPlanModifiers(attrs, "domain_id", useStateForUnknownString())
	addStringPlanModifiers(attrs, "region", useStateForUnknownString())
	addStringPlanModifiers(attrs, "uri", useStateForUnknownString())
	addStringPlanModifiers(attrs, "created_at", useStateForUnknownString())
	addStringPlanModifiers(attrs, "updated_at", useStateForUnknownString())
	addStringPlanModifiers(attrs, "url", requiresReplaceString())
	addStringPlanModifiers(attrs, "description", useStateForUnknownString())
	addStringPlanModifiers(attrs, "metadata", useStateForUnknownString())

	// Defaults
	setListDefault(attrs, "bindings", types.ListValueMust(types.StringType, []attr.Value{types.StringValue("public")}))
	setBoolDefault(attrs, "pooling_enabled", false)

	resp.Schema = s
}

func (r *cloudEndpointResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Skip on create or destroy
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return
	}

	var plan cloudEndpointResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state cloudEndpointResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If any user-configurable field changed, mark updated_at as unknown
	// so the provider can return the new value from the API.
	hasChanges := !plan.Description.Equal(state.Description) ||
		!plan.Metadata.Equal(state.Metadata) ||
		!plan.URL.Equal(state.URL) ||
		!plan.TrafficPolicy.Equal(state.TrafficPolicy) ||
		!plan.PoolingEnabled.Equal(state.PoolingEnabled) ||
		!stringListEqual(plan.Bindings, state.Bindings)

	if hasChanges {
		resp.Plan.SetAttribute(ctx, path.Root("updated_at"), types.StringUnknown())
	}
}

func (r *cloudEndpointResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = endpoints.NewClient(clientConfig)
}

func (r *cloudEndpointResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan cloudEndpointResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq := &ngrok.EndpointCreate{
		URL:            plan.URL.ValueString(),
		TrafficPolicy:  plan.TrafficPolicy.ValueString(),
		Description:    stringPtrFromFramework(plan.Description),
		Metadata:       stringPtrFromFramework(plan.Metadata),
		Bindings:       expandStringList(plan.Bindings),
		PoolingEnabled: boolPtrFromFramework(plan.PoolingEnabled),
	}

	endpoint, err := r.client.Create(ctx, createReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating cloud endpoint", err.Error())
		return
	}

	flattenCloudEndpoint(endpoint, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *cloudEndpointResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state cloudEndpointResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	endpoint, err := r.client.Get(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading cloud endpoint", err.Error())
		return
	}

	flattenCloudEndpoint(endpoint, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *cloudEndpointResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan cloudEndpointResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state cloudEndpointResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq := &ngrok.EndpointUpdate{
		ID:             state.ID.ValueString(),
		TrafficPolicy:  stringPtrFromFramework(plan.TrafficPolicy),
		Description:    stringPtrFromFramework(plan.Description),
		Metadata:       stringPtrFromFramework(plan.Metadata),
		Bindings:       expandStringList(plan.Bindings),
		PoolingEnabled: boolPtrFromFramework(plan.PoolingEnabled),
	}

	endpoint, err := r.client.Update(ctx, updateReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating cloud endpoint", err.Error())
		return
	}

	flattenCloudEndpoint(endpoint, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *cloudEndpointResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state cloudEndpointResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			return
		}
		resp.Diagnostics.AddError("Error deleting cloud endpoint", err.Error())
	}
}

func (r *cloudEndpointResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func flattenCloudEndpoint(endpoint *ngrok.Endpoint, model *cloudEndpointResourceModel) {
	model.ID = types.StringValue(endpoint.ID)
	model.URL = types.StringValue(endpoint.URL)
	model.Type = types.StringValue(endpoint.Type)
	model.TrafficPolicy = types.StringValue(endpoint.TrafficPolicy)
	model.Description = types.StringValue(endpoint.Description)
	model.Metadata = types.StringValue(endpoint.Metadata)
	model.Bindings = flattenStringList(endpoint.Bindings)
	model.PoolingEnabled = types.BoolValue(endpoint.PoolingEnabled)
	model.DomainID = types.StringValue(flattenRef(endpoint.Domain))
	model.Region = types.StringValue(endpoint.Region)
	model.URI = types.StringValue(endpoint.URI)
	model.CreatedAt = types.StringValue(endpoint.CreatedAt)
	model.UpdatedAt = types.StringValue(endpoint.UpdatedAt)
}
