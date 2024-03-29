package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func resourceSSHCertificateAuthorities() *schema.Resource {
	return &schema.Resource{
		Create:      resourceSSHCertificateAuthoritiesCreate,
		Read:        resourceSSHCertificateAuthoritiesGet,
		Update:      resourceSSHCertificateAuthoritiesUpdate,
		Delete:      resourceSSHCertificateAuthoritiesDelete,
		Description: "An SSH Certificate Authority is a pair of an SSH Certificate and its private\n key that can be used to sign other SSH host and user certificates.",
		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "human-readable description of this SSH Certificate Authority. optional, max 255 bytes.",
			},
			"elliptic_curve": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "the type of elliptic curve to use when creating an ECDSA key",
			},
			"id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    false,
				Description: "unique identifier for this SSH Certificate Authority",
			},
			"key_size": {
				Type:        schema.TypeInt,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "the key size to use when creating an RSA key. one of `2048` or `4096`",
			},
			"key_type": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "the type of private key for this SSH Certificate Authority",
			},
			"metadata": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "arbitrary user-defined machine-readable data of this SSH Certificate Authority. optional, max 4096 bytes.",
			},
			"private_key_type": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "the type of private key to generate. one of `rsa`, `ecdsa`, `ed25519`",
			},
			"public_key": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "raw public key for this SSH Certificate Authority",
			},
		},
	}
}

func resourceSSHCertificateAuthoritiesCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.SSHCertificateAuthorityCreate
	if v, ok := d.GetOk("description"); ok {
		arg.Description = *expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = *expandString(v)
	}
	if v, ok := d.GetOk("private_key_type"); ok {
		arg.PrivateKeyType = *expandString(v)
	}
	if v, ok := d.GetOk("elliptic_curve"); ok {
		arg.EllipticCurve = *expandString(v)
	}
	if v, ok := d.GetOk("key_size"); ok {
		arg.KeySize = *expandInt64(v)
	}

	res, _, err := b.client.SSHCertificateAuthoritiesCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] SSHCertificateAuthoritiesCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceSSHCertificateAuthoritiesGet(d, m)
}

func resourceSSHCertificateAuthoritiesGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.SSHCertificateAuthoritiesGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceSSHCertificateAuthoritiesGetDecode(d, res, resp, err)
}

func resourceSSHCertificateAuthoritiesGetDecode(d *schema.ResourceData, res *restapi.SSHCertificateAuthority, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		log.Printf("[ERROR] SSHCertificateAuthoritiesGet: %s", err)
		return err
	default:
		d.Set("description", res.Description)
		d.Set("id", res.ID)
		d.Set("key_type", res.KeyType)
		d.Set("metadata", res.Metadata)
		d.Set("public_key", res.PublicKey)
	}
	return nil
}

func resourceSSHCertificateAuthoritiesUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.SSHCertificateAuthorityUpdate
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

	res, _, err := b.client.SSHCertificateAuthoritiesUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] SSHCertificateAuthoritiesUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceSSHCertificateAuthoritiesGet(d, m)
}

func resourceSSHCertificateAuthoritiesDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.SSHCertificateAuthoritiesDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] SSHCertificateAuthoritiesDelete: %s", err)
	}
	return err
}
