package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func resourceVaults() *schema.Resource {
	return &schema.Resource{
		Create:      resourceVaultsCreate,
		Read:        resourceVaultsGet,
		Update:      resourceVaultsUpdate,
		Delete:      resourceVaultsDelete,
		Description: "Vaults is an api service for securely storing and managing sensitive data such as secrets, credentials, and tokens.",
		Schema: map[string]*schema.Schema{
			"created_by": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "Reference to who created this Vault",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "description of Vault",
			},
			"id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    false,
				Description: "identifier for Vault",
			},
			"last_updated_by": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "Reference to who created this Vault",
			},
			"metadata": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "Arbitrary user-defined metadata for this Vault",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "Name of vault",
			},
		},
	}
}

func resourceVaultsCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.VaultCreate
	if v, ok := d.GetOk("name"); ok {
		arg.Name = *expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = *expandString(v)
	}
	if v, ok := d.GetOk("description"); ok {
		arg.Description = *expandString(v)
	}

	res, _, err := b.client.VaultsCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] VaultsCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceVaultsGet(d, m)
}

func resourceVaultsGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.VaultsGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceVaultsGetDecode(d, res, resp, err)
}

func resourceVaultsGetDecode(d *schema.ResourceData, res *restapi.Vault, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		log.Printf("[ERROR] VaultsGet: %s", err)
		return err
	default:
		d.Set("created_by", res.CreatedBy)
		d.Set("description", res.Description)
		d.Set("id", res.ID)
		d.Set("last_updated_by", res.LastUpdatedBy)
		d.Set("metadata", res.Metadata)
		d.Set("name", res.Name)
	}
	return nil
}

func resourceVaultsUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.VaultUpdate
	arg.ID = d.Id()
	if v, ok := d.GetOk("id"); ok {
		arg.ID = *expandString(v)
	}
	if v, ok := d.GetOk("name"); ok {
		arg.Name = expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = expandString(v)
	}
	if v, ok := d.GetOk("description"); ok {
		arg.Description = expandString(v)
	}

	res, _, err := b.client.VaultsUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] VaultsUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceVaultsGet(d, m)
}

func resourceVaultsDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.VaultsDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] VaultsDelete: %s", err)
	}
	return err
}
