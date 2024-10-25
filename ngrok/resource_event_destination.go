package ngrok

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	restapi "github.com/ngrok/terraform-provider-ngrok/restapi"
)

func resourceEventDestinations() *schema.Resource {
	return &schema.Resource{
		Create: resourceEventDestinationsCreate,
		Read:   resourceEventDestinationsGet,
		Update: resourceEventDestinationsUpdate,
		Delete: resourceEventDestinationsDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "Human-readable description of the Event Destination. Optional, max 255 bytes.",
			},
			"format": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "The output format you would like to serialize events into when sending to their target. Currently the only accepted value is `JSON`.",
			},
			"id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Optional:    false,
				Sensitive:   false,
				ForceNew:    false,
				Description: "Unique identifier for this Event Destination.",
			},
			"metadata": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "Arbitrary user-defined machine-readable data of this Event Destination. Optional, max 4096 bytes.",
			},
			"target": {
				Type:        schema.TypeSet,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "An object that encapsulates where and how to send your events. An event destination must contain exactly one of the following objects, leaving the rest null: `kinesis`, `firehose`, `cloudwatch_logs`, or `s3`.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"firehose": {
							Type:        schema.TypeSet,
							Required:    false,
							Computed:    false,
							Optional:    true,
							Sensitive:   false,
							ForceNew:    false,
							Description: "Configuration used to send events to Amazon Kinesis Data Firehose.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth": {
										Type:        schema.TypeSet,
										Required:    false,
										Computed:    false,
										Optional:    true,
										Sensitive:   false,
										ForceNew:    false,
										Description: "Configuration for how to authenticate into your AWS account. Exactly one of `role` or `creds` should be configured.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"role": {
													Type:        schema.TypeSet,
													Required:    false,
													Computed:    false,
													Optional:    true,
													Sensitive:   false,
													ForceNew:    false,
													Description: "A role for ngrok to assume on your behalf to deposit events into your AWS account.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"role_arn": {
																Type:        schema.TypeString,
																Required:    true,
																Computed:    false,
																Optional:    false,
																Sensitive:   false,
																ForceNew:    false,
																Description: "An ARN that specifies the role that ngrok should use to deliver to the configured target.",
															},
														},
													},
												},
												"creds": {
													Type:        schema.TypeSet,
													Required:    false,
													Computed:    false,
													Optional:    true,
													Sensitive:   false,
													ForceNew:    false,
													Description: "Credentials to your AWS account if you prefer ngrok to sign in with long-term access keys.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"aws_access_key_id": {
																Type:        schema.TypeString,
																Required:    true,
																Computed:    false,
																Optional:    false,
																Sensitive:   false,
																ForceNew:    false,
																Description: "The ID portion of an AWS access key.",
															},
															"aws_secret_access_key": {
																Type:        schema.TypeString,
																Required:    true,
																Computed:    false,
																Optional:    false,
																Sensitive:   false,
																ForceNew:    false,
																Description: "The secret portion of an AWS access key.",
															},
														},
													},
												},
											},
										},
									},
									"delivery_stream_arn": {
										Type:        schema.TypeString,
										Required:    false,
										Computed:    false,
										Optional:    true,
										Sensitive:   false,
										ForceNew:    false,
										Description: "An Amazon Resource Name specifying the Firehose delivery stream to deposit events into.",
									},
								},
							},
						},
						"kinesis": {
							Type:        schema.TypeSet,
							Required:    false,
							Computed:    false,
							Optional:    true,
							Sensitive:   false,
							ForceNew:    false,
							Description: "Configuration used to send events to Amazon Kinesis.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth": {
										Type:        schema.TypeSet,
										Required:    false,
										Computed:    false,
										Optional:    true,
										Sensitive:   false,
										ForceNew:    false,
										Description: "Configuration for how to authenticate into your AWS account. Exactly one of `role` or `creds` should be configured.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"role": {
													Type:        schema.TypeSet,
													Required:    false,
													Computed:    false,
													Optional:    true,
													Sensitive:   false,
													ForceNew:    false,
													Description: "A role for ngrok to assume on your behalf to deposit events into your AWS account.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"role_arn": {
																Type:        schema.TypeString,
																Required:    true,
																Computed:    false,
																Optional:    false,
																Sensitive:   false,
																ForceNew:    false,
																Description: "An ARN that specifies the role that ngrok should use to deliver to the configured target.",
															},
														},
													},
												},
												"creds": {
													Type:        schema.TypeSet,
													Required:    false,
													Computed:    false,
													Optional:    true,
													Sensitive:   false,
													ForceNew:    false,
													Description: "Credentials to your AWS account if you prefer ngrok to sign in with long-term access keys.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"aws_access_key_id": {
																Type:        schema.TypeString,
																Required:    true,
																Computed:    false,
																Optional:    false,
																Sensitive:   false,
																ForceNew:    false,
																Description: "The ID portion of an AWS access key.",
															},
															"aws_secret_access_key": {
																Type:        schema.TypeString,
																Required:    true,
																Computed:    false,
																Optional:    false,
																Sensitive:   false,
																ForceNew:    false,
																Description: "The secret portion of an AWS access key.",
															},
														},
													},
												},
											},
										},
									},
									"stream_arn": {
										Type:        schema.TypeString,
										Required:    false,
										Computed:    false,
										Optional:    true,
										Sensitive:   false,
										ForceNew:    false,
										Description: "An Amazon Resource Name specifying the Kinesis stream to deposit events into.",
									},
								},
							},
						},
						"cloudwatch_logs": {
							Type:        schema.TypeSet,
							Required:    false,
							Computed:    false,
							Optional:    true,
							Sensitive:   false,
							ForceNew:    false,
							Description: "Configuration used to send events to Amazon CloudWatch Logs.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth": {
										Type:        schema.TypeSet,
										Required:    false,
										Computed:    false,
										Optional:    true,
										Sensitive:   false,
										ForceNew:    false,
										Description: "Configuration for how to authenticate into your AWS account. Exactly one of `role` or `creds` should be configured.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"role": {
													Type:        schema.TypeSet,
													Required:    false,
													Computed:    false,
													Optional:    true,
													Sensitive:   false,
													ForceNew:    false,
													Description: "A role for ngrok to assume on your behalf to deposit events into your AWS account.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"role_arn": {
																Type:        schema.TypeString,
																Required:    true,
																Computed:    false,
																Optional:    false,
																Sensitive:   false,
																ForceNew:    false,
																Description: "An ARN that specifies the role that ngrok should use to deliver to the configured target.",
															},
														},
													},
												},
												"creds": {
													Type:        schema.TypeSet,
													Required:    false,
													Computed:    false,
													Optional:    true,
													Sensitive:   false,
													ForceNew:    false,
													Description: "Credentials to your AWS account if you prefer ngrok to sign in with long-term access keys.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"aws_access_key_id": {
																Type:        schema.TypeString,
																Required:    true,
																Computed:    false,
																Optional:    false,
																Sensitive:   false,
																ForceNew:    false,
																Description: "The ID portion of an AWS access key.",
															},
															"aws_secret_access_key": {
																Type:        schema.TypeString,
																Required:    true,
																Computed:    false,
																Optional:    false,
																Sensitive:   false,
																ForceNew:    false,
																Description: "The secret portion of an AWS access key.",
															},
														},
													},
												},
											},
										},
									},
									"log_group_arn": {
										Type:        schema.TypeString,
										Required:    false,
										Computed:    false,
										Optional:    true,
										Sensitive:   false,
										ForceNew:    false,
										Description: "An Amazon Resource Name specifying the CloudWatch Logs group to deposit events into.",
									},
								},
							},
						},
						"debug": {
							Type:        schema.TypeSet,
							Required:    false,
							Computed:    false,
							Optional:    true,
							Sensitive:   false,
							ForceNew:    false,
							Description: "Configuration used for internal debugging.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"log": {
										Type:        schema.TypeBool,
										Required:    false,
										Computed:    false,
										Optional:    true,
										Sensitive:   false,
										ForceNew:    false,
										Description: "Whether or not to output to publisher service logs.",
									},
									"callback_url": {
										Type:        schema.TypeString,
										Required:    false,
										Computed:    false,
										Optional:    true,
										Sensitive:   false,
										ForceNew:    false,
										Description: "URL to send events to.",
									},
								},
							},
						},
						"datadog": {
							Type:        schema.TypeSet,
							Required:    false,
							Computed:    false,
							Optional:    true,
							Sensitive:   false,
							ForceNew:    false,
							Description: "Configuration used to send events to Datadog.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"api_key": {
										Type:        schema.TypeString,
										Required:    false,
										Computed:    false,
										Optional:    true,
										Sensitive:   false,
										ForceNew:    false,
										Description: "Datadog API key to use.",
									},
									"ddtags": {
										Type:        schema.TypeString,
										Required:    false,
										Computed:    false,
										Optional:    true,
										Sensitive:   false,
										ForceNew:    false,
										Description: "Tags to send with the event.",
									},
									"service": {
										Type:        schema.TypeString,
										Required:    false,
										Computed:    false,
										Optional:    true,
										Sensitive:   false,
										ForceNew:    false,
										Description: "Service name to send with the event.",
									},
									"ddsite": {
										Type:        schema.TypeString,
										Required:    false,
										Computed:    false,
										Optional:    true,
										Sensitive:   false,
										ForceNew:    false,
										Description: "Datadog site to send event to.",
									},
								},
							},
						},
						"azure_logs_ingestion": {
							Type:        schema.TypeSet,
							Required:    false,
							Computed:    false,
							Optional:    true,
							Sensitive:   false,
							ForceNew:    false,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tenant_id": {
										Type:        schema.TypeString,
										Required:    true,
										Computed:    false,
										Optional:    false,
										Sensitive:   false,
										ForceNew:    false,
										Description: "Tenant ID for the Azure account",
									},
									"client_id": {
										Type:        schema.TypeString,
										Required:    true,
										Computed:    false,
										Optional:    false,
										Sensitive:   false,
										ForceNew:    false,
										Description: "Client ID for the application client",
									},
									"client_secret": {
										Type:        schema.TypeString,
										Required:    true,
										Computed:    false,
										Optional:    false,
										Sensitive:   false,
										ForceNew:    false,
										Description: "Client Secret for the application client",
									},
									"logs_ingestion_uri": {
										Type:        schema.TypeString,
										Required:    true,
										Computed:    false,
										Optional:    false,
										Sensitive:   false,
										ForceNew:    false,
										Description: "Data collection endpoint logs ingestion URI",
									},
									"data_collection_rule_id": {
										Type:        schema.TypeString,
										Required:    true,
										Computed:    false,
										Optional:    false,
										Sensitive:   false,
										ForceNew:    false,
										Description: "Data collection rule immutable ID",
									},
									"data_collection_stream_name": {
										Type:        schema.TypeString,
										Required:    true,
										Computed:    false,
										Optional:    false,
										Sensitive:   false,
										ForceNew:    false,
										Description: "Data collection stream name to use as destination, located inside the DCR",
									},
								},
							},
						},
					},
				},
			},
			"verify_with_test_event": {
				Type:        schema.TypeBool,
				Required:    false,
				Computed:    false,
				Optional:    true,
				Sensitive:   false,
				ForceNew:    false,
				Description: "",
			},
		},
	}
}

func resourceEventDestinationsCreate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.EventDestinationCreate
	if v, ok := d.GetOk("metadata"); ok {
		arg.Metadata = *expandString(v)
	}
	if v, ok := d.GetOk("description"); ok {
		arg.Description = *expandString(v)
	}
	if v, ok := d.GetOk("format"); ok {
		arg.Format = *expandString(v)
	}
	if v, ok := d.GetOk("target"); ok {
		arg.Target = *expandEventTarget(v)
	}
	if v, ok := d.GetOk("verify_with_test_event"); ok {
		arg.VerifyWithTestEvent = expandBool(v)
	}

	res, _, err := b.client.EventDestinationsCreate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] EventDestinationsCreate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceEventDestinationsGet(d, m)
}

func resourceEventDestinationsGet(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	res, resp, err := b.client.EventDestinationsGet(context.Background(), &restapi.Item{
		ID: d.Id(),
	})
	return resourceEventDestinationsGetDecode(d, res, resp, err)
}

func resourceEventDestinationsGetDecode(d *schema.ResourceData, res *restapi.EventDestination, resp *http.Response, err error) error {
	switch {
	case resp != nil && resp.StatusCode == 404:
		d.SetId("")
	case err != nil:
		log.Printf("[ERROR] EventDestinationsGet: %s", err)
		return err
	default:
		d.Set("description", res.Description)
		d.Set("format", res.Format)
		d.Set("id", res.ID)
		d.Set("metadata", res.Metadata)
		d.Set("target", flattenEventTarget(&res.Target))
	}
	return nil
}

func resourceEventDestinationsUpdate(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)

	var arg restapi.EventDestinationUpdate
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
	if v, ok := d.GetOk("format"); ok {
		arg.Format = expandString(v)
	}
	if v, ok := d.GetOk("target"); ok {
		arg.Target = expandEventTarget(v)
	}
	if v, ok := d.GetOk("verify_with_test_event"); ok {
		arg.VerifyWithTestEvent = expandBool(v)
	}

	res, _, err := b.client.EventDestinationsUpdate(context.Background(), &arg)
	if err != nil {
		log.Printf("[ERROR] EventDestinationsUpdate: %s", err)
		return err
	}
	d.SetId(res.ID)

	return resourceEventDestinationsGet(d, m)
}

func resourceEventDestinationsDelete(d *schema.ResourceData, m interface{}) (err error) {
	b := m.(*base)
	_, _, err = b.client.EventDestinationsDelete(context.Background(), &restapi.Item{ID: d.Id()})
	if err != nil {
		log.Printf("[ERROR] EventDestinationsDelete: %s", err)
	}
	return err
}
