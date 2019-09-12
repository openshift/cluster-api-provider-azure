module sigs.k8s.io/cluster-api-provider-azure

go 1.12

require (
	contrib.go.opencensus.io/exporter/ocagent v0.2.0 // indirect
	github.com/Azure/azure-sdk-for-go v26.6.0+incompatible
	github.com/Azure/go-autorest v11.5.2+incompatible
	github.com/census-instrumentation/opencensus-proto v0.1.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dimchansky/utfbom v1.1.0 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/go-log/log v0.0.0-20181211034820-a514cf01a3eb // indirect
	github.com/golang/groupcache v0.0.0-20190129154638-5b532d6fd5ef // indirect
	github.com/golang/mock v1.3.1
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/lithammer/dedent v1.1.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/openshift/cluster-api v0.0.0-20190829140302-072f7d777dc8
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_model v0.0.0-20190129233127-fd36f4220a90 // indirect
	github.com/prometheus/common v0.2.0 // indirect
	github.com/prometheus/procfs v0.0.0-20190227231451-bbced9601137 // indirect
	github.com/spf13/cobra v0.0.3
	go.opencensus.io v0.22.0 // indirect
	golang.org/x/crypto v0.0.0-20190611184440-5c40567a22f8
	golang.org/x/sys v0.0.0-20190624142023-c5567b49c5d0 // indirect
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	google.golang.org/api v0.9.0 // indirect
	google.golang.org/appengine v1.6.1 // indirect
	google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55 // indirect
	google.golang.org/grpc v1.21.1 // indirect
	k8s.io/api v0.0.0-20190831074750-7364b6bdad65
	k8s.io/apimachinery v0.0.0-20190831074630-461753078381
	k8s.io/apiserver v0.0.0-20190409021813-1ec86e4da56c // indirect
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/cloud-provider v0.0.0-20190314002645-c892ea32361a // indirect
	k8s.io/cluster-bootstrap v0.0.0-20190301180307-31878142fd8d
	k8s.io/klog v0.4.0
	k8s.io/kube-proxy v0.0.0-20190831080623-67697732d2b9 // indirect
	k8s.io/kubelet v0.0.0-20190831080713-e78366423f3e // indirect
	k8s.io/kubernetes v1.14.2
	k8s.io/utils v0.0.0-20190801114015-581e00157fb1
	sigs.k8s.io/controller-runtime v0.2.0-beta.2
	sigs.k8s.io/controller-tools v0.2.1
	sigs.k8s.io/yaml v1.1.0
)

replace git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999
