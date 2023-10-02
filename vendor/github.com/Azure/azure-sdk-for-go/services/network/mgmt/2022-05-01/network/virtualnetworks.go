package network

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// VirtualNetworksClient is the network Client
type VirtualNetworksClient struct {
	BaseClient
}

// NewVirtualNetworksClient creates an instance of the VirtualNetworksClient client.
func NewVirtualNetworksClient(subscriptionID string) VirtualNetworksClient {
	return NewVirtualNetworksClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualNetworksClientWithBaseURI creates an instance of the VirtualNetworksClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewVirtualNetworksClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworksClient {
	return VirtualNetworksClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckIPAddressAvailability checks whether a private IP address is available for use.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualNetworkName - the name of the virtual network.
// IPAddress - the private IP address to be verified.
func (client VirtualNetworksClient) CheckIPAddressAvailability(ctx context.Context, resourceGroupName string, virtualNetworkName string, IPAddress string) (result IPAddressAvailabilityResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworksClient.CheckIPAddressAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CheckIPAddressAvailabilityPreparer(ctx, resourceGroupName, virtualNetworkName, IPAddress)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "CheckIPAddressAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckIPAddressAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "CheckIPAddressAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckIPAddressAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "CheckIPAddressAvailability", resp, "Failure responding to request")
		return
	}

	return
}

// CheckIPAddressAvailabilityPreparer prepares the CheckIPAddressAvailability request.
func (client VirtualNetworksClient) CheckIPAddressAvailabilityPreparer(ctx context.Context, resourceGroupName string, virtualNetworkName string, IPAddress string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName": autorest.Encode("path", virtualNetworkName),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"ipAddress":   autorest.Encode("query", IPAddress),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/CheckIPAddressAvailability", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckIPAddressAvailabilitySender sends the CheckIPAddressAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) CheckIPAddressAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckIPAddressAvailabilityResponder handles the response to the CheckIPAddressAvailability request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) CheckIPAddressAvailabilityResponder(resp *http.Response) (result IPAddressAvailabilityResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdate creates or updates a virtual network in the specified resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualNetworkName - the name of the virtual network.
// parameters - parameters supplied to the create or update virtual network operation.
func (client VirtualNetworksClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters VirtualNetwork) (result VirtualNetworksCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworksClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.VirtualNetworkPropertiesFormat", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.VirtualNetworkPropertiesFormat.BgpCommunities", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.VirtualNetworkPropertiesFormat.BgpCommunities.VirtualNetworkCommunity", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "parameters.VirtualNetworkPropertiesFormat.Encryption", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.VirtualNetworkPropertiesFormat.Encryption.Enabled", Name: validation.Null, Rule: true, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("network.VirtualNetworksClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, virtualNetworkName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualNetworksClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters VirtualNetwork) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName": autorest.Encode("path", virtualNetworkName),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) CreateOrUpdateSender(req *http.Request) (future VirtualNetworksCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) CreateOrUpdateResponder(resp *http.Response) (result VirtualNetwork, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified virtual network.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualNetworkName - the name of the virtual network.
func (client VirtualNetworksClient) Delete(ctx context.Context, resourceGroupName string, virtualNetworkName string) (result VirtualNetworksDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworksClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, virtualNetworkName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VirtualNetworksClient) DeletePreparer(ctx context.Context, resourceGroupName string, virtualNetworkName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName": autorest.Encode("path", virtualNetworkName),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) DeleteSender(req *http.Request) (future VirtualNetworksDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified virtual network by resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualNetworkName - the name of the virtual network.
// expand - expands referenced resources.
func (client VirtualNetworksClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, expand string) (result VirtualNetwork, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworksClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, virtualNetworkName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualNetworksClient) GetPreparer(ctx context.Context, resourceGroupName string, virtualNetworkName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName": autorest.Encode("path", virtualNetworkName),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) GetResponder(resp *http.Response) (result VirtualNetwork, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all virtual networks in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client VirtualNetworksClient) List(ctx context.Context, resourceGroupName string) (result VirtualNetworkListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworksClient.List")
		defer func() {
			sc := -1
			if result.vnlr.Response.Response != nil {
				sc = result.vnlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.vnlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "List", resp, "Failure sending request")
		return
	}

	result.vnlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "List", resp, "Failure responding to request")
		return
	}
	if result.vnlr.hasNextLink() && result.vnlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualNetworksClient) ListPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) ListResponder(resp *http.Response) (result VirtualNetworkListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client VirtualNetworksClient) listNextResults(ctx context.Context, lastResults VirtualNetworkListResult) (result VirtualNetworkListResult, err error) {
	req, err := lastResults.virtualNetworkListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualNetworksClient) ListComplete(ctx context.Context, resourceGroupName string) (result VirtualNetworkListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworksClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName)
	return
}

// ListAll gets all virtual networks in a subscription.
func (client VirtualNetworksClient) ListAll(ctx context.Context) (result VirtualNetworkListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworksClient.ListAll")
		defer func() {
			sc := -1
			if result.vnlr.Response.Response != nil {
				sc = result.vnlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listAllNextResults
	req, err := client.ListAllPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.vnlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListAll", resp, "Failure sending request")
		return
	}

	result.vnlr, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListAll", resp, "Failure responding to request")
		return
	}
	if result.vnlr.hasNextLink() && result.vnlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client VirtualNetworksClient) ListAllPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualNetworks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) ListAllResponder(resp *http.Response) (result VirtualNetworkListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAllNextResults retrieves the next set of results, if any.
func (client VirtualNetworksClient) listAllNextResults(ctx context.Context, lastResults VirtualNetworkListResult) (result VirtualNetworkListResult, err error) {
	req, err := lastResults.virtualNetworkListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "listAllNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "listAllNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "listAllNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAllComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualNetworksClient) ListAllComplete(ctx context.Context) (result VirtualNetworkListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworksClient.ListAll")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAll(ctx)
	return
}

// ListDdosProtectionStatus gets the Ddos Protection Status of all IP Addresses under the Virtual Network
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualNetworkName - the name of the virtual network.
// top - the max number of ip addresses to return.
// skipToken - the skipToken that is given with nextLink.
func (client VirtualNetworksClient) ListDdosProtectionStatus(ctx context.Context, resourceGroupName string, virtualNetworkName string, top *int32, skipToken string) (result VirtualNetworksListDdosProtectionStatusFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworksClient.ListDdosProtectionStatus")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListDdosProtectionStatusPreparer(ctx, resourceGroupName, virtualNetworkName, top, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListDdosProtectionStatus", nil, "Failure preparing request")
		return
	}

	result, err = client.ListDdosProtectionStatusSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListDdosProtectionStatus", result.Response(), "Failure sending request")
		return
	}

	return
}

// ListDdosProtectionStatusPreparer prepares the ListDdosProtectionStatus request.
func (client VirtualNetworksClient) ListDdosProtectionStatusPreparer(ctx context.Context, resourceGroupName string, virtualNetworkName string, top *int32, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName": autorest.Encode("path", virtualNetworkName),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["top"] = autorest.Encode("query", *top)
	}
	if len(skipToken) > 0 {
		queryParameters["skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/ddosProtectionStatus", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListDdosProtectionStatusSender sends the ListDdosProtectionStatus request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) ListDdosProtectionStatusSender(req *http.Request) (future VirtualNetworksListDdosProtectionStatusFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// ListDdosProtectionStatusResponder handles the response to the ListDdosProtectionStatus request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) ListDdosProtectionStatusResponder(resp *http.Response) (result VirtualNetworkDdosProtectionStatusResultPage, err error) {
	result.vndpsr, err = client.listDdosProtectionStatusResponder(resp)
	result.fn = client.listDdosProtectionStatusNextResults
	return
}

func (client VirtualNetworksClient) listDdosProtectionStatusResponder(resp *http.Response) (result VirtualNetworkDdosProtectionStatusResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listDdosProtectionStatusNextResults retrieves the next set of results, if any.
func (client VirtualNetworksClient) listDdosProtectionStatusNextResults(ctx context.Context, lastResults VirtualNetworkDdosProtectionStatusResult) (result VirtualNetworkDdosProtectionStatusResult, err error) {
	req, err := lastResults.virtualNetworkDdosProtectionStatusResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "listDdosProtectionStatusNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "listDdosProtectionStatusNextResults", resp, "Failure sending next results request")
	}
	return client.listDdosProtectionStatusResponder(resp)
}

// ListDdosProtectionStatusComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualNetworksClient) ListDdosProtectionStatusComplete(ctx context.Context, resourceGroupName string, virtualNetworkName string, top *int32, skipToken string) (result VirtualNetworksListDdosProtectionStatusAllFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworksClient.ListDdosProtectionStatus")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	var future VirtualNetworksListDdosProtectionStatusFuture
	future, err = client.ListDdosProtectionStatus(ctx, resourceGroupName, virtualNetworkName, top, skipToken)
	result.FutureAPI = future.FutureAPI
	return
}

// ListUsage lists usage stats.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualNetworkName - the name of the virtual network.
func (client VirtualNetworksClient) ListUsage(ctx context.Context, resourceGroupName string, virtualNetworkName string) (result VirtualNetworkListUsageResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworksClient.ListUsage")
		defer func() {
			sc := -1
			if result.vnlur.Response.Response != nil {
				sc = result.vnlur.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listUsageNextResults
	req, err := client.ListUsagePreparer(ctx, resourceGroupName, virtualNetworkName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListUsage", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListUsageSender(req)
	if err != nil {
		result.vnlur.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListUsage", resp, "Failure sending request")
		return
	}

	result.vnlur, err = client.ListUsageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListUsage", resp, "Failure responding to request")
		return
	}
	if result.vnlur.hasNextLink() && result.vnlur.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListUsagePreparer prepares the ListUsage request.
func (client VirtualNetworksClient) ListUsagePreparer(ctx context.Context, resourceGroupName string, virtualNetworkName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName": autorest.Encode("path", virtualNetworkName),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/usages", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListUsageSender sends the ListUsage request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) ListUsageSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListUsageResponder handles the response to the ListUsage request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) ListUsageResponder(resp *http.Response) (result VirtualNetworkListUsageResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listUsageNextResults retrieves the next set of results, if any.
func (client VirtualNetworksClient) listUsageNextResults(ctx context.Context, lastResults VirtualNetworkListUsageResult) (result VirtualNetworkListUsageResult, err error) {
	req, err := lastResults.virtualNetworkListUsageResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "listUsageNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListUsageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "listUsageNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListUsageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "listUsageNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListUsageComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualNetworksClient) ListUsageComplete(ctx context.Context, resourceGroupName string, virtualNetworkName string) (result VirtualNetworkListUsageResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworksClient.ListUsage")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListUsage(ctx, resourceGroupName, virtualNetworkName)
	return
}

// UpdateTags updates a virtual network tags.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualNetworkName - the name of the virtual network.
// parameters - parameters supplied to update virtual network tags.
func (client VirtualNetworksClient) UpdateTags(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters TagsObject) (result VirtualNetwork, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworksClient.UpdateTags")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateTagsPreparer(ctx, resourceGroupName, virtualNetworkName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "UpdateTags", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateTagsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "UpdateTags", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateTagsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "UpdateTags", resp, "Failure responding to request")
		return
	}

	return
}

// UpdateTagsPreparer prepares the UpdateTags request.
func (client VirtualNetworksClient) UpdateTagsPreparer(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters TagsObject) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName": autorest.Encode("path", virtualNetworkName),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateTagsSender sends the UpdateTags request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) UpdateTagsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateTagsResponder handles the response to the UpdateTags request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) UpdateTagsResponder(resp *http.Response) (result VirtualNetwork, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
