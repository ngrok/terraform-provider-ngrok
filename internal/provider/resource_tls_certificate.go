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

	ngrok "github.com/ngrok/ngrok-api-go/v9"
	"github.com/ngrok/ngrok-api-go/v9/tls_certificates"
	"github.com/ngrok/terraform-provider-ngrok/internal/resource_tls_certificate"
)

var (
	_ resource.Resource                = &tlsCertificateResource{}
	_ resource.ResourceWithImportState = &tlsCertificateResource{}
)

type tlsCertificateResourceModel struct {
	ID                        types.String `tfsdk:"id"`
	URI                       types.String `tfsdk:"uri"`
	CreatedAt                 types.String `tfsdk:"created_at"`
	Description               types.String `tfsdk:"description"`
	Metadata                  types.String `tfsdk:"metadata"`
	CertificatePEM            types.String `tfsdk:"certificate_pem"`
	PrivateKeyPEM             types.String `tfsdk:"private_key_pem"`
	SubjectCommonName         types.String `tfsdk:"subject_common_name"`
	SubjectAlternativeNames   types.Object `tfsdk:"subject_alternative_names"`
	IssuedAt                  types.String `tfsdk:"issued_at"`
	NotBefore                 types.String `tfsdk:"not_before"`
	NotAfter                  types.String `tfsdk:"not_after"`
	KeyUsages                 types.List   `tfsdk:"key_usages"`
	ExtendedKeyUsages         types.List   `tfsdk:"extended_key_usages"`
	PrivateKeyType            types.String `tfsdk:"private_key_type"`
	IssuerCommonName          types.String `tfsdk:"issuer_common_name"`
	SerialNumber              types.String `tfsdk:"serial_number"`
	SubjectOrganization       types.String `tfsdk:"subject_organization"`
	SubjectOrganizationalUnit types.String `tfsdk:"subject_organizational_unit"`
	SubjectLocality           types.String `tfsdk:"subject_locality"`
	SubjectProvince           types.String `tfsdk:"subject_province"`
	SubjectCountry            types.String `tfsdk:"subject_country"`
}

type tlsCertificateResource struct {
	client *tls_certificates.Client
}

func NewTLSCertificateResource() resource.Resource {
	return &tlsCertificateResource{}
}

func (r *tlsCertificateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tls_certificate"
}

func (r *tlsCertificateResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	s := resource_tls_certificate.TlsCertificateResourceSchema(ctx)
	attrs := s.Attributes

	// Replace generated CustomType subject_alternative_names with plain SingleNestedAttribute
	// so the hand-written flatten function (which uses types.Object) works correctly.
	delete(attrs, "subject_alternative_names")
	attrs["subject_alternative_names"] = schema.SingleNestedAttribute{
		Description: "Subject alternative names (SANs) from the leaf of this TLS certificate.",
		Computed:    true,
		Attributes: map[string]schema.Attribute{
			"dns_names": schema.ListAttribute{
				Description: "Set of additional domains (including wildcards) this TLS certificate is valid for.",
				Computed:    true,
				ElementType: types.StringType,
			},
			"ips": schema.ListAttribute{
				Description: "Set of IP addresses this TLS certificate is also valid for.",
				Computed:    true,
				ElementType: types.StringType,
			},
		},
	}

	addStringPlanModifiers(attrs, "id", useStateForUnknownString())
	addStringPlanModifiers(attrs, "uri", useStateForUnknownString())
	addStringPlanModifiers(attrs, "created_at", useStateForUnknownString())
	addStringPlanModifiers(attrs, "description", useStateForUnknownString())
	addStringPlanModifiers(attrs, "metadata", useStateForUnknownString())
	addStringPlanModifiers(attrs, "certificate_pem", requiresReplaceString())
	addStringPlanModifiers(attrs, "private_key_pem", requiresReplaceString())
	markSensitive(attrs, "private_key_pem")
	addStringPlanModifiers(attrs, "subject_common_name", useStateForUnknownString())
	addObjectPlanModifiers(attrs, "subject_alternative_names", useStateForUnknownObject())
	addStringPlanModifiers(attrs, "issued_at", useStateForUnknownString())
	addStringPlanModifiers(attrs, "not_before", useStateForUnknownString())
	addStringPlanModifiers(attrs, "not_after", useStateForUnknownString())
	addListPlanModifiers(attrs, "key_usages", useStateForUnknownList())
	addListPlanModifiers(attrs, "extended_key_usages", useStateForUnknownList())
	addStringPlanModifiers(attrs, "private_key_type", useStateForUnknownString())
	addStringPlanModifiers(attrs, "issuer_common_name", useStateForUnknownString())
	addStringPlanModifiers(attrs, "serial_number", useStateForUnknownString())
	addStringPlanModifiers(attrs, "subject_organization", useStateForUnknownString())
	addStringPlanModifiers(attrs, "subject_organizational_unit", useStateForUnknownString())
	addStringPlanModifiers(attrs, "subject_locality", useStateForUnknownString())
	addStringPlanModifiers(attrs, "subject_province", useStateForUnknownString())
	addStringPlanModifiers(attrs, "subject_country", useStateForUnknownString())

	resp.Schema = s
}

func (r *tlsCertificateResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = tls_certificates.NewClient(clientConfig)
}

func (r *tlsCertificateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan tlsCertificateResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq := &ngrok.TLSCertificateCreate{
		Description:    plan.Description.ValueString(),
		Metadata:       plan.Metadata.ValueString(),
		CertificatePEM: plan.CertificatePEM.ValueString(),
		PrivateKeyPEM:  plan.PrivateKeyPEM.ValueString(),
	}

	cert, err := r.client.Create(ctx, createReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating TLS certificate", err.Error())
		return
	}

	flattenTLSCertificate(ctx, cert, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *tlsCertificateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state tlsCertificateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	cert, err := r.client.Get(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading TLS certificate", err.Error())
		return
	}

	// Preserve private_key_pem from state since the API does not return it.
	privateKeyPEM := state.PrivateKeyPEM

	flattenTLSCertificate(ctx, cert, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	state.PrivateKeyPEM = privateKeyPEM
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *tlsCertificateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan tlsCertificateResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state tlsCertificateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq := &ngrok.TLSCertificateUpdate{
		ID:          state.ID.ValueString(),
		Description: stringPtrFromFramework(plan.Description),
		Metadata:    stringPtrFromFramework(plan.Metadata),
	}

	cert, err := r.client.Update(ctx, updateReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating TLS certificate", err.Error())
		return
	}

	// Preserve private_key_pem from plan since the API does not return it.
	privateKeyPEM := plan.PrivateKeyPEM

	flattenTLSCertificate(ctx, cert, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.PrivateKeyPEM = privateKeyPEM
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *tlsCertificateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state tlsCertificateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, state.ID.ValueString())
	if err != nil {
		if isNotFound(err) {
			return
		}
		resp.Diagnostics.AddError("Error deleting TLS certificate", err.Error())
	}
}

func (r *tlsCertificateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func sanAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"dns_names": types.ListType{ElemType: types.StringType},
		"ips":       types.ListType{ElemType: types.StringType},
	}
}

func flattenSANs(ctx context.Context, sans ngrok.TLSCertificateSANs, diags *diag.Diagnostics) types.Object {
	dnsNames, d := types.ListValueFrom(ctx, types.StringType, sans.DNSNames)
	diags.Append(d...)
	ips, d := types.ListValueFrom(ctx, types.StringType, sans.IPs)
	diags.Append(d...)
	if diags.HasError() {
		return types.ObjectNull(sanAttrTypes())
	}

	obj, d := types.ObjectValue(sanAttrTypes(), map[string]attr.Value{
		"dns_names": dnsNames,
		"ips":       ips,
	})
	diags.Append(d...)
	return obj
}

func flattenTLSCertificate(ctx context.Context, cert *ngrok.TLSCertificate, model *tlsCertificateResourceModel, diags *diag.Diagnostics) {
	model.ID = types.StringValue(cert.ID)
	model.URI = types.StringValue(cert.URI)
	model.CreatedAt = types.StringValue(cert.CreatedAt)
	model.Description = types.StringValue(cert.Description)
	model.Metadata = types.StringValue(cert.Metadata)
	model.CertificatePEM = types.StringValue(cert.CertificatePEM)
	model.SubjectCommonName = types.StringValue(cert.SubjectCommonName)
	model.SubjectAlternativeNames = flattenSANs(ctx, cert.SubjectAlternativeNames, diags)

	if cert.IssuedAt != nil {
		model.IssuedAt = types.StringValue(*cert.IssuedAt)
	} else {
		model.IssuedAt = types.StringNull()
	}

	model.NotBefore = types.StringValue(cert.NotBefore)
	model.NotAfter = types.StringValue(cert.NotAfter)
	keyUsages, d := types.ListValueFrom(ctx, types.StringType, cert.KeyUsages)
	diags.Append(d...)
	model.KeyUsages = keyUsages
	extKeyUsages, d := types.ListValueFrom(ctx, types.StringType, cert.ExtendedKeyUsages)
	diags.Append(d...)
	model.ExtendedKeyUsages = extKeyUsages
	model.PrivateKeyType = types.StringValue(cert.PrivateKeyType)
	model.IssuerCommonName = types.StringValue(cert.IssuerCommonName)
	model.SerialNumber = types.StringValue(cert.SerialNumber)
	model.SubjectOrganization = types.StringValue(cert.SubjectOrganization)
	model.SubjectOrganizationalUnit = types.StringValue(cert.SubjectOrganizationalUnit)
	model.SubjectLocality = types.StringValue(cert.SubjectLocality)
	model.SubjectProvince = types.StringValue(cert.SubjectProvince)
	model.SubjectCountry = types.StringValue(cert.SubjectCountry)
}
