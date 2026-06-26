package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/kubernetes_operators"
)

var _ datasource.DataSource = &kubernetesOperatorDataSource{}

type kubernetesOperatorDataSourceModel struct {
	ID              types.String   `tfsdk:"id"`
	URI             types.String   `tfsdk:"uri"`
	CreatedAt       types.String   `tfsdk:"created_at"`
	UpdatedAt       types.String   `tfsdk:"updated_at"`
	Description     types.String   `tfsdk:"description"`
	Metadata        types.String   `tfsdk:"metadata"`
	EnabledFeatures []types.String `tfsdk:"enabled_features"`
	Region          types.String   `tfsdk:"region"`
	Deployment      types.Object   `tfsdk:"deployment"`
}

type kubernetesOperatorDataSource struct {
	client *kubernetes_operators.Client
}

func NewKubernetesOperatorDataSource() datasource.DataSource {
	return &kubernetesOperatorDataSource{}
}

func (d *kubernetesOperatorDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kubernetes_operator"
}

func (d *kubernetesOperatorDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Use this data source to look up a Kubernetes Operator by ID.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier for this Kubernetes Operator.",
				Required:    true,
			},
			"uri": schema.StringAttribute{
				Description: "URI of this Kubernetes Operator API resource.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "Timestamp when the Kubernetes Operator was created. RFC 3339 format.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Description: "Timestamp when the Kubernetes Operator was last updated. RFC 3339 format.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "Human-readable description of this Kubernetes Operator.",
				Computed:    true,
			},
			"metadata": schema.StringAttribute{
				Description: "Arbitrary user-defined machine-readable data of this Kubernetes Operator.",
				Computed:    true,
			},
			"enabled_features": schema.ListAttribute{
				Description: "Features enabled for this Kubernetes Operator (e.g. bindings, ingress, gateway).",
				Computed:    true,
				ElementType: types.StringType,
			},
			"region": schema.StringAttribute{
				Description: "The ngrok region in which the ingress for this operator is served.",
				Computed:    true,
			},
			"deployment": schema.SingleNestedAttribute{
				Description: "Information about the deployment of this Kubernetes Operator.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description: "The deployment name.",
						Computed:    true,
					},
					"namespace": schema.StringAttribute{
						Description: "The namespace this Kubernetes Operator is deployed to.",
						Computed:    true,
					},
					"version": schema.StringAttribute{
						Description: "The version of this Kubernetes Operator.",
						Computed:    true,
					},
					"cluster_name": schema.StringAttribute{
						Description: "User-given name for the cluster the Kubernetes Operator is deployed to.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *kubernetesOperatorDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	d.client = kubernetes_operators.NewClient(clientConfig)
}

func (d *kubernetesOperatorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config kubernetesOperatorDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	op, err := d.client.Get(ctx, config.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading Kubernetes Operator", err.Error())
		return
	}

	var model kubernetesOperatorDataSourceModel
	model.ID = types.StringValue(op.ID)
	model.URI = types.StringValue(op.URI)
	model.CreatedAt = types.StringValue(op.CreatedAt)
	model.UpdatedAt = types.StringValue(op.UpdatedAt)
	model.Description = types.StringValue(op.Description)
	model.Metadata = types.StringValue(op.Metadata)
	model.EnabledFeatures = flattenStringList(op.EnabledFeatures)
	model.Region = types.StringValue(op.Region)

	deploymentAttrs := map[string]types.String{
		"name":         types.StringValue(op.Deployment.Name),
		"namespace":    types.StringValue(op.Deployment.Namespace),
		"version":      types.StringValue(op.Deployment.Version),
		"cluster_name": types.StringValue(op.Deployment.ClusterName),
	}
	deploymentObj, diags := types.ObjectValueFrom(ctx, map[string]attr.Type{
		"name":         types.StringType,
		"namespace":    types.StringType,
		"version":      types.StringType,
		"cluster_name": types.StringType,
	}, deploymentAttrs)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	model.Deployment = deploymentObj

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
