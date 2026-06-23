package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/ssh_user_certificates"
	"github.com/ngrok/terraform-provider-ngrok/internal/datasource_ssh_user_certificate"
)

var _ datasource.DataSource = &sshUserCertificateDataSource{}

type sshUserCertificateDataSourceModel struct {
	ID                        types.String   `tfsdk:"id"`
	URI                       types.String   `tfsdk:"uri"`
	CreatedAt                 types.String   `tfsdk:"created_at"`
	Description               types.String   `tfsdk:"description"`
	Metadata                  types.String   `tfsdk:"metadata"`
	PublicKey                 types.String   `tfsdk:"public_key"`
	KeyType                   types.String   `tfsdk:"key_type"`
	SSHCertificateAuthorityID types.String   `tfsdk:"ssh_certificate_authority_id"`
	Principals                []types.String `tfsdk:"principals"`
	CriticalOptions           types.Map      `tfsdk:"critical_options"`
	Extensions                types.Map      `tfsdk:"extensions"`
	ValidAfter                types.String   `tfsdk:"valid_after"`
	ValidUntil                types.String   `tfsdk:"valid_until"`
	Certificate               types.String   `tfsdk:"certificate"`
}

type sshUserCertificateDataSource struct {
	client *ssh_user_certificates.Client
}

func NewSSHUserCertificateDataSource() datasource.DataSource {
	return &sshUserCertificateDataSource{}
}

func (d *sshUserCertificateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssh_user_certificate"
}

func (d *sshUserCertificateDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_ssh_user_certificate.SshUserCertificateDataSourceSchema(ctx)
	resp.Schema.Description = "Use this data source to look up an SSH User Certificate by ID."

	attrs := resp.Schema.Attributes

	attrs["critical_options"] = schema.MapAttribute{
		Description: "A map of critical options included in the certificate.",
		Computed:    true,
		ElementType: types.StringType,
	}
	attrs["extensions"] = schema.MapAttribute{
		Description: "A map of extensions included in the certificate.",
		Computed:    true,
		ElementType: types.StringType,
	}
}

func (d *sshUserCertificateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	d.client = ssh_user_certificates.NewClient(clientConfig)
}

func (d *sshUserCertificateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config sshUserCertificateDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	cert, err := d.client.Get(ctx, config.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading SSH user certificate", err.Error())
		return
	}

	var model sshUserCertificateDataSourceModel
	model.ID = types.StringValue(cert.ID)
	model.URI = types.StringValue(cert.URI)
	model.CreatedAt = types.StringValue(cert.CreatedAt)
	model.Description = types.StringValue(cert.Description)
	model.Metadata = types.StringValue(cert.Metadata)
	model.PublicKey = types.StringValue(cert.PublicKey)
	model.KeyType = types.StringValue(cert.KeyType)
	model.SSHCertificateAuthorityID = types.StringValue(cert.SSHCertificateAuthorityID)
	model.Principals = flattenStringList(cert.Principals)
	model.CriticalOptions = flattenStringMap(ctx, cert.CriticalOptions)
	model.Extensions = flattenStringMap(ctx, cert.Extensions)
	model.ValidAfter = types.StringValue(cert.ValidAfter)
	model.ValidUntil = types.StringValue(cert.ValidUntil)
	model.Certificate = types.StringValue(cert.Certificate)

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
