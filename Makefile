TEST?=$$(go list ./... | grep -v 'vendor')
SWEEP?=us
HOSTNAME=ngrok.com
NAMESPACE=ngrok
NAME=ngrok
BINARY=terraform-provider-${NAME}
VERSION=0.0.0

default: install

build:
		go build -o ${BINARY}

install: build
		mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
		mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test:
		echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
		TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

sweep:
		@echo "WARNING: This will destroy infrastructure. Use only in development accounts."
		go test $(TEST) -sweep=$(SWEEP) -v $(TESTARGS) -timeout 60m
