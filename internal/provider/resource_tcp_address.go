package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/reserved_addrs"
	"github.com/ngrok/terraform-provider-ngrok/internal/resource_reserved_addr"
)

var (
	_ resource.Resource                = &reservedAddrResource{}
	_ resource.ResourceWithImportState = &reservedAddrResource{}
)

type reservedAddrResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Addr        types.String `tfsdk:"addr"`
	Region      types.String `tfsdk:"region"`
	Description types.String `tfsdk:"description"`
	Metadata    types.String `tfsdk:"metadata"`
	URI         types.String `tfsdk:"uri"`
	CreatedAt   types.String `tfsdk:"created_at"`
}

type reservedAddrResource struct {
	client *reserved_addrs.Client
}

func NewReservedAddrResource() resource.Resource {
	return &reservedAddrResource{}
}

func (r *reservedAddrResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tcp_address"
}

func (r *reservedAddrResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	s := resource_reserved_addr.ReservedAddrResourceSchema(ctx)
	attrs := s.Attributes

	addStringPlanModifiers(attrs, "id", useStateForUnknownString())
	addStringPlanModifiers(attrs, "addr", useStateForUnknownString())
	addStringPlanModifiers(attrs, "region", requiresReplaceString(), useStateForUnknownString())
	addStringPlanModifiers(attrs, "description", useStateForUnknownString())
	addStringPlanModifiers(attrs, "metadata", useStateForUnknownString())
	addStringPlanModifiers(attrs, "uri", useStateForUnknownString())
	addStringPlanModifiers(attrs, "created_at", useStateForUnknownString())

	resp.Schema = s
}

func (r *reservedAddrResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = reserved_addrs.NewClient(clientConfig)
}

func (r *reservedAddrResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan reservedAddrResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq := &ngrok.ReservedAddrCreate{
		Description: plan.Description.ValueString(),
		Metadata:    plan.Metadata.ValueString(),
	}
	if !plan.Region.IsNull() && !plan.Region.IsUnknown() {
		createReq.Region = plan.Region.ValueString()
	}

	addr, err := r.client.Create(ctx, createReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating reserved address", err.Error())
		return
	}

	flattenReservedAddr(addr, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *reservedAddrResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state reservedAddrResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	addr, err := r.client.Get(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading reserved address", err.Error())
		return
	}

	flattenReservedAddr(addr, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *reservedAddrResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan reservedAddrResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state reservedAddrResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq := &ngrok.ReservedAddrUpdate{
		ID:          state.ID.ValueString(),
		Description: stringPtrFromFramework(plan.Description),
		Metadata:    stringPtrFromFramework(plan.Metadata),
	}

	addr, err := r.client.Update(ctx, updateReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating reserved address", err.Error())
		return
	}

	flattenReservedAddr(addr, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *reservedAddrResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state reservedAddrResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			return
		}
		resp.Diagnostics.AddError("Error deleting reserved address", err.Error())
	}
}

func (r *reservedAddrResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func flattenReservedAddr(addr *ngrok.ReservedAddr, model *reservedAddrResourceModel) {
	model.ID = types.StringValue(addr.ID)
	model.Addr = types.StringValue(addr.Addr)
	model.Region = types.StringValue(addr.Region)
	model.Description = types.StringValue(addr.Description)
	model.Metadata = types.StringValue(addr.Metadata)
	model.URI = types.StringValue(addr.URI)
	model.CreatedAt = types.StringValue(addr.CreatedAt)
}
