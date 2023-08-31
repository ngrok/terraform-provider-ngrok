// Code generated for API Clients. DO NOT EDIT.

package ngrok

import (
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("NGROK_API_KEY", ""),
				Description: "ngrok API key used for authenticating to the API",
				Required:    true,
			},
			"api_base_url": {
				Type:        schema.TypeString,
				Description: "base URL for the ngrok API. only overridden for ngrok's internal testing purposes.",
				DefaultFunc: schema.EnvDefaultFunc("NGROK_API_BASE_URL", "https://api.ngrok.com"),
				Optional:    true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ngrok_agent_ingress":             resourceAgentIngresses(),
			"ngrok_api_key":                   resourceAPIKeys(),
			"ngrok_failover_backend":          resourceFailoverBackends(),
			"ngrok_http_response_backend":     resourceHTTPResponseBackends(),
			"ngrok_tunnel_group_backend":      resourceTunnelGroupBackends(),
			"ngrok_certificate_authority":     resourceCertificateAuthorities(),
			"ngrok_credential":                resourceCredentials(),
			"ngrok_event_destination":         resourceEventDestinations(),
			"ngrok_event_subscription":        resourceEventSubscriptions(),
			"ngrok_ip_policy":                 resourceIPPolicies(),
			"ngrok_ip_policy_rule":            resourceIPPolicyRules(),
			"ngrok_ip_restriction":            resourceIPRestrictions(),
			"ngrok_reserved_addr":             resourceReservedAddrs(),
			"ngrok_reserved_domain":           resourceReservedDomains(),
			"ngrok_ssh_certificate_authority": resourceSSHCertificateAuthorities(),
			"ngrok_ssh_credential":            resourceSSHCredentials(),
			"ngrok_ssh_host_certificate":      resourceSSHHostCertificates(),
			"ngrok_ssh_user_certificate":      resourceSSHUserCertificates(),
			"ngrok_tls_certificate":           resourceTLSCertificates(),
		},
		ConfigureFunc: configureProvider,
	}
}

type base struct {
	client *restapi.Client
}

func configureProvider(data *schema.ResourceData) (_ interface{}, err error) {
	var (
		b base

		cfg = restapi.ClientConfig{
			APIKey:  data.Get("api_key").(string),
			BaseURL: data.Get("api_base_url").(string),
		}
	)

	if cfg.APIKey != "" {
		b.client, err = restapi.NewClient(cfg)
	} else {
		err = errors.New("you must specify an 'api_key'")
	}

	if err != nil {
		return nil, err
	}

	return &b, nil
}
