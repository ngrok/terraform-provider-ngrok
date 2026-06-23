package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/event_destinations"
	"github.com/ngrok/terraform-provider-ngrok/internal/datasource_event_destination"
)

var _ datasource.DataSource = &eventDestinationDataSource{}

type eventDestinationDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	Description types.String `tfsdk:"description"`
	Metadata    types.String `tfsdk:"metadata"`
	Format      types.String `tfsdk:"format"`
	Target      types.Object `tfsdk:"target"`
	URI         types.String `tfsdk:"uri"`
	CreatedAt   types.String `tfsdk:"created_at"`
}

type eventDestinationDataSource struct {
	client *event_destinations.Client
}

func NewEventDestinationDataSource() datasource.DataSource {
	return &eventDestinationDataSource{}
}

func (d *eventDestinationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_event_destination"
}

func awsAuthDataSourceSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Description: "Configuration for how to authenticate into your AWS account.",
		Computed:    true,
		Attributes: map[string]schema.Attribute{
			"role": schema.SingleNestedAttribute{
				Description: "A role for ngrok to assume on your behalf.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"role_arn": schema.StringAttribute{Description: "An ARN that specifies the role.", Computed: true},
				},
			},
			"creds": schema.SingleNestedAttribute{
				Description: "Credentials to your AWS account.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"aws_access_key_id":     schema.StringAttribute{Description: "The ID portion of an AWS access key.", Computed: true},
					"aws_secret_access_key": schema.StringAttribute{Description: "The secret portion of an AWS access key.", Computed: true, Sensitive: true},
				},
			},
		},
	}
}

func (d *eventDestinationDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_event_destination.EventDestinationDataSourceSchema(ctx)
	resp.Schema.Description = "Use this data source to look up an event destination by ID."

	attrs := resp.Schema.Attributes

	// Make id required for lookup
	attrs["id"] = schema.StringAttribute{
		Description: "Unique event destination resource identifier.",
		Required:    true,
	}

	// Add target attribute (excluded from codegen due to duplicate CustomType collision)
	attrs["target"] = schema.SingleNestedAttribute{
		Description: "An object that encapsulates where and how to send your events.",
		Computed:    true,
		Attributes: map[string]schema.Attribute{
			"firehose": schema.SingleNestedAttribute{
				Description: "Configuration used to send events to Amazon Kinesis Data Firehose.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"auth":                awsAuthDataSourceSchema(),
					"delivery_stream_arn": schema.StringAttribute{Description: "Firehose delivery stream ARN.", Computed: true},
				},
			},
			"kinesis": schema.SingleNestedAttribute{
				Description: "Configuration used to send events to Amazon Kinesis.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"auth":       awsAuthDataSourceSchema(),
					"stream_arn": schema.StringAttribute{Description: "Kinesis stream ARN.", Computed: true},
				},
			},
			"cloudwatch_logs": schema.SingleNestedAttribute{
				Description: "Configuration used to send events to Amazon CloudWatch Logs.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"auth":          awsAuthDataSourceSchema(),
					"log_group_arn": schema.StringAttribute{Description: "CloudWatch Logs group ARN.", Computed: true},
				},
			},
			"datadog": schema.SingleNestedAttribute{
				Description: "Configuration used to send events to Datadog.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"api_key": schema.StringAttribute{Description: "Datadog API key.", Computed: true, Sensitive: true},
					"ddtags":  schema.StringAttribute{Description: "Tags to send with the event.", Computed: true},
					"service": schema.StringAttribute{Description: "Service name.", Computed: true},
					"ddsite":  schema.StringAttribute{Description: "Datadog site.", Computed: true},
				},
			},
			"azure_logs_ingestion": schema.SingleNestedAttribute{
				Description: "Configuration used to send events to Azure Logs Ingestion.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"tenant_id":                   schema.StringAttribute{Description: "Tenant ID for the Azure account.", Computed: true},
					"client_id":                   schema.StringAttribute{Description: "Client ID.", Computed: true},
					"client_secret":               schema.StringAttribute{Description: "Client Secret.", Computed: true, Sensitive: true},
					"logs_ingestion_uri":          schema.StringAttribute{Description: "Logs ingestion URI.", Computed: true},
					"data_collection_rule_id":     schema.StringAttribute{Description: "Data collection rule ID.", Computed: true},
					"data_collection_stream_name": schema.StringAttribute{Description: "Data collection stream name.", Computed: true},
				},
			},
		},
	}
}

func (d *eventDestinationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	d.client = event_destinations.NewClient(clientConfig)
}

func (d *eventDestinationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config eventDestinationDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dest, err := d.client.Get(ctx, config.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading event destination", err.Error())
		return
	}

	var model eventDestinationDataSourceModel
	flattenEventDestinationDataSource(ctx, dest, &model, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func flattenEventDestinationDataSource(ctx context.Context, dest *ngrok.EventDestination, model *eventDestinationDataSourceModel, diags *diag.Diagnostics) {
	model.ID = types.StringValue(dest.ID)
	model.Description = types.StringValue(dest.Description)
	model.Metadata = types.StringValue(dest.Metadata)
	model.Format = types.StringValue(dest.Format)
	model.URI = types.StringValue(dest.URI)
	model.CreatedAt = types.StringValue(dest.CreatedAt)
	model.Target = flattenEventTarget(ctx, dest.Target, types.ObjectNull(eventTargetAttrTypes()), diags)
}
