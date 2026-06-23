# ngrok Terraform Provider (v2)

A Terraform provider for managing [ngrok](https://ngrok.com) resources, built with the [Terraform Plugin Framework](https://developer.hashicorp.com/terraform/plugin/framework).

This is a ground-up rewrite of the [original ngrok provider](https://github.com/ngrok/terraform-provider-ngrok), replacing the legacy SDKv2 with the Plugin Framework and the hand-rolled REST client with the official [ngrok-api-go/v9](https://github.com/ngrok/ngrok-api-go) client.

## Usage

```hcl
terraform {
  required_providers {
    ngrok = {
      source  = "ngrok/ngrok"
      version = "~> 1.0"
    }
  }
}

provider "ngrok" {}

resource "ngrok_reserved_domain" "example" {
  domain      = "app.example.com"
  description = "Production app domain"
}

resource "ngrok_cloud_endpoint" "example" {
  url = "https://${ngrok_reserved_domain.example.domain}"

  traffic_policy = jsonencode({
    on_http_request = [
      {
        actions = [
          {
            type   = "custom-response"
            config = { status_code = 200, content = "Hello from ngrok!" }
          }
        ]
      }
    ]
  })
}
```

Configure the API key via the `NGROK_API_KEY` environment variable or in the provider block.

## Resources

| Resource | Description |
|---|---|
| `ngrok_api_key` | API keys for authenticating to the ngrok API |
| `ngrok_agent_ingress` | Custom agent ingress domains |
| `ngrok_certificate_authority` | Certificate authorities for mTLS |
| `ngrok_cloud_endpoint` | Cloud endpoints with traffic policy |
| `ngrok_credential` | Tunnel authtokens |
| `ngrok_event_destination` | Event destination targets (Datadog, AWS, Azure) |
| `ngrok_event_subscription` | Event subscriptions |
| `ngrok_ip_policy` | IP policy groups |
| `ngrok_ip_policy_rule` | IP policy CIDR rules |
| `ngrok_ip_restriction` | IP restrictions on API/dashboard/agent/endpoints |
| `ngrok_reserved_addr` | Reserved TCP addresses |
| `ngrok_reserved_domain` | Reserved domains |
| `ngrok_secret` | Secrets stored in vaults |
| `ngrok_service_user` | Service users (bot users) |
| `ngrok_ssh_certificate_authority` | SSH certificate authorities |
| `ngrok_ssh_credential` | SSH credentials |
| `ngrok_ssh_host_certificate` | SSH host certificates |
| `ngrok_ssh_user_certificate` | SSH user certificates |
| `ngrok_tls_certificate` | TLS certificates |
| `ngrok_vault` | Secret management vaults |

Every resource also has a corresponding **data source** for lookups. Several data sources support dual lookup (by ID or by name/domain):

- `ngrok_reserved_domain` — by `id` or `domain`
- `ngrok_secret` — by `id` or `name`
- `ngrok_vault` — by `id` or `name`

There are also **read-only data sources** with no corresponding resource:

| Data Source | Description |
|---|---|
| `ngrok_application_session` | Active application sessions |
| `ngrok_application_user` | Application users from OAuth/OIDC |
| `ngrok_tunnel_session` | Active tunnel/agent sessions |

## What changed from v0.x

- **Terraform Plugin Framework** replaces legacy SDKv2
- **Official `ngrok-api-go/v9` client** replaces hand-rolled REST client
- **Cloud endpoints** replace edges — all edge, backend, and endpoint configuration resources are removed
- **New resources**: `ngrok_cloud_endpoint`, `ngrok_vault`, `ngrok_secret`
- **Data sources** for every resource, plus read-only data sources for sessions and application users
- **Import support** for every resource
- **OpenAPI-driven code generation** for schema definitions — see [Project Structure](#project-structure)

See the [upgrade guide](docs/guides/version-1-upgrade.md) for migration details.

## Project Structure

```
├── codegen/                          # Code generation inputs
│   ├── openapi_spec.yaml             # ngrok OpenAPI spec (source of truth)
│   ├── generator_config.yml          # tfplugingen-openapi config
│   └── provider_code_spec.json       # Intermediate spec (generated)
├── internal/
│   ├── provider/                     # Hand-written provider, resource, and data source logic
│   │   ├── provider.go               # Provider definition and configuration
│   │   ├── resource_*.go             # Resource CRUD + flatten/expand (uses generated schemas)
│   │   ├── datasource_*.go           # Data source read + flatten (uses generated schemas)
│   │   ├── helpers.go                # Shared flatten/expand utilities
│   │   └── schema_overrides.go       # Helpers for plan modifiers, defaults, attribute overrides
│   ├── resource_*/                   # Generated resource schema packages (one per resource)
│   ├── datasource_*/                 # Generated data source schema packages (one per data source)
│   └── provider_ngrok/               # Generated provider schema package
├── docs/                             # Terraform registry documentation
│   └── design/                       # Architecture decision records
├── test-manual/                      # Manual testing configs (not committed)
└── Makefile                          # Build, install, codegen targets
```

**Schema generation flow:**

```
openapi_spec.yaml → tfplugingen-openapi → provider_code_spec.json → tfplugingen-framework → internal/{resource,datasource}_*/*_gen.go
```

The generated `*_gen.go` files provide base schema definitions and model structs. The hand-written `resource_*.go` and `datasource_*.go` files import these and apply overrides: replacing CustomType attributes with standard types, flattening `Ref` objects to `_id` fields, and adding plan modifiers.

See [docs/design/openapi-codegen.md](docs/design/openapi-codegen.md) for the full design.

## Development

```bash
# Build and install locally
make install

# Regenerate schemas from OpenAPI spec
make codegen

# Run tests
go test ./...

# Run acceptance tests (requires NGROK_API_KEY)
TF_ACC=1 NGROK_API_KEY=your-key go test ./... -v -timeout 120m
```

After `make install`, configure your test HCL to use the local provider:

```hcl
terraform {
  required_providers {
    ngrok = {
      source  = "ngrok/ngrok"
      version = "0.0.1"
    }
  }
}
```

Then run `rm -f .terraform.lock.hcl && terraform init` to pick up the local build.
