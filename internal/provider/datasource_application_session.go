package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/application_sessions"
	"github.com/ngrok/terraform-provider-ngrok/internal/datasource_application_session"
)

var _ datasource.DataSource = &applicationSessionDataSource{}

type applicationSessionDataSourceModel struct {
	ID                types.String `tfsdk:"id"`
	URI               types.String `tfsdk:"uri"`
	PublicURL         types.String `tfsdk:"public_url"`
	ApplicationUserID types.String `tfsdk:"application_user_id"`
	CreatedAt         types.String `tfsdk:"created_at"`
	LastActive        types.String `tfsdk:"last_active"`
	ExpiresAt         types.String `tfsdk:"expires_at"`
	EndpointID        types.String `tfsdk:"endpoint_id"`
	EdgeID            types.String `tfsdk:"edge_id"`
	RouteID           types.String `tfsdk:"route_id"`
}

type applicationSessionDataSource struct {
	client *application_sessions.Client
}

func NewApplicationSessionDataSource() datasource.DataSource {
	return &applicationSessionDataSource{}
}

func (d *applicationSessionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_session"
}

func (d *applicationSessionDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	s := datasource_application_session.ApplicationSessionDataSourceSchema(ctx)
	attrs := s.Attributes
	// Delete Ref nested objects and complex nested objects not in hand-written model
	delete(attrs, "application_user")
	delete(attrs, "endpoint")
	delete(attrs, "edge")
	delete(attrs, "route")
	delete(attrs, "browser_session")
	// Add flat _id fields for Ref objects
	attrs["application_user_id"] = schema.StringAttribute{
		Computed:    true,
		Description: "The ID of the application user associated with this session.",
	}
	attrs["endpoint_id"] = schema.StringAttribute{
		Computed:    true,
		Description: "The ID of the endpoint associated with this session.",
	}
	attrs["edge_id"] = schema.StringAttribute{
		Computed:    true,
		Description: "The ID of the edge associated with this session.",
	}
	attrs["route_id"] = schema.StringAttribute{
		Computed:    true,
		Description: "The ID of the route associated with this session.",
	}
	resp.Schema = s
}

func (d *applicationSessionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	d.client = application_sessions.NewClient(clientConfig)
}

func (d *applicationSessionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config applicationSessionDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	session, err := d.client.Get(ctx, config.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading application session", err.Error())
		return
	}

	var model applicationSessionDataSourceModel
	flattenApplicationSessionDataSource(session, &model)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func flattenApplicationSessionDataSource(s *ngrok.ApplicationSession, model *applicationSessionDataSourceModel) {
	model.ID = types.StringValue(s.ID)
	model.URI = types.StringValue(s.URI)
	model.PublicURL = types.StringValue(s.PublicURL)
	model.ApplicationUserID = types.StringValue(flattenRef(s.ApplicationUser))
	model.CreatedAt = types.StringValue(s.CreatedAt)
	model.LastActive = types.StringValue(s.LastActive)
	model.ExpiresAt = types.StringValue(s.ExpiresAt)
	model.EndpointID = types.StringValue(flattenRef(s.Endpoint))
	model.EdgeID = types.StringValue(flattenRef(s.Edge))
	model.RouteID = types.StringValue(flattenRef(s.Route))
}
