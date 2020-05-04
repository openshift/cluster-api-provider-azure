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

package applicationsecuritygroups

import (
	"context"
	"errors"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-12-01/network"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/klog"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
)

// Spec specification for network application security groups
type Spec struct {
	Name string
}

// Get provides information about a route table.
func (s *Service) Get(ctx context.Context, spec azure.Spec) (interface{}, error) {
	asgSpec, ok := spec.(*Spec)
	if !ok {
		return network.ApplicationSecurityGroup{}, errors.New("invalid application security groups specification")
	}
	applicationSecurityGroup, err := s.Client.Get(ctx, s.Scope.MachineConfig.ResourceGroup, asgSpec.Name)
	if err != nil && azure.ResourceNotFound(err) {
		return nil, fmt.Errorf("application security group %s not found: %w", asgSpec.Name, err)
	} else if err != nil {
		return applicationSecurityGroup, err
	}
	return applicationSecurityGroup, nil
}

// CreateOrUpdate creates or updates a route table.
func (s *Service) CreateOrUpdate(ctx context.Context, spec azure.Spec) error {
	asgSpec, ok := spec.(*Spec)
	if !ok {
		return errors.New("invalid application security groups specification")
	}

	klog.V(2).Infof("creating application security group %s", asgSpec.Name)
	f, err := s.Client.CreateOrUpdate(
		ctx,
		s.Scope.MachineConfig.ResourceGroup,
		asgSpec.Name,
		network.ApplicationSecurityGroup{
			Location: to.StringPtr(s.Scope.MachineConfig.Location),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to create application security group %s in resource group %s: %w", asgSpec.Name, s.Scope.MachineConfig.ResourceGroup, err)
	}

	err = f.WaitForCompletionRef(ctx, s.Client.Client)
	if err != nil {
		return fmt.Errorf("cannot create, future response: %w", err)
	}

	_, err = f.Result(s.Client)
	if err != nil {
		return fmt.Errorf("result error: %w", err)
	}
	klog.V(2).Infof("created application security group %s", asgSpec.Name)
	return err
}

// Delete deletes the route table with the provided name.
func (s *Service) Delete(ctx context.Context, spec azure.Spec) error {
	asgSpec, ok := spec.(*Spec)
	if !ok {
		return errors.New("invalid application security groups specification")
	}
	klog.V(2).Infof("deleting application security group %s", asgSpec.Name)
	f, err := s.Client.Delete(ctx, s.Scope.MachineConfig.ResourceGroup, asgSpec.Name)
	if err != nil && azure.ResourceNotFound(err) {
		// already deleted
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to delete application security group %s in resource group %s: %w", asgSpec.Name, s.Scope.MachineConfig.ResourceGroup, err)
	}

	err = f.WaitForCompletionRef(ctx, s.Client.Client)
	if err != nil {
		return fmt.Errorf("cannot create, future response: %w", err)
	}

	_, err = f.Result(s.Client)
	if err != nil {
		return err
	}

	klog.V(2).Infof("deleted application security group %s", asgSpec.Name)
	return err
}
