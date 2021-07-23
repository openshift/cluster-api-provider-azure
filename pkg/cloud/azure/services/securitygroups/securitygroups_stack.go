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

package securitygroups

import (
	"context"
	"errors"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/profiles/2019-03-01/network/mgmt/network"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/klog/v2"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
)

// StackHubSpec specification for network security groups
type StackHubSpec struct {
	Name           string
	IsControlPlane bool
}

// Get provides information about a route table.
func (s *StackHubService) Get(ctx context.Context, spec azure.Spec) (interface{}, error) {
	nsgSpec, ok := spec.(*StackHubSpec)
	if !ok {
		return network.SecurityGroup{}, errors.New("invalid security groups specification")
	}
	securityGroup, err := s.Client.Get(ctx, s.Scope.MachineConfig.ResourceGroup, nsgSpec.Name, "")
	if err != nil && azure.ResourceNotFound(err) {
		return nil, fmt.Errorf("security group %s not found: %w", nsgSpec.Name, err)
	} else if err != nil {
		return securityGroup, err
	}
	return securityGroup, nil
}

// CreateOrUpdate creates or updates a route table.
func (s *StackHubService) CreateOrUpdate(ctx context.Context, spec azure.Spec) error {
	nsgSpec, ok := spec.(*StackHubSpec)
	if !ok {
		return errors.New("invalid security groups specification")
	}

	securityRules := &[]network.SecurityRule{}

	if nsgSpec.IsControlPlane {
		klog.V(2).Infof("using additional rules for control plane %s", nsgSpec.Name)
		securityRules = &[]network.SecurityRule{
			{
				Name: to.StringPtr("allow_ssh"),
				SecurityRulePropertiesFormat: &network.SecurityRulePropertiesFormat{
					Protocol:                 network.SecurityRuleProtocolTCP,
					SourceAddressPrefix:      to.StringPtr("*"),
					SourcePortRange:          to.StringPtr("*"),
					DestinationAddressPrefix: to.StringPtr("*"),
					DestinationPortRange:     to.StringPtr("22"),
					Access:                   network.SecurityRuleAccessAllow,
					Direction:                network.SecurityRuleDirectionInbound,
					Priority:                 to.Int32Ptr(100),
				},
			},
			{
				Name: to.StringPtr("allow_6443"),
				SecurityRulePropertiesFormat: &network.SecurityRulePropertiesFormat{
					Protocol:                 network.SecurityRuleProtocolTCP,
					SourceAddressPrefix:      to.StringPtr("*"),
					SourcePortRange:          to.StringPtr("*"),
					DestinationAddressPrefix: to.StringPtr("*"),
					DestinationPortRange:     to.StringPtr("6443"),
					Access:                   network.SecurityRuleAccessAllow,
					Direction:                network.SecurityRuleDirectionInbound,
					Priority:                 to.Int32Ptr(101),
				},
			},
		}
	}

	klog.V(2).Infof("creating security group %s", nsgSpec.Name)
	f, err := s.Client.CreateOrUpdate(
		ctx,
		s.Scope.MachineConfig.ResourceGroup,
		nsgSpec.Name,
		network.SecurityGroup{
			Location: to.StringPtr(s.Scope.MachineConfig.Location),
			SecurityGroupPropertiesFormat: &network.SecurityGroupPropertiesFormat{
				SecurityRules: securityRules,
			},
		},
	)
	if err != nil {
		return fmt.Errorf("failed to create security group %s in resource group %s: %w", nsgSpec.Name, s.Scope.MachineConfig.ResourceGroup, err)
	}

	err = f.WaitForCompletionRef(ctx, s.Client.Client)
	if err != nil {
		return fmt.Errorf("cannot create, future response: %w", err)
	}

	_, err = f.Result(s.Client)
	if err != nil {
		return fmt.Errorf("result error: %w", err)
	}
	klog.V(2).Infof("created security group %s", nsgSpec.Name)
	return err
}

// Delete deletes the route table with the provided name.
func (s *StackHubService) Delete(ctx context.Context, spec azure.Spec) error {
	nsgSpec, ok := spec.(*StackHubSpec)
	if !ok {
		return errors.New("invalid security groups specification")
	}
	klog.V(2).Infof("deleting security group %s", nsgSpec.Name)
	f, err := s.Client.Delete(ctx, s.Scope.MachineConfig.ResourceGroup, nsgSpec.Name)
	if err != nil && azure.ResourceNotFound(err) {
		// already deleted
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to delete security group %s in resource group %s: %w", nsgSpec.Name, s.Scope.MachineConfig.ResourceGroup, err)
	}

	err = f.WaitForCompletionRef(ctx, s.Client.Client)
	if err != nil {
		return fmt.Errorf("cannot create, future response: %w", err)
	}

	_, err = f.Result(s.Client)
	if err != nil {
		return err
	}

	klog.V(2).Infof("deleted security group %s", nsgSpec.Name)
	return err
}
