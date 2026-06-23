package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/tunnel_sessions"
	"github.com/ngrok/terraform-provider-ngrok/internal/datasource_tunnel_session"
)

var _ datasource.DataSource = &tunnelSessionDataSource{}

type tunnelSessionDataSourceModel struct {
	ID           types.String `tfsdk:"id"`
	AgentVersion types.String `tfsdk:"agent_version"`
	CredentialID types.String `tfsdk:"credential_id"`
	IP           types.String `tfsdk:"ip"`
	Metadata     types.String `tfsdk:"metadata"`
	OS           types.String `tfsdk:"os"`
	Region       types.String `tfsdk:"region"`
	StartedAt    types.String `tfsdk:"started_at"`
	Transport    types.String `tfsdk:"transport"`
	URI          types.String `tfsdk:"uri"`
}

type tunnelSessionDataSource struct {
	client *tunnel_sessions.Client
}

func NewTunnelSessionDataSource() datasource.DataSource {
	return &tunnelSessionDataSource{}
}

func (d *tunnelSessionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tunnel_session"
}

func (d *tunnelSessionDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	s := datasource_tunnel_session.TunnelSessionDataSourceSchema(ctx)
	attrs := s.Attributes
	// Delete Ref nested object
	delete(attrs, "credential")
	// Add flat _id field for credential Ref
	attrs["credential_id"] = schema.StringAttribute{
		Computed:    true,
		Description: "The ID of the credential used to start this tunnel session.",
	}
	resp.Schema = s
}

func (d *tunnelSessionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	d.client = tunnel_sessions.NewClient(clientConfig)
}

func (d *tunnelSessionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config tunnelSessionDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	session, err := d.client.Get(ctx, config.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading tunnel session", err.Error())
		return
	}

	var model tunnelSessionDataSourceModel
	flattenTunnelSessionDataSource(session, &model)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func flattenTunnelSessionDataSource(s *ngrok.TunnelSession, model *tunnelSessionDataSourceModel) {
	model.ID = types.StringValue(s.ID)
	model.AgentVersion = types.StringValue(s.AgentVersion)
	model.CredentialID = types.StringValue(s.Credential.ID)
	model.IP = types.StringValue(s.IP)
	model.Metadata = types.StringValue(s.Metadata)
	model.OS = types.StringValue(s.OS)
	model.Region = types.StringValue(s.Region)
	model.StartedAt = types.StringValue(s.StartedAt)
	model.Transport = types.StringValue(s.Transport)
	model.URI = types.StringValue(s.URI)
}
