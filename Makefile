.PHONY: build clean

VERSION=0.1.0
REVISION := $(shell git rev-parse --short HEAD;)

BUILDPATH=$(CURDIR)
MAKEPATH=$(BUILDPATH)/make
PLATFORM_FILE="./cmd/artikube/main.go"

$(shell $(GO) version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f1)

build: build-mac

build-mac: 
	go build -v --ldflags="-w -X main.Version=$(VERSION) -X main.Revision=$(REVISION) " \
		-o bin/darwin/amd64/artikube $(PLATFORM_FILE) 
	sha256sum bin/darwin/amd64/artikube || shasum -a 256 bin/darwin/amd64/artikube

clean:
	@git status --ignored --short | grep '^!! ' | sed 's/!! //' | xargs rm -rf

validate-go-version: ## Validates the installed version of go against Mattermost's minimum requirement.
	@if [ $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f1) -gt 1 ]; then \
		exit 0 ;\
	elif [ $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f1) -lt 1 ]; then \
		echo 'Golang version is not supported, please update to at least 1.12';\
		exit 1; \
	elif [ $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f2) -lt 12 ] ; then \
		echo 'Golang version is not supported, please update to at least 1.12';\
		exit 1; \
	fi