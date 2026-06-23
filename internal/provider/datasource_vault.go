package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/vaults"
	"github.com/ngrok/terraform-provider-ngrok/internal/datasource_vault"
)

var _ datasource.DataSource = &vaultDataSource{}

type vaultDataSourceModel struct {
	ID            types.String `tfsdk:"id"`
	URI           types.String `tfsdk:"uri"`
	CreatedAt     types.String `tfsdk:"created_at"`
	UpdatedAt     types.String `tfsdk:"updated_at"`
	Name          types.String `tfsdk:"name"`
	Description   types.String `tfsdk:"description"`
	Metadata      types.String `tfsdk:"metadata"`
	CreatedBy     types.String `tfsdk:"created_by"`
	LastUpdatedBy types.String `tfsdk:"last_updated_by"`
}

type vaultDataSource struct {
	client *vaults.Client
}

func NewVaultDataSource() datasource.DataSource {
	return &vaultDataSource{}
}

func (d *vaultDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vault"
}

func (d *vaultDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	s := datasource_vault.VaultDataSourceSchema(ctx)
	attrs := s.Attributes
	idAttr := attrs["id"].(schema.StringAttribute)
	idAttr.Required = false
	idAttr.Optional = true
	idAttr.Computed = true
	attrs["id"] = idAttr
	nameAttr := attrs["name"].(schema.StringAttribute)
	nameAttr.Optional = true
	attrs["name"] = nameAttr
	resp.Schema = s
}

func (d *vaultDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	d.client = vaults.NewClient(clientConfig)
}

func (d *vaultDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config vaultDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var vault *ngrok.Vault

	if !config.ID.IsNull() && config.ID.ValueString() != "" {
		var err error
		vault, err = d.client.Get(ctx, config.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Error reading vault", err.Error())
			return
		}
	} else if !config.Name.IsNull() && config.Name.ValueString() != "" {
		filter := fmt.Sprintf(`obj.name == %q`, config.Name.ValueString())
		iter := d.client.List(&ngrok.FilteredPaging{
			Filter: &filter,
		})

		if !iter.Next(ctx) {
			if err := iter.Err(); err != nil {
				resp.Diagnostics.AddError("Error listing vaults", err.Error())
				return
			}
			resp.Diagnostics.AddError(
				"Vault not found",
				fmt.Sprintf("No vault found with name %q.", config.Name.ValueString()),
			)
			return
		}
		vault = iter.Item()

		if iter.Next(ctx) {
			resp.Diagnostics.AddError(
				"Multiple vaults found",
				fmt.Sprintf("More than one vault found with name %q. Use id instead.", config.Name.ValueString()),
			)
			return
		}
		if err := iter.Err(); err != nil {
			resp.Diagnostics.AddError("Error listing vaults", err.Error())
			return
		}
	} else {
		resp.Diagnostics.AddError(
			"Missing lookup attribute",
			"Either id or name must be specified.",
		)
		return
	}

	var model vaultDataSourceModel
	model.ID = types.StringValue(vault.ID)
	model.URI = types.StringValue(vault.URI)
	model.CreatedAt = types.StringValue(vault.CreatedAt)
	model.UpdatedAt = types.StringValue(vault.UpdatedAt)
	model.Name = types.StringValue(vault.Name)
	model.Description = types.StringValue(vault.Description)
	model.Metadata = types.StringValue(vault.Metadata)
	model.CreatedBy = types.StringValue(vault.CreatedBy)
	model.LastUpdatedBy = types.StringValue(vault.LastUpdatedBy)

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
