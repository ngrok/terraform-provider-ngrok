package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/ip_restrictions"
	"github.com/ngrok/terraform-provider-ngrok/internal/datasource_ip_restriction"
)

var _ datasource.DataSource = &ipRestrictionDataSource{}

type ipRestrictionDataSourceModel struct {
	ID          types.String   `tfsdk:"id"`
	URI         types.String   `tfsdk:"uri"`
	CreatedAt   types.String   `tfsdk:"created_at"`
	Description types.String   `tfsdk:"description"`
	Metadata    types.String   `tfsdk:"metadata"`
	Enforced    types.Bool     `tfsdk:"enforced"`
	Type        types.String   `tfsdk:"type"`
	IPPolicyIDs []types.String `tfsdk:"ip_policy_ids"`
}

type ipRestrictionDataSource struct {
	client *ip_restrictions.Client
}

func NewIPRestrictionDataSource() datasource.DataSource {
	return &ipRestrictionDataSource{}
}

func (d *ipRestrictionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_restriction"
}

func (d *ipRestrictionDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	s := datasource_ip_restriction.IpRestrictionDataSourceSchema(ctx)
	attrs := s.Attributes
	// Delete ListNestedAttribute of Ref objects
	delete(attrs, "ip_policies")
	// Add flat _ids list field for ip_policies Ref
	attrs["ip_policy_ids"] = schema.ListAttribute{
		Computed:    true,
		Description: "The set of IP policy identifiers that are used to enforce the restriction.",
		ElementType: types.StringType,
	}
	resp.Schema = s
}

func (d *ipRestrictionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	d.client = ip_restrictions.NewClient(clientConfig)
}

func (d *ipRestrictionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ipRestrictionDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	restriction, err := d.client.Get(ctx, config.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading IP restriction", err.Error())
		return
	}

	var model ipRestrictionDataSourceModel
	model.ID = types.StringValue(restriction.ID)
	model.URI = types.StringValue(restriction.URI)
	model.CreatedAt = types.StringValue(restriction.CreatedAt)
	model.Description = types.StringValue(restriction.Description)
	model.Metadata = types.StringValue(restriction.Metadata)
	model.Enforced = types.BoolValue(restriction.Enforced)
	model.Type = types.StringValue(restriction.Type)
	model.IPPolicyIDs = flattenRefList(restriction.IPPolicies)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
