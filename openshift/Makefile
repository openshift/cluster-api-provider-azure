# Copyright The Kubernetes Authors.
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

BIN_DIR := bin
TOOLS_DIR := tools

include provider-version.mk

MANIFESTS_GEN := go run ./vendor/github.com/openshift/cluster-capi-operator/manifests-gen/

.PHONY: verify
verify: verify-ocp-manifests

.PHONY: verify-%
verify-%:
	$(MAKE) $*
	./verify-diff.sh

.PHONY: update-manifests-gen
update-manifests-gen:
	cd tools && go get github.com/openshift/cluster-capi-operator/manifests-gen@main github.com/openshift/cluster-capi-operator@main && go mod tidy && go mod vendor

.PHONY: ocp-manifests
ocp-manifests: ## Builds openshift specific manifests
	cd tools && $(MANIFESTS_GEN) \
		--manifests-path "../capi-operator-manifests" \
		--kustomize-dir "../../openshift" \
		--name cluster-api-provider-azure \
		--install-order 20 \
		--attribute type=infrastructure \
		--attribute version="${PROVIDER_VERSION}" \
		--self-image-ref registry.ci.openshift.org/openshift:azure-cluster-api-controllers \
		--platform Azure \
		--protect-cluster-resource azurecluster
