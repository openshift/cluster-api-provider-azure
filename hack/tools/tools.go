// This package imports things required by build scripts, to force `go mod` to see them as dependencies
package tools

import (
	_ "github.com/golang/mock/mockgen"                  //nolint
	_ "k8s.io/code-generator"                           //nolint
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen" //nolint
)
