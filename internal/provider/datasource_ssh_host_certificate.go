package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/ssh_host_certificates"
	"github.com/ngrok/terraform-provider-ngrok/internal/datasource_ssh_host_certificate"
)

var _ datasource.DataSource = &sshHostCertificateDataSource{}

type sshHostCertificateDataSourceModel struct {
	ID                        types.String   `tfsdk:"id"`
	URI                       types.String   `tfsdk:"uri"`
	CreatedAt                 types.String   `tfsdk:"created_at"`
	Description               types.String   `tfsdk:"description"`
	Metadata                  types.String   `tfsdk:"metadata"`
	PublicKey                 types.String   `tfsdk:"public_key"`
	KeyType                   types.String   `tfsdk:"key_type"`
	SSHCertificateAuthorityID types.String   `tfsdk:"ssh_certificate_authority_id"`
	Principals                []types.String `tfsdk:"principals"`
	ValidAfter                types.String   `tfsdk:"valid_after"`
	ValidUntil                types.String   `tfsdk:"valid_until"`
	Certificate               types.String   `tfsdk:"certificate"`
}

type sshHostCertificateDataSource struct {
	client *ssh_host_certificates.Client
}

func NewSSHHostCertificateDataSource() datasource.DataSource {
	return &sshHostCertificateDataSource{}
}

func (d *sshHostCertificateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssh_host_certificate"
}

func (d *sshHostCertificateDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_ssh_host_certificate.SshHostCertificateDataSourceSchema(ctx)
}

func (d *sshHostCertificateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	d.client = ssh_host_certificates.NewClient(clientConfig)
}

func (d *sshHostCertificateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config sshHostCertificateDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	cert, err := d.client.Get(ctx, config.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading SSH host certificate", err.Error())
		return
	}

	var model sshHostCertificateDataSourceModel
	model.ID = types.StringValue(cert.ID)
	model.URI = types.StringValue(cert.URI)
	model.CreatedAt = types.StringValue(cert.CreatedAt)
	model.Description = types.StringValue(cert.Description)
	model.Metadata = types.StringValue(cert.Metadata)
	model.PublicKey = types.StringValue(cert.PublicKey)
	model.KeyType = types.StringValue(cert.KeyType)
	model.SSHCertificateAuthorityID = types.StringValue(cert.SSHCertificateAuthorityID)
	model.Principals = flattenStringList(cert.Principals)
	model.ValidAfter = types.StringValue(cert.ValidAfter)
	model.ValidUntil = types.StringValue(cert.ValidUntil)
	model.Certificate = types.StringValue(cert.Certificate)

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
