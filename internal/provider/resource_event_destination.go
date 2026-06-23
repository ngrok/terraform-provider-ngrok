package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/event_destinations"
	"github.com/ngrok/terraform-provider-ngrok/internal/resource_event_destination"
)

var (
	_ resource.Resource                = &eventDestinationResource{}
	_ resource.ResourceWithImportState = &eventDestinationResource{}
)

type eventDestinationResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Description types.String `tfsdk:"description"`
	Metadata    types.String `tfsdk:"metadata"`
	Format      types.String `tfsdk:"format"`
	Target      types.Object `tfsdk:"target"`
	URI         types.String `tfsdk:"uri"`
	CreatedAt   types.String `tfsdk:"created_at"`
}

type eventDestinationResource struct {
	client *event_destinations.Client
}

func NewEventDestinationResource() resource.Resource {
	return &eventDestinationResource{}
}

func (r *eventDestinationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_event_destination"
}

func awsAuthSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Description: "Configuration for how to authenticate into your AWS account.",
		Optional:    true,
		Attributes: map[string]schema.Attribute{
			"role": schema.SingleNestedAttribute{
				Description: "A role for ngrok to assume on your behalf.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"role_arn": schema.StringAttribute{
						Description: "An ARN that specifies the role that ngrok should use to deliver to the configured target.",
						Required:    true,
					},
				},
			},
			"creds": schema.SingleNestedAttribute{
				Description: "Credentials to your AWS account if you prefer ngrok to sign in with long-term access keys.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"aws_access_key_id": schema.StringAttribute{
						Description: "The ID portion of an AWS access key.",
						Required:    true,
					},
					"aws_secret_access_key": schema.StringAttribute{
						Description: "The secret portion of an AWS access key.",
						Optional:    true,
						Sensitive:   true,
					},
				},
			},
		},
	}
}

func (r *eventDestinationResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	s := resource_event_destination.EventDestinationResourceSchema(ctx)
	attrs := s.Attributes

	// Add target attribute (excluded from codegen due to duplicate CustomType collision)
	attrs["target"] = schema.SingleNestedAttribute{
		Description: "An object that encapsulates where and how to send your events.",
		Optional:    true,
		Attributes: map[string]schema.Attribute{
			"firehose": schema.SingleNestedAttribute{
				Description: "Configuration used to send events to Amazon Kinesis Data Firehose.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"auth":                awsAuthSchema(),
					"delivery_stream_arn": schema.StringAttribute{Description: "An Amazon Resource Name specifying the Firehose delivery stream.", Required: true},
				},
			},
			"kinesis": schema.SingleNestedAttribute{
				Description: "Configuration used to send events to Amazon Kinesis.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"auth":       awsAuthSchema(),
					"stream_arn": schema.StringAttribute{Description: "An Amazon Resource Name specifying the Kinesis stream.", Required: true},
				},
			},
			"cloudwatch_logs": schema.SingleNestedAttribute{
				Description: "Configuration used to send events to Amazon CloudWatch Logs.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"auth":          awsAuthSchema(),
					"log_group_arn": schema.StringAttribute{Description: "An Amazon Resource Name specifying the CloudWatch Logs group.", Required: true},
				},
			},
			"datadog": schema.SingleNestedAttribute{
				Description: "Configuration used to send events to Datadog.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"api_key": schema.StringAttribute{Description: "Datadog API key to use.", Optional: true, Sensitive: true},
					"ddtags":  schema.StringAttribute{Description: "Tags to send with the event.", Optional: true},
					"service": schema.StringAttribute{Description: "Service name to send with the event.", Optional: true},
					"ddsite":  schema.StringAttribute{Description: "Datadog site to send event to.", Optional: true},
				},
			},
			"azure_logs_ingestion": schema.SingleNestedAttribute{
				Description: "Configuration used to send events to Azure Logs Ingestion.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"tenant_id":                   schema.StringAttribute{Description: "Tenant ID for the Azure account.", Required: true},
					"client_id":                   schema.StringAttribute{Description: "Client ID for the application client.", Required: true},
					"client_secret":               schema.StringAttribute{Description: "Client Secret for the application client.", Required: true, Sensitive: true},
					"logs_ingestion_uri":          schema.StringAttribute{Description: "Data collection endpoint logs ingestion URI.", Required: true},
					"data_collection_rule_id":     schema.StringAttribute{Description: "Data collection rule immutable ID.", Required: true},
					"data_collection_stream_name": schema.StringAttribute{Description: "Data collection stream name to use as destination.", Required: true},
				},
			},
		},
	}

	addStringPlanModifiers(attrs, "id", useStateForUnknownString())
	addStringPlanModifiers(attrs, "description", useStateForUnknownString())
	addStringPlanModifiers(attrs, "metadata", useStateForUnknownString())
	addStringPlanModifiers(attrs, "format", useStateForUnknownString())
	addStringPlanModifiers(attrs, "uri", useStateForUnknownString())
	addStringPlanModifiers(attrs, "created_at", useStateForUnknownString())

	resp.Schema = s
}

func (r *eventDestinationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clientConfig, ok := req.ProviderData.(*ngrok.ClientConfig)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *ngrok.ClientConfig, got: %T.", req.ProviderData),
		)
		return
	}
	r.client = event_destinations.NewClient(clientConfig)
}

func (r *eventDestinationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan eventDestinationResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq := &ngrok.EventDestinationCreate{
		Description: plan.Description.ValueString(),
		Metadata:    plan.Metadata.ValueString(),
		Format:      plan.Format.ValueString(),
		Target:      expandEventTarget(ctx, plan.Target, &resp.Diagnostics),
	}
	if resp.Diagnostics.HasError() {
		return
	}

	dest, err := r.client.Create(ctx, createReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating event destination", err.Error())
		return
	}

	flattenEventDestination(ctx, dest, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *eventDestinationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state eventDestinationResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dest, err := r.client.Get(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading event destination", err.Error())
		return
	}

	flattenEventDestination(ctx, dest, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *eventDestinationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan eventDestinationResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state eventDestinationResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq := &ngrok.EventDestinationUpdate{
		ID:          state.ID.ValueString(),
		Description: stringPtrFromFramework(plan.Description),
		Metadata:    stringPtrFromFramework(plan.Metadata),
		Format:      stringPtrFromFramework(plan.Format),
	}

	if !plan.Target.IsNull() && !plan.Target.IsUnknown() {
		target := expandEventTarget(ctx, plan.Target, &resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}
		updateReq.Target = &target
	}

	dest, err := r.client.Update(ctx, updateReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating event destination", err.Error())
		return
	}

	flattenEventDestination(ctx, dest, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *eventDestinationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state eventDestinationResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			return
		}
		resp.Diagnostics.AddError("Error deleting event destination", err.Error())
	}
}

func (r *eventDestinationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// --- Flatten / Expand helpers for EventTarget ---

type eventTargetModel struct {
	Firehose           types.Object `tfsdk:"firehose"`
	Kinesis            types.Object `tfsdk:"kinesis"`
	CloudwatchLogs     types.Object `tfsdk:"cloudwatch_logs"`
	Datadog            types.Object `tfsdk:"datadog"`
	AzureLogsIngestion types.Object `tfsdk:"azure_logs_ingestion"`
}

type awsAuthModel struct {
	Role  types.Object `tfsdk:"role"`
	Creds types.Object `tfsdk:"creds"`
}

type awsRoleModel struct {
	RoleARN types.String `tfsdk:"role_arn"`
}

type awsCredsModel struct {
	AWSAccessKeyID     types.String `tfsdk:"aws_access_key_id"`
	AWSSecretAccessKey types.String `tfsdk:"aws_secret_access_key"`
}

type firehoseModel struct {
	Auth              types.Object `tfsdk:"auth"`
	DeliveryStreamARN types.String `tfsdk:"delivery_stream_arn"`
}

type kinesisModel struct {
	Auth      types.Object `tfsdk:"auth"`
	StreamARN types.String `tfsdk:"stream_arn"`
}

type cloudwatchLogsModel struct {
	Auth        types.Object `tfsdk:"auth"`
	LogGroupARN types.String `tfsdk:"log_group_arn"`
}

type datadogModel struct {
	ApiKey  types.String `tfsdk:"api_key"`
	Ddtags  types.String `tfsdk:"ddtags"`
	Service types.String `tfsdk:"service"`
	Ddsite  types.String `tfsdk:"ddsite"`
}

type azureLogsIngestionModel struct {
	TenantId                 types.String `tfsdk:"tenant_id"`
	ClientId                 types.String `tfsdk:"client_id"`
	ClientSecret             types.String `tfsdk:"client_secret"`
	LogsIngestionURI         types.String `tfsdk:"logs_ingestion_uri"`
	DataCollectionRuleId     types.String `tfsdk:"data_collection_rule_id"`
	DataCollectionStreamName types.String `tfsdk:"data_collection_stream_name"`
}

func awsRoleAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"role_arn": types.StringType,
	}
}

func awsCredsAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"aws_access_key_id":     types.StringType,
		"aws_secret_access_key": types.StringType,
	}
}

func awsAuthAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"role":  types.ObjectType{AttrTypes: awsRoleAttrTypes()},
		"creds": types.ObjectType{AttrTypes: awsCredsAttrTypes()},
	}
}

func firehoseAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"auth":                types.ObjectType{AttrTypes: awsAuthAttrTypes()},
		"delivery_stream_arn": types.StringType,
	}
}

func kinesisAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"auth":       types.ObjectType{AttrTypes: awsAuthAttrTypes()},
		"stream_arn": types.StringType,
	}
}

func cloudwatchLogsAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"auth":          types.ObjectType{AttrTypes: awsAuthAttrTypes()},
		"log_group_arn": types.StringType,
	}
}

func datadogAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"api_key": types.StringType,
		"ddtags":  types.StringType,
		"service": types.StringType,
		"ddsite":  types.StringType,
	}
}

func azureLogsIngestionAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tenant_id":                   types.StringType,
		"client_id":                   types.StringType,
		"client_secret":               types.StringType,
		"logs_ingestion_uri":          types.StringType,
		"data_collection_rule_id":     types.StringType,
		"data_collection_stream_name": types.StringType,
	}
}

func eventTargetAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"firehose":             types.ObjectType{AttrTypes: firehoseAttrTypes()},
		"kinesis":              types.ObjectType{AttrTypes: kinesisAttrTypes()},
		"cloudwatch_logs":      types.ObjectType{AttrTypes: cloudwatchLogsAttrTypes()},
		"datadog":              types.ObjectType{AttrTypes: datadogAttrTypes()},
		"azure_logs_ingestion": types.ObjectType{AttrTypes: azureLogsIngestionAttrTypes()},
	}
}

func flattenAWSAuth(ctx context.Context, auth ngrok.AWSAuth, priorAuth types.Object, diags *diag.Diagnostics) types.Object {
	// Extract prior creds secret for preservation
	var priorSecret types.String
	if !priorAuth.IsNull() && !priorAuth.IsUnknown() {
		var priorModel awsAuthModel
		diags.Append(priorAuth.As(ctx, &priorModel, basetypes.ObjectAsOptions{})...)
		if !priorModel.Creds.IsNull() && !priorModel.Creds.IsUnknown() {
			var priorCreds awsCredsModel
			diags.Append(priorModel.Creds.As(ctx, &priorCreds, basetypes.ObjectAsOptions{})...)
			priorSecret = priorCreds.AWSSecretAccessKey
		}
	}

	roleObj := types.ObjectNull(awsRoleAttrTypes())
	if auth.Role != nil {
		obj, d := types.ObjectValueFrom(ctx, awsRoleAttrTypes(), &awsRoleModel{
			RoleARN: types.StringValue(auth.Role.RoleARN),
		})
		diags.Append(d...)
		roleObj = obj
	}

	credsObj := types.ObjectNull(awsCredsAttrTypes())
	if auth.Creds != nil {
		secretVal := types.StringValue(stringFromPtr(auth.Creds.AWSSecretAccessKey))
		obj, d := types.ObjectValueFrom(ctx, awsCredsAttrTypes(), &awsCredsModel{
			AWSAccessKeyID:     types.StringValue(auth.Creds.AWSAccessKeyID),
			AWSSecretAccessKey: preserveSensitive(secretVal, priorSecret),
		})
		diags.Append(d...)
		credsObj = obj
	}

	obj, d := types.ObjectValueFrom(ctx, awsAuthAttrTypes(), &awsAuthModel{
		Role:  roleObj,
		Creds: credsObj,
	})
	diags.Append(d...)
	return obj
}

func flattenEventTarget(ctx context.Context, target ngrok.EventTarget, priorTarget types.Object, diags *diag.Diagnostics) types.Object {
	// Extract prior sub-objects for sensitive field preservation
	var prior eventTargetModel
	if !priorTarget.IsNull() && !priorTarget.IsUnknown() {
		diags.Append(priorTarget.As(ctx, &prior, basetypes.ObjectAsOptions{})...)
	}

	firehoseObj := types.ObjectNull(firehoseAttrTypes())
	if target.Firehose != nil {
		var priorAuth types.Object
		if !prior.Firehose.IsNull() && !prior.Firehose.IsUnknown() {
			var priorFH firehoseModel
			diags.Append(prior.Firehose.As(ctx, &priorFH, basetypes.ObjectAsOptions{})...)
			priorAuth = priorFH.Auth
		}
		authObj := flattenAWSAuth(ctx, target.Firehose.Auth, priorAuth, diags)
		obj, d := types.ObjectValueFrom(ctx, firehoseAttrTypes(), &firehoseModel{
			Auth:              authObj,
			DeliveryStreamARN: types.StringValue(target.Firehose.DeliveryStreamARN),
		})
		diags.Append(d...)
		firehoseObj = obj
	}

	kinesisObj := types.ObjectNull(kinesisAttrTypes())
	if target.Kinesis != nil {
		var priorAuth types.Object
		if !prior.Kinesis.IsNull() && !prior.Kinesis.IsUnknown() {
			var priorK kinesisModel
			diags.Append(prior.Kinesis.As(ctx, &priorK, basetypes.ObjectAsOptions{})...)
			priorAuth = priorK.Auth
		}
		authObj := flattenAWSAuth(ctx, target.Kinesis.Auth, priorAuth, diags)
		obj, d := types.ObjectValueFrom(ctx, kinesisAttrTypes(), &kinesisModel{
			Auth:      authObj,
			StreamARN: types.StringValue(target.Kinesis.StreamARN),
		})
		diags.Append(d...)
		kinesisObj = obj
	}

	cwObj := types.ObjectNull(cloudwatchLogsAttrTypes())
	if target.CloudwatchLogs != nil {
		var priorAuth types.Object
		if !prior.CloudwatchLogs.IsNull() && !prior.CloudwatchLogs.IsUnknown() {
			var priorCW cloudwatchLogsModel
			diags.Append(prior.CloudwatchLogs.As(ctx, &priorCW, basetypes.ObjectAsOptions{})...)
			priorAuth = priorCW.Auth
		}
		authObj := flattenAWSAuth(ctx, target.CloudwatchLogs.Auth, priorAuth, diags)
		obj, d := types.ObjectValueFrom(ctx, cloudwatchLogsAttrTypes(), &cloudwatchLogsModel{
			Auth:        authObj,
			LogGroupARN: types.StringValue(target.CloudwatchLogs.LogGroupARN),
		})
		diags.Append(d...)
		cwObj = obj
	}

	ddObj := types.ObjectNull(datadogAttrTypes())
	if target.Datadog != nil {
		var priorApiKey types.String
		if !prior.Datadog.IsNull() && !prior.Datadog.IsUnknown() {
			var priorDD datadogModel
			diags.Append(prior.Datadog.As(ctx, &priorDD, basetypes.ObjectAsOptions{})...)
			priorApiKey = priorDD.ApiKey
		}
		apiKeyVal := stringValueFromPtr(target.Datadog.ApiKey)
		obj, d := types.ObjectValueFrom(ctx, datadogAttrTypes(), &datadogModel{
			ApiKey:  preserveSensitive(apiKeyVal, priorApiKey),
			Ddtags:  stringValueFromPtr(target.Datadog.Ddtags),
			Service: stringValueFromPtr(target.Datadog.Service),
			Ddsite:  stringValueFromPtr(target.Datadog.Ddsite),
		})
		diags.Append(d...)
		ddObj = obj
	}

	azureObj := types.ObjectNull(azureLogsIngestionAttrTypes())
	if target.AzureLogsIngestion != nil {
		var priorClientSecret types.String
		if !prior.AzureLogsIngestion.IsNull() && !prior.AzureLogsIngestion.IsUnknown() {
			var priorAz azureLogsIngestionModel
			diags.Append(prior.AzureLogsIngestion.As(ctx, &priorAz, basetypes.ObjectAsOptions{})...)
			priorClientSecret = priorAz.ClientSecret
		}
		clientSecretVal := types.StringValue(target.AzureLogsIngestion.ClientSecret)
		obj, d := types.ObjectValueFrom(ctx, azureLogsIngestionAttrTypes(), &azureLogsIngestionModel{
			TenantId:                 types.StringValue(target.AzureLogsIngestion.TenantId),
			ClientId:                 types.StringValue(target.AzureLogsIngestion.ClientId),
			ClientSecret:             preserveSensitive(clientSecretVal, priorClientSecret),
			LogsIngestionURI:         types.StringValue(target.AzureLogsIngestion.LogsIngestionURI),
			DataCollectionRuleId:     types.StringValue(target.AzureLogsIngestion.DataCollectionRuleId),
			DataCollectionStreamName: types.StringValue(target.AzureLogsIngestion.DataCollectionStreamName),
		})
		diags.Append(d...)
		azureObj = obj
	}

	obj, d := types.ObjectValueFrom(ctx, eventTargetAttrTypes(), &eventTargetModel{
		Firehose:           firehoseObj,
		Kinesis:            kinesisObj,
		CloudwatchLogs:     cwObj,
		Datadog:            ddObj,
		AzureLogsIngestion: azureObj,
	})
	diags.Append(d...)
	return obj
}

func expandAWSAuth(ctx context.Context, obj types.Object, diags *diag.Diagnostics) ngrok.AWSAuth {
	if obj.IsNull() || obj.IsUnknown() {
		return ngrok.AWSAuth{}
	}
	var model awsAuthModel
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return ngrok.AWSAuth{}
	}

	var auth ngrok.AWSAuth
	if !model.Role.IsNull() && !model.Role.IsUnknown() {
		var role awsRoleModel
		diags.Append(model.Role.As(ctx, &role, basetypes.ObjectAsOptions{})...)
		auth.Role = &ngrok.AWSRole{
			RoleARN: role.RoleARN.ValueString(),
		}
	}
	if !model.Creds.IsNull() && !model.Creds.IsUnknown() {
		var creds awsCredsModel
		diags.Append(model.Creds.As(ctx, &creds, basetypes.ObjectAsOptions{})...)
		auth.Creds = &ngrok.AWSCredentials{
			AWSAccessKeyID:     creds.AWSAccessKeyID.ValueString(),
			AWSSecretAccessKey: stringPtrFromFramework(creds.AWSSecretAccessKey),
		}
	}
	return auth
}

func expandEventTarget(ctx context.Context, obj types.Object, diags *diag.Diagnostics) ngrok.EventTarget {
	if obj.IsNull() || obj.IsUnknown() {
		return ngrok.EventTarget{}
	}

	var model eventTargetModel
	diags.Append(obj.As(ctx, &model, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return ngrok.EventTarget{}
	}

	var target ngrok.EventTarget

	if !model.Firehose.IsNull() && !model.Firehose.IsUnknown() {
		var fh firehoseModel
		diags.Append(model.Firehose.As(ctx, &fh, basetypes.ObjectAsOptions{})...)
		target.Firehose = &ngrok.EventTargetFirehose{
			Auth:              expandAWSAuth(ctx, fh.Auth, diags),
			DeliveryStreamARN: fh.DeliveryStreamARN.ValueString(),
		}
	}

	if !model.Kinesis.IsNull() && !model.Kinesis.IsUnknown() {
		var k kinesisModel
		diags.Append(model.Kinesis.As(ctx, &k, basetypes.ObjectAsOptions{})...)
		target.Kinesis = &ngrok.EventTargetKinesis{
			Auth:      expandAWSAuth(ctx, k.Auth, diags),
			StreamARN: k.StreamARN.ValueString(),
		}
	}

	if !model.CloudwatchLogs.IsNull() && !model.CloudwatchLogs.IsUnknown() {
		var cw cloudwatchLogsModel
		diags.Append(model.CloudwatchLogs.As(ctx, &cw, basetypes.ObjectAsOptions{})...)
		target.CloudwatchLogs = &ngrok.EventTargetCloudwatchLogs{
			Auth:        expandAWSAuth(ctx, cw.Auth, diags),
			LogGroupARN: cw.LogGroupARN.ValueString(),
		}
	}

	if !model.Datadog.IsNull() && !model.Datadog.IsUnknown() {
		var dd datadogModel
		diags.Append(model.Datadog.As(ctx, &dd, basetypes.ObjectAsOptions{})...)
		target.Datadog = &ngrok.EventTargetDatadog{
			ApiKey:  stringPtrFromFramework(dd.ApiKey),
			Ddtags:  stringPtrFromFramework(dd.Ddtags),
			Service: stringPtrFromFramework(dd.Service),
			Ddsite:  stringPtrFromFramework(dd.Ddsite),
		}
	}

	if !model.AzureLogsIngestion.IsNull() && !model.AzureLogsIngestion.IsUnknown() {
		var az azureLogsIngestionModel
		diags.Append(model.AzureLogsIngestion.As(ctx, &az, basetypes.ObjectAsOptions{})...)
		target.AzureLogsIngestion = &ngrok.EventTargetAzureLogsIngestion{
			TenantId:                 az.TenantId.ValueString(),
			ClientId:                 az.ClientId.ValueString(),
			ClientSecret:             az.ClientSecret.ValueString(),
			LogsIngestionURI:         az.LogsIngestionURI.ValueString(),
			DataCollectionRuleId:     az.DataCollectionRuleId.ValueString(),
			DataCollectionStreamName: az.DataCollectionStreamName.ValueString(),
		}
	}

	return target
}

func flattenEventDestination(ctx context.Context, dest *ngrok.EventDestination, model *eventDestinationResourceModel, diags *diag.Diagnostics) {
	model.ID = types.StringValue(dest.ID)
	model.Description = types.StringValue(dest.Description)
	model.Metadata = types.StringValue(dest.Metadata)
	model.Format = types.StringValue(dest.Format)
	model.URI = types.StringValue(dest.URI)
	model.CreatedAt = types.StringValue(dest.CreatedAt)
	// For normal CRUD, preserve target from plan/state since the API redacts
	// sensitive fields (api_key, aws_secret_access_key, client_secret).
	// For import (model.Target is unknown/null), flatten from the API so
	// state is populated — sensitive fields will be redacted but at least
	// the non-sensitive structure is present.
	if model.Target.IsNull() || model.Target.IsUnknown() {
		model.Target = flattenEventTarget(ctx, dest.Target, model.Target, diags)
	}
}
