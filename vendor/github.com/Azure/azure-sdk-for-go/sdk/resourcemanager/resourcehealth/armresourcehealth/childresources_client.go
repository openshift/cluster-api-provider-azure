//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcehealth

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// ChildResourcesClient contains the methods for the ChildResources group.
// Don't use this type directly, use NewChildResourcesClient() instead.
type ChildResourcesClient struct {
	internal *arm.Client
}

// NewChildResourcesClient creates a new instance of ChildResourcesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewChildResourcesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ChildResourcesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ChildResourcesClient{
		internal: cl,
	}
	return client, nil
}

// NewListPager - Lists the all the children and its current health status for a parent resource. Use the nextLink property
// in the response to get the next page of children current health
//
// Generated from API version 2022-10-01
//   - resourceURI - The fully qualified ID of the resource, including the resource name and resource type. Currently the API
//     only support not nested parent resource type:
//     /subscriptions/{subscriptionId}/resourceGroups/{resource-group-name}/providers/{resource-provider-name}/{resource-type}/{resource-name}
//   - options - ChildResourcesClientListOptions contains the optional parameters for the ChildResourcesClient.NewListPager method.
func (client *ChildResourcesClient) NewListPager(resourceURI string, options *ChildResourcesClientListOptions) *runtime.Pager[ChildResourcesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ChildResourcesClientListResponse]{
		More: func(page ChildResourcesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ChildResourcesClientListResponse) (ChildResourcesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ChildResourcesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceURI, options)
			}, nil)
			if err != nil {
				return ChildResourcesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ChildResourcesClient) listCreateRequest(ctx context.Context, resourceURI string, options *ChildResourcesClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ResourceHealth/childResources"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ChildResourcesClient) listHandleResponse(resp *http.Response) (ChildResourcesClientListResponse, error) {
	result := ChildResourcesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailabilityStatusListResult); err != nil {
		return ChildResourcesClientListResponse{}, err
	}
	return result, nil
}
