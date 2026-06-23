package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/credentials"
	"github.com/ngrok/terraform-provider-ngrok/internal/datasource_credential"
)

var _ datasource.DataSource = &credentialDataSource{}

type credentialDataSourceModel struct {
	ID          types.String   `tfsdk:"id"`
	URI         types.String   `tfsdk:"uri"`
	CreatedAt   types.String   `tfsdk:"created_at"`
	Description types.String   `tfsdk:"description"`
	Metadata    types.String   `tfsdk:"metadata"`
	ACL         []types.String `tfsdk:"acl"`
	OwnerID     types.String   `tfsdk:"owner_id"`
}

type credentialDataSource struct {
	client *credentials.Client
}

func NewCredentialDataSource() datasource.DataSource {
	return &credentialDataSource{}
}

func (d *credentialDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_credential"
}

func (d *credentialDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	s := datasource_credential.CredentialDataSourceSchema(ctx)
	attrs := s.Attributes
	delete(attrs, "token")
	resp.Schema = s
}

func (d *credentialDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	d.client = credentials.NewClient(clientConfig)
}

func (d *credentialDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config credentialDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	cred, err := d.client.Get(ctx, config.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading credential", err.Error())
		return
	}

	var model credentialDataSourceModel
	model.ID = types.StringValue(cred.ID)
	model.URI = types.StringValue(cred.URI)
	model.CreatedAt = types.StringValue(cred.CreatedAt)
	model.Description = types.StringValue(cred.Description)
	model.Metadata = types.StringValue(cred.Metadata)
	model.ACL = flattenStringList(cred.ACL)

	if cred.OwnerID != nil {
		model.OwnerID = types.StringValue(*cred.OwnerID)
	} else {
		model.OwnerID = types.StringNull()
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
