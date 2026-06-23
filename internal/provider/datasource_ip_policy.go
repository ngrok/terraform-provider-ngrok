package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/ip_policies"
	"github.com/ngrok/terraform-provider-ngrok/internal/datasource_ip_policy"
)

var _ datasource.DataSource = &ipPolicyDataSource{}

type ipPolicyDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	URI         types.String `tfsdk:"uri"`
	CreatedAt   types.String `tfsdk:"created_at"`
	Description types.String `tfsdk:"description"`
	Metadata    types.String `tfsdk:"metadata"`
}

type ipPolicyDataSource struct {
	client *ip_policies.Client
}

func NewIPPolicyDataSource() datasource.DataSource {
	return &ipPolicyDataSource{}
}

func (d *ipPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_policy"
}

func (d *ipPolicyDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_ip_policy.IpPolicyDataSourceSchema(ctx)
}

func (d *ipPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	d.client = ip_policies.NewClient(clientConfig)
}

func (d *ipPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ipPolicyDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	policy, err := d.client.Get(ctx, config.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading IP policy", err.Error())
		return
	}

	var model ipPolicyDataSourceModel
	model.ID = types.StringValue(policy.ID)
	model.URI = types.StringValue(policy.URI)
	model.CreatedAt = types.StringValue(policy.CreatedAt)
	model.Description = types.StringValue(policy.Description)
	model.Metadata = types.StringValue(policy.Metadata)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
