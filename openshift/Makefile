include ../Makefile

RELEASE_DIR := openshift
GO_INSTALL = ../scripts/go_install.sh

.PHONY: ocp-manifests
ocp-manifests:  $(KUSTOMIZE)  ## Builds openshift specific manifests
	$(KUSTOMIZE) build ../config/default > infrastructure-components.yaml
