package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func resourceIPRestrictions() *schema.Resource {
	return &schema.Resource{
		Create:      resourceIPRestrictionsCreate,
		Read:        resourceIPRestrictionsGet,
		Update:      resourceIPRestrictionsUpdate,
		Delete:      resourceIPRestrictionsDelete,
		Description: "An IP restriction is a restriction placed on the CIDRs that are allowed to\n initiate traffic to a specific aspect of your ngrok account. An IP\n restriction has a type which defines the ingress it applies to. IP\n restrictions can be used to enforce the source IPs that can make API\n requests, log in to the dashboard, start ngrok agents, and connect to your\n public-facing endpoints.",
		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "human-readable description of this IP restriction. optional, max 255 bytes.",
			},
			"enforced": {
				Type:        schema.TypeBool,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "true if the IP restriction will be enforced. if false, only warnings will be issued",
			},
			"id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    false,
				Description: "unique identifier for this IP restriction",
			},
			"ip_policy_ids": {
				Type:        schema.TypeList,
				Required:    true,
				Computed:    false,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    false,
				Description: "the set of IP policy identifiers that are used to enforce the restriction",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"metadata": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "arbitrary user-defined machine-readable data of this IP restriction. optional, max 4096 bytes.",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Computed:    false,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "the type of IP restriction. this defines what traffic will be restricted with the attached policies. four values are currently supported: `dashboard`, `api`, `agent`, and `endpoints`",
			},
		},
	}
}

func resourceIPRestrictionsCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.IPRestrictionCreate
	if v, ok := d.GetOk("description"); ok {
		arg.Description = *expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = *expandString(v)
	}
	if v, ok := d.GetOk("enforced"); ok {
		arg.Enforced = *expandBool(v)
	}
	if v, ok := d.GetOk("type"); ok {
		arg.Type = *expandString(v)
	}
	if v, ok := d.GetOk("ip_policy_ids"); ok {
		arg.IPPolicyIDs = *expandStringSlice(v)
	}

	res, _, err := b.client.IPRestrictionsCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] IPRestrictionsCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceIPRestrictionsGet(d, m)
}

func resourceIPRestrictionsGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.IPRestrictionsGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceIPRestrictionsGetDecode(d, res, resp, err)
}

func resourceIPRestrictionsGetDecode(d *schema.ResourceData, res *restapi.IPRestriction, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		log.Printf("[ERROR] IPRestrictionsGet: %s", err)
		return err
	default:
		d.Set("description", res.Description)
		d.Set("enforced", res.Enforced)
		d.Set("id", res.ID)
		d.Set("metadata", res.Metadata)
		d.Set("type", res.Type)
	}
	return nil
}

func resourceIPRestrictionsUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.IPRestrictionUpdate
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
	if v, ok := d.GetOk("enforced"); ok {
		arg.Enforced = expandBool(v)
	}
	if v, ok := d.GetOk("ip_policy_ids"); ok {
		arg.IPPolicyIDs = *expandStringSlice(v)
	}

	res, _, err := b.client.IPRestrictionsUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] IPRestrictionsUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceIPRestrictionsGet(d, m)
}

func resourceIPRestrictionsDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.IPRestrictionsDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] IPRestrictionsDelete: %s", err)
	}
	return err
}
