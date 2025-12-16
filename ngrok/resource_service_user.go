package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func resourceServiceUsers() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceUsersCreate,
		Read:   resourceServiceUsersGet,
		Update: resourceServiceUsersUpdate,
		Delete: resourceServiceUsersDelete,

		Schema: map[string]*schema.Schema{
			"active": {
				Type:        schema.TypeBool,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "whether or not the service is active",
			},
			"id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    false,
				Description: "unique API key resource identifier",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "human-readable name used to identify the service",
			},
		},
	}
}

func resourceServiceUsersCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.ServiceUserCreate
	if v, ok := d.GetOk("name"); ok {
		arg.Name = *expandString(v)
	}
	if v, ok := d.GetOk("active"); ok {
		arg.Active = expandBool(v)
	}

	res, _, err := b.client.ServiceUsersCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] ServiceUsersCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceServiceUsersGet(d, m)
}

func resourceServiceUsersGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.ServiceUsersGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceServiceUsersGetDecode(d, res, resp, err)
}

func resourceServiceUsersGetDecode(d *schema.ResourceData, res *restapi.ServiceUser, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		log.Printf("[ERROR] ServiceUsersGet: %s", err)
		return err
	default:
		d.Set("active", res.Active)
		d.Set("id", res.ID)
		d.Set("name", res.Name)
	}
	return nil
}

func resourceServiceUsersUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.ServiceUserUpdate
	arg.ID = d.Id()
	if v, ok := d.GetOk("id"); ok {
		arg.ID = *expandString(v)
	}
	if v, ok := d.GetOk("name"); ok {
		arg.Name = expandString(v)
	}
	if v, ok := d.GetOk("active"); ok {
		arg.Active = expandBool(v)
	}

	res, _, err := b.client.ServiceUsersUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] ServiceUsersUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceServiceUsersGet(d, m)
}

func resourceServiceUsersDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.ServiceUsersDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] ServiceUsersDelete: %s", err)
	}
	return err
}
