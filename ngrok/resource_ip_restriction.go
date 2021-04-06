// Code generated by apic. DO NOT EDIT.

package ngrok

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
	transform "github.com/ngrok/terraform-provider-ngrok/transform"
)

func resourceIPRestrictions() *schema.Resource {
	return &schema.Resource{
		Create: resourceIPRestrictionsCreate,
		Read:   resourceIPRestrictionsGet,
		Update: resourceIPRestrictionsUpdate,
		Delete: resourceIPRestrictionsDelete,

		Schema: map[string]*schema.Schema{
			"created_at": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "timestamp when the IP restriction was created, RFC 3339 format",
			},
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
				Description: "true if the IP restriction will be enforce. if false, only warnings will be issued",
			},
			"ip_policy_ids": {
				Type:        schema.TypeList,
				Required:    false,
				Computed:    false,
				Optional:    true,
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
			"ngrok_id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "unique identifier for this IP restriction",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "the type of IP restriction. this defines what traffic will be restricted with the attached policies. four values are currently supported: dashboard, api, agent, and endpoints",
			},
			"uri": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "URI of the IP restriction API resource",
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
	if err == nil {
		d.SetId(res.ID)
	}
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
		return err
	default:
		d.Set("created_at", res.CreatedAt)
		d.Set("description", res.Description)
		d.Set("enforced", res.Enforced)
		d.Set("ip_policy_ids", transform.ConvertRefSliceToStringSlice(&res.IPPolicies))
		d.Set("metadata", res.Metadata)
		d.Set("ngrok_id", res.ID)
		d.Set("type", res.Type)
		d.Set("uri", res.URI)
	}
	return nil
}

func resourceIPRestrictionsUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.IPRestrictionUpdate
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
	if v, ok := d.GetOk("enforced"); ok {
		arg.Enforced = expandBool(v)
	}
	if v, ok := d.GetOk("ip_policy_ids"); ok {
		arg.IPPolicyIDs = *expandStringSlice(v)
	}

	res, _, err := b.client.IPRestrictionsUpdate(context.Background(), &arg)
	if err != nil {
		return err
	}
	d.SetId(res.ID)

	return resourceIPRestrictionsGet(d, m)
}

func resourceIPRestrictionsDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.IPRestrictionsDelete(context.Background(), &restapi.Item{ID: d.Id()})
	return err
}
