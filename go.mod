module sigs.k8s.io/cluster-api-provider-azure

go 1.16

require (
	github.com/Azure/azure-sdk-for-go v48.0.0+incompatible
	github.com/Azure/go-autorest/autorest v0.11.18
	github.com/Azure/go-autorest/autorest/adal v0.9.13
	github.com/Azure/go-autorest/autorest/to v0.3.0
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/go-logr/logr v0.4.0
	github.com/golang/mock v1.5.0
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.14.0
	github.com/openshift/api v0.0.0-20210816181336-8ff39b776da3
	github.com/openshift/machine-api-operator v0.2.1-0.20210820103535-d50698c302f5
	github.com/spf13/cobra v1.2.1
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83

	// kube 1.22
	k8s.io/api v0.22.0
	k8s.io/apimachinery v0.22.0
	k8s.io/client-go v0.22.0
	k8s.io/code-generator v0.22.0
	k8s.io/klog/v2 v2.9.0
	k8s.io/utils v0.0.0-20210819203725-bdf08cb9a70a
	sigs.k8s.io/controller-runtime v0.9.6
	sigs.k8s.io/controller-tools v0.6.2
	sigs.k8s.io/yaml v1.2.0
)

replace (
	sigs.k8s.io/cluster-api-provider-aws => github.com/openshift/cluster-api-provider-aws v0.2.1-0.20210622023641-c69a3acaee27
	sigs.k8s.io/cluster-api-provider-azure => github.com/openshift/cluster-api-provider-azure v0.1.0-alpha.3.0.20210816141152-a7c40345b994
)
