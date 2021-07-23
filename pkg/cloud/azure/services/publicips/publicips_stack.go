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

package publicips

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/profiles/2019-03-01/network/mgmt/network"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/klog/v2"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
)

// StackHubSpec specification for public ip
type StackHubSpec struct {
	Name string
}

// Get provides information about a route table.
func (s *StackHubService) Get(ctx context.Context, spec azure.Spec) (interface{}, error) {
	publicIPSpec, ok := spec.(*StackHubSpec)
	if !ok {
		return network.PublicIPAddress{}, errors.New("Invalid PublicIP Specification")
	}
	publicIP, err := s.Client.Get(ctx, s.Scope.MachineConfig.ResourceGroup, publicIPSpec.Name, "")
	if err != nil && azure.ResourceNotFound(err) {
		return nil, fmt.Errorf("publicip %s not found: %w", publicIPSpec.Name, err)
	} else if err != nil {
		return publicIP, err
	}
	return publicIP, nil
}

// CreateOrUpdate creates or updates a public ip
func (s *StackHubService) CreateOrUpdate(ctx context.Context, spec azure.Spec) error {
	publicIPSpec, ok := spec.(*StackHubSpec)
	if !ok {
		return errors.New("Invalid PublicIP Specification")
	}
	ipName := publicIPSpec.Name
	klog.V(2).Infof("creating public ip %s", ipName)

	// https://docs.microsoft.com/en-us/azure/load-balancer/load-balancer-standard-availability-zones#zone-redundant-by-default
	f, err := s.Client.CreateOrUpdate(
		ctx,
		s.Scope.MachineConfig.ResourceGroup,
		ipName,
		network.PublicIPAddress{
			Sku:      &network.PublicIPAddressSku{Name: network.PublicIPAddressSkuNameStandard},
			Name:     to.StringPtr(ipName),
			Location: to.StringPtr(s.Scope.MachineConfig.Location),
			PublicIPAddressPropertiesFormat: &network.PublicIPAddressPropertiesFormat{
				PublicIPAddressVersion:   network.IPv4,
				PublicIPAllocationMethod: network.Static,
				DNSSettings: &network.PublicIPAddressDNSSettings{
					DomainNameLabel: to.StringPtr(strings.ToLower(ipName)),
				},
			},
		},
	)

	if err != nil {
		return fmt.Errorf("cannot create public ip: %w", err)
	}

	err = f.WaitForCompletionRef(ctx, s.Client.Client)
	if err != nil {
		return fmt.Errorf("cannot create, future response: %w", err)
	}

	_, err = f.Result(s.Client)
	if err != nil {
		return fmt.Errorf("result error: %w", err)
	}
	klog.V(2).Infof("successfully created public ip %s", ipName)
	return err
}

// Delete deletes the public ip with the provided scope.
func (s *StackHubService) Delete(ctx context.Context, spec azure.Spec) error {
	publicIPSpec, ok := spec.(*StackHubSpec)
	if !ok {
		return errors.New("Invalid PublicIP Specification")
	}
	klog.V(2).Infof("deleting public ip %s", publicIPSpec.Name)
	f, err := s.Client.Delete(ctx, s.Scope.MachineConfig.ResourceGroup, publicIPSpec.Name)
	if err != nil && azure.ResourceNotFound(err) {
		// already deleted
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to delete public ip %s in resource group %s: %w", publicIPSpec.Name, s.Scope.MachineConfig.ResourceGroup, err)
	}

	err = f.WaitForCompletionRef(ctx, s.Client.Client)
	if err != nil {
		return fmt.Errorf("cannot create, future response: %w", err)
	}

	_, err = f.Result(s.Client)
	if err != nil {
		return fmt.Errorf("result error: %w", err)
	}
	klog.V(2).Infof("deleted public ip %s", publicIPSpec.Name)
	return err
}
