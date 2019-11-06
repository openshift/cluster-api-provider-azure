module sigs.k8s.io/cluster-api-provider-azure

go 1.13

require (
	cloud.google.com/go v0.45.1 // indirect
	contrib.go.opencensus.io/exporter/ocagent v0.2.0 // indirect
	github.com/Azure/azure-sdk-for-go v26.6.0+incompatible
	github.com/Azure/go-autorest v11.5.2+incompatible
	github.com/Azure/go-autorest/autorest v0.9.0 // indirect
	github.com/census-instrumentation/opencensus-proto v0.1.0 // indirect
	github.com/coreos/etcd v3.3.15+incompatible // indirect
	github.com/coreos/go-oidc v2.1.0+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/dimchansky/utfbom v1.1.0 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/go-log/log v0.1.1-0.20181211034820-a514cf01a3eb // indirect
	github.com/go-openapi/validate v0.19.2 // indirect
	github.com/gobuffalo/envy v1.6.15 // indirect
	github.com/golang/groupcache v0.0.0-20190129154638-5b532d6fd5ef // indirect
	github.com/golang/mock v1.3.1
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/gophercloud/gophercloud v0.1.0 // indirect
	github.com/gorilla/websocket v1.4.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.5.0 // indirect
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/jonboulle/clockwork v0.1.0 // indirect
	github.com/markbates/inflect v1.0.4 // indirect
	github.com/openshift/cluster-api v0.0.0-20191008120530-c4a461a19efb
	github.com/pkg/errors v0.8.1
	github.com/rogpeppe/go-internal v1.2.2 // indirect
	github.com/sirupsen/logrus v1.4.2 // indirect
	github.com/spf13/cobra v0.0.5
	golang.org/x/crypto v0.0.0-20190611184440-5c40567a22f8
	gomodules.xyz/jsonpatch v2.0.1+incompatible // indirect
	google.golang.org/grpc v1.23.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/square/go-jose.v2 v2.2.2 // indirect
	k8s.io/api v0.0.0-20190918195907-bd6ac527cfd2
	k8s.io/apimachinery v0.0.0-20190913080033-27d36303b655
	k8s.io/client-go v0.0.0-20190918200256-06eb1244587a
	k8s.io/code-generator v0.0.0-20190912054826-cd179ad6a269 // indirect
	k8s.io/klog v0.4.0
	k8s.io/utils v0.0.0-20190801114015-581e00157fb1
	sigs.k8s.io/controller-runtime v0.3.0
	sigs.k8s.io/controller-tools v0.1.10 // indirect
	sigs.k8s.io/structured-merge-diff v0.0.0-20190817042607-6149e4549fca // indirect
	sigs.k8s.io/yaml v1.1.0
)

replace gopkg.in/fsnotify.v1 v1.4.7 => github.com/fsnotify/fsnotify v1.4.7
