package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
	transform "github.com/ngrok/terraform-provider-ngrok/transform"
)

func resourceTLSCertificates() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTLSCertificatesCreate,
		Read:        resourceTLSCertificatesGet,
		Update:      resourceTLSCertificatesUpdate,
		Delete:      resourceTLSCertificatesDelete,
		Description: "TLS Certificates are pairs of x509 certificates and their matching private\n key that can be used to terminate TLS traffic. TLS certificates are unused\n until they are attached to a Domain. TLS Certificates may also be\n provisioned by ngrok automatically for domains on which you have enabled\n automated certificate provisioning.",
		Schema: map[string]*schema.Schema{
			"certificate_pem": {
				Type:             schema.TypeString,
				Required:         true,
				Computed:         false,
				Optional:         false,
				Sensitive:        false,
				ForceNew:         true,
				Description:      "chain of PEM-encoded certificates, leaf first. See [Certificate Bundles](https://ngrok.com/docs/cloud-edge/endpoints#certificate-chains).",
				DiffSuppressFunc: transform.DiffSuppressWhitespace,
			},
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "human-readable description of this TLS certificate. optional, max 255 bytes.",
			},
			"id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    false,
				Description: "unique identifier for this TLS certificate",
			},
			"metadata": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "arbitrary user-defined machine-readable data of this TLS certificate. optional, max 4096 bytes.",
			},
			"private_key_pem": {
				Type:        schema.TypeString,
				Required:    true,
				Computed:    false,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "private key for the TLS certificate, PEM-encoded. See [Private Keys](https://ngrok.com/docs/cloud-edge/endpoints#private-keys).",
			},
			"subject_alternative_names": {
				Type:        schema.TypeSet,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "subject alternative names (SANs) from the leaf of this TLS certificate",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_names": {
							Type:        schema.TypeList,
							Required:    false,
							Computed:    true,
							Optional:    false,
							Sensitive:   false,
							ForceNew:    true,
							Description: "set of additional domains (including wildcards) this TLS certificate is valid for",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"ips": {
							Type:        schema.TypeList,
							Required:    false,
							Computed:    true,
							Optional:    false,
							Sensitive:   false,
							ForceNew:    true,
							Description: "set of IP addresses this TLS certificate is also valid for",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
		},
	}
}

func resourceTLSCertificatesCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.TLSCertificateCreate
	if v, ok := d.GetOk("description"); ok {
		arg.Description = *expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = *expandString(v)
	}
	if v, ok := d.GetOk("certificate_pem"); ok {
		arg.CertificatePEM = *expandString(v)
	}
	if v, ok := d.GetOk("private_key_pem"); ok {
		arg.PrivateKeyPEM = *expandString(v)
	}

	res, _, err := b.client.TLSCertificatesCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] TLSCertificatesCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceTLSCertificatesGet(d, m)
}

func resourceTLSCertificatesGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.TLSCertificatesGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceTLSCertificatesGetDecode(d, res, resp, err)
}

func resourceTLSCertificatesGetDecode(d *schema.ResourceData, res *restapi.TLSCertificate, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		log.Printf("[ERROR] TLSCertificatesGet: %s", err)
		return err
	default:
		d.Set("certificate_pem", res.CertificatePEM)
		d.Set("description", res.Description)
		d.Set("id", res.ID)
		d.Set("metadata", res.Metadata)
		d.Set("subject_alternative_names", flattenTLSCertificateSANs(&res.SubjectAlternativeNames))
	}
	return nil
}

func resourceTLSCertificatesUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.TLSCertificateUpdate
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

	res, _, err := b.client.TLSCertificatesUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] TLSCertificatesUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceTLSCertificatesGet(d, m)
}

func resourceTLSCertificatesDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.TLSCertificatesDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] TLSCertificatesDelete: %s", err)
	}
	return err
}
