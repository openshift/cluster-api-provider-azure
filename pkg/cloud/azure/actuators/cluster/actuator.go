/*
Copyright 2018 The Kubernetes Authors.

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

package cluster

import (
	clusterv1 "github.com/openshift/cluster-api/pkg/apis/cluster/v1alpha1"
	client "github.com/openshift/cluster-api/pkg/client/clientset_generated/clientset/typed/cluster/v1alpha1"
	controllerError "github.com/openshift/cluster-api/pkg/controller/error"
	"github.com/pkg/errors"
	"k8s.io/klog"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/deployer"
)

//+kubebuilder:rbac:groups=azureprovider.k8s.io,resources=azureclusterproviderconfigs;azureclusterproviderstatuses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cluster.k8s.io,resources=clusters;clusters/status,verbs=get;list;watch;create;update;patch;delete

// Actuator is responsible for performing cluster reconciliation
type Actuator struct {
	*deployer.Deployer

	client client.ClusterV1alpha1Interface
}

// ActuatorParams holds parameter information for Actuator
type ActuatorParams struct {
	Client client.ClusterV1alpha1Interface
}

// NewActuator creates a new Actuator
func NewActuator(params ActuatorParams) *Actuator {
	return &Actuator{
		Deployer: deployer.New(deployer.Params{ScopeGetter: actuators.DefaultScopeGetter}),
		client:   params.Client,
	}
}

// Reconcile reconciles a cluster and is invoked by the Cluster Controller
func (a *Actuator) Reconcile(cluster *clusterv1.Cluster) error {
	klog.Infof("Reconciling cluster %v", cluster.Name)

	scope, err := actuators.NewScope(actuators.ScopeParams{Cluster: cluster, Client: a.client})
	if err != nil {
		return errors.Errorf("failed to create scope: %+v", err)
	}

	defer scope.Close()

	err = NewReconciler(scope).Reconcile()
	if err != nil {
		return errors.Wrap(err, "failed to reconcile cluster services")
	}

	return nil
}

// Delete deletes a cluster and is invoked by the Cluster Controller.
func (a *Actuator) Delete(cluster *clusterv1.Cluster) error {
	klog.Infof("Reconciling cluster %v", cluster.Name)

	scope, err := actuators.NewScope(actuators.ScopeParams{Cluster: cluster, Client: a.client})
	if err != nil {
		return errors.Errorf("failed to create scope: %+v", err)
	}

	defer scope.Close()

	if err := NewReconciler(scope).Delete(); err != nil {
		klog.Errorf("Error deleting resource group: %v.", err)
		return &controllerError.RequeueAfterError{
			RequeueAfter: 5 * 1000 * 1000 * 1000,
		}
	}

	return nil
}
