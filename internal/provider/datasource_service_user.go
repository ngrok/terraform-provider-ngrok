package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/bot_users"
	"github.com/ngrok/terraform-provider-ngrok/internal/datasource_service_user"
)

var _ datasource.DataSource = &serviceUserDataSource{}

type serviceUserDataSourceModel struct {
	ID        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Active    types.Bool   `tfsdk:"active"`
	URI       types.String `tfsdk:"uri"`
	CreatedAt types.String `tfsdk:"created_at"`
}

type serviceUserDataSource struct {
	client *bot_users.Client
}

func NewServiceUserDataSource() datasource.DataSource {
	return &serviceUserDataSource{}
}

func (d *serviceUserDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_user"
}

func (d *serviceUserDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_service_user.ServiceUserDataSourceSchema(ctx)
}

func (d *serviceUserDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	d.client = bot_users.NewClient(clientConfig)
}

func (d *serviceUserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config serviceUserDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	user, err := d.client.Get(ctx, config.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading service user", err.Error())
		return
	}

	var model serviceUserDataSourceModel
	model.ID = types.StringValue(user.ID)
	model.Name = types.StringValue(user.Name)
	model.Active = types.BoolValue(user.Active)
	model.URI = types.StringValue(user.URI)
	model.CreatedAt = types.StringValue(user.CreatedAt)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
