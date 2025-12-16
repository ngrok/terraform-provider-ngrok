package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func resourceSecrets() *schema.Resource {
	return &schema.Resource{
		Create:      resourceSecretsCreate,
		Read:        resourceSecretsGet,
		Update:      resourceSecretsUpdate,
		Delete:      resourceSecretsDelete,
		Description: "Secrets is an api service for securely storing and managing sensitive data such as secrets, credentials, and tokens.",
		Schema: map[string]*schema.Schema{
			"created_by": {
				Type:        schema.TypeSet,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "Reference to who created this Secret",
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
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "description of Secret",
			},
			"id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    false,
				Description: "identifier for Secret",
			},
			"last_updated_by": {
				Type:        schema.TypeSet,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "Reference to who created this Secret",
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
			"metadata": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "Arbitrary user-defined metadata for this Secret",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "Name of secret",
			},
			"value": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "Value of secret",
			},
			"vault": {
				Type:        schema.TypeSet,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "Reference to the vault the secret is stored in",
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
			"vault_id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "unique identifier of the referenced vault",
			},
			"vault_name": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "Name of the vault the secret is stored in",
			},
		},
	}
}

func resourceSecretsCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.SecretCreate
	if v, ok := d.GetOk("name"); ok {
		arg.Name = *expandString(v)
	}
	if v, ok := d.GetOk("value"); ok {
		arg.Value = *expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = *expandString(v)
	}
	if v, ok := d.GetOk("description"); ok {
		arg.Description = *expandString(v)
	}
	if v, ok := d.GetOk("vault_id"); ok {
		arg.VaultID = *expandString(v)
	}
	if v, ok := d.GetOk("vault_name"); ok {
		arg.VaultName = *expandString(v)
	}

	res, _, err := b.client.SecretsCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] SecretsCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceSecretsGet(d, m)
}

func resourceSecretsGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.SecretsGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceSecretsGetDecode(d, res, resp, err)
}

func resourceSecretsGetDecode(d *schema.ResourceData, res *restapi.Secret, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		log.Printf("[ERROR] SecretsGet: %s", err)
		return err
	default:
		d.Set("created_by", flattenRef(&res.CreatedBy))
		d.Set("description", res.Description)
		d.Set("id", res.ID)
		d.Set("last_updated_by", flattenRef(&res.LastUpdatedBy))
		d.Set("metadata", res.Metadata)
		d.Set("name", res.Name)
		d.Set("vault", flattenRef(&res.Vault))
		d.Set("vault_id", res.Vault.ID)
		d.Set("vault_name", res.VaultName)
	}
	return nil
}

func resourceSecretsUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.SecretUpdate
	arg.ID = d.Id()
	if v, ok := d.GetOk("id"); ok {
		arg.ID = *expandString(v)
	}
	if v, ok := d.GetOk("name"); ok {
		arg.Name = expandString(v)
	}
	if v, ok := d.GetOk("value"); ok {
		arg.Value = expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = expandString(v)
	}
	if v, ok := d.GetOk("description"); ok {
		arg.Description = expandString(v)
	}

	res, _, err := b.client.SecretsUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] SecretsUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceSecretsGet(d, m)
}

func resourceSecretsDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.SecretsDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] SecretsDelete: %s", err)
	}
	return err
}
