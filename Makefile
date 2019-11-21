export CGO_ENABLED=0
export GO111MODULE=on

.PHONY: build

ONOS_CERTS_VERSION := latest
ONOS_CERTS_DEBUG_VERSION := debug
ONOS_BUILD_VERSION := stable

build: # @HELP build the Go binaries and run all validations (default)
build:
	CGO_ENABLED=1 go build -o build/_output/onos-certs ./cmd/onos-certs
	CGO_ENABLED=1 go build -gcflags "all=-N -l" -o build/_output/onos-certs-debug ./cmd/onos-certs

test: # @HELP run the unit tests and source code validation
test: build deps license_check linters
	go test github.com/onosproject/onos-certs/pkg/...
	go test github.com/onosproject/onos-certs/cmd/...

coverage: # @HELP generate unit test coverage data
coverage: build deps linters license_check
	./build/bin/coveralls-coverage

deps: # @HELP ensure that the required dependencies are in place
	go build -v ./...
	bash -c "diff -u <(echo -n) <(git diff go.mod)"
	bash -c "diff -u <(echo -n) <(git diff go.sum)"

linters: # @HELP examines Go source code and reports coding problems
	golangci-lint run

license_check: # @HELP examine and ensure license headers exist
	@if [ ! -d "../build-tools" ]; then cd .. && git clone https://github.com/onosproject/build-tools.git; fi
	./../build-tools/licensing/boilerplate.py -v --rootdir=${CURDIR}

onos-certs-base-docker: # @HELP build onos-topo base Docker image
	@go mod vendor
	docker build . -f build/base/Dockerfile \
		--build-arg ONOS_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t onosproject/onos-certs-base:${ONOS_CERTS_VERSION}
	@rm -rf vendor


onos-certs-debug-docker: onos-certs-base-docker # @HELP build onos-topo Docker debug image
	docker build . -f build/onos-certs-debug/Dockerfile \
		--build-arg ONOS_CERTS_BASE_VERSION=${ONOS_CERTS_VERSION} \
		-t onosproject/onos-certs:${ONOS_CERTS_DEBUG_VERSION}

onos-certs-docker: # @HELP build onos certs Docker image
onos-certs-docker: onos-certs-base-docker
	@go mod vendor
	docker build . -f build/onos-certs/Dockerfile \
		--build-arg ONOS_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t onosproject/onos-certs:${ONOS_CERTS_VERSION}
	@rm -rf vendor

images: # @HELP build all Docker images
images: build onos-certs-docker onos-certs-debug-docker

all: build images


clean: # @HELP remove all the build artifacts
	rm -rf ./build/_output ./vendor ./cmd/onos-certs/onos-certs

help:
	@grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST) \
    | sort \
    | awk ' \
        BEGIN {FS = ": *# *@HELP"}; \
        {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}; \
    '
