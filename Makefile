IMAGE ?= srijanlabs/spinnaker-servicebroker
TAG ?= 0.1.0

build:
	go build -o spinnaker-servicebroker -i github.com/srijanaravali/spinnaker-servicebroker/cmd/servicebroker

linux: ## Builds a Linux executable
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
	go build -o servicebroker-linux --ldflags="-s" github.com/srijanaravali/spinnaker-servicebroker/cmd/servicebroker

image: ## Builds docker image
	docker build . -t $(IMAGE):$(TAG)

clean: ## Cleans up build artifacts
	rm -f spinnaker-servicebroker
	rm -f servicebroker-linux
	rm -f packaging/helm/index.yaml
	rm -f packaging/helm/spinnaker-servicebroker-*.tgz
	rm -rf release/

help: ## Shows the help
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
        awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ''