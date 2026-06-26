package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/reserved_addrs"
	"github.com/ngrok/terraform-provider-ngrok/internal/datasource_reserved_addr"
)

var _ datasource.DataSource = &reservedAddrDataSource{}

type reservedAddrDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	Addr        types.String `tfsdk:"addr"`
	Region      types.String `tfsdk:"region"`
	Description types.String `tfsdk:"description"`
	Metadata    types.String `tfsdk:"metadata"`
	URI         types.String `tfsdk:"uri"`
	CreatedAt   types.String `tfsdk:"created_at"`
}

type reservedAddrDataSource struct {
	client *reserved_addrs.Client
}

func NewReservedAddrDataSource() datasource.DataSource {
	return &reservedAddrDataSource{}
}

func (d *reservedAddrDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tcp_address"
}

func (d *reservedAddrDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	s := datasource_reserved_addr.ReservedAddrDataSourceSchema(ctx)
	attrs := s.Attributes
	// id: change from Required to Optional+Computed for lookup-by-addr
	idAttr := attrs["id"].(schema.StringAttribute)
	idAttr.Required = false
	idAttr.Optional = true
	idAttr.Computed = true
	attrs["id"] = idAttr
	// addr: change from Computed to Optional+Computed for lookup-by-addr
	addrAttr := attrs["addr"].(schema.StringAttribute)
	addrAttr.Optional = true
	attrs["addr"] = addrAttr
	resp.Schema = s
}

func (d *reservedAddrDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clientConfig, ok := req.ProviderData.(*ngrok.ClientConfig)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *ngrok.ClientConfig, got: %T.", req.ProviderData),
		)
		return
	}
	d.client = reserved_addrs.NewClient(clientConfig)
}

func (d *reservedAddrDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config reservedAddrDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var addr *ngrok.ReservedAddr

	if !config.ID.IsNull() && config.ID.ValueString() != "" {
		var err error
		addr, err = d.client.Get(ctx, config.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Error reading reserved address", err.Error())
			return
		}
	} else if !config.Addr.IsNull() && config.Addr.ValueString() != "" {
		filter := fmt.Sprintf(`obj.addr == %q`, config.Addr.ValueString())
		iter := d.client.List(&ngrok.FilteredPaging{
			Filter: &filter,
		})

		if !iter.Next(ctx) {
			if err := iter.Err(); err != nil {
				resp.Diagnostics.AddError("Error listing reserved addresses", err.Error())
				return
			}
			resp.Diagnostics.AddError(
				"Reserved address not found",
				fmt.Sprintf("No reserved address found with addr %q.", config.Addr.ValueString()),
			)
			return
		}
		addr = iter.Item()

		if iter.Next(ctx) {
			resp.Diagnostics.AddError(
				"Multiple reserved addresses found",
				fmt.Sprintf("More than one reserved address found with addr %q. Use id instead.", config.Addr.ValueString()),
			)
			return
		}
		if err := iter.Err(); err != nil {
			resp.Diagnostics.AddError("Error listing reserved addresses", err.Error())
			return
		}
	} else {
		resp.Diagnostics.AddError(
			"Missing lookup attribute",
			"Either id or addr must be specified.",
		)
		return
	}

	var model reservedAddrDataSourceModel
	flattenReservedAddrDataSource(addr, &model)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func flattenReservedAddrDataSource(addr *ngrok.ReservedAddr, model *reservedAddrDataSourceModel) {
	model.ID = types.StringValue(addr.ID)
	model.Addr = types.StringValue(addr.Addr)
	model.Region = types.StringValue(addr.Region)
	model.Description = types.StringValue(addr.Description)
	model.Metadata = types.StringValue(addr.Metadata)
	model.URI = types.StringValue(addr.URI)
	model.CreatedAt = types.StringValue(addr.CreatedAt)
}
