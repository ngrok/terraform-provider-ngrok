package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func resourceIPPolicyRules() *schema.Resource {
	return &schema.Resource{
		Create:      resourceIPPolicyRulesCreate,
		Read:        resourceIPPolicyRulesGet,
		Update:      resourceIPPolicyRulesUpdate,
		Delete:      resourceIPPolicyRulesDelete,
		Description: "IP Policy Rules are the IPv4 or IPv6 CIDRs entries that\n make up an IP Policy.",
		Schema: map[string]*schema.Schema{
			"action": {
				Type:        schema.TypeString,
				Required:    true,
				Computed:    false,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "the action to apply to the policy rule, either `allow` or `deny`",
			},
			"cidr": {
				Type:        schema.TypeString,
				Required:    true,
				Computed:    false,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    false,
				Description: "an IP or IP range specified in CIDR notation. IPv4 and IPv6 are both supported.",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "human-readable description of the source IPs of this IP rule. optional, max 255 bytes.",
			},
			"id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "unique identifier for this IP policy rule",
			},
			"ip_policy_id": {
				Type:        schema.TypeString,
				Required:    true,
				Computed:    false,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "ID of the IP policy this IP policy rule will be attached to",
			},
			"metadata": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "arbitrary user-defined machine-readable data of this IP policy rule. optional, max 4096 bytes.",
			},
		},
	}
}

func resourceIPPolicyRulesCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.IPPolicyRuleCreate
	if v, ok := d.GetOk("description"); ok {
		arg.Description = *expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = *expandString(v)
	}
	if v, ok := d.GetOk("cidr"); ok {
		arg.CIDR = *expandString(v)
	}
	if v, ok := d.GetOk("ip_policy_id"); ok {
		arg.IPPolicyID = *expandString(v)
	}
	if v, ok := d.GetOk("action"); ok {
		arg.Action = expandString(v)
	}

	res, _, err := b.client.IPPolicyRulesCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] IPPolicyRulesCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceIPPolicyRulesGet(d, m)
}

func resourceIPPolicyRulesGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.IPPolicyRulesGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceIPPolicyRulesGetDecode(d, res, resp, err)
}

func resourceIPPolicyRulesGetDecode(d *schema.ResourceData, res *restapi.IPPolicyRule, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		log.Printf("[ERROR] IPPolicyRulesGet: %s", err)
		return err
	default:
		d.Set("action", res.Action)
		d.Set("cidr", res.CIDR)
		d.Set("description", res.Description)
		d.Set("id", res.ID)
		d.Set("ip_policy_id", res.IPPolicy.ID)
		d.Set("metadata", res.Metadata)
	}
	return nil
}

func resourceIPPolicyRulesUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.IPPolicyRuleUpdate
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
	if v, ok := d.GetOk("cidr"); ok {
		arg.CIDR = expandString(v)
	}

	res, _, err := b.client.IPPolicyRulesUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] IPPolicyRulesUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceIPPolicyRulesGet(d, m)
}

func resourceIPPolicyRulesDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.IPPolicyRulesDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] IPPolicyRulesDelete: %s", err)
	}
	return err
}
