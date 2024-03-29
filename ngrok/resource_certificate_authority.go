package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
	transform "github.com/ngrok/terraform-provider-ngrok/transform"
)

func resourceCertificateAuthorities() *schema.Resource {
	return &schema.Resource{
		Create:      resourceCertificateAuthoritiesCreate,
		Read:        resourceCertificateAuthoritiesGet,
		Update:      resourceCertificateAuthoritiesUpdate,
		Delete:      resourceCertificateAuthoritiesDelete,
		Description: "Certificate Authorities are x509 certificates that are used to sign other\n x509 certificates. Attach a Certificate Authority to the Mutual TLS module\n to verify that the TLS certificate presented by a client has been signed by\n this CA. Certificate Authorities  are used only for mTLS validation only and\n thus a private key is not included in the resource.",
		Schema: map[string]*schema.Schema{
			"ca_pem": {
				Type:             schema.TypeString,
				Required:         true,
				Computed:         false,
				Optional:         false,
				Sensitive:        false,
				ForceNew:         true,
				Description:      "raw PEM of the Certificate Authority",
				DiffSuppressFunc: transform.DiffSuppressWhitespace,
			},
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "human-readable description of this Certificate Authority. optional, max 255 bytes.",
			},
			"id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    false,
				Description: "unique identifier for this Certificate Authority",
			},
			"metadata": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "arbitrary user-defined machine-readable data of this Certificate Authority. optional, max 4096 bytes.",
			},
		},
	}
}

func resourceCertificateAuthoritiesCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.CertificateAuthorityCreate
	if v, ok := d.GetOk("description"); ok {
		arg.Description = *expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = *expandString(v)
	}
	if v, ok := d.GetOk("ca_pem"); ok {
		arg.CAPEM = *expandString(v)
	}

	res, _, err := b.client.CertificateAuthoritiesCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] CertificateAuthoritiesCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceCertificateAuthoritiesGet(d, m)
}

func resourceCertificateAuthoritiesGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.CertificateAuthoritiesGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceCertificateAuthoritiesGetDecode(d, res, resp, err)
}

func resourceCertificateAuthoritiesGetDecode(d *schema.ResourceData, res *restapi.CertificateAuthority, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		log.Printf("[ERROR] CertificateAuthoritiesGet: %s", err)
		return err
	default:
		d.Set("ca_pem", res.CAPEM)
		d.Set("description", res.Description)
		d.Set("id", res.ID)
		d.Set("metadata", res.Metadata)
	}
	return nil
}

func resourceCertificateAuthoritiesUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.CertificateAuthorityUpdate
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

	res, _, err := b.client.CertificateAuthoritiesUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] CertificateAuthoritiesUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceCertificateAuthoritiesGet(d, m)
}

func resourceCertificateAuthoritiesDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.CertificateAuthoritiesDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] CertificateAuthoritiesDelete: %s", err)
	}
	return err
}
