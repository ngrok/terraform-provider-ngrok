package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func resourceEventSubscriptions() *schema.Resource {
	return &schema.Resource{
		Create: resourceEventSubscriptionsCreate,
		Read:   resourceEventSubscriptionsGet,
		Update: resourceEventSubscriptionsUpdate,
		Delete: resourceEventSubscriptionsDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "Arbitrary customer supplied information intended to be human readable. Optional, max 255 chars.",
			},
			"destination_ids": {
				Type:        schema.TypeList,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "A list of Event Destination IDs which should be used for this Event Subscription.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "Unique identifier for this Event Subscription.",
			},
			"metadata": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "Arbitrary customer supplied information intended to be machine readable. Optional, max 4096 chars.",
			},
			"sources": {
				Type:        schema.TypeList,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "Sources containing the types for which this event subscription will trigger",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Required:    false,
							Computed:    false,
							Optional:    true,
							Sensitive:   false,
							ForceNew:    false,
							Description: "Type of event for which an event subscription will trigger",
						},
						"filter": {
							Type:        schema.TypeString,
							Required:    false,
							Computed:    false,
							Optional:    true,
							Sensitive:   false,
							ForceNew:    false,
							Description: "TODO",
						},
						"fields": {
							Type:        schema.TypeList,
							Required:    false,
							Computed:    false,
							Optional:    true,
							Sensitive:   false,
							ForceNew:    false,
							Description: "TODO",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"uri": {
							Type:        schema.TypeString,
							Required:    false,
							Computed:    true,
							Optional:    true,
							Sensitive:   false,
							ForceNew:    true,
							Description: "URI of the Event Source API resource.",
						},
					},
				},
			},
		},
	}
}

func resourceEventSubscriptionsCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.EventSubscriptionCreate
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = *expandString(v)
	}
	if v, ok := d.GetOk("description"); ok {
		arg.Description = *expandString(v)
	}
	if v, ok := d.GetOk("sources"); ok {
		arg.Sources = *expandEventSourceReplaceSlice(v)
	}
	if v, ok := d.GetOk("destination_ids"); ok {
		arg.DestinationIDs = *expandStringSlice(v)
	}

	res, _, err := b.client.EventSubscriptionsCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] EventSubscriptionsCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceEventSubscriptionsGet(d, m)
}

func resourceEventSubscriptionsGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.EventSubscriptionsGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceEventSubscriptionsGetDecode(d, res, resp, err)
}

func resourceEventSubscriptionsGetDecode(d *schema.ResourceData, res *restapi.EventSubscription, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		log.Printf("[ERROR] EventSubscriptionsGet: %s", err)
		return err
	default:
		d.Set("description", res.Description)
		d.Set("id", res.ID)
		d.Set("metadata", res.Metadata)
		d.Set("sources", flattenEventSourceSlice(&res.Sources))
	}
	return nil
}

func resourceEventSubscriptionsUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.EventSubscriptionUpdate
	arg.ID = d.Id()
	if v, ok := d.GetOk("id"); ok {
		arg.ID = *expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = expandString(v)
	}
	if v, ok := d.GetOk("description"); ok {
		arg.Description = expandString(v)
	}
	if v, ok := d.GetOk("sources"); ok {
		arg.Sources = expandEventSourceReplaceSlice(v)
	}
	if v, ok := d.GetOk("destination_ids"); ok {
		arg.DestinationIDs = expandStringSlice(v)
	}

	res, _, err := b.client.EventSubscriptionsUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] EventSubscriptionsUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceEventSubscriptionsGet(d, m)
}

func resourceEventSubscriptionsDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.EventSubscriptionsDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] EventSubscriptionsDelete: %s", err)
	}
	return err
}
