default: build

# Source of the ngrok OpenAPI spec. Override OPENAPI_REF to pull a specific
# branch, tag, or commit SHA (e.g. `make update-openapi OPENAPI_REF=v1.2.3`).
OPENAPI_REF ?= main
OPENAPI_URL ?= https://raw.githubusercontent.com/ngrok/ngrok-openapi/$(OPENAPI_REF)/ngrok.yaml

build:
	go build -o terraform-provider-ngrok .

install: build
	mkdir -p ~/.terraform.d/plugins/registry.terraform.io/ngrok/ngrok/0.0.1/$(shell go env GOOS)_$(shell go env GOARCH)
	cp terraform-provider-ngrok ~/.terraform.d/plugins/registry.terraform.io/ngrok/ngrok/0.0.1/$(shell go env GOOS)_$(shell go env GOARCH)/

test:
	go test ./... -v

testacc:
	TF_ACC=1 go test ./... -v -timeout 120m

lint:
	golangci-lint run ./...

generate:
	go generate ./...

# Codegen and docs tools are pinned in go.mod via `tool` directives and run
# with `go tool` — no separate `go install` step or PATH wrangling required.
codegen: codegen-spec codegen-framework

codegen-spec:
	go tool tfplugingen-openapi generate \
		--config codegen/generator_config.yml \
		--output codegen/provider_code_spec.json \
		codegen/openapi_spec.yaml

codegen-framework:
	go tool tfplugingen-framework generate all \
		--input codegen/provider_code_spec.json \
		--output internal

docs:
	go tool tfplugindocs generate --provider-name ngrok

docs-validate:
	go tool tfplugindocs validate --provider-name ngrok

update-openapi:
	# Pull the latest apic-generated spec from the ngrok-openapi repo
	curl -fsSL "$(OPENAPI_URL)" -o codegen/openapi_spec.yaml

dev: install
	@echo "Installed. Run 'rm -f test-manual/.terraform.lock.hcl && terraform -chdir=test-manual init' to test."

.PHONY: build install test testacc lint generate codegen codegen-spec codegen-framework docs docs-validate update-openapi dev
