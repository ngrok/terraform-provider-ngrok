package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func resourceBotUsers() *schema.Resource {
	return &schema.Resource{
		Create: resourceBotUsersCreate,
		Read:   resourceBotUsersGet,
		Update: resourceBotUsersUpdate,
		Delete: resourceBotUsersDelete,

		Schema: map[string]*schema.Schema{
			"active": {
				Type:        schema.TypeBool,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "whether or not the bot is active",
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
				Description: "human-readable name used to identify the bot",
			},
		},
	}
}

func resourceBotUsersCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.BotUserCreate
	if v, ok := d.GetOk("name"); ok {
		arg.Name = *expandString(v)
	}
	if v, ok := d.GetOk("active"); ok {
		arg.Active = expandBool(v)
	}

	res, _, err := b.client.BotUsersCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] BotUsersCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceBotUsersGet(d, m)
}

func resourceBotUsersGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.BotUsersGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceBotUsersGetDecode(d, res, resp, err)
}

func resourceBotUsersGetDecode(d *schema.ResourceData, res *restapi.BotUser, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		log.Printf("[ERROR] BotUsersGet: %s", err)
		return err
	default:
		d.Set("active", res.Active)
		d.Set("id", res.ID)
		d.Set("name", res.Name)
	}
	return nil
}

func resourceBotUsersUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.BotUserUpdate
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

	res, _, err := b.client.BotUsersUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] BotUsersUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceBotUsersGet(d, m)
}

func resourceBotUsersDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.BotUsersDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] BotUsersDelete: %s", err)
	}
	return err
}
