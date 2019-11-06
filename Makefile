# Copyright 2018 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# If you update this file, please follow
# https://suva.sh/posts/well-documented-makefiles

.DEFAULT_GOAL:=help

## Image URL to use all building/pushing image targets
STABLE_DOCKER_REPO ?= quay.io/k8s
MANAGER_IMAGE_NAME ?= cluster-api-azure-controller
MANAGER_IMAGE_TAG ?= 0.1.0-alpha.3
MANAGER_IMAGE ?= $(STABLE_DOCKER_REPO)/$(MANAGER_IMAGE_NAME):$(MANAGER_IMAGE_TAG)
DEV_DOCKER_REPO ?= quay.io/k8s
DEV_MANAGER_IMAGE ?= $(DEV_DOCKER_REPO)/$(MANAGER_IMAGE_NAME):$(MANAGER_IMAGE_TAG)-dev

# determine the OS
HOSTOS := $(shell go env GOHOSTOS)
HOSTARCH := $(shell go env GOARCH)
BINARYPATHPATTERN :=${HOSTOS}_${HOSTARCH}_*

.PHONY: all
all: check-install test manager

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: dep-ensure
dep-ensure: check-install ## Ensure dependencies are up to date
	@echo Checking status of dependencies
	@dep status 2>&1 > /dev/null || make dep-install
	@echo Finished verifying dependencies

.PHONY: dep-install
dep-install: ## Force install go dependencies
	dep ensure

.PHONY: check-install
check-install: ## Checks that you've installed this repository correctly
	@./hack/check-install.sh

.PHONY: manager
manager: generate ## Build manager binary.
	bazel build //cmd/manager
	install bazel-bin/cmd/manager/${BINARYPATHPATTERN}/manager $(shell go env GOPATH)/bin/azure-manager

.PHONY: test verify
test: generate verify ## Run tests
	bazel test --nosandbox_debug //pkg/... //cmd/...

.PHONY: docker-build
docker-build: generate ## Build the production docker image
	docker build . -t $(MANAGER_IMAGE)

.PHONY: crds
crds:
	bazel build //config
	cp -R bazel-genfiles/config/crds/* config/crds/
	cp -R bazel-genfiles/config/rbac/* config/rbac/

# TODO(vincepri): This should move to rebuild Bazel binaries once every
# make target uses Bazel bins to run operations.
.PHONY: binaries-dev
binaries-dev: ## Builds and installs the binaries on the local GOPATH
	go get -v ./...
	go install -v ./...

## Old make targets
# TODO: Migrate old make targets

vendor:
	dep version || go get -u github.com/golang/dep/cmd/dep
	dep ensure -v
vendor-update:
	dep version || go get -u github.com/golang/dep/cmd/dep
	dep ensure -v -update
vendor-validate:
	dep check

machine-unit-tests:
	go test ./pkg/cloud/azure/actuators/machine -coverprofile machine-actuator-cover.out

cluster-unit-tests:
	go test ./pkg/cloud/azure/actuators/cluster -coverprofile cluster-actuator-cover.out

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet
	go run ./cmd/manager/main.go

# Install CRDs into a cluster
install:
	kubectl apply -f config/crds

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy:
	kubectl apply -f config/crds
	kustomize build config/default | kubectl apply -f -

# Run go fmt against code
fmt:
	go fmt ./pkg/... ./cmd/...

.PHONY: goimports
goimports: ## Go fmt your code
	hack/goimports.sh .

# Run go vet against code
vet:
	go vet -composites=false ./pkg/... ./cmd/... 

check: bootstrap vendor-validate lint

# TODO: IMPORTANT
.PHONY: test-e2e
test-e2e:
	hack/e2e.sh
