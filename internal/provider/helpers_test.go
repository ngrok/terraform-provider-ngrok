package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestPreserveEquivalentDomain(t *testing.T) {
	tests := []struct {
		name     string
		apiValue string
		prior    types.String
		want     string
	}{
		{"trailing dot preserved", "example.com", types.StringValue("example.com."), "example.com."},
		{"uppercase preserved", "example.com", types.StringValue("Example.Com"), "Example.Com"},
		{"whitespace preserved", "example.com", types.StringValue(" example.com "), " example.com "},
		{"different domain not preserved", "other.com", types.StringValue("example.com"), "other.com"},
		{"null prior returns api", "example.com", types.StringNull(), "example.com"},
		{"unknown prior returns api", "example.com", types.StringUnknown(), "example.com"},
		{"exact match preserved", "example.com", types.StringValue("example.com"), "example.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := preserveEquivalentDomain(tt.apiValue, tt.prior)
			if got.ValueString() != tt.want {
				t.Errorf("preserveEquivalentDomain(%q, %q) = %q, want %q",
					tt.apiValue, tt.prior.ValueString(), got.ValueString(), tt.want)
			}
		})
	}
}

func TestPreserveEquivalentURL(t *testing.T) {
	tests := []struct {
		name     string
		apiValue string
		prior    types.String
		want     string
	}{
		{"uppercase scheme preserved", "https://example.com", types.StringValue("HTTPS://example.com"), "HTTPS://example.com"},
		{"explicit 443 preserved", "https://example.com", types.StringValue("https://example.com:443"), "https://example.com:443"},
		{"explicit 80 preserved", "http://example.com", types.StringValue("http://example.com:80"), "http://example.com:80"},
		{"non-default port not preserved", "https://example.com:8443", types.StringValue("https://example.com"), "https://example.com:8443"},
		{"different host not preserved", "https://other.com", types.StringValue("https://example.com"), "https://other.com"},
		{"null prior returns api", "https://example.com", types.StringNull(), "https://example.com"},
		{"exact match preserved", "https://example.com", types.StringValue("https://example.com"), "https://example.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := preserveEquivalentURL(tt.apiValue, tt.prior)
			if got.ValueString() != tt.want {
				t.Errorf("preserveEquivalentURL(%q, %q) = %q, want %q",
					tt.apiValue, tt.prior.ValueString(), got.ValueString(), tt.want)
			}
		})
	}
}

func TestPreserveEquivalentCIDR(t *testing.T) {
	tests := []struct {
		name     string
		apiValue string
		prior    types.String
		want     string
	}{
		{"uppercase IPv6 preserved", "2001:db8::/32", types.StringValue("2001:DB8::/32"), "2001:DB8::/32"},
		{"whitespace preserved", "10.0.0.0/8", types.StringValue(" 10.0.0.0/8 "), " 10.0.0.0/8 "},
		{"different CIDR not preserved", "192.168.0.0/16", types.StringValue("10.0.0.0/8"), "192.168.0.0/16"},
		{"null prior returns api", "10.0.0.0/8", types.StringNull(), "10.0.0.0/8"},
		{"exact match preserved", "10.0.0.0/8", types.StringValue("10.0.0.0/8"), "10.0.0.0/8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := preserveEquivalentCIDR(tt.apiValue, tt.prior)
			if got.ValueString() != tt.want {
				t.Errorf("preserveEquivalentCIDR(%q, %q) = %q, want %q",
					tt.apiValue, tt.prior.ValueString(), got.ValueString(), tt.want)
			}
		})
	}
}

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want string
	}{
		{"strips default https port", "https://example.com:443", "https://example.com"},
		{"strips default http port", "http://example.com:80", "http://example.com"},
		{"keeps non-default port", "https://example.com:8443", "https://example.com:8443"},
		{"lowercases scheme", "HTTPS://example.com", "https://example.com"},
		{"no port unchanged", "https://example.com", "https://example.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeURL(tt.raw)
			if got != tt.want {
				t.Errorf("normalizeURL(%q) = %q, want %q", tt.raw, got, tt.want)
			}
		})
	}
}
