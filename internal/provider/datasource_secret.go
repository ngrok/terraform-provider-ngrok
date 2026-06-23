package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/secrets"
	"github.com/ngrok/terraform-provider-ngrok/internal/datasource_secret"
)

var _ datasource.DataSource = &secretDataSource{}

type secretDataSourceModel struct {
	ID              types.String `tfsdk:"id"`
	URI             types.String `tfsdk:"uri"`
	CreatedAt       types.String `tfsdk:"created_at"`
	UpdatedAt       types.String `tfsdk:"updated_at"`
	Name            types.String `tfsdk:"name"`
	Description     types.String `tfsdk:"description"`
	Metadata        types.String `tfsdk:"metadata"`
	VaultID         types.String `tfsdk:"vault_id"`
	VaultName       types.String `tfsdk:"vault_name"`
	CreatedByID     types.String `tfsdk:"created_by_id"`
	LastUpdatedByID types.String `tfsdk:"last_updated_by_id"`
}

type secretDataSource struct {
	client *secrets.Client
}

func NewSecretDataSource() datasource.DataSource {
	return &secretDataSource{}
}

func (d *secretDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_secret"
}

func (d *secretDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_secret.SecretDataSourceSchema(ctx)
	resp.Schema.Description = "Use this data source to look up a secret by ID or name."

	attrs := resp.Schema.Attributes

	attrs["id"] = schema.StringAttribute{
		Description: "Unique secret resource identifier. Provide either id or name.",
		Optional:    true,
		Computed:    true,
	}
	attrs["name"] = schema.StringAttribute{
		Description: "Human-readable name of the secret. Provide either id or name.",
		Optional:    true,
		Computed:    true,
	}

	delete(attrs, "created_by")
	delete(attrs, "last_updated_by")
	delete(attrs, "vault")

	attrs["vault_id"] = schema.StringAttribute{
		Description: "ID of the vault that this secret belongs to.",
		Computed:    true,
	}
	attrs["created_by_id"] = schema.StringAttribute{
		Description: "The ID of the user or bot that created the secret.",
		Computed:    true,
	}
	attrs["last_updated_by_id"] = schema.StringAttribute{
		Description: "The ID of the user or bot that last updated the secret.",
		Computed:    true,
	}
}

func (d *secretDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	d.client = secrets.NewClient(clientConfig)
}

func (d *secretDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config secretDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var secret *ngrok.Secret

	if !config.ID.IsNull() && config.ID.ValueString() != "" {
		var err error
		secret, err = d.client.Get(ctx, config.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Error reading secret", err.Error())
			return
		}
	} else if !config.Name.IsNull() && config.Name.ValueString() != "" {
		filter := fmt.Sprintf(`obj.name == %q`, config.Name.ValueString())
		iter := d.client.List(&ngrok.FilteredPaging{
			Filter: &filter,
		})

		if !iter.Next(ctx) {
			if err := iter.Err(); err != nil {
				resp.Diagnostics.AddError("Error listing secrets", err.Error())
				return
			}
			resp.Diagnostics.AddError(
				"Secret not found",
				fmt.Sprintf("No secret found with name %q.", config.Name.ValueString()),
			)
			return
		}
		secret = iter.Item()

		if iter.Next(ctx) {
			resp.Diagnostics.AddError(
				"Multiple secrets found",
				fmt.Sprintf("More than one secret found with name %q. Use id instead.", config.Name.ValueString()),
			)
			return
		}
		if err := iter.Err(); err != nil {
			resp.Diagnostics.AddError("Error listing secrets", err.Error())
			return
		}
	} else {
		resp.Diagnostics.AddError(
			"Missing lookup attribute",
			"Either id or name must be specified.",
		)
		return
	}

	var model secretDataSourceModel
	model.ID = types.StringValue(secret.ID)
	model.URI = types.StringValue(secret.URI)
	model.CreatedAt = types.StringValue(secret.CreatedAt)
	model.UpdatedAt = types.StringValue(secret.UpdatedAt)
	model.Name = types.StringValue(secret.Name)
	model.Description = types.StringValue(secret.Description)
	model.Metadata = types.StringValue(secret.Metadata)
	model.VaultID = types.StringValue(secret.Vault.ID)
	model.VaultName = types.StringValue(secret.VaultName)
	model.CreatedByID = types.StringValue(secret.CreatedBy.ID)
	model.LastUpdatedByID = types.StringValue(secret.LastUpdatedBy.ID)

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
