package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func resourceAPIKeys() *schema.Resource {
	return &schema.Resource{
		Create:      resourceAPIKeysCreate,
		Read:        resourceAPIKeysGet,
		Update:      resourceAPIKeysUpdate,
		Delete:      resourceAPIKeysDelete,
		Description: "API Keys are used to authenticate to the [ngrok\n API](https://ngrok.com/docs/api#authentication). You may use the API itself\n to provision and manage API Keys but you'll need to provision your first API\n key from the [API Keys page](https://dashboard.ngrok.com/api/keys) on your\n ngrok.com dashboard.",
		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "human-readable description of what uses the API key to authenticate. optional, max 255 bytes.",
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
			"metadata": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "arbitrary user-defined data of this API key. optional, max 4096 bytes",
			},
			"owner_id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "If supplied at credential creation, ownership will be assigned to the specified User or Bot. Only admins may specify an owner other than themselves. Defaults to the authenticated User or Bot.",
			},
			"token": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   true,
				ForceNew:    true,
				Description: "the bearer token that can be placed into the Authorization header to authenticate request to the ngrok API. **This value is only available one time, on the API response from key creation. Otherwise it is null.**",
			},
		},
	}
}

func resourceAPIKeysCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.APIKeyCreate
	if v, ok := d.GetOk("description"); ok {
		arg.Description = *expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = *expandString(v)
	}
	if v, ok := d.GetOk("owner_id"); ok {
		arg.OwnerID = expandString(v)
	}
	if v, ok := d.GetOk("owner_email"); ok {
		arg.OwnerEmail = *expandString(v)
	}

	res, _, err := b.client.APIKeysCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] APIKeysCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceAPIKeysGet(d, m)
}

func resourceAPIKeysGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.APIKeysGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceAPIKeysGetDecode(d, res, resp, err)
}

func resourceAPIKeysGetDecode(d *schema.ResourceData, res *restapi.APIKey, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		log.Printf("[ERROR] APIKeysGet: %s", err)
		return err
	default:
		d.Set("description", res.Description)
		d.Set("id", res.ID)
		d.Set("metadata", res.Metadata)
		d.Set("owner_id", res.OwnerID)
		d.Set("token", res.Token)
	}
	return nil
}

func resourceAPIKeysUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.APIKeyUpdate
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

	res, _, err := b.client.APIKeysUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] APIKeysUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceAPIKeysGet(d, m)
}

func resourceAPIKeysDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.APIKeysDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] APIKeysDelete: %s", err)
	}
	return err
}
