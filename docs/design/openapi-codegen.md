# Design: OpenAPI → Terraform Code Generation

## Problem

Each Terraform resource and data source follows a nearly identical pattern with boilerplate for schema definitions, models, flatten/expand functions, and CRUD methods. When the ngrok API changes, every affected file must be updated manually.

Code generation eliminates this duplication by deriving schema definitions and model structs directly from the ngrok OpenAPI spec. Schemas stay in sync with the API automatically, descriptions are always accurate, and validators (enum, minLength, etc.) come for free.

## Pipeline

The provider uses HashiCorp's official [OpenAPI Provider Spec Generator](https://github.com/hashicorp/terraform-plugin-codegen-openapi) and [Framework Code Generator](https://github.com/hashicorp/terraform-plugin-codegen-framework) to generate schema definitions and models from the ngrok OpenAPI spec:

```
ngrok.yaml (OpenAPI 3.0) → tfplugingen-openapi → provider_code_spec.json → tfplugingen-framework → *_gen.go
```

1. **`ngrok.yaml`** — The ngrok OpenAPI 3.0 spec, pulled from [ngrok/ngrok-openapi](https://github.com/ngrok/ngrok-openapi) and copied into `codegen/openapi_spec.yaml`.
2. **`tfplugingen-openapi`** — Reads the OpenAPI spec and the generator config (`codegen/generator_config.yml`), produces an intermediate `provider_code_spec.json`.
3. **`provider_code_spec.json`** — A framework-agnostic JSON description of every resource and data source schema. Committed to the repo so changes are reviewable in PRs.
4. **`tfplugingen-framework`** — Reads the provider code spec, writes `*_gen.go` files into per-resource/data-source packages under `internal/`.

The pipeline is agnostic to what produces the OpenAPI spec — it only consumes the spec file.

## What Gets Generated

The HashiCorp tools generate **schemas and models only**:
- `Schema()` function with all attributes, descriptions, computed/optional/required flags
- Go struct models with `tfsdk` tags
- Validators (from `enum`, `minLength`, `maxItems`, etc. in the OpenAPI spec)
- Default values (from `default` in the OpenAPI spec)

## What Remains Hand-Written

- **CRUD methods** (Create, Read, Update, Delete) — the actual API calls
- **Provider configuration** (API key, base URL)
- **Plan modifiers** (UseStateForUnknown, RequiresReplace) — not expressible in OpenAPI
- **Flatten/expand logic** for complex types (Ref → string ID, nested objects)
- **Sensitive field preservation** (write-only fields like `api_key`, `token`, `value`)
- **ModifyPlan hooks** (marking `updated_at` as unknown on changes)
- **Import state** logic

## Architecture

### Directory Structure

```
terraform-provider-ngrok/
├── codegen/
│   ├── generator_config.yml    # Maps OpenAPI paths → TF resources/data sources
│   ├── provider_code_spec.json # Generated intermediate spec (committed for review)
│   └── openapi_spec.yaml       # Copied from ngrok-openapi repo
├── internal/
│   ├── provider/
│   │   ├── provider.go                         # Hand-written
│   │   ├── helpers.go                          # Hand-written
│   │   ├── schema_overrides.go                 # Hand-written: plan modifier helpers
│   │   ├── validators.go                       # Hand-written: custom validators
│   │   ├── resource_certificate_authority.go   # Hand-written CRUD + overrides
│   │   ├── resource_cloud_endpoint.go          # Hand-written CRUD + overrides
│   │   ├── datasource_certificate_authority.go # Hand-written read logic
│   │   └── ...
│   ├── resource_certificate_authority/
│   │   └── certificate_authority_resource_gen.go  # Generated schema + model
│   ├── resource_cloud_endpoint/
│   │   └── cloud_endpoint_resource_gen.go         # Generated schema + model
│   ├── datasource_certificate_authority/
│   │   └── certificate_authority_data_source_gen.go
│   └── ...
├── Makefile                    # `make codegen` target
└── ...
```

Generated code lives in per-resource/data-source packages (`internal/resource_*/`, `internal/datasource_*/`). Hand-written resource logic lives in `internal/provider/`.

### Generator Config

The `generator_config.yml` maps OpenAPI operations to Terraform resources and data sources:

```yaml
provider:
  name: ngrok

resources:
  certificate_authority:
    create:
      path: /certificate_authorities
      method: POST
    read:
      path: /certificate_authorities/{id}
      method: GET
    update:
      path: /certificate_authorities/{id}
      method: PATCH
    delete:
      path: /certificate_authorities/{id}
      method: DELETE

  cloud_endpoint:
    create:
      path: /endpoints
      method: POST
    read:
      path: /endpoints/{id}
      method: GET
    update:
      path: /endpoints/{id}
      method: PATCH
    delete:
      path: /endpoints/{id}
      method: DELETE

  # ... all resources mapped similarly

data_sources:
  certificate_authority:
    read:
      path: /certificate_authorities/{id}
      method: GET

  # ... all data sources
```

Fields can be excluded from generation using the `schema.ignores` key:

```yaml
  event_destination:
    create:
      path: /event_destinations
      method: POST
    # ...
    schema:
      ignores:
        - target
```

### Resource File Pattern (Post-Codegen)

Each resource file imports the generated schema and applies overrides:

```go
package provider

import (
    "context"
    "terraform-provider-ngrok/internal/resource_certificate_authority"
    // ...
)

func (r *certificateAuthorityResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    // Start from generated schema
    s := resource_certificate_authority.CertificateAuthorityResourceSchema(ctx)
    attrs := s.Attributes

    // Apply TF-specific overrides not expressible in OpenAPI
    addStringPlanModifiers(attrs, "id", useStateForUnknownString())
    addStringPlanModifiers(attrs, "uri", useStateForUnknownString())
    addStringPlanModifiers(attrs, "created_at", useStateForUnknownString())
    addStringPlanModifiers(attrs, "ca_pem", requiresReplaceString())

    resp.Schema = s
}
```

The override helpers (`addStringPlanModifiers`, `addListPlanModifiers`, `markSensitive`, etc.) live in `internal/provider/schema_overrides.go` and provide a consistent way to layer plan modifiers, sensitive flags, defaults, and deprecation messages onto the generated schema.

The CRUD methods, flatten/expand, and provider wiring remain hand-written in the same resource file.

## Mapping: OpenAPI Spec → Terraform Schema

### What maps cleanly

| OpenAPI | Terraform Schema |
|---|---|
| `type: string` | `schema.StringAttribute` |
| `type: integer` | `schema.Int64Attribute` |
| `type: boolean` | `schema.BoolAttribute` |
| `type: array, items: {type: string}` | `schema.ListAttribute{ElementType: types.StringType}` |
| `type: object` (with properties) | `schema.SingleNestedAttribute` |
| `$ref: "#/components/schemas/Ref"` | `schema.SingleNestedAttribute` (with `id` field) |
| `required` in create requestBody | `Required: true` |
| Field only in response (not in create) | `Computed: true` |
| Field in both create and response, not required | `Optional: true, Computed: true` |
| `description` | `Description` |
| `enum` | `Validators` |

### What needs manual overrides

| Concern | Example | Override Approach |
|---|---|---|
| **RequiresReplace** | `ca_pem`, `public_key`, `domain` | Post-process schema to add plan modifier |
| **UseStateForUnknown** | `id`, `uri`, `created_at` | Post-process computed attrs |
| **Sensitive** | `api_key.token`, `secret.value` | Override in resource file |
| **Write-only preservation** | API returns `""` for secrets | Hand-written in flatten |
| **Ref flattening** | `domain: {$ref: Ref}` → `domain_id: string` | Schema override + custom flatten |
| **ModifyPlan** | `updated_at` unknown on changes | Hand-written as today |
| **Field aliasing** | OpenAPI `id` path param vs body `id` | Generator config `aliases` |
| **Field exclusion** | Endpoint has ephemeral-only fields | `schema.ignores` in generator config |

### Ref Type Challenge

The ngrok OpenAPI spec uses `$ref: "#/components/schemas/Ref"` for references, which has a single `id` property. In the TF provider, these are flattened to just the ID string (e.g., `domain_id` instead of a nested `domain.id` object). Options:

1. **Override in generator config** using `aliases` to map `domain.id` → `domain_id`
2. **Post-process the generated schema** to replace `SingleNestedAttribute` with `StringAttribute`
3. **Accept nested objects** and change TF schema to `domain { id = "..." }` (breaking change)

**Decision**: Option 2 — post-process in a schema override function per resource. This keeps the flat `_id` pattern that users expect.

## Build Pipeline

Both codegen tools are pinned in `go.mod` via `tool` directives and invoked with `go tool` — no separate install step required.

```makefile
# Makefile targets

# Pull the latest OpenAPI spec from the ngrok-openapi repo.
# Override OPENAPI_REF to target a specific branch/tag/SHA.
OPENAPI_REF ?= main
OPENAPI_URL ?= https://raw.githubusercontent.com/ngrok/ngrok-openapi/$(OPENAPI_REF)/ngrok.yaml

.PHONY: codegen
codegen: codegen-spec codegen-framework

.PHONY: codegen-spec
codegen-spec:
	go tool tfplugingen-openapi generate \
		--config codegen/generator_config.yml \
		--output codegen/provider_code_spec.json \
		codegen/openapi_spec.yaml

.PHONY: codegen-framework
codegen-framework:
	go tool tfplugingen-framework generate all \
		--input codegen/provider_code_spec.json \
		--output internal

.PHONY: update-openapi
update-openapi:
	curl -fsSL "$(OPENAPI_URL)" -o codegen/openapi_spec.yaml
```

## Trade-offs

### Benefits
- **Schema consistency**: Generated schemas always match the OpenAPI spec
- **Reduced boilerplate**: ~40% of each resource file is schema definition
- **Automatic updates**: When ngrok API adds fields, regenerate to pick them up
- **Validation for free**: `enum`, `minLength`, etc. from OpenAPI become TF validators
- **Descriptions stay in sync**: API docs = TF docs

### Costs
- **Two-layer schema**: Generated base + manual overrides adds indirection
- **Override complexity**: Plan modifiers, sensitive flags, Ref flattening all need manual post-processing
- **Tech preview**: HashiCorp's codegen tools are in tech preview, API may change
- **Partial generation**: Only schemas/models are generated; CRUD logic stays hand-written (~60% of code)
- **Ref type mismatch**: OpenAPI `Ref` objects don't map cleanly to our flat ID pattern

### Effort Reduction Per Resource

| Component | Lines (Hand-Written) | Lines (With Codegen) | Savings |
|---|---|---|---|
| Schema definition | ~80 | ~5 (import + overrides) | ~94% |
| Model struct | ~15 | 0 (generated) | 100% |
| CRUD methods | ~100 | ~100 (unchanged) | 0% |
| Flatten/expand | ~30 | ~30 (unchanged) | 0% |
| Provider wiring | ~5 | ~5 (unchanged) | 0% |
| **Total per resource** | **~230** | **~140** | **~39%** |

For data sources (read-only), savings are higher (~60%) since there's no CRUD boilerplate.

## Decisions

1. **Commit `provider_code_spec.json`**: Yes. It makes PRs reviewable — you can see "this API change added these fields" in readable JSON before looking at generated Go. Helps debug when generated output looks wrong.

2. **Schema override pattern**: Inline in each resource's `Schema()` method. Call the generated schema function, then modify the returned schema to add plan modifiers, mark fields sensitive, swap Ref objects for flat ID strings, etc. Keeps overrides co-located with the resource they belong to. Shared helpers live in `schema_overrides.go`.

3. **Ref handling**: Keep flattening to `_id` strings (e.g., `certificate_id`, not `certificate { id = "..." }`). This is simpler for users and is the established pattern. Post-process the generated schema to replace `SingleNestedAttribute` with `StringAttribute` for Ref-typed fields.

4. **OpenAPI spec source**: Copy into this repo via `make update-openapi`. Submodules add friction (forgot to init, wrong commit checked out). An explicit copy is predictable and the file is small (~11k lines).

5. **CI integration**: Run `make codegen` in CI and fail if there's a diff. This catches cases where someone updates the generator config or OpenAPI spec but forgets to regenerate.
