package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func resourceSSHUserCertificates() *schema.Resource {
	return &schema.Resource{
		Create:      resourceSSHUserCertificatesCreate,
		Read:        resourceSSHUserCertificatesGet,
		Update:      resourceSSHUserCertificatesUpdate,
		Delete:      resourceSSHUserCertificatesDelete,
		Description: "SSH User Certificates are presented by SSH clients when connecting to an SSH\n server to authenticate their connection. The SSH server must trust the SSH\n Certificate Authority used to sign the certificate.",
		Schema: map[string]*schema.Schema{
			"certificate": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "the signed SSH certificate in OpenSSH Authorized Keys Format. this value should be placed in a `-cert.pub` certificate file on disk that should be referenced in your `sshd_config` configuration file with a `HostCertificate` directive",
			},
			"critical_options": {
				Type:        schema.TypeMap,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "A map of critical options included in the certificate. Only two critical options are currently defined by OpenSSH: `force-command` and `source-address`. See [the OpenSSH certificate protocol spec](https://github.com/openssh/openssh-portable/blob/master/PROTOCOL.certkeys) for additional details.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "human-readable description of this SSH User Certificate. optional, max 255 bytes.",
			},
			"extensions": {
				Type:        schema.TypeMap,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "A map of extensions included in the certificate. Extensions are additional metadata that can be interpreted by the SSH server for any purpose. These can be used to permit or deny the ability to open a terminal, do port forwarding, x11 forwarding, and more. If unspecified, the certificate will include limited permissions with the following extension map: `{\"permit-pty\": \"\", \"permit-user-rc\": \"\"}` OpenSSH understands a number of predefined extensions. See [the OpenSSH certificate protocol spec](https://github.com/openssh/openssh-portable/blob/master/PROTOCOL.certkeys) for additional details.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "unique identifier for this SSH User Certificate",
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
				Description: "arbitrary user-defined machine-readable data of this SSH User Certificate. optional, max 4096 bytes.",
			},
			"principals": {
				Type:        schema.TypeList,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "the list of principals included in the ssh user certificate. This is the list of usernames that the certificate holder may sign in as on a machine authorizing the signing certificate authority. Dangerously, if no principals are specified, this certificate may be used to log in as any user.",
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
				Description: "the ssh certificate authority that is used to sign this ssh user certificate",
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

func resourceSSHUserCertificatesCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.SSHUserCertificateCreate
	if v, ok := d.GetOk("ssh_certificate_authority_id"); ok {
		arg.SSHCertificateAuthorityID = *expandString(v)
	}
	if v, ok := d.GetOk("public_key"); ok {
		arg.PublicKey = *expandString(v)
	}
	if v, ok := d.GetOk("principals"); ok {
		arg.Principals = *expandStringSlice(v)
	}
	if v, ok := d.GetOk("critical_options"); ok {
		arg.CriticalOptions = *expandStringMap(v)
	}
	if v, ok := d.GetOk("extensions"); ok {
		arg.Extensions = *expandStringMap(v)
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

	res, _, err := b.client.SSHUserCertificatesCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] SSHUserCertificatesCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceSSHUserCertificatesGet(d, m)
}

func resourceSSHUserCertificatesGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.SSHUserCertificatesGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceSSHUserCertificatesGetDecode(d, res, resp, err)
}

func resourceSSHUserCertificatesGetDecode(d *schema.ResourceData, res *restapi.SSHUserCertificate, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		log.Printf("[ERROR] SSHUserCertificatesGet: %s", err)
		return err
	default:
		d.Set("certificate", res.Certificate)
		d.Set("critical_options", res.CriticalOptions)
		d.Set("description", res.Description)
		d.Set("extensions", res.Extensions)
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

func resourceSSHUserCertificatesUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.SSHUserCertificateUpdate
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

	res, _, err := b.client.SSHUserCertificatesUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] SSHUserCertificatesUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceSSHUserCertificatesGet(d, m)
}

func resourceSSHUserCertificatesDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.SSHUserCertificatesDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] SSHUserCertificatesDelete: %s", err)
	}
	return err
}
