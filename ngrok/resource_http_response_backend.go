package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func resourceHTTPResponseBackends() *schema.Resource {
	return &schema.Resource{
		Create: resourceHTTPResponseBackendsCreate,
		Read:   resourceHTTPResponseBackendsGet,
		Update: resourceHTTPResponseBackendsUpdate,
		Delete: resourceHTTPResponseBackendsDelete,

		Schema: map[string]*schema.Schema{
			"body": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "body to return as fixed content",
			},
			"created_at": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
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
			"headers": {
				Type:        schema.TypeMap,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "headers to return",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    false,
				Description: "",
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
			"status_code": {
				Type:        schema.TypeInt,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "status code to return",
			},
			"uri": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "URI of the HTTPResponseBackend API resource",
			},
		},
	}
}

func resourceHTTPResponseBackendsCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.HTTPResponseBackendCreate
	if v, ok := d.GetOk("description"); ok {
		arg.Description = *expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = *expandString(v)
	}
	if v, ok := d.GetOk("body"); ok {
		arg.Body = *expandString(v)
	}
	if v, ok := d.GetOk("headers"); ok {
		arg.Headers = *expandStringMap(v)
	}
	if v, ok := d.GetOk("status_code"); ok {
		arg.StatusCode = expandInt32(v)
	}

	res, _, err := b.client.HTTPResponseBackendsCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] HTTPResponseBackendsCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceHTTPResponseBackendsGet(d, m)
}

func resourceHTTPResponseBackendsGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.HTTPResponseBackendsGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceHTTPResponseBackendsGetDecode(d, res, resp, err)
}

func resourceHTTPResponseBackendsGetDecode(d *schema.ResourceData, res *restapi.HTTPResponseBackend, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		log.Printf("[ERROR] HTTPResponseBackendsGet: %s", err)
		return err
	default:
		d.Set("body", res.Body)
		d.Set("created_at", res.CreatedAt)
		d.Set("description", res.Description)
		d.Set("headers", res.Headers)
		d.Set("id", res.ID)
		d.Set("metadata", res.Metadata)
		d.Set("status_code", res.StatusCode)
		d.Set("uri", res.URI)
	}
	return nil
}

func resourceHTTPResponseBackendsUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.HTTPResponseBackendUpdate
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
	if v, ok := d.GetOk("body"); ok {
		arg.Body = expandString(v)
	}
	if v, ok := d.GetOk("headers"); ok {
		arg.Headers = expandStringMap(v)
	}
	if v, ok := d.GetOk("status_code"); ok {
		arg.StatusCode = expandInt32(v)
	}

	res, _, err := b.client.HTTPResponseBackendsUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] HTTPResponseBackendsUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceHTTPResponseBackendsGet(d, m)
}

func resourceHTTPResponseBackendsDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.HTTPResponseBackendsDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] HTTPResponseBackendsDelete: %s", err)
	}
	return err
}
