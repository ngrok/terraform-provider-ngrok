package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/agent_ingresses"
	"github.com/ngrok/terraform-provider-ngrok/internal/datasource_agent_ingress"
)

var _ datasource.DataSource = &agentIngressDataSource{}

type agentIngressDataSourceModel struct {
	ID                          types.String `tfsdk:"id"`
	Domain                      types.String `tfsdk:"domain"`
	Description                 types.String `tfsdk:"description"`
	Metadata                    types.String `tfsdk:"metadata"`
	NSTargets                   types.List   `tfsdk:"ns_targets"`
	RegionDomains               types.List   `tfsdk:"region_domains"`
	CertificateManagementPolicy types.Object `tfsdk:"certificate_management_policy"`
	URI                         types.String `tfsdk:"uri"`
	CreatedAt                   types.String `tfsdk:"created_at"`
}

type agentIngressDataSource struct {
	client *agent_ingresses.Client
}

func NewAgentIngressDataSource() datasource.DataSource {
	return &agentIngressDataSource{}
}

func (d *agentIngressDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_agent_ingress"
}

func (d *agentIngressDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_agent_ingress.AgentIngressDataSourceSchema(ctx)
	resp.Schema.Description = "Use this data source to look up an agent ingress by ID or domain."

	attrs := resp.Schema.Attributes

	attrs["id"] = schema.StringAttribute{
		Description: "Unique agent ingress resource identifier. Provide either id or domain.",
		Optional:    true,
		Computed:    true,
	}
	attrs["domain"] = schema.StringAttribute{
		Description: "The domain of the agent ingress. Provide either id or domain.",
		Optional:    true,
		Computed:    true,
	}

	delete(attrs, "certificate_management_status")

	attrs["certificate_management_policy"] = schema.SingleNestedAttribute{
		Description: "Configuration for automatic management of TLS certificates for this domain.",
		Computed:    true,
		Attributes: map[string]schema.Attribute{
			"authority": schema.StringAttribute{
				Description: "Certificate authority to request certificates from.",
				Computed:    true,
			},
			"private_key_type": schema.StringAttribute{
				Description: "Type of private key to use when requesting certificates.",
				Computed:    true,
			},
		},
	}
}

func (d *agentIngressDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	d.client = agent_ingresses.NewClient(clientConfig)
}

func (d *agentIngressDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config agentIngressDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var ingress *ngrok.AgentIngress

	if !config.ID.IsNull() && config.ID.ValueString() != "" {
		var err error
		ingress, err = d.client.Get(ctx, config.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Error reading agent ingress", err.Error())
			return
		}
	} else if !config.Domain.IsNull() && config.Domain.ValueString() != "" {
		filter := fmt.Sprintf(`obj.domain == %q`, config.Domain.ValueString())
		iter := d.client.List(&ngrok.FilteredPaging{
			Filter: &filter,
		})

		if !iter.Next(ctx) {
			if err := iter.Err(); err != nil {
				resp.Diagnostics.AddError("Error listing agent ingresses", err.Error())
				return
			}
			resp.Diagnostics.AddError(
				"Agent ingress not found",
				fmt.Sprintf("No agent ingress found with domain %q.", config.Domain.ValueString()),
			)
			return
		}
		ingress = iter.Item()

		if iter.Next(ctx) {
			resp.Diagnostics.AddError(
				"Multiple agent ingresses found",
				fmt.Sprintf("More than one agent ingress found with domain %q. Use id instead.", config.Domain.ValueString()),
			)
			return
		}
		if err := iter.Err(); err != nil {
			resp.Diagnostics.AddError("Error listing agent ingresses", err.Error())
			return
		}
	} else {
		resp.Diagnostics.AddError(
			"Missing lookup attribute",
			"Either id or domain must be specified.",
		)
		return
	}

	var model agentIngressDataSourceModel
	flattenAgentIngressDataSource(ctx, ingress, &model, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func flattenAgentIngressDataSource(ctx context.Context, ingress *ngrok.AgentIngress, model *agentIngressDataSourceModel, diags *diag.Diagnostics) {
	model.ID = types.StringValue(ingress.ID)
	model.Domain = types.StringValue(ingress.Domain)
	model.Description = types.StringValue(ingress.Description)
	model.Metadata = types.StringValue(ingress.Metadata)
	model.URI = types.StringValue(ingress.URI)
	model.CreatedAt = types.StringValue(ingress.CreatedAt)

	nsTargets, d := types.ListValueFrom(ctx, types.StringType, ingress.NSTargets)
	diags.Append(d...)
	model.NSTargets = nsTargets

	regionDomains, d := types.ListValueFrom(ctx, types.StringType, ingress.RegionDomains)
	diags.Append(d...)
	model.RegionDomains = regionDomains

	model.CertificateManagementPolicy = flattenAgentIngressCertPolicy(ctx, ingress.CertificateManagementPolicy, diags)
}
