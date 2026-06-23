package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/ip_policy_rules"
	"github.com/ngrok/terraform-provider-ngrok/internal/datasource_ip_policy_rule"
)

var _ datasource.DataSource = &ipPolicyRuleDataSource{}

type ipPolicyRuleDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	URI         types.String `tfsdk:"uri"`
	CreatedAt   types.String `tfsdk:"created_at"`
	Description types.String `tfsdk:"description"`
	Metadata    types.String `tfsdk:"metadata"`
	CIDR        types.String `tfsdk:"cidr"`
	IPPolicyID  types.String `tfsdk:"ip_policy_id"`
	Action      types.String `tfsdk:"action"`
}

type ipPolicyRuleDataSource struct {
	client *ip_policy_rules.Client
}

func NewIPPolicyRuleDataSource() datasource.DataSource {
	return &ipPolicyRuleDataSource{}
}

func (d *ipPolicyRuleDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_policy_rule"
}

func (d *ipPolicyRuleDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	s := datasource_ip_policy_rule.IpPolicyRuleDataSourceSchema(ctx)
	attrs := s.Attributes
	// Delete Ref nested object
	delete(attrs, "ip_policy")
	// Add flat _id field for ip_policy Ref
	attrs["ip_policy_id"] = schema.StringAttribute{
		Computed:    true,
		Description: "ID of the IP policy this rule belongs to.",
	}
	resp.Schema = s
}

func (d *ipPolicyRuleDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	d.client = ip_policy_rules.NewClient(clientConfig)
}

func (d *ipPolicyRuleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ipPolicyRuleDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	rule, err := d.client.Get(ctx, config.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading IP policy rule", err.Error())
		return
	}

	var model ipPolicyRuleDataSourceModel
	model.ID = types.StringValue(rule.ID)
	model.URI = types.StringValue(rule.URI)
	model.CreatedAt = types.StringValue(rule.CreatedAt)
	model.Description = types.StringValue(rule.Description)
	model.Metadata = types.StringValue(rule.Metadata)
	model.CIDR = types.StringValue(rule.CIDR)
	model.IPPolicyID = types.StringValue(rule.IPPolicy.ID)
	model.Action = types.StringValue(rule.Action)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
