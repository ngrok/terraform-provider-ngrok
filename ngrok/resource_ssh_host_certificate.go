package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func resourceSSHHostCertificates() *schema.Resource {
	return &schema.Resource{
		Create:      resourceSSHHostCertificatesCreate,
		Read:        resourceSSHHostCertificatesGet,
		Update:      resourceSSHHostCertificatesUpdate,
		Delete:      resourceSSHHostCertificatesDelete,
		Description: "SSH Host Certificates along with the corresponding private key allows an SSH\n server to assert its authenticity to connecting SSH clients who trust the\n SSH Certificate Authority that was used to sign the certificate.",
		Schema: map[string]*schema.Schema{
			"certificate": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "the signed SSH certificate in OpenSSH Authorized Keys format. this value should be placed in a `-cert.pub` certificate file on disk that should be referenced in your `sshd_config` configuration file with a `HostCertificate` directive",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "human-readable description of this SSH Host Certificate. optional, max 255 bytes.",
			},
			"id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "unique identifier for this SSH Host Certificate",
			},
			"key_type": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "the key type of the `public_key`, one of `rsa`, `ecdsa` or `ed25519`",
			},
			"metadata": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "arbitrary user-defined machine-readable data of this SSH Host Certificate. optional, max 4096 bytes.",
			},
			"principals": {
				Type:        schema.TypeList,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "the list of principals included in the ssh host certificate. This is the list of hostnames and/or IP addresses that are authorized to serve SSH traffic with this certificate. Dangerously, if no principals are specified, this certificate is considered valid for all hosts.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"public_key": {
				Type:        schema.TypeString,
				Required:    true,
				Computed:    false,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "a public key in OpenSSH Authorized Keys format that this certificate signs",
			},
			"ssh_certificate_authority_id": {
				Type:        schema.TypeString,
				Required:    true,
				Computed:    false,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "the ssh certificate authority that is used to sign this ssh host certificate",
			},
			"valid_after": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "the time when the ssh host certificate becomes valid, in RFC 3339 format.",
			},
			"valid_until": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "the time after which the ssh host certificate becomes invalid, in RFC 3339 format. the OpenSSH certificates RFC calls this `valid_before`.",
			},
		},
	}
}

func resourceSSHHostCertificatesCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.SSHHostCertificateCreate
	if v, ok := d.GetOk("ssh_certificate_authority_id"); ok {
		arg.SSHCertificateAuthorityID = *expandString(v)
	}
	if v, ok := d.GetOk("public_key"); ok {
		arg.PublicKey = *expandString(v)
	}
	if v, ok := d.GetOk("principals"); ok {
		arg.Principals = *expandStringSlice(v)
	}
	if v, ok := d.GetOk("valid_after"); ok {
		arg.ValidAfter = *expandString(v)
	}
	if v, ok := d.GetOk("valid_until"); ok {
		arg.ValidUntil = *expandString(v)
	}
	if v, ok := d.GetOk("description"); ok {
		arg.Description = *expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = *expandString(v)
	}

	res, _, err := b.client.SSHHostCertificatesCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] SSHHostCertificatesCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceSSHHostCertificatesGet(d, m)
}

func resourceSSHHostCertificatesGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.SSHHostCertificatesGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceSSHHostCertificatesGetDecode(d, res, resp, err)
}

func resourceSSHHostCertificatesGetDecode(d *schema.ResourceData, res *restapi.SSHHostCertificate, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		log.Printf("[ERROR] SSHHostCertificatesGet: %s", err)
		return err
	default:
		d.Set("certificate", res.Certificate)
		d.Set("description", res.Description)
		d.Set("id", res.ID)
		d.Set("key_type", res.KeyType)
		d.Set("metadata", res.Metadata)
		d.Set("principals", res.Principals)
		d.Set("public_key", res.PublicKey)
		d.Set("ssh_certificate_authority_id", res.SSHCertificateAuthorityID)
		d.Set("valid_after", res.ValidAfter)
		d.Set("valid_until", res.ValidUntil)
	}
	return nil
}

func resourceSSHHostCertificatesUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.SSHHostCertificateUpdate
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

	res, _, err := b.client.SSHHostCertificatesUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] SSHHostCertificatesUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceSSHHostCertificatesGet(d, m)
}

func resourceSSHHostCertificatesDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.SSHHostCertificatesDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] SSHHostCertificatesDelete: %s", err)
	}
	return err
}
