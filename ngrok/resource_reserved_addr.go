// Code generated by apic. DO NOT EDIT.

package ngrok

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func resourceReservedAddrs() *schema.Resource {
	return &schema.Resource{
		Create: resourceReservedAddrsCreate,
		Read:   resourceReservedAddrsGet,
		Update: resourceReservedAddrsUpdate,
		Delete: resourceReservedAddrsDelete,

		Schema: map[string]*schema.Schema{
			"addr": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "hostname:port of the reserved address that was assigned at creation time",
			},
			"created_at": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "timestamp when the reserved address was created, RFC 3339 format",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "human-readable description of what this reserved address will be used for",
			},
			"endpoint_configuration": {
				Type:        schema.TypeSet,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "object reference to the endpoint configuration that will be applied to traffic to this address",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ngrok_id": {
							Type:        schema.TypeString,
							Required:    false,
							Computed:    true,
							Optional:    true,
							Sensitive:   false,
							ForceNew:    false,
							Description: "a resource identifier",
						},
						"uri": {
							Type:        schema.TypeString,
							Required:    false,
							Computed:    true,
							Optional:    true,
							Sensitive:   false,
							ForceNew:    true,
							Description: "a uri for locating a resource",
						},
					},
				},
			},
			"endpoint_configuration_id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "ID of an endpoint configuration of type tcp that will be used to handle inbound traffic to this address",
			},
			"metadata": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "arbitrary user-defined machine-readable data of this reserved address. Optional, max 4096 bytes.",
			},
			"ngrok_id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "unique reserved address resource identifier",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Default:     "us",
				Sensitive:   false,
				ForceNew:    true,
				Description: "reserve the address in this geographic ngrok datacenter. Optional, default is us. (au, eu, ap, us, jp, in, sa)",
			},
			"uri": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    true,
				Description: "URI of the reserved address API resource",
			},
		},
	}
}

func resourceReservedAddrsCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.ReservedAddrCreate
	if v, ok := d.GetOk("description"); ok {
		arg.Description = *expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = *expandString(v)
	}
	if v, ok := d.GetOk("region"); ok {
		arg.Region = *expandString(v)
	}
	if v, ok := d.GetOk("endpoint_configuration_id"); ok {
		arg.EndpointConfigurationID = *expandString(v)
	}

	res, _, err := b.client.ReservedAddrsCreate(context.Background(), &arg)
	if err == nil {
		d.SetId(res.ID)
	}
	return resourceReservedAddrsGet(d, m)
}

func resourceReservedAddrsGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.ReservedAddrsGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceReservedAddrsGetDecode(d, res, resp, err)
}

func resourceReservedAddrsGetDecode(d *schema.ResourceData, res *restapi.ReservedAddr, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		return err
	default:
		d.Set("addr", res.Addr)
		d.Set("created_at", res.CreatedAt)
		d.Set("description", res.Description)
		d.Set("endpoint_configuration", flattenRef(res.EndpointConfiguration))
		if res.EndpointConfiguration != nil {
			d.Set("endpoint_configuration_id", res.EndpointConfiguration.ID)
		}
		d.Set("metadata", res.Metadata)
		d.Set("ngrok_id", res.ID)
		d.Set("region", res.Region)
		d.Set("uri", res.URI)
	}
	return nil
}

func resourceReservedAddrsUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.ReservedAddrUpdate
	arg.ID = d.Id()
	if v, ok := d.GetOk("ngrok_id"); ok {
		arg.ID = *expandString(v)
	}
	if v, ok := d.GetOk("description"); ok {
		arg.Description = expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = expandString(v)
	}
	if v, ok := d.GetOk("endpoint_configuration_id"); ok {
		arg.EndpointConfigurationID = expandString(v)
	}

	res, _, err := b.client.ReservedAddrsUpdate(context.Background(), &arg)
	if err != nil {
		return err
	}
	d.SetId(res.ID)

	return resourceReservedAddrsGet(d, m)
}

func resourceReservedAddrsDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.ReservedAddrsDelete(context.Background(), &restapi.Item{ID: d.Id()})
	return err
}
