BIN_DIR := bin
TOOLS_DIR := tools

include provider-version.mk

MANIFESTS_GEN := go run ./vendor/github.com/openshift/cluster-capi-operator/manifests-gen/
# manifests-gen resets container resource requests after running kustomize, so
# ocp-manifests applies capi-operator-manifests-overrides to its output.
KUSTOMIZE_BUILD := go run ./kustomize-build

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
	cd tools && $(KUSTOMIZE_BUILD) ../capi-operator-manifests-overrides ../capi-operator-manifests/default/manifests.yaml
