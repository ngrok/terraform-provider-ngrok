---
page_title: "Upgrading to Version 1.0"
description: |-
  Guide for migrating from terraform-provider-ngrok v0.x to v1.0.
---

# Upgrading to v1.0

Version 1.0 of the ngrok Terraform provider is a major release that rebuilds the provider from the ground up. This guide walks you through upgrading from v0.x.

## What Changed and Why

The v1.0 provider is a complete rewrite:

- **Terraform Plugin Framework** — Replaces the legacy `terraform-plugin-sdk/v2` (SDKv2), which is now in maintenance mode. The Plugin Framework provides better type safety, native nested objects, plan modifiers, and validators.
- **Official `ngrok-api-go/v9` client** — Replaces the hand-rolled `restapi/` package. The provider now uses ngrok's official, maintained Go API client.
- **Cloud Endpoints replace Edges** — The ngrok platform has deprecated HTTPS, TCP, and TLS edges in favor of **cloud endpoints** with **traffic policy**. All edge, backend, and endpoint configuration resources have been removed.
- **Data sources** — Every resource now has a corresponding data source for looking up existing resources.
- **Import support** — Every resource supports `terraform import`.

## Upgrade Steps

### Step 1: Update the Version Constraint

In your `required_providers` block, update the version constraint to `~> 1.0`:

```hcl
terraform {
  required_providers {
    ngrok = {
      source  = "ngrok/ngrok"
      version = "~> 1.0"
    }
  }
}
```

-> **Note:** If you had `version = "~> 0.7"` or similar, Terraform will **not** auto-upgrade. You must explicitly change the constraint.

### Step 2: Run `terraform init -upgrade`

```shell
terraform init -upgrade
```

This downloads the v1.0 provider binary.

### Step 3: Remove Deprecated Resource Blocks

The following resources have been removed in v1.0. For each one present in your configuration:

1. Remove the resource block from your `.tf` files.
2. Remove it from Terraform state with `terraform state rm`.

```shell
# Edges
terraform state rm ngrok_edge_https.example
terraform state rm ngrok_edge_tcp.example
terraform state rm ngrok_edge_tls.example

# Backends
terraform state rm ngrok_failover_backend.example
terraform state rm ngrok_weighted_backend.example
terraform state rm ngrok_tunnel_group_backend.example
terraform state rm ngrok_http_response_backend.example

# Endpoint Configurations
terraform state rm ngrok_endpoint_configuration.example
```

~> **Important:** `terraform state rm` removes the resource from Terraform's tracking only — it does **not** delete the actual infrastructure. Your existing edges and backends continue to function on ngrok until you migrate them or they are removed from the platform.

### Step 4: Remove Deprecated Fields from Carried-Over Resources

Some resources that carry over to v1.0 have removed or deprecated fields. Update your HCL:

**`ngrok_reserved_domain`:**

- `region` — Deprecated. ngrok now handles traffic globally via the Global Network. Remove this field. If present, the provider will ignore it with a deprecation warning.
- `http_endpoint_configuration_id` — Removed. Delete this field.
- `https_endpoint_configuration_id` — Removed. Delete this field.

```hcl
# Before (v0.x)
resource "ngrok_reserved_domain" "example" {
  domain = "app.example.com"
  region = "us"
  http_endpoint_configuration_id = ngrok_endpoint_configuration.example.id
}

# After (v1.0)
resource "ngrok_reserved_domain" "example" {
  domain = "app.example.com"
}
```

### Step 5: Run `terraform plan` to Verify

```shell
terraform plan
```

For resources that carry over (reserved domains, reserved addresses, API keys, credentials, certificates, IP policies, etc.), the plan should show **no changes**. If it does, review the diff — it likely indicates a field that needs to be removed.

### Step 6: Replace Edge Workflows with Cloud Endpoints

The core migration task is replacing your edge + backend + endpoint configuration setup with the new `ngrok_cloud_endpoint` resource and traffic policy. See the [example migration](#example-migration) below.

## Resource Mapping

The following table maps every v0.x resource to its v1.0 equivalent:

| v0.x Resource | v1.0 Resource | Notes |
|---|---|---|
| `ngrok_reserved_domain` | `ngrok_reserved_domain` | Carried over. Remove `region` and edge-related fields. |
| `ngrok_reserved_addr` | `ngrok_reserved_addr` | Carried over. |
| `ngrok_api_key` | `ngrok_api_key` | Carried over. |
| `ngrok_credential` | `ngrok_credential` | Carried over. |
| `ngrok_certificate_authority` | `ngrok_certificate_authority` | Carried over. |
| `ngrok_tls_certificate` | `ngrok_tls_certificate` | Carried over. |
| `ngrok_ip_policy` | `ngrok_ip_policy` | Carried over. |
| `ngrok_ip_policy_rule` | `ngrok_ip_policy_rule` | Carried over. |
| `ngrok_ip_restriction` | `ngrok_ip_restriction` | Carried over. |
| `ngrok_service_user` | `ngrok_service_user` | Carried over. |
| `ngrok_agent_ingress` | `ngrok_agent_ingress` | Carried over. |
| `ngrok_event_destination` | `ngrok_event_destination` | Carried over. |
| `ngrok_event_subscription` | `ngrok_event_subscription` | Carried over. |
| `ngrok_ssh_certificate_authority` | `ngrok_ssh_certificate_authority` | Carried over. |
| `ngrok_ssh_credential` | `ngrok_ssh_credential` | Carried over. |
| `ngrok_ssh_host_certificate` | `ngrok_ssh_host_certificate` | Carried over. |
| `ngrok_ssh_user_certificate` | `ngrok_ssh_user_certificate` | Carried over. |
| `ngrok_edge_https` | **Removed** | Replace with `ngrok_cloud_endpoint`. |
| `ngrok_edge_tcp` | **Removed** | Replace with `ngrok_cloud_endpoint`. |
| `ngrok_edge_tls` | **Removed** | Replace with `ngrok_cloud_endpoint`. |
| `ngrok_failover_backend` | **Removed** | Use traffic policy on `ngrok_cloud_endpoint`. |
| `ngrok_weighted_backend` | **Removed** | Use traffic policy on `ngrok_cloud_endpoint`. |
| `ngrok_tunnel_group_backend` | **Removed** | Use traffic policy on `ngrok_cloud_endpoint`. |
| `ngrok_http_response_backend` | **Removed** | Use traffic policy on `ngrok_cloud_endpoint`. |
| `ngrok_endpoint_configuration` | **Removed** | Use traffic policy on `ngrok_cloud_endpoint`. |
| — | **`ngrok_cloud_endpoint`** | New. Cloud endpoints with traffic policy. |
| — | **`ngrok_vault`** | New. Secret management vaults. |
| — | **`ngrok_secret`** | New. Secret management secrets. |

## Example Migration

The following example shows how to convert a typical v0.x HTTPS edge setup into a v1.0 cloud endpoint with traffic policy.

### Before: v0.x (Edge + Backend + Endpoint Configuration)

```hcl
resource "ngrok_reserved_domain" "app" {
  domain = "app.example.com"
  region = "us"
}

resource "ngrok_tunnel_group_backend" "app" {
  description = "app backend"
  labels = {
    env = "production"
  }
}

resource "ngrok_endpoint_configuration" "app" {
  type        = "https"
  description = "app endpoint config"

  rate_limit {
    algorithm   = "sliding_window"
    capacity    = 100
    rate        = "60s"
  }

  ip_restriction {
    ip_policies = [ngrok_ip_policy.office.id]
  }
}

resource "ngrok_edge_https" "app" {
  description = "Production API edge"
  hostports   = ["app.example.com:443"]

  https_endpoint_configuration_id = ngrok_endpoint_configuration.app.id

  route {
    match      = "/"
    match_type = "path_prefix"
    backend {
      backend_id = ngrok_tunnel_group_backend.app.id
    }
  }
}
```

### After: v1.0 (Cloud Endpoint + Traffic Policy)

```hcl
resource "ngrok_reserved_domain" "app" {
  domain = "app.example.com"
}

resource "ngrok_cloud_endpoint" "app" {
  url = "https://${ngrok_reserved_domain.app.domain}"

  description = "Production API endpoint"

  traffic_policy = jsonencode({
    on_http_request = [
      {
        actions = [
          {
            type = "rate-limit"
            config = {
              name      = "global"
              algorithm = "sliding_window"
              capacity  = 100
              rate      = "60s"
            }
          }
        ]
      },
      {
        actions = [
          {
            type = "restrict-ips"
            config = {
              enforce = true
              ip_policies = [ngrok_ip_policy.office.id]
            }
          }
        ]
      }
    ]
  })
}
```

### Migration Commands

After updating your configuration files, remove the old resources from state:

```shell
terraform state rm ngrok_tunnel_group_backend.app
terraform state rm ngrok_endpoint_configuration.app
terraform state rm ngrok_edge_https.app
```

Then apply the new configuration:

```shell
terraform plan    # Review the changes
terraform apply   # Create the cloud endpoint
```

-> **Tip:** Use `jsonencode()` for traffic policy instead of `file()`. It provides readable diffs when the policy changes and supports Terraform resource interpolation (e.g., referencing `ngrok_ip_policy.office.id` directly in the policy).
