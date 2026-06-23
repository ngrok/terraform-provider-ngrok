package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"net/url"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/terraform-provider-ngrok/internal/datasource_event_subscription"
)

var _ datasource.DataSource = &eventSubscriptionDataSource{}

type eventSubscriptionDataSourceModel struct {
	ID             types.String `tfsdk:"id"`
	Description    types.String `tfsdk:"description"`
	Metadata       types.String `tfsdk:"metadata"`
	Sources        types.List   `tfsdk:"sources"`
	DestinationIDs types.List   `tfsdk:"destination_ids"`
	URI            types.String `tfsdk:"uri"`
	CreatedAt      types.String `tfsdk:"created_at"`
}

type eventSubscriptionDataSource struct {
	client *ngrok.BaseClient
}

func NewEventSubscriptionDataSource() datasource.DataSource {
	return &eventSubscriptionDataSource{}
}

func (d *eventSubscriptionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_event_subscription"
}

func (d *eventSubscriptionDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_event_subscription.EventSubscriptionDataSourceSchema(ctx)
	resp.Schema.Description = "Use this data source to look up an event subscription by ID."

	attrs := resp.Schema.Attributes

	delete(attrs, "destinations")
	attrs["destination_ids"] = schema.ListAttribute{
		Description: "A list of Event Destination IDs which should be used for this Event Subscription.",
		Computed:    true,
		ElementType: types.StringType,
	}

	attrs["sources"] = schema.ListNestedAttribute{
		Description: "Sources containing the types for which this event subscription will trigger.",
		Computed:    true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"type": schema.StringAttribute{
					Description: "Type of event for which an event subscription will trigger.",
					Computed:    true,
				},
				"fields": schema.ListAttribute{
					Description: "The fields included in events for this source.",
					Computed:    true,
					ElementType: types.StringType,
				},
				"uri": schema.StringAttribute{
					Description: "URI of the Event Source API resource.",
					Computed:    true,
				},
			},
		},
	}
}

func (d *eventSubscriptionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clientConfig, ok := req.ProviderData.(*ngrok.ClientConfig)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *ngrok.ClientConfig, got: %T.", req.ProviderData),
		)
		return
	}
	d.client = ngrok.NewBaseClient(clientConfig)
}

func (d *eventSubscriptionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config eventSubscriptionDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var sub eventSubscriptionAPI
	err := d.client.Do(ctx, "GET", &url.URL{Path: "/event_subscriptions/" + config.ID.ValueString()}, nil, &sub)
	if err != nil {
		resp.Diagnostics.AddError("Error reading event subscription", err.Error())
		return
	}

	var model eventSubscriptionDataSourceModel
	flattenEventSubscriptionDataSource(ctx, &sub, &model, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func flattenEventSubscriptionDataSource(ctx context.Context, sub *eventSubscriptionAPI, model *eventSubscriptionDataSourceModel, diags *diag.Diagnostics) {
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
	destIDs, dd := types.ListValueFrom(ctx, types.StringType, ids)
	diags.Append(dd...)
	model.DestinationIDs = destIDs
}
