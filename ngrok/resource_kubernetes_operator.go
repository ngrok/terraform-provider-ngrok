package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func resourceKubernetesOperators() *schema.Resource {
	return &schema.Resource{
		Create:      resourceKubernetesOperatorsCreate,
		Read:        resourceKubernetesOperatorsGet,
		Update:      resourceKubernetesOperatorsUpdate,
		Delete:      resourceKubernetesOperatorsDelete,
		Description: "KubernetesOperators is used by the Kubernetes Operator to register and\n manage its own resource, as well as for users to see active kubernetes\n clusters.",
		Schema: map[string]*schema.Schema{
			"binding": {
				Type:        schema.TypeSet,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "information about the Bindings feature of this Kubernetes Operator, if enabled",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    false,
							Computed:    false,
							Optional:    true,
							Sensitive:   false,
							ForceNew:    false,
							Description: "the name by which endpoints can be bound to this Kubernetes Operator. starts with \"k8s/\"",
						},
						"allowed_urls": {
							Type:        schema.TypeList,
							Required:    false,
							Computed:    false,
							Optional:    true,
							Sensitive:   false,
							ForceNew:    false,
							Description: "the regexes for urls allowed to be bound to this operator",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"cert": {
							Type:        schema.TypeSet,
							Required:    false,
							Computed:    true,
							Optional:    false,
							Sensitive:   false,
							ForceNew:    true,
							Description: "the binding certificate information",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cert": {
										Type:        schema.TypeString,
										Required:    false,
										Computed:    true,
										Optional:    false,
										Sensitive:   false,
										ForceNew:    true,
										Description: "the public client certificate generated for this Kubernetes Operator from the CSR supplied when enabling the Bindings feature",
									},
								},
							},
						},
						"ingress_endpoint": {
							Type:        schema.TypeString,
							Required:    false,
							Computed:    false,
							Optional:    true,
							Sensitive:   false,
							ForceNew:    false,
							Description: "the public ingress endpoint for this Kubernetes Operator",
						},
					},
				},
			},
			"deployment": {
				Type:        schema.TypeSet,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "information about the deployment of this Kubernetes Operator",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    false,
							Computed:    false,
							Optional:    true,
							Sensitive:   false,
							ForceNew:    false,
							Description: "the deployment name",
						},
						"namespace": {
							Type:        schema.TypeString,
							Required:    false,
							Computed:    false,
							Optional:    true,
							Sensitive:   false,
							ForceNew:    true,
							Description: "the namespace this Kubernetes Operator is deployed to",
						},
						"version": {
							Type:        schema.TypeString,
							Required:    false,
							Computed:    false,
							Optional:    true,
							Sensitive:   false,
							ForceNew:    true,
							Description: "the version of this Kubernetes Operator",
						},
						"cluster_name": {
							Type:        schema.TypeString,
							Required:    false,
							Computed:    false,
							Optional:    true,
							Sensitive:   false,
							ForceNew:    true,
							Description: "user-given name for the cluster the Kubernetes Operator is deployed to",
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
				Description: "human-readable description of this Kubernetes Operator. optional, max 255 bytes.",
			},
			"enabled_features": {
				Type:        schema.TypeList,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "features enabled for this Kubernetes Operator. a subset of {\"bindings\", \"ingress\", and \"gateway\"}",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    false,
				Description: "unique identifier for this Kubernetes Operator",
			},
			"metadata": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "arbitrary user-defined machine-readable data of this Kubernetes Operator. optional, max 4096 bytes.",
			},
			"principal": {
				Type:        schema.TypeSet,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    true,
				Description: "the principal who created this Kubernetes Operator",
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
			"region": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "the ngrok region in which the ingress for this operator is served. defaults to \"global\"",
			},
		},
	}
}

func resourceKubernetesOperatorsCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.KubernetesOperatorCreate
	if v, ok := d.GetOk("description"); ok {
		arg.Description = *expandString(v)
	}
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = *expandString(v)
	}
	if v, ok := d.GetOk("enabled_features"); ok {
		arg.EnabledFeatures = *expandStringSlice(v)
	}
	if v, ok := d.GetOk("region"); ok {
		arg.Region = *expandString(v)
	}
	if v, ok := d.GetOk("deployment"); ok {
		arg.Deployment = *expandKubernetesOperatorDeployment(v)
	}
	if v, ok := d.GetOk("binding"); ok {
		arg.Binding = expandKubernetesOperatorBindingCreate(v)
	}

	res, _, err := b.client.KubernetesOperatorsCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] KubernetesOperatorsCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceKubernetesOperatorsGet(d, m)
}

func resourceKubernetesOperatorsGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.KubernetesOperatorsGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceKubernetesOperatorsGetDecode(d, res, resp, err)
}

func resourceKubernetesOperatorsGetDecode(d *schema.ResourceData, res *restapi.KubernetesOperator, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		log.Printf("[ERROR] KubernetesOperatorsGet: %s", err)
		return err
	default:
		d.Set("binding", flattenKubernetesOperatorBinding(res.Binding))
		d.Set("deployment", flattenKubernetesOperatorDeployment(&res.Deployment))
		d.Set("description", res.Description)
		d.Set("enabled_features", res.EnabledFeatures)
		d.Set("id", res.ID)
		d.Set("metadata", res.Metadata)
		d.Set("principal", flattenRef(&res.Principal))
		d.Set("region", res.Region)
	}
	return nil
}

func resourceKubernetesOperatorsUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.KubernetesOperatorUpdate
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
	if v, ok := d.GetOk("enabled_features"); ok {
		arg.EnabledFeatures = expandStringSlice(v)
	}
	if v, ok := d.GetOk("region"); ok {
		arg.Region = expandString(v)
	}
	if v, ok := d.GetOk("binding"); ok {
		arg.Binding = expandKubernetesOperatorBindingUpdate(v)
	}
	if v, ok := d.GetOk("deployment"); ok {
		arg.Deployment = expandKubernetesOperatorDeploymentUpdate(v)
	}

	res, _, err := b.client.KubernetesOperatorsUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] KubernetesOperatorsUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceKubernetesOperatorsGet(d, m)
}

func resourceKubernetesOperatorsDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.KubernetesOperatorsDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] KubernetesOperatorsDelete: %s", err)
	}
	return err
}
