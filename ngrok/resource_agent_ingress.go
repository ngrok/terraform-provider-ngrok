package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func resourceAgentIngresses() *schema.Resource {
	return &schema.Resource{
		Create: resourceAgentIngressesCreate,
		Read:   resourceAgentIngressesGet,
		Update: resourceAgentIngressesUpdate,
		Delete: resourceAgentIngressesDelete,

		Schema: map[string]*schema.Schema{
			"created_at": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "timestamp when the Agent Ingress was created, RFC 3339 format",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "human-readable description of the use of this Agent Ingress. optional, max 255 bytes.",
			},
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				Computed:    false,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "the domain that you own to be used as the base domain name to generate regional agent ingress domains.",
			},
			"id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    false,
				Description: "unique Agent Ingress resource identifier",
			},
			"metadata": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "arbitrary user-defined machine-readable data of this Agent Ingress. optional, max 4096 bytes",
			},
			"ns_targets": {
				Type:        schema.TypeList,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "a list of target values to use as the values of NS records for the domain property these values will delegate control over the domain to ngrok",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"region_domains": {
				Type:        schema.TypeList,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "a list of regional agent ingress domains that are subdomains of the value of domain this value may increase over time as ngrok adds more regions",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"uri": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "URI to the API resource of this Agent ingress",
			},
		},
	}
}

func resourceAgentIngressesCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.AgentIngressCreate
	if v, ok := d.GetOk("description"); ok {
		arg.Description = *expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = *expandString(v)
	}
	if v, ok := d.GetOk("domain"); ok {
		arg.Domain = *expandString(v)
	}

	res, _, err := b.client.AgentIngressesCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] AgentIngressesCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceAgentIngressesGet(d, m)
}

func resourceAgentIngressesGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.AgentIngressesGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceAgentIngressesGetDecode(d, res, resp, err)
}

func resourceAgentIngressesGetDecode(d *schema.ResourceData, res *restapi.AgentIngress, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		log.Printf("[ERROR] AgentIngressesGet: %s", err)
		return err
	default:
		d.Set("created_at", res.CreatedAt)
		d.Set("description", res.Description)
		d.Set("domain", res.Domain)
		d.Set("id", res.ID)
		d.Set("metadata", res.Metadata)
		d.Set("ns_targets", res.NSTargets)
		d.Set("region_domains", res.RegionDomains)
		d.Set("uri", res.URI)
	}
	return nil
}

func resourceAgentIngressesUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.AgentIngressUpdate
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

	res, _, err := b.client.AgentIngressesUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] AgentIngressesUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceAgentIngressesGet(d, m)
}

func resourceAgentIngressesDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.AgentIngressesDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] AgentIngressesDelete: %s", err)
	}
	return err
}
