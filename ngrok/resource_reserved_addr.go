package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func resourceReservedAddrs() *schema.Resource {
	return &schema.Resource{
		Create:      resourceReservedAddrsCreate,
		Read:        resourceReservedAddrsGet,
		Update:      resourceReservedAddrsUpdate,
		Delete:      resourceReservedAddrsDelete,
		Description: "Reserved Addresses are TCP addresses that can be used to listen for traffic.\n TCP address hostnames and ports are assigned by ngrok, they cannot be\n chosen.",
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
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "human-readable description of what this reserved address will be used for",
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
			"id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "unique reserved address resource identifier",
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
		arg.EndpointConfigurationID = expandString(v)
	}

	res, _, err := b.client.ReservedAddrsCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] ReservedAddrsCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

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
		log.Printf("[ERROR] ReservedAddrsGet: %s", err)
		return err
	default:
		d.Set("addr", res.Addr)
		d.Set("description", res.Description)
		if res.EndpointConfiguration != nil {
			d.Set("endpoint_configuration_id", res.EndpointConfiguration.ID)
		}
		d.Set("id", res.ID)
		d.Set("metadata", res.Metadata)
		d.Set("region", res.Region)
	}
	return nil
}

func resourceReservedAddrsUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.ReservedAddrUpdate
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
	if v, ok := d.GetOk("endpoint_configuration_id"); ok {
		arg.EndpointConfigurationID = expandString(v)
	}

	res, _, err := b.client.ReservedAddrsUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] ReservedAddrsUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceReservedAddrsGet(d, m)
}

func resourceReservedAddrsDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.ReservedAddrsDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] ReservedAddrsDelete: %s", err)
	}
	return err
}
