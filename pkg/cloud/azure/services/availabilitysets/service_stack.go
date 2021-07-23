package availabilitysets

import (
	"github.com/Azure/azure-sdk-for-go/profiles/2019-03-01/compute/mgmt/compute"

	"github.com/Azure/go-autorest/autorest"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
)

// StackHubService provides operations on availability zones
type StackHubService struct {
	Client compute.AvailabilitySetsClient
	Scope  *actuators.MachineScope
}

// getAvailabilitySetsClient creates a new availability zones client from subscriptionid.
func getAvailabilitySetsClientStackHub(resourceManagerEndpoint, subscriptionID string, authorizer autorest.Authorizer) compute.AvailabilitySetsClient {
	availabilitySetClient := compute.NewAvailabilitySetsClientWithBaseURI(resourceManagerEndpoint, subscriptionID)
	availabilitySetClient.Authorizer = authorizer
	availabilitySetClient.AddToUserAgent(azure.UserAgent)
	return availabilitySetClient
}

// NewStackHubService creates a new availability zones service.
func NewStackHubService(scope *actuators.MachineScope) azure.Service {
	return &StackHubService{
		Client: getAvailabilitySetsClientStackHub(scope.ResourceManagerEndpoint, scope.SubscriptionID, scope.Authorizer),
		Scope:  scope,
	}
}
