package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/endpoints"
	"github.com/ngrok/terraform-provider-ngrok/internal/datasource_cloud_endpoint"
)

var _ datasource.DataSource = &cloudEndpointDataSource{}

type cloudEndpointDataSourceModel struct {
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

type cloudEndpointDataSource struct {
	client *endpoints.Client
}

func NewCloudEndpointDataSource() datasource.DataSource {
	return &cloudEndpointDataSource{}
}

func (d *cloudEndpointDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloud_endpoint"
}

func (d *cloudEndpointDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	s := datasource_cloud_endpoint.CloudEndpointDataSourceSchema(ctx)
	attrs := s.Attributes
	// Delete Ref nested objects and extra OpenAPI fields not in hand-written model
	delete(attrs, "domain")
	delete(attrs, "edge")
	delete(attrs, "tunnel")
	delete(attrs, "tunnel_session")
	delete(attrs, "tcp_addr")
	delete(attrs, "principal")
	delete(attrs, "host")
	delete(attrs, "hostport")
	delete(attrs, "name")
	delete(attrs, "port")
	delete(attrs, "proto")
	delete(attrs, "public_url")
	delete(attrs, "scheme")
	delete(attrs, "upstream_protocol")
	delete(attrs, "upstream_url")
	// Add flat _id field for domain Ref
	attrs["domain_id"] = schema.StringAttribute{
		Computed:    true,
		Description: "ID of the domain reserved for this endpoint.",
	}
	resp.Schema = s
}

func (d *cloudEndpointDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	d.client = endpoints.NewClient(clientConfig)
}

func (d *cloudEndpointDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config cloudEndpointDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	endpoint, err := d.client.Get(ctx, config.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading cloud endpoint", err.Error())
		return
	}

	var model cloudEndpointDataSourceModel
	flattenCloudEndpointDataSource(endpoint, &model)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func flattenCloudEndpointDataSource(endpoint *ngrok.Endpoint, model *cloudEndpointDataSourceModel) {
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
