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

GO111MODULE = on
export GO111MODULE
GOFLAGS += -mod=vendor
export GOFLAGS
GOPROXY ?=
export GOPROXY

.DEFAULT_GOAL:=help

FASTBUILD ?= n ## Set FASTBUILD=y (case-sensitive) to skip some slow tasks

## Image URL to use all building/pushing image targets
STABLE_DOCKER_REPO ?= quay.io/k8s
MANAGER_IMAGE_NAME ?= cluster-api-azure-controller
MANAGER_IMAGE_TAG ?= 0.1.0-alpha.3
MANAGER_IMAGE ?= $(STABLE_DOCKER_REPO)/$(MANAGER_IMAGE_NAME):$(MANAGER_IMAGE_TAG)
DEV_DOCKER_REPO ?= quay.io/k8s
DEV_MANAGER_IMAGE ?= $(DEV_DOCKER_REPO)/$(MANAGER_IMAGE_NAME):$(MANAGER_IMAGE_TAG)-dev

DEPCACHEAGE ?= 24h # Enables caching for Dep
BAZEL_ARGS ?=

# Bazel variables
DEP ?= bazel run dep

# determine the OS
HOSTOS := $(shell go env GOHOSTOS)
HOSTARCH := $(shell go env GOARCH)
BINARYPATHPATTERN :=${HOSTOS}_${HOSTARCH}_*

.PHONY: all
all: check-install test manager clusterctl #clusterazureadm

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

verify:
	./hack/verify_boilerplate.py

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor
	go mod verify

machine-unit-tests:
	go test ./pkg/cloud/azure/actuators/machine -coverprofile machine-actuator-cover.out

cluster-unit-tests:
	go test ./pkg/cloud/azure/actuators/cluster -coverprofile cluster-actuator-cover.out

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet
	go run ./cmd/manager/main.go

# Install CRDs into a cluster
install: manifests
	kubectl apply -f config/crds

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy: manifests
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

verify-boilerplate:
	./hack/verify-boilerplate.sh

check: verify-boilerplate bootstrap vendor-validate lint

.PHONY: test-e2e
test-e2e:
	hack/e2e.sh
