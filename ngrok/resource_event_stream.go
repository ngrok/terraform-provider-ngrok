// Code generated by apic. DO NOT EDIT.

package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func resourceEventStreams() *schema.Resource {
	return &schema.Resource{
		Create: resourceEventStreamsCreate,
		Read:   resourceEventStreamsGet,
		Update: resourceEventStreamsUpdate,
		Delete: resourceEventStreamsDelete,

		Schema: map[string]*schema.Schema{
			"created_at": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "Timestamp when the Event Stream was created, RFC 3339 format.",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "Human-readable description of the Event Stream. Optional, max 255 bytes.",
			},
			"destination_ids": {
				Type:        schema.TypeList,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "A list of Event Destination IDs which should be used for this Event Stream. Event Streams are required to have at least one Event Destination.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"event_type": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "The protocol that determines which events will be collected. Supported values are `tcp_connection_closed` and `http_request_complete`.",
			},
			"fields": {
				Type:        schema.TypeList,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "A list of protocol-specific fields you want to collect on each event.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"metadata": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "Arbitrary user-defined machine-readable data of this Event Stream. Optional, max 4096 bytes.",
			},
			"ngrok_id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    false,
				Description: "Unique identifier for this Event Stream.",
			},
			"sampling_rate": {
				Type:        schema.TypeFloat,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "The percentage of all events you would like to capture. Valid values range from 0.01, representing 1% of all events to 1.00, representing 100% of all events.",
			},
			"uri": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "URI of the Event Stream API resource.",
			},
		},
	}
}

func resourceEventStreamsCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.EventStreamCreate
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = *expandString(v)
	}
	if v, ok := d.GetOk("description"); ok {
		arg.Description = *expandString(v)
	}
	if v, ok := d.GetOk("fields"); ok {
		arg.Fields = *expandStringSlice(v)
	}
	if v, ok := d.GetOk("event_type"); ok {
		arg.EventType = *expandString(v)
	}
	if v, ok := d.GetOk("destination_ids"); ok {
		arg.DestinationIDs = *expandStringSlice(v)
	}
	if v, ok := d.GetOk("sampling_rate"); ok {
		arg.SamplingRate = *expandFloat64(v)
	}

	res, _, err := b.client.EventStreamsCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] EventStreamsCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceEventStreamsGet(d, m)
}

func resourceEventStreamsGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.EventStreamsGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceEventStreamsGetDecode(d, res, resp, err)
}

func resourceEventStreamsGetDecode(d *schema.ResourceData, res *restapi.EventStream, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		log.Printf("[ERROR] EventStreamsGet: %s", err)
		return err
	default:
		d.Set("created_at", res.CreatedAt)
		d.Set("description", res.Description)
		d.Set("destination_ids", res.DestinationIDs)
		d.Set("event_type", res.EventType)
		d.Set("fields", res.Fields)
		d.Set("metadata", res.Metadata)
		d.Set("ngrok_id", res.ID)
		d.Set("sampling_rate", res.SamplingRate)
		d.Set("uri", res.URI)
	}
	return nil
}

func resourceEventStreamsUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.EventStreamUpdate
	arg.ID = d.Id()
	if v, ok := d.GetOk("ngrok_id"); ok {
		arg.ID = *expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = expandString(v)
	}
	if v, ok := d.GetOk("description"); ok {
		arg.Description = expandString(v)
	}
	if v, ok := d.GetOk("fields"); ok {
		arg.Fields = expandStringSlice(v)
	}
	if v, ok := d.GetOk("destination_ids"); ok {
		arg.DestinationIDs = expandStringSlice(v)
	}
	if v, ok := d.GetOk("sampling_rate"); ok {
		arg.SamplingRate = expandFloat64(v)
	}

	res, _, err := b.client.EventStreamsUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] EventStreamsUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceEventStreamsGet(d, m)
}

func resourceEventStreamsDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.EventStreamsDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] EventStreamsDelete: %s", err)
	}
	return err
}
