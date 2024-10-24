package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
	transform "github.com/ngrok/terraform-provider-ngrok/transform"
)

func resourceReservedDomains() *schema.Resource {
	return &schema.Resource{
		Create:      resourceReservedDomainsCreate,
		Read:        resourceReservedDomainsGet,
		Update:      resourceReservedDomainsUpdate,
		Delete:      resourceReservedDomainsDelete,
		Description: "Reserved Domains are hostnames that you can listen for traffic on. Domains\n can be used to listen for http, https or tls traffic. You may use a domain\n that you own by creating a CNAME record specified in the returned resource.\n This CNAME record points traffic for that domain to ngrok's edge servers.",
		Schema: map[string]*schema.Schema{
			"acme_challenge_cname_target": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "DNS CNAME target for the host _acme-challenge.example.com, where example.com is your reserved domain name. This is required to issue certificates for wildcard, non-ngrok reserved domains. Must be null for non-wildcard domains and ngrok subdomains.",
			},
			"certificate_id": {
				Type:             schema.TypeString,
				Required:         false,
				Computed:         false,
				Optional:         true,
				Sensitive:        false,
				ForceNew:         false,
				Description:      "ID of a user-uploaded TLS certificate to use for connections to targeting this domain. Optional, mutually exclusive with `certificate_management_policy`.",
				DiffSuppressFunc: transform.DiffSuppressAutoCertId,
			},
			"certificate_management_policy": {
				Type:        schema.TypeSet,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "configuration for automatic management of TLS certificates for this domain, or null if automatic management is disabled",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authority": {
							Type:        schema.TypeString,
							Required:    false,
							Computed:    false,
							Optional:    true,
							Sensitive:   false,
							ForceNew:    false,
							Description: "certificate authority to request certificates from. The only supported value is letsencrypt.",
						},
						"private_key_type": {
							Type:        schema.TypeString,
							Required:    false,
							Computed:    false,
							Optional:    true,
							Default:     "ecdsa",
							Sensitive:   false,
							ForceNew:    false,
							Description: "type of private key to use when requesting certificates. Defaults to ecdsa, can be either rsa or ecdsa.",
						},
					},
				},
			},
			"cname_target": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "DNS CNAME target for a custom hostname, or null if the reserved domain is a subdomain of an ngrok owned domain (e.g. *.ngrok.app)",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "human-readable description of what this reserved domain will be used for",
			},
			"domain": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "hostname of the reserved domain",
			},
			"error_redirect_url": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "Custom URL with CEL Expression Variable support for redirecting when an ngrok error occurs. Max 10000 bytes.",
			},
			"http_endpoint_configuration_id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "ID of an endpoint configuration of type http that will be used to handle inbound http traffic to this domain",
			},
			"https_endpoint_configuration_id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "ID of an endpoint configuration of type https that will be used to handle inbound https traffic to this domain",
			},
			"id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    false,
				Description: "unique reserved domain resource identifier",
			},
			"metadata": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "arbitrary user-defined machine-readable data of this reserved domain. Optional, max 4096 bytes.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "the domain name to reserve. It may be a full domain name like app.example.com. If the name does not contain a '.' it will reserve that subdomain on ngrok.io.",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "deprecated: With the launch of the ngrok Global Network domains traffic is now handled globally. This field applied only to endpoints. Note that agents may still connect to specific regions. Optional, null by default. (au, eu, ap, us, jp, in, sa)",
			},
		},
	}
}

func resourceReservedDomainsCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.ReservedDomainCreate
	if v, ok := d.GetOk("name"); ok {
		arg.Name = *expandString(v)
	}
	if v, ok := d.GetOk("domain"); ok {
		arg.Domain = *expandString(v)
	}
	if v, ok := d.GetOk("region"); ok {
		arg.Region = *expandString(v)
	}
	if v, ok := d.GetOk("description"); ok {
		arg.Description = *expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = *expandString(v)
	}
	if v, ok := d.GetOk("http_endpoint_configuration_id"); ok {
		arg.HTTPEndpointConfigurationID = expandString(v)
	}
	if v, ok := d.GetOk("https_endpoint_configuration_id"); ok {
		arg.HTTPSEndpointConfigurationID = expandString(v)
	}
	if v, ok := d.GetOk("certificate_id"); ok {
		arg.CertificateID = expandString(v)
	}
	if v, ok := d.GetOk("certificate_management_policy"); ok {
		arg.CertificateManagementPolicy = expandReservedDomainCertPolicy(v)
	}
	if v, ok := d.GetOk("error_redirect_url"); ok {
		arg.ErrorRedirectUrl = expandString(v)
	}

	res, _, err := b.client.ReservedDomainsCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] ReservedDomainsCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceReservedDomainsGet(d, m)
}

func resourceReservedDomainsGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.ReservedDomainsGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceReservedDomainsGetDecode(d, res, resp, err)
}

func resourceReservedDomainsGetDecode(d *schema.ResourceData, res *restapi.ReservedDomain, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		log.Printf("[ERROR] ReservedDomainsGet: %s", err)
		return err
	default:
		d.Set("acme_challenge_cname_target", res.ACMEChallengeCNAMETarget)
		if res.Certificate != nil {
			d.Set("certificate_id", res.Certificate.ID)
		}
		d.Set("certificate_management_policy", flattenReservedDomainCertPolicy(res.CertificateManagementPolicy))
		d.Set("cname_target", res.CNAMETarget)
		d.Set("description", res.Description)
		d.Set("domain", res.Domain)
		d.Set("error_redirect_url", res.ErrorRedirectURL)
		if res.HTTPEndpointConfiguration != nil {
			d.Set("http_endpoint_configuration_id", res.HTTPEndpointConfiguration.ID)
		}
		if res.HTTPSEndpointConfiguration != nil {
			d.Set("https_endpoint_configuration_id", res.HTTPSEndpointConfiguration.ID)
		}
		d.Set("id", res.ID)
		d.Set("metadata", res.Metadata)
		d.Set("region", res.Region)
	}
	return nil
}

func resourceReservedDomainsUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.ReservedDomainUpdate
	arg.ID = d.Id()
	if v, ok := d.GetOk("id"); ok {
		arg.ID = *expandString(v)
	}
	if v, ok := d.GetOk("description"); ok {
		arg.Description = expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = expandString(v)
	}
	if v, ok := d.GetOk("http_endpoint_configuration_id"); ok {
		arg.HTTPEndpointConfigurationID = expandString(v)
	}
	if v, ok := d.GetOk("https_endpoint_configuration_id"); ok {
		arg.HTTPSEndpointConfigurationID = expandString(v)
	}
	if v, ok := d.GetOk("certificate_id"); ok {
		arg.CertificateID = expandString(v)
	}
	if v, ok := d.GetOk("certificate_management_policy"); ok {
		arg.CertificateManagementPolicy = expandReservedDomainCertPolicy(v)
	}
	if v, ok := d.GetOk("region"); ok {
		arg.Region = expandString(v)
	}
	if v, ok := d.GetOk("error_redirect_url"); ok {
		arg.ErrorRedirectUrl = expandString(v)
	}

	res, _, err := b.client.ReservedDomainsUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] ReservedDomainsUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceReservedDomainsGet(d, m)
}

func resourceReservedDomainsDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.ReservedDomainsDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] ReservedDomainsDelete: %s", err)
	}
	return err
}
