package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
)

var _ provider.Provider = &ngrokProvider{}

type ngrokProvider struct {
	version string
}

type ngrokProviderModel struct {
	APIKey     types.String `tfsdk:"api_key"`
	APIBaseURL types.String `tfsdk:"api_base_url"`
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ngrokProvider{
			version: version,
		}
	}
}

func (p *ngrokProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "ngrok"
	resp.Version = p.version
}

func (p *ngrokProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Interact with the ngrok API.",
		Attributes: map[string]schema.Attribute{
			"api_key": schema.StringAttribute{
				Description: "The ngrok API key. Can also be set via the NGROK_API_KEY environment variable.",
				Optional:    true,
				Sensitive:   true,
			},
			"api_base_url": schema.StringAttribute{
				Description: "The base URL for the ngrok API. Defaults to https://api.ngrok.com. Can also be set via the NGROK_API_BASE_URL environment variable.",
				Optional:    true,
			},
		},
	}
}

func (p *ngrokProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config ngrokProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiKey := os.Getenv("NGROK_API_KEY")
	if !config.APIKey.IsNull() {
		apiKey = config.APIKey.ValueString()
	}
	if apiKey == "" {
		resp.Diagnostics.AddError(
			"Missing API Key",
			"The ngrok API key must be set in the provider configuration or via the NGROK_API_KEY environment variable.",
		)
		return
	}

	var opts []ngrok.ClientConfigOption

	apiBaseURL := os.Getenv("NGROK_API_BASE_URL")
	if !config.APIBaseURL.IsNull() {
		apiBaseURL = config.APIBaseURL.ValueString()
	}
	if apiBaseURL != "" {
		opts = append(opts, ngrok.WithBaseURL(apiBaseURL))
	}

	userAgent := "terraform-provider-ngrok/" + p.version
	opts = append(opts, ngrok.WithUserAgent(userAgent))

	clientConfig := ngrok.NewClientConfig(apiKey, opts...)

	resp.DataSourceData = clientConfig
	resp.ResourceData = clientConfig
}

func (p *ngrokProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewAPIKeyResource,
		NewCredentialResource,
		NewReservedDomainResource,
		NewReservedAddrResource,
		NewCloudEndpointResource,
		NewIPPolicyResource,
		NewIPPolicyRuleResource,
		NewIPRestrictionResource,
		NewCertificateAuthorityResource,
		NewTLSCertificateResource,
		NewServiceUserResource,
		NewAgentIngressResource,
		NewEventDestinationResource,
		NewEventSubscriptionResource,
		NewVaultResource,
		NewSecretResource,
		NewSSHCertificateAuthorityResource,
		NewSSHCredentialResource,
		NewSSHHostCertificateResource,
		NewSSHUserCertificateResource,
	}
}

func (p *ngrokProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAPIKeyDataSource,
		NewCredentialDataSource,
		NewReservedDomainDataSource,
		NewReservedAddrDataSource,
		NewCloudEndpointDataSource,
		NewIPPolicyDataSource,
		NewIPPolicyRuleDataSource,
		NewIPRestrictionDataSource,
		NewCertificateAuthorityDataSource,
		NewTLSCertificateDataSource,
		NewServiceUserDataSource,
		NewAgentIngressDataSource,
		NewEventDestinationDataSource,
		NewEventSubscriptionDataSource,
		NewVaultDataSource,
		NewSecretDataSource,
		NewSSHCertificateAuthorityDataSource,
		NewSSHCredentialDataSource,
		NewSSHHostCertificateDataSource,
		NewSSHUserCertificateDataSource,
		NewTunnelSessionDataSource,
		NewApplicationSessionDataSource,
		NewApplicationUserDataSource,
	}
}
