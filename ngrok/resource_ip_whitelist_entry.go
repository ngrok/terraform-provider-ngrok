// Code generated by apic. DO NOT EDIT.

package ngrok

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func resourceIPWhitelist() *schema.Resource {
	return &schema.Resource{
		Create: resourceIPWhitelistCreate,
		Read:   resourceIPWhitelistGet,
		Update: resourceIPWhitelistUpdate,
		Delete: resourceIPWhitelistDelete,

		Schema: map[string]*schema.Schema{
			"created_at": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "timestamp when the IP whitelist entry was created, RFC 3339 format",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "human-readable description of the source IPs for this IP whitelist entry. optional, max 255 bytes.",
			},
			"ip_net": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "an IP address or IP network range in CIDR notation (e.g. 10.1.1.1 or 10.1.0.0/16) of addresses that will be whitelisted to communicate with your tunnel endpoints",
			},
			"metadata": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "arbitrary user-defined machine-readable data of this IP whitelist entry. optional, max 4096 bytes.",
			},
			"ngrok_id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "unique identifier for this IP whitelist entry",
			},
			"uri": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "URI of the IP whitelist entry API resource",
			},
		},
	}
}

func resourceIPWhitelistCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.IPWhitelistEntryCreate
	if v, ok := d.GetOk("description"); ok {
		arg.Description = *expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = *expandString(v)
	}
	if v, ok := d.GetOk("ip_net"); ok {
		arg.IPNet = *expandString(v)
	}

	res, _, err := b.client.IPWhitelistCreate(context.Background(), &arg)
	if err == nil {
		d.SetId(res.ID)
	}
	return resourceIPWhitelistGet(d, m)
}

func resourceIPWhitelistGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.IPWhitelistGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceIPWhitelistGetDecode(d, res, resp, err)
}

func resourceIPWhitelistGetDecode(d *schema.ResourceData, res *restapi.IPWhitelistEntry, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		return err
	default:
		d.Set("created_at", res.CreatedAt)
		d.Set("description", res.Description)
		d.Set("ip_net", res.IPNet)
		d.Set("metadata", res.Metadata)
		d.Set("ngrok_id", res.ID)
		d.Set("uri", res.URI)
	}
	return nil
}

func resourceIPWhitelistUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.IPWhitelistEntryUpdate
	arg.ID = d.Id()
	if v, ok := d.GetOk("ngrok_id"); ok {
		arg.ID = *expandString(v)
	}
	if v, ok := d.GetOk("description"); ok {
		arg.Description = expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = expandString(v)
	}

	res, _, err := b.client.IPWhitelistUpdate(context.Background(), &arg)
	if err != nil {
		return err
	}
	d.SetId(res.ID)

	return resourceIPWhitelistGet(d, m)
}

func resourceIPWhitelistDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.IPWhitelistDelete(context.Background(), &restapi.Item{ID: d.Id()})
	return err
}
