package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func resourceFailoverBackends() *schema.Resource {
	return &schema.Resource{
		Create:      resourceFailoverBackendsCreate,
		Read:        resourceFailoverBackendsGet,
		Update:      resourceFailoverBackendsUpdate,
		Delete:      resourceFailoverBackendsDelete,
		Description: "A Failover backend defines failover behavior within a list of referenced\n backends. Traffic is sent to the first backend in the list. If that backend\n is offline or no connection can be established, ngrok attempts to connect to\n the next backend in the list until one is successful.",
		Schema: map[string]*schema.Schema{
			"backends": {
				Type:        schema.TypeList,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "the ids of the child backends in order",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
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
				Description: "unique identifier for this Failover backend",
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
			"uri": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "URI of the FailoverBackend API resource",
			},
		},
	}
}

func resourceFailoverBackendsCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.FailoverBackendCreate
	if v, ok := d.GetOk("description"); ok {
		arg.Description = *expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = *expandString(v)
	}
	if v, ok := d.GetOk("backends"); ok {
		arg.Backends = *expandStringSlice(v)
	}

	res, _, err := b.client.FailoverBackendsCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] FailoverBackendsCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceFailoverBackendsGet(d, m)
}

func resourceFailoverBackendsGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.FailoverBackendsGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceFailoverBackendsGetDecode(d, res, resp, err)
}

func resourceFailoverBackendsGetDecode(d *schema.ResourceData, res *restapi.FailoverBackend, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		log.Printf("[ERROR] FailoverBackendsGet: %s", err)
		return err
	default:
		d.Set("backends", res.Backends)
		d.Set("created_at", res.CreatedAt)
		d.Set("description", res.Description)
		d.Set("id", res.ID)
		d.Set("metadata", res.Metadata)
		d.Set("uri", res.URI)
	}
	return nil
}

func resourceFailoverBackendsUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.FailoverBackendUpdate
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
	if v, ok := d.GetOk("backends"); ok {
		arg.Backends = *expandStringSlice(v)
	}

	res, _, err := b.client.FailoverBackendsUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] FailoverBackendsUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceFailoverBackendsGet(d, m)
}

func resourceFailoverBackendsDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.FailoverBackendsDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] FailoverBackendsDelete: %s", err)
	}
	return err
}
