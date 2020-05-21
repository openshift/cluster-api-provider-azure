package applicationsecuritygroups

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-12-01/network"
	"github.com/Azure/go-autorest/autorest"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
)

// Client wraps go-sdk
type Client interface {
	Get(context.Context, string, string) (network.ApplicationSecurityGroup, error)
	CreateOrUpdate(context.Context, string, string, network.ApplicationSecurityGroup) error
	Delete(context.Context, string, string) error
}

// AzureClient contains the Azure go-sdk Client
type AzureClient struct {
	applicationsecuritygroups network.ApplicationSecurityGroupsClient
}

var _ Client = &AzureClient{}

// NewClient creates a new application security groups client from subscription ID.
func NewClient(subscriptionID string, authorizer autorest.Authorizer) *AzureClient {
	c := newApplicationSecurityGroups(subscriptionID, authorizer)
	return &AzureClient{c}
}

// newApplicationSecurityGroups creates a new application security groups client from subscription ID.
func newApplicationSecurityGroups(subscriptionID string, authorizer autorest.Authorizer) network.ApplicationSecurityGroupsClient {
	applicationSecurityGroupsClient := network.NewApplicationSecurityGroupsClient(subscriptionID)
	applicationSecurityGroupsClient.Authorizer = authorizer
	applicationSecurityGroupsClient.AddToUserAgent(azure.UserAgent)
	return applicationSecurityGroupsClient
}

// Get gets the specified application security group.
func (ac *AzureClient) Get(ctx context.Context, resourceGroupName, appGroupName string) (network.ApplicationSecurityGroup, error) {
	return ac.applicationsecuritygroups.Get(ctx, resourceGroupName, appGroupName)
}

// CreateOrUpdate creates or updates a application security group.
func (ac *AzureClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, appGroupName string, sg network.ApplicationSecurityGroup) error {
	future, err := ac.applicationsecuritygroups.CreateOrUpdate(ctx, resourceGroupName, appGroupName, sg)
	if err != nil {
		return err
	}
	err = future.WaitForCompletionRef(ctx, ac.applicationsecuritygroups.Client)
	if err != nil {
		return err
	}
	_, err = future.Result(ac.applicationsecuritygroups)
	return err
}

// Delete deletes the specified application security group.
func (ac *AzureClient) Delete(ctx context.Context, resourceGroupName, appGroupName string) error {
	future, err := ac.applicationsecuritygroups.Delete(ctx, resourceGroupName, appGroupName)
	if err != nil {
		return err
	}
	err = future.WaitForCompletionRef(ctx, ac.applicationsecuritygroups.Client)
	if err != nil {
		return err
	}
	_, err = future.Result(ac.applicationsecuritygroups)
	return err
}
