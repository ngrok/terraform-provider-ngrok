# Design: OpenAPI → Terraform Code Generation

## Status: Proposal

## Problem

The current Terraform provider v2 has 21 resources and 24 data sources, all hand-written. Each resource file follows a nearly identical pattern (~250 lines) with boilerplate for schema definitions, models, flatten/expand functions, and CRUD methods. When the ngrok API changes, every affected file must be updated manually.

### Current v1 Pipeline (apic → everything)

The existing v1 provider at [ngrok/terraform-provider-ngrok](https://github.com/ngrok/terraform-provider-ngrok) is **almost entirely code-generated** by the `apic` tool from proto definitions:

```
proto/api/*.proto → apic tool → {
  1. restapi/          (REST client, API structs, methods — 26k lines)
  2. ngrok/provider.go (resource map wiring)
  3. ngrok/resource_*.go (schema + CRUD for each resource)
  4. ngrok/flatten_expand.go (17k lines of flatten/expand for every API struct)
  5. transform/convert.go (Ref→string helper)
}
```

The proto files contain rich Terraform-specific annotations that drive this generation:
- `terraform_method: Create/Read/Update/Delete` on RPCs
- `TerraformField.sensitive`, `TerraformField.computed` on fields
- `TerraformOutputMapping.skip_schema` to exclude computed-only fields from the schema
- `TerraformOutputMapping.diff_suppress_func` for custom diff suppression (e.g., `DiffSuppressWhitespace` on `ca_pem`)
- `TerraformOutputMapping.replace` and `flatten_func` for field remapping
- `apic.field.required` and `apic.field.nullable` to control Required/Optional/Computed

The generated resource files use **SDKv2** (`hashicorp/terraform-plugin-sdk/v2`), a hand-rolled `restapi` package (not the official `ngrok-api-go` client), and include everything: schemas, CRUD methods, expand/flatten — all in one generated file per resource.

**The OpenAPI spec (`ngrok.yaml`) is a separate side output** of the same `apic` tool — it is generated alongside the TF provider, not consumed by it. It lives at [ngrok/ngrok-openapi](https://github.com/ngrok/ngrok-openapi).

### What the v1 generator produces vs. what we need

| Aspect | v1 (apic-generated, SDKv2) | v2 (hand-written, Plugin Framework) |
|---|---|---|
| Framework | `terraform-plugin-sdk/v2` | `terraform-plugin-framework` |
| API client | Hand-rolled `restapi` package | Official `ngrok-api-go/v9` |
| Schema attrs | `schema.TypeString`, `ForceNew`, `DiffSuppressFunc` | `schema.StringAttribute`, `RequiresReplace()`, validators |
| Data model | `d.Set("field", value)` | Typed struct with `tfsdk` tags |
| Null handling | No null/unknown distinction | `types.String` with null/unknown semantics |
| Data sources | None | 24 data sources |
| Plan modifiers | N/A (SDKv2 doesn't have them) | `UseStateForUnknown()`, `ModifyPlan` |

The v1 generator cannot be reused because it targets the wrong framework, wrong client library, and lacks Plugin Framework concepts (plan modifiers, typed models, null semantics).

## Proposed Approach

Decouple TF generation from apic by consuming the OpenAPI spec that apic already produces. Use HashiCorp's official [OpenAPI Provider Spec Generator](https://github.com/hashicorp/terraform-plugin-codegen-openapi) and [Framework Code Generator](https://github.com/hashicorp/terraform-plugin-codegen-framework) to generate schema definitions and models from it.

### Full pipeline (near-term)

Proto and apic remain upstream — they produce the OpenAPI spec as they do today. We add a new downstream pipeline that consumes it:

```
proto/api/*.proto → apic → ngrok.yaml (OpenAPI 3.0) → tfplugingen-openapi → provider_code_spec.json → tfplugingen-framework → *_gen.go
|___________________________________|                |___________________________________________________________________________|
        Already exists today                              New: what we're building
```

This replaces the old flow where apic generated the TF provider directly from proto. The apic TF generator (`go/cmd/apic/gen/terraform_provider/`) is retired.

### Future state (long-term)

When proto and apic are eventually retired and OpenAPI becomes the system of record for the API, the left side of the pipeline changes but the TF codegen side stays exactly the same:

```
ngrok.yaml (OpenAPI 3.0, hand-authored) → tfplugingen-openapi → provider_code_spec.json → tfplugingen-framework → *_gen.go
```

The TF pipeline is intentionally agnostic to what produces the OpenAPI spec.

### What Gets Generated

The HashiCorp tools generate **schemas and models only**:
- `Schema()` function with all attributes, descriptions, computed/optional/required flags
- Go struct models with `tfsdk` tags
- Validators (from `enum`, `minLength`, `maxItems`, etc. in the OpenAPI spec)
- Default values (from `default` in the OpenAPI spec)

### What Remains Hand-Written

- **CRUD methods** (Create, Read, Update, Delete) — the actual API calls
- **Provider configuration** (API key, base URL)
- **Plan modifiers** (UseStateForUnknown, RequiresReplace) — not expressible in OpenAPI
- **Flatten/expand logic** for complex types (Ref → string ID, nested objects)
- **Sensitive field preservation** (write-only fields like `api_key`, `token`, `value`)
- **ModifyPlan hooks** (marking `updated_at` as unknown on changes)
- **Import state** logic

### v1 vs Proposed: What's Generated

| Component | v1 (apic from proto) | Proposed (OpenAPI codegen) |
|---|---|---|
| REST client / API structs | ✅ Generated (`restapi/`) | ❌ Use `ngrok-api-go/v9` |
| Resource schemas | ✅ Generated (incl. ForceNew, Sensitive, DiffSuppress) | ✅ Generated (descriptions, types, required/computed only) |
| Data models | ✅ Generated (implicit via `d.Set`) | ✅ Generated (typed structs with `tfsdk` tags) |
| Flatten/expand | ✅ Generated (17k lines) | ❌ Hand-written (uses `ngrok-api-go` types directly) |
| CRUD methods | ✅ Generated | ❌ Hand-written |
| Plan modifiers | N/A (SDKv2) | ❌ Hand-written (override layer) |
| Validators | ❌ (only DiffSuppress) | ✅ Generated (from `enum`, `minLength`, etc.) |
| Provider wiring | ✅ Generated | ❌ Hand-written |

**Key insight**: The v1 apic tool generates ~95% of the provider because the proto annotations encode Terraform-specific concerns (ForceNew, Sensitive, DiffSuppress). The OpenAPI spec does not carry this information, so the proposed approach generates less (~schemas and models) but what it generates is higher quality (typed models, validators) and targets the correct framework.

## Architecture

### Directory Structure

```
terraform-provider-ngrok/
├── codegen/
│   ├── generator_config.yml    # Maps OpenAPI paths → TF resources/data sources
│   ├── provider_code_spec.json # Generated intermediate spec (committed for review)
│   └── openapi_spec.yaml       # Copied/symlinked from ngrok-openapi repo
├── internal/
│   ├── provider/
│   │   ├── provider.go                         # Hand-written (unchanged)
│   │   ├── helpers.go                          # Hand-written (unchanged)
│   │   ├── resource_certificate_authority.go   # Hand-written CRUD + overrides
│   │   ├── resource_cloud_endpoint.go          # Hand-written CRUD + overrides
│   │   └── ...
│   ├── resource_certificate_authority/
│   │   └── certificate_authority_resource_gen.go  # Generated schema + model
│   ├── resource_cloud_endpoint/
│   │   └── cloud_endpoint_resource_gen.go         # Generated schema + model
│   ├── datasource_certificate_authority/
│   │   └── certificate_authority_data_source_gen.go
│   └── ...
├── Makefile                    # `make generate` target
└── ...
```

### Generator Config

The `generator_config.yml` maps OpenAPI operations to Terraform resources/data sources:

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

  api_key:
    create:
      path: /api_keys
      method: POST
    read:
      path: /api_keys/{id}
      method: GET
    update:
      path: /api_keys/{id}
      method: PATCH
    delete:
      path: /api_keys/{id}
      method: DELETE

  # ... all 21 resources mapped similarly

data_sources:
  certificate_authority:
    read:
      path: /certificate_authorities/{id}
      method: GET

  cloud_endpoint:
    read:
      path: /endpoints/{id}
      method: GET

  # ... all 24 data sources
```

### Resource File Pattern (Post-Codegen)

Each resource file becomes simpler — it imports the generated schema and uses it:

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

    // Apply TF-specific overrides not expressible in OpenAPI
    // (plan modifiers, ForceNew, sensitive, etc.)
    applyPlanModifiers(&s)

    resp.Schema = s
}
```

The CRUD methods, flatten/expand, and provider wiring remain hand-written as today.

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
| **Field exclusion** | Endpoint has ephemeral-only fields (`tunnel`, `edge`, `hostport`, `proto`) | Schema override to remove |

### Ref Type Challenge

The ngrok OpenAPI spec uses `$ref: "#/components/schemas/Ref"` for references, which has a single `id` property. In our current TF provider, we flatten these to just the ID string (e.g., `domain_id` instead of a nested `domain.id` object). Options:

1. **Override in generator config** using `aliases` to map `domain.id` → `domain_id`
2. **Post-process the generated schema** to replace `SingleNestedAttribute` with `StringAttribute`
3. **Accept nested objects** and change TF schema to `domain { id = "..." }` (breaking change)

**Recommendation**: Option 2 — post-process in a schema override function per resource.

## Build Pipeline

```makefile
# Makefile targets

.PHONY: generate
generate: generate-spec generate-code

.PHONY: generate-spec
generate-spec:
	tfplugingen-openapi generate \
		--config codegen/generator_config.yml \
		--output codegen/provider_code_spec.json \
		codegen/openapi_spec.yaml

.PHONY: generate-code
generate-code:
	tfplugingen-framework generate all \
		--input codegen/provider_code_spec.json \
		--output internal

.PHONY: update-openapi
update-openapi:
	# Pull the latest apic-generated spec from the ngrok-openapi repo
	cp ../ngrok-openapi/ngrok.yaml codegen/openapi_spec.yaml
```

## Migration Strategy

### Phase 1: Proof of Concept (1-2 resources)
1. Set up the codegen pipeline with `certificate_authority` and `api_key`
2. Generate schemas, verify they match current hand-written schemas
3. Write the schema override layer for plan modifiers
4. Validate CRUD still works with generated schemas

### Phase 2: Migrate Simple Resources
- Resources with only string fields, no nested objects, no sensitive fields
- `ip_policy`, `ssh_certificate_authority`, `reserved_addr`, `vault`, etc.

### Phase 3: Migrate Complex Resources
- Resources with `Ref` types: `cloud_endpoint`, `reserved_domain`, `ip_restriction`
- Resources with sensitive/write-only fields: `api_key`, `credential`, `secret`
- Resources with nested objects: `event_destination`, `reserved_domain`

### Phase 4: Migrate Data Sources
- Simpler than resources (read-only, no plan modifiers)
- Can largely use generated schemas as-is

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

### Comparison: Effort Reduction Per Resource

| Component | Lines Today | Lines After Codegen | Savings |
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

2. **Schema override pattern**: Inline in each resource's `Schema()` method. Call the generated schema function, then modify the returned schema to add plan modifiers, mark fields sensitive, swap Ref objects for flat ID strings, etc. Keeps overrides co-located with the resource they belong to.

3. **Ref handling**: Keep flattening to `_id` strings (e.g., `certificate_id`, not `certificate { id = "..." }`). This matches the v1 provider, is simpler for users, and is the established pattern. Post-process the generated schema to replace `SingleNestedAttribute` with `StringAttribute` for Ref-typed fields.

4. **OpenAPI spec source**: Copy into this repo via `make update-openapi`. Submodules add friction (forgot to init, wrong commit checked out). An explicit copy is predictable and the file is small (~11k lines).

5. **CI integration**: Yes. Run `make generate` in CI and fail if there's a diff. This catches cases where someone updates the generator config or OpenAPI spec but forgets to regenerate.

6. **Kubernetes operator resource**: Keep fully hand-written for now. It has deeply nested schemas (`ingress`, `bindings` with sub-objects) that may hit codegen limitations. Attempt migration in Phase 3 once the pattern is proven on simpler resources; if the codegen can't handle it, it stays hand-written.
