/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/openshift/cluster-api/cmd/clusterctl/cmd"
	"github.com/openshift/cluster-api/pkg/apis/cluster/common"
	"sigs.k8s.io/cluster-api-provider-azure/cmd/versioninfo"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators/cluster"
)

func registerCustomCommands() {
	cmd.RootCmd.AddCommand(versioninfo.VersionCmd())
}

func main() {
	clusterActuator := cluster.NewActuator(cluster.ActuatorParams{})
	common.RegisterClusterProvisioner("azure", clusterActuator)
	registerCustomCommands()
	cmd.Execute()
}
