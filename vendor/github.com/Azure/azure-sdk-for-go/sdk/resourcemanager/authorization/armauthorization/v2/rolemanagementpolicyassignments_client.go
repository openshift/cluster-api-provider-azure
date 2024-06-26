//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RoleManagementPolicyAssignmentsClient contains the methods for the RoleManagementPolicyAssignments group.
// Don't use this type directly, use NewRoleManagementPolicyAssignmentsClient() instead.
type RoleManagementPolicyAssignmentsClient struct {
	internal *arm.Client
}

// NewRoleManagementPolicyAssignmentsClient creates a new instance of RoleManagementPolicyAssignmentsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRoleManagementPolicyAssignmentsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*RoleManagementPolicyAssignmentsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RoleManagementPolicyAssignmentsClient{
		internal: cl,
	}
	return client, nil
}

// Create - Create a role management policy assignment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-01
//   - scope - The scope of the role management policy assignment to upsert.
//   - roleManagementPolicyAssignmentName - The name of format {guid_guid} the role management policy assignment to upsert.
//   - parameters - Parameters for the role management policy assignment.
//   - options - RoleManagementPolicyAssignmentsClientCreateOptions contains the optional parameters for the RoleManagementPolicyAssignmentsClient.Create
//     method.
func (client *RoleManagementPolicyAssignmentsClient) Create(ctx context.Context, scope string, roleManagementPolicyAssignmentName string, parameters RoleManagementPolicyAssignment, options *RoleManagementPolicyAssignmentsClientCreateOptions) (RoleManagementPolicyAssignmentsClientCreateResponse, error) {
	var err error
	const operationName = "RoleManagementPolicyAssignmentsClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, scope, roleManagementPolicyAssignmentName, parameters, options)
	if err != nil {
		return RoleManagementPolicyAssignmentsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoleManagementPolicyAssignmentsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return RoleManagementPolicyAssignmentsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *RoleManagementPolicyAssignmentsClient) createCreateRequest(ctx context.Context, scope string, roleManagementPolicyAssignmentName string, parameters RoleManagementPolicyAssignment, options *RoleManagementPolicyAssignmentsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementPolicyAssignments/{roleManagementPolicyAssignmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleManagementPolicyAssignmentName == "" {
		return nil, errors.New("parameter roleManagementPolicyAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleManagementPolicyAssignmentName}", url.PathEscape(roleManagementPolicyAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *RoleManagementPolicyAssignmentsClient) createHandleResponse(resp *http.Response) (RoleManagementPolicyAssignmentsClientCreateResponse, error) {
	result := RoleManagementPolicyAssignmentsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleManagementPolicyAssignment); err != nil {
		return RoleManagementPolicyAssignmentsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a role management policy assignment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-01
//   - scope - The scope of the role management policy assignment to delete.
//   - roleManagementPolicyAssignmentName - The name of format {guid_guid} the role management policy assignment to delete.
//   - options - RoleManagementPolicyAssignmentsClientDeleteOptions contains the optional parameters for the RoleManagementPolicyAssignmentsClient.Delete
//     method.
func (client *RoleManagementPolicyAssignmentsClient) Delete(ctx context.Context, scope string, roleManagementPolicyAssignmentName string, options *RoleManagementPolicyAssignmentsClientDeleteOptions) (RoleManagementPolicyAssignmentsClientDeleteResponse, error) {
	var err error
	const operationName = "RoleManagementPolicyAssignmentsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, scope, roleManagementPolicyAssignmentName, options)
	if err != nil {
		return RoleManagementPolicyAssignmentsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoleManagementPolicyAssignmentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return RoleManagementPolicyAssignmentsClientDeleteResponse{}, err
	}
	return RoleManagementPolicyAssignmentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RoleManagementPolicyAssignmentsClient) deleteCreateRequest(ctx context.Context, scope string, roleManagementPolicyAssignmentName string, options *RoleManagementPolicyAssignmentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementPolicyAssignments/{roleManagementPolicyAssignmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleManagementPolicyAssignmentName == "" {
		return nil, errors.New("parameter roleManagementPolicyAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleManagementPolicyAssignmentName}", url.PathEscape(roleManagementPolicyAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the specified role management policy assignment for a resource scope
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-01
//   - scope - The scope of the role management policy.
//   - roleManagementPolicyAssignmentName - The name of format {guid_guid} the role management policy assignment to get.
//   - options - RoleManagementPolicyAssignmentsClientGetOptions contains the optional parameters for the RoleManagementPolicyAssignmentsClient.Get
//     method.
func (client *RoleManagementPolicyAssignmentsClient) Get(ctx context.Context, scope string, roleManagementPolicyAssignmentName string, options *RoleManagementPolicyAssignmentsClientGetOptions) (RoleManagementPolicyAssignmentsClientGetResponse, error) {
	var err error
	const operationName = "RoleManagementPolicyAssignmentsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, scope, roleManagementPolicyAssignmentName, options)
	if err != nil {
		return RoleManagementPolicyAssignmentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoleManagementPolicyAssignmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RoleManagementPolicyAssignmentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RoleManagementPolicyAssignmentsClient) getCreateRequest(ctx context.Context, scope string, roleManagementPolicyAssignmentName string, options *RoleManagementPolicyAssignmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementPolicyAssignments/{roleManagementPolicyAssignmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleManagementPolicyAssignmentName == "" {
		return nil, errors.New("parameter roleManagementPolicyAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleManagementPolicyAssignmentName}", url.PathEscape(roleManagementPolicyAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoleManagementPolicyAssignmentsClient) getHandleResponse(resp *http.Response) (RoleManagementPolicyAssignmentsClientGetResponse, error) {
	result := RoleManagementPolicyAssignmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleManagementPolicyAssignment); err != nil {
		return RoleManagementPolicyAssignmentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListForScopePager - Gets role management assignment policies for a resource scope.
//
// Generated from API version 2020-10-01
//   - scope - The scope of the role management policy.
//   - options - RoleManagementPolicyAssignmentsClientListForScopeOptions contains the optional parameters for the RoleManagementPolicyAssignmentsClient.NewListForScopePager
//     method.
func (client *RoleManagementPolicyAssignmentsClient) NewListForScopePager(scope string, options *RoleManagementPolicyAssignmentsClientListForScopeOptions) *runtime.Pager[RoleManagementPolicyAssignmentsClientListForScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[RoleManagementPolicyAssignmentsClientListForScopeResponse]{
		More: func(page RoleManagementPolicyAssignmentsClientListForScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RoleManagementPolicyAssignmentsClientListForScopeResponse) (RoleManagementPolicyAssignmentsClientListForScopeResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RoleManagementPolicyAssignmentsClient.NewListForScopePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listForScopeCreateRequest(ctx, scope, options)
			}, nil)
			if err != nil {
				return RoleManagementPolicyAssignmentsClientListForScopeResponse{}, err
			}
			return client.listForScopeHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *RoleManagementPolicyAssignmentsClient) listForScopeCreateRequest(ctx context.Context, scope string, options *RoleManagementPolicyAssignmentsClientListForScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementPolicyAssignments"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForScopeHandleResponse handles the ListForScope response.
func (client *RoleManagementPolicyAssignmentsClient) listForScopeHandleResponse(resp *http.Response) (RoleManagementPolicyAssignmentsClientListForScopeResponse, error) {
	result := RoleManagementPolicyAssignmentsClientListForScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleManagementPolicyAssignmentListResult); err != nil {
		return RoleManagementPolicyAssignmentsClientListForScopeResponse{}, err
	}
	return result, nil
}
