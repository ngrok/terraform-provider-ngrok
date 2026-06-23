package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/application_users"
	"github.com/ngrok/terraform-provider-ngrok/internal/datasource_application_user"
)

var _ datasource.DataSource = &applicationUserDataSource{}

type applicationUserDataSourceModel struct {
	ID             types.String `tfsdk:"id"`
	URI            types.String `tfsdk:"uri"`
	ProviderUserID types.String `tfsdk:"provider_user_id"`
	Username       types.String `tfsdk:"username"`
	Email          types.String `tfsdk:"email"`
	Name           types.String `tfsdk:"name"`
	CreatedAt      types.String `tfsdk:"created_at"`
	LastActive     types.String `tfsdk:"last_active"`
	LastLogin      types.String `tfsdk:"last_login"`
}

type applicationUserDataSource struct {
	client *application_users.Client
}

func NewApplicationUserDataSource() datasource.DataSource {
	return &applicationUserDataSource{}
}

func (d *applicationUserDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_user"
}

func (d *applicationUserDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	s := datasource_application_user.ApplicationUserDataSourceSchema(ctx)
	attrs := s.Attributes
	delete(attrs, "identity_provider")
	resp.Schema = s
}

func (d *applicationUserDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	d.client = application_users.NewClient(clientConfig)
}

func (d *applicationUserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config applicationUserDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	user, err := d.client.Get(ctx, config.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading application user", err.Error())
		return
	}

	var model applicationUserDataSourceModel
	flattenApplicationUserDataSource(user, &model)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func flattenApplicationUserDataSource(u *ngrok.ApplicationUser, model *applicationUserDataSourceModel) {
	model.ID = types.StringValue(u.ID)
	model.URI = types.StringValue(u.URI)
	model.ProviderUserID = types.StringValue(u.ProviderUserID)
	model.Username = types.StringValue(u.Username)
	model.Email = types.StringValue(u.Email)
	model.Name = types.StringValue(u.Name)
	model.CreatedAt = types.StringValue(u.CreatedAt)
	model.LastActive = types.StringValue(u.LastActive)
	model.LastLogin = types.StringValue(u.LastLogin)
}
