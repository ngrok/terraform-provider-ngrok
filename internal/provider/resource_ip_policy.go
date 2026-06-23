package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/ip_policies"
	"github.com/ngrok/terraform-provider-ngrok/internal/resource_ip_policy"
)

var (
	_ resource.Resource                = &ipPolicyResource{}
	_ resource.ResourceWithImportState = &ipPolicyResource{}
)

type ipPolicyResourceModel struct {
	ID          types.String `tfsdk:"id"`
	URI         types.String `tfsdk:"uri"`
	CreatedAt   types.String `tfsdk:"created_at"`
	Description types.String `tfsdk:"description"`
	Metadata    types.String `tfsdk:"metadata"`
}

type ipPolicyResource struct {
	client *ip_policies.Client
}

func NewIPPolicyResource() resource.Resource {
	return &ipPolicyResource{}
}

func (r *ipPolicyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_policy"
}

func (r *ipPolicyResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	s := resource_ip_policy.IpPolicyResourceSchema(ctx)
	attrs := s.Attributes

	addStringPlanModifiers(attrs, "id", useStateForUnknownString())
	addStringPlanModifiers(attrs, "uri", useStateForUnknownString())
	addStringPlanModifiers(attrs, "created_at", useStateForUnknownString())
	addStringPlanModifiers(attrs, "description", useStateForUnknownString())
	addStringPlanModifiers(attrs, "metadata", useStateForUnknownString())

	resp.Schema = s
}

func (r *ipPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = ip_policies.NewClient(clientConfig)
}

func (r *ipPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ipPolicyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	policy, err := r.client.Create(ctx, &ngrok.IPPolicyCreate{
		Description: plan.Description.ValueString(),
		Metadata:    plan.Metadata.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Error creating IP policy", err.Error())
		return
	}

	flattenIPPolicy(policy, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ipPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ipPolicyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	policy, err := r.client.Get(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading IP policy", err.Error())
		return
	}

	flattenIPPolicy(policy, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ipPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ipPolicyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state ipPolicyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	policy, err := r.client.Update(ctx, &ngrok.IPPolicyUpdate{
		ID:          state.ID.ValueString(),
		Description: stringPtrFromFramework(plan.Description),
		Metadata:    stringPtrFromFramework(plan.Metadata),
	})
	if err != nil {
		resp.Diagnostics.AddError("Error updating IP policy", err.Error())
		return
	}

	flattenIPPolicy(policy, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ipPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ipPolicyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			return
		}
		resp.Diagnostics.AddError("Error deleting IP policy", err.Error())
	}
}

func (r *ipPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func flattenIPPolicy(policy *ngrok.IPPolicy, model *ipPolicyResourceModel) {
	model.ID = types.StringValue(policy.ID)
	model.URI = types.StringValue(policy.URI)
	model.CreatedAt = types.StringValue(policy.CreatedAt)
	model.Description = types.StringValue(policy.Description)
	model.Metadata = types.StringValue(policy.Metadata)
}
