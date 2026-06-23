package provider

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/terraform-provider-ngrok/internal/resource_event_subscription"
)

var (
	_ resource.Resource                = &eventSubscriptionResource{}
	_ resource.ResourceWithImportState = &eventSubscriptionResource{}
)

// Local API types that include the undocumented "fields" parameter.

type eventSourceReplaceAPI struct {
	Type   string   `json:"type,omitzero"`
	Fields []string `json:"fields,omitempty"`
}

type eventSourceAPI struct {
	Type   string   `json:"type,omitzero"`
	Fields []string `json:"fields,omitempty"`
	URI    string   `json:"uri,omitzero"`
}

type eventSubscriptionCreateAPI struct {
	Metadata       string                  `json:"metadata,omitzero"`
	Description    string                  `json:"description,omitzero"`
	Sources        []eventSourceReplaceAPI `json:"sources,omitzero"`
	DestinationIDs []string                `json:"destination_ids,omitzero"`
}

type eventSubscriptionUpdateAPI struct {
	ID             string                  `json:"id,omitzero"`
	Metadata       *string                 `json:"metadata,omitzero"`
	Description    *string                 `json:"description,omitzero"`
	Sources        []eventSourceReplaceAPI `json:"sources,omitzero"`
	DestinationIDs []string                `json:"destination_ids,omitzero"`
}

type refAPI struct {
	ID  string `json:"id,omitzero"`
	URI string `json:"uri,omitzero"`
}

type eventSubscriptionAPI struct {
	ID           string           `json:"id,omitzero"`
	URI          string           `json:"uri,omitzero"`
	CreatedAt    string           `json:"created_at,omitzero"`
	Metadata     string           `json:"metadata,omitzero"`
	Description  string           `json:"description,omitzero"`
	Sources      []eventSourceAPI `json:"sources,omitzero"`
	Destinations []refAPI         `json:"destinations,omitzero"`
}

type eventSubscriptionResourceModel struct {
	ID             types.String `tfsdk:"id"`
	Description    types.String `tfsdk:"description"`
	Metadata       types.String `tfsdk:"metadata"`
	Sources        types.List   `tfsdk:"sources"`
	DestinationIDs types.List   `tfsdk:"destination_ids"`
	URI            types.String `tfsdk:"uri"`
	CreatedAt      types.String `tfsdk:"created_at"`
}

type eventSubscriptionResource struct {
	client *ngrok.BaseClient
}

func NewEventSubscriptionResource() resource.Resource {
	return &eventSubscriptionResource{}
}

func (r *eventSubscriptionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_event_subscription"
}

func (r *eventSubscriptionResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	s := resource_event_subscription.EventSubscriptionResourceSchema(ctx)
	attrs := s.Attributes

	// Remove Ref nested objects not in hand-written model
	delete(attrs, "destinations")

	// Replace generated CustomType sources with standard ListNestedAttribute
	attrs["sources"] = schema.ListNestedAttribute{
		Description: "Sources containing the types for which this event subscription will trigger.",
		Required:    true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"type":   schema.StringAttribute{Description: "Type of event for which an event subscription will trigger.", Required: true},
				"fields": schema.ListAttribute{Description: "The fields to include in events for this source.", Optional: true, ElementType: types.StringType},
				"uri":    schema.StringAttribute{Description: "URI of the Event Source API resource.", Computed: true},
			},
		},
	}

	addStringPlanModifiers(attrs, "id", useStateForUnknownString())
	addStringPlanModifiers(attrs, "uri", useStateForUnknownString())
	addStringPlanModifiers(attrs, "created_at", useStateForUnknownString())
	addStringPlanModifiers(attrs, "description", useStateForUnknownString())
	addStringPlanModifiers(attrs, "metadata", useStateForUnknownString())

	resp.Schema = s
}

func (r *eventSubscriptionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = ngrok.NewBaseClient(clientConfig)
}

func (r *eventSubscriptionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan eventSubscriptionResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	sources := expandEventSourcesAPI(ctx, plan.Sources, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	var destIDs []types.String
	resp.Diagnostics.Append(plan.DestinationIDs.ElementsAs(ctx, &destIDs, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq := &eventSubscriptionCreateAPI{
		Description:    plan.Description.ValueString(),
		Metadata:       plan.Metadata.ValueString(),
		Sources:        sources,
		DestinationIDs: expandStringList(destIDs),
	}

	var sub eventSubscriptionAPI
	err := r.client.Do(ctx, "POST", &url.URL{Path: "/event_subscriptions"}, createReq, &sub)
	if err != nil {
		resp.Diagnostics.AddError("Error creating event subscription", err.Error())
		return
	}

	flattenEventSubscriptionAPI(ctx, &sub, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *eventSubscriptionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state eventSubscriptionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var sub eventSubscriptionAPI
	err := r.client.Do(ctx, "GET", &url.URL{Path: "/event_subscriptions/" + state.ID.ValueString()}, nil, &sub)
	if err != nil {
		if isNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading event subscription", err.Error())
		return
	}

	flattenEventSubscriptionAPI(ctx, &sub, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *eventSubscriptionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan eventSubscriptionResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state eventSubscriptionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	sources := expandEventSourcesAPI(ctx, plan.Sources, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	var destIDs []types.String
	resp.Diagnostics.Append(plan.DestinationIDs.ElementsAs(ctx, &destIDs, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq := &eventSubscriptionUpdateAPI{
		Description:    stringPtrFromFramework(plan.Description),
		Metadata:       stringPtrFromFramework(plan.Metadata),
		Sources:        sources,
		DestinationIDs: expandStringList(destIDs),
	}

	var sub eventSubscriptionAPI
	err := r.client.Do(ctx, "PATCH", &url.URL{Path: "/event_subscriptions/" + state.ID.ValueString()}, updateReq, &sub)
	if err != nil {
		resp.Diagnostics.AddError("Error updating event subscription", err.Error())
		return
	}

	flattenEventSubscriptionAPI(ctx, &sub, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *eventSubscriptionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state eventSubscriptionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Do(ctx, "DELETE", &url.URL{Path: "/event_subscriptions/" + state.ID.ValueString()}, nil, nil)
	if err != nil {
		if isNotFound(err) {
			return
		}
		resp.Diagnostics.AddError("Error deleting event subscription", err.Error())
	}
}

func (r *eventSubscriptionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

type eventSourceModel struct {
	Type   types.String `tfsdk:"type"`
	Fields types.List   `tfsdk:"fields"`
	URI    types.String `tfsdk:"uri"`
}

func eventSourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"type":   types.StringType,
		"fields": types.ListType{ElemType: types.StringType},
		"uri":    types.StringType,
	}
}

func expandEventSourcesAPI(ctx context.Context, list types.List, diags *diag.Diagnostics) []eventSourceReplaceAPI {
	if list.IsNull() || list.IsUnknown() {
		return nil
	}

	var models []eventSourceModel
	diags.Append(list.ElementsAs(ctx, &models, false)...)
	if diags.HasError() {
		return nil
	}

	sources := make([]eventSourceReplaceAPI, len(models))
	for i, m := range models {
		src := eventSourceReplaceAPI{
			Type: m.Type.ValueString(),
		}
		if !m.Fields.IsNull() && !m.Fields.IsUnknown() {
			var fields []types.String
			diags.Append(m.Fields.ElementsAs(ctx, &fields, false)...)
			src.Fields = expandStringList(fields)
		}
		sources[i] = src
	}
	return sources
}

func flattenEventSourcesAPI(ctx context.Context, sources []eventSourceAPI, diags *diag.Diagnostics) types.List {
	if sources == nil {
		return types.ListNull(types.ObjectType{AttrTypes: eventSourceAttrTypes()})
	}

	models := make([]eventSourceModel, len(sources))
	for i, s := range sources {
		var fields types.List
		if len(s.Fields) > 0 {
			f, d := types.ListValueFrom(ctx, types.StringType, s.Fields)
			diags.Append(d...)
			fields = f
		} else {
			fields = types.ListNull(types.StringType)
		}
		models[i] = eventSourceModel{
			Type:   types.StringValue(s.Type),
			Fields: fields,
			URI:    types.StringValue(s.URI),
		}
	}

	list, d := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: eventSourceAttrTypes()}, models)
	diags.Append(d...)
	return list
}

func flattenEventSubscriptionAPI(ctx context.Context, sub *eventSubscriptionAPI, model *eventSubscriptionResourceModel, diags *diag.Diagnostics) {
	model.ID = types.StringValue(sub.ID)
	model.Description = types.StringValue(sub.Description)
	model.Metadata = types.StringValue(sub.Metadata)
	model.URI = types.StringValue(sub.URI)
	model.CreatedAt = types.StringValue(sub.CreatedAt)
	model.Sources = flattenEventSourcesAPI(ctx, sub.Sources, diags)

	ids := make([]string, len(sub.Destinations))
	for i, d := range sub.Destinations {
		ids[i] = d.ID
	}
	destIDs, d := types.ListValueFrom(ctx, types.StringType, ids)
	diags.Append(d...)
	model.DestinationIDs = destIDs
}
