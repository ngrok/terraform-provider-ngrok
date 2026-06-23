package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/certificate_authorities"
	"github.com/ngrok/terraform-provider-ngrok/internal/datasource_certificate_authority"
)

var _ datasource.DataSource = &certificateAuthorityDataSource{}

type certificateAuthorityDataSourceModel struct {
	ID                types.String   `tfsdk:"id"`
	URI               types.String   `tfsdk:"uri"`
	CreatedAt         types.String   `tfsdk:"created_at"`
	Description       types.String   `tfsdk:"description"`
	Metadata          types.String   `tfsdk:"metadata"`
	CAPEM             types.String   `tfsdk:"ca_pem"`
	SubjectCommonName types.String   `tfsdk:"subject_common_name"`
	NotBefore         types.String   `tfsdk:"not_before"`
	NotAfter          types.String   `tfsdk:"not_after"`
	KeyUsages         []types.String `tfsdk:"key_usages"`
	ExtendedKeyUsages []types.String `tfsdk:"extended_key_usages"`
}

type certificateAuthorityDataSource struct {
	client *certificate_authorities.Client
}

func NewCertificateAuthorityDataSource() datasource.DataSource {
	return &certificateAuthorityDataSource{}
}

func (d *certificateAuthorityDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_certificate_authority"
}

func (d *certificateAuthorityDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_certificate_authority.CertificateAuthorityDataSourceSchema(ctx)
}

func (d *certificateAuthorityDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	d.client = certificate_authorities.NewClient(clientConfig)
}

func (d *certificateAuthorityDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config certificateAuthorityDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ca, err := d.client.Get(ctx, config.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading certificate authority", err.Error())
		return
	}

	var model certificateAuthorityDataSourceModel
	flattenCertificateAuthorityDataSource(ca, &model)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func flattenCertificateAuthorityDataSource(ca *ngrok.CertificateAuthority, model *certificateAuthorityDataSourceModel) {
	model.ID = types.StringValue(ca.ID)
	model.URI = types.StringValue(ca.URI)
	model.CreatedAt = types.StringValue(ca.CreatedAt)
	model.Description = types.StringValue(ca.Description)
	model.Metadata = types.StringValue(ca.Metadata)
	model.CAPEM = types.StringValue(ca.CAPEM)
	model.SubjectCommonName = types.StringValue(ca.SubjectCommonName)
	model.NotBefore = types.StringValue(ca.NotBefore)
	model.NotAfter = types.StringValue(ca.NotAfter)
	model.KeyUsages = flattenStringList(ca.KeyUsages)
	model.ExtendedKeyUsages = flattenStringList(ca.ExtendedKeyUsages)
}
