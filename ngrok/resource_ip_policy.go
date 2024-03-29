package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func resourceIPPolicies() *schema.Resource {
	return &schema.Resource{
		Create:      resourceIPPoliciesCreate,
		Read:        resourceIPPoliciesGet,
		Update:      resourceIPPoliciesUpdate,
		Delete:      resourceIPPoliciesDelete,
		Description: "IP Policies are reusable groups of CIDR ranges with an `allow` or `deny`\n action. They can be attached to endpoints via the Endpoint Configuration IP\n Policy module. They can also be used with IP Restrictions to control source\n IP ranges that can start tunnel sessions and connect to the API and dashboard.",
		Schema: map[string]*schema.Schema{
			"action": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "this field is deprecated. Please leave it empty and use the ip policy rule object's \"action\" field instead. It is temporarily retained for backwards compatibility reasons.",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "human-readable description of the source IPs of this IP policy. optional, max 255 bytes.",
			},
			"id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    false,
				Description: "unique identifier for this IP policy",
			},
			"metadata": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "arbitrary user-defined machine-readable data of this IP policy. optional, max 4096 bytes.",
			},
		},
	}
}

func resourceIPPoliciesCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.IPPolicyCreate
	if v, ok := d.GetOk("description"); ok {
		arg.Description = *expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = *expandString(v)
	}
	if v, ok := d.GetOk("action"); ok {
		arg.Action = expandString(v)
	}

	res, _, err := b.client.IPPoliciesCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] IPPoliciesCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceIPPoliciesGet(d, m)
}

func resourceIPPoliciesGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.IPPoliciesGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceIPPoliciesGetDecode(d, res, resp, err)
}

func resourceIPPoliciesGetDecode(d *schema.ResourceData, res *restapi.IPPolicy, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		log.Printf("[ERROR] IPPoliciesGet: %s", err)
		return err
	default:
		d.Set("action", res.Action)
		d.Set("description", res.Description)
		d.Set("id", res.ID)
		d.Set("metadata", res.Metadata)
	}
	return nil
}

func resourceIPPoliciesUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.IPPolicyUpdate
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

	res, _, err := b.client.IPPoliciesUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] IPPoliciesUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceIPPoliciesGet(d, m)
}

func resourceIPPoliciesDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.IPPoliciesDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] IPPoliciesDelete: %s", err)
	}
	return err
}
