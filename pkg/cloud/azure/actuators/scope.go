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

package actuators

import (
	"context"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

const (
	// AzureCredsSubscriptionIDKey subcription ID
	AzureCredsSubscriptionIDKey = "azure_subscription_id"
	// AzureCredsClientIDKey client id
	AzureCredsClientIDKey = "azure_client_id"
	// AzureCredsClientSecretKey client secret
	AzureCredsClientSecretKey = "azure_client_secret"
	// AzureCredsTenantIDKey tenant ID
	AzureCredsTenantIDKey = "azure_tenant_id"
	// AzureCredsResourceGroupKey resource group
	AzureCredsResourceGroupKey = "azure_resourcegroup"
	// AzureCredsRegionKey region
	AzureCredsRegionKey = "azure_region"
	// AzureResourcePrefix resource prefix for created azure resources
	AzureResourcePrefix = "azure_resource_prefix"
)

func NewScopeFromSecret(ctx context.Context, secret *corev1.Secret) (*Scope, error) {
	secretType := types.NamespacedName{Namespace: secret.Namespace, Name: secret.Name}

	subscriptionID, ok := secret.Data[AzureCredsSubscriptionIDKey]
	if !ok {
		return nil, errors.Errorf("Azure subscription id %v did not contain key %v",
			secretType.String(), AzureCredsSubscriptionIDKey)
	}
	clientID, ok := secret.Data[AzureCredsClientIDKey]
	if !ok {
		return nil, errors.Errorf("Azure client id %v did not contain key %v",
			secretType.String(), AzureCredsClientIDKey)
	}
	clientSecret, ok := secret.Data[AzureCredsClientSecretKey]
	if !ok {
		return nil, errors.Errorf("Azure client secret %v did not contain key %v",
			secretType.String(), AzureCredsClientSecretKey)
	}
	tenantID, ok := secret.Data[AzureCredsTenantIDKey]
	if !ok {
		return nil, errors.Errorf("Azure tenant id %v did not contain key %v",
			secretType.String(), AzureCredsTenantIDKey)
	}
	resourceGroup, ok := secret.Data[AzureCredsResourceGroupKey]
	if !ok {
		return nil, errors.Errorf("Azure resource group %v did not contain key %v",
			secretType.String(), AzureCredsResourceGroupKey)
	}
	region, ok := secret.Data[AzureCredsRegionKey]
	if !ok {
		return nil, errors.Errorf("Azure region %v did not contain key %v",
			secretType.String(), AzureCredsRegionKey)
	}
	clusterName, ok := secret.Data[AzureResourcePrefix]
	if !ok {
		return nil, errors.Errorf("Azure resource prefix %v did not contain key %v",
			secretType.String(), AzureResourcePrefix)
	}

	config := auth.NewClientCredentialsConfig(string(clientID), string(clientSecret), string(tenantID))
	config.Resource = azure.PublicCloud.ResourceManagerEndpoint
	authorizer, err := config.Authorizer()
	if err != nil {
		return nil, errors.Errorf("failed to create azure session: %v", err)
	}

	return &Scope{
		ClusterName: string(clusterName),
		AzureClients: AzureClients{
			Authorizer:     authorizer,
			SubscriptionID: string(subscriptionID),
		},
		Location:             string(region),
		ResourceGroup:        string(resourceGroup),
		NetworkResourceGroup: string(resourceGroup),
	}, nil
}

// AzureClients contains all the Azure clients used by the scopes.
type AzureClients struct {
	SubscriptionID string
	Authorizer     autorest.Authorizer
}

// Scope defines the basic context for an actuator to operate upon.
type Scope struct {
	AzureClients
	ClusterName          string
	Location             string
	ResourceGroup        string
	NetworkResourceGroup string
}
