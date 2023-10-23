package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func resourceTunnelGroupBackends() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTunnelGroupBackendsCreate,
		Read:        resourceTunnelGroupBackendsGet,
		Update:      resourceTunnelGroupBackendsUpdate,
		Delete:      resourceTunnelGroupBackendsDelete,
		Description: "A Tunnel Group Backend balances traffic among all online tunnels that match\n a label selector.",
		Schema: map[string]*schema.Schema{
			"created_at": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "timestamp when the backend was created, RFC 3339 format",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "human-readable description of this backend. Optional",
			},
			"id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    false,
				Description: "unique identifier for this TunnelGroup backend",
			},
			"labels": {
				Type:        schema.TypeMap,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "labels to watch for tunnels on, e.g. app->foo, dc->bar",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"metadata": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "arbitrary user-defined machine-readable data of this backend. Optional",
			},
			"tunnels": {
				Type:        schema.TypeList,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "tunnels matching this backend",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    false,
							Computed:    true,
							Optional:    false,
							Sensitive:   false,
							ForceNew:    false,
							Description: "a resource identifier",
						},
						"uri": {
							Type:        schema.TypeString,
							Required:    false,
							Computed:    true,
							Optional:    false,
							Sensitive:   false,
							ForceNew:    true,
							Description: "a uri for locating a resource",
						},
					},
				},
			},
			"uri": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "URI of the TunnelGroupBackend API resource",
			},
		},
	}
}

func resourceTunnelGroupBackendsCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.TunnelGroupBackendCreate
	if v, ok := d.GetOk("description"); ok {
		arg.Description = *expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = *expandString(v)
	}
	if v, ok := d.GetOk("labels"); ok {
		arg.Labels = *expandStringMap(v)
	}

	res, _, err := b.client.TunnelGroupBackendsCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] TunnelGroupBackendsCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceTunnelGroupBackendsGet(d, m)
}

func resourceTunnelGroupBackendsGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.TunnelGroupBackendsGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceTunnelGroupBackendsGetDecode(d, res, resp, err)
}

func resourceTunnelGroupBackendsGetDecode(d *schema.ResourceData, res *restapi.TunnelGroupBackend, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		log.Printf("[ERROR] TunnelGroupBackendsGet: %s", err)
		return err
	default:
		d.Set("created_at", res.CreatedAt)
		d.Set("description", res.Description)
		d.Set("id", res.ID)
		d.Set("labels", res.Labels)
		d.Set("metadata", res.Metadata)
		d.Set("tunnels", flattenRefSlice(&res.Tunnels))
		d.Set("uri", res.URI)
	}
	return nil
}

func resourceTunnelGroupBackendsUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.TunnelGroupBackendUpdate
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
	if v, ok := d.GetOk("labels"); ok {
		arg.Labels = *expandStringMap(v)
	}

	res, _, err := b.client.TunnelGroupBackendsUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] TunnelGroupBackendsUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceTunnelGroupBackendsGet(d, m)
}

func resourceTunnelGroupBackendsDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.TunnelGroupBackendsDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] TunnelGroupBackendsDelete: %s", err)
	}
	return err
}
