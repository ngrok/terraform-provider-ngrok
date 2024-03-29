package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func resourceSSHCredentials() *schema.Resource {
	return &schema.Resource{
		Create:      resourceSSHCredentialsCreate,
		Read:        resourceSSHCredentialsGet,
		Update:      resourceSSHCredentialsUpdate,
		Delete:      resourceSSHCredentialsDelete,
		Description: "SSH Credentials are SSH public keys that can be used to start SSH tunnels\n via the ngrok SSH tunnel gateway.",
		Schema: map[string]*schema.Schema{
			"acl": {
				Type:        schema.TypeList,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "optional list of ACL rules. If unspecified, the credential will have no restrictions. The only allowed ACL rule at this time is the `bind` rule. The `bind` rule allows the caller to restrict what domains, addresses, and labels the token is allowed to bind. For example, to allow the token to open a tunnel on example.ngrok.io your ACL would include the rule `bind:example.ngrok.io`. Bind rules for domains may specify a leading wildcard to match multiple domains with a common suffix. For example, you may specify a rule of `bind:*.example.com` which will allow `x.example.com`, `y.example.com`, `*.example.com`, etc. Bind rules for labels may specify a wildcard key and/or value to match multiple labels. For example, you may specify a rule of `bind:*=example` which will allow `x=example`, `y=example`, etc. A rule of `'*'` is equivalent to no acl at all and will explicitly permit all actions.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "human-readable description of who or what will use the ssh credential to authenticate. Optional, max 255 bytes.",
			},
			"id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    false,
				Description: "unique ssh credential resource identifier",
			},
			"metadata": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "arbitrary user-defined machine-readable data of this ssh credential. Optional, max 4096 bytes.",
			},
			"owner_id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "If supplied at credential creation, ownership will be assigned to the specified User or Bot. Only admins may specify an owner other than themselves. Defaults to the authenticated User or Bot.",
			},
			"public_key": {
				Type:        schema.TypeString,
				Required:    true,
				Computed:    false,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "the PEM-encoded public key of the SSH keypair that will be used to authenticate",
			},
		},
	}
}

func resourceSSHCredentialsCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.SSHCredentialCreate
	if v, ok := d.GetOk("description"); ok {
		arg.Description = *expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = *expandString(v)
	}
	if v, ok := d.GetOk("acl"); ok {
		arg.ACL = *expandStringSlice(v)
	}
	if v, ok := d.GetOk("public_key"); ok {
		arg.PublicKey = *expandString(v)
	}
	if v, ok := d.GetOk("owner_id"); ok {
		arg.OwnerID = expandString(v)
	}
	if v, ok := d.GetOk("owner_email"); ok {
		arg.OwnerEmail = *expandString(v)
	}

	res, _, err := b.client.SSHCredentialsCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] SSHCredentialsCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceSSHCredentialsGet(d, m)
}

func resourceSSHCredentialsGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.SSHCredentialsGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceSSHCredentialsGetDecode(d, res, resp, err)
}

func resourceSSHCredentialsGetDecode(d *schema.ResourceData, res *restapi.SSHCredential, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		log.Printf("[ERROR] SSHCredentialsGet: %s", err)
		return err
	default:
		d.Set("acl", res.ACL)
		d.Set("description", res.Description)
		d.Set("id", res.ID)
		d.Set("metadata", res.Metadata)
		d.Set("owner_id", res.OwnerID)
		d.Set("public_key", res.PublicKey)
	}
	return nil
}

func resourceSSHCredentialsUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.SSHCredentialUpdate
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
	if v, ok := d.GetOk("acl"); ok {
		arg.ACL = expandStringSlice(v)
	}

	res, _, err := b.client.SSHCredentialsUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] SSHCredentialsUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceSSHCredentialsGet(d, m)
}

func resourceSSHCredentialsDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.SSHCredentialsDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] SSHCredentialsDelete: %s", err)
	}
	return err
}
