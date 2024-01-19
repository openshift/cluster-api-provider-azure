//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

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

// PredictiveMetricClient contains the methods for the PredictiveMetric group.
// Don't use this type directly, use NewPredictiveMetricClient() instead.
type PredictiveMetricClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPredictiveMetricClient creates a new instance of PredictiveMetricClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPredictiveMetricClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PredictiveMetricClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PredictiveMetricClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - get predictive autoscale metric future data
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - autoscaleSettingName - The autoscale setting name.
//   - timespan - The timespan of the query. It is a string with the following format 'startDateTimeISO/endDateTimeISO'.
//   - interval - The interval (i.e. timegrain) of the query.
//   - metricNamespace - Metric namespace to query metric definitions for.
//   - metricName - The names of the metrics (comma separated) to retrieve. Special case: If a metricname itself has a comma in
//     it then use %2 to indicate it. Eg: 'Metric,Name1' should be 'Metric%2Name1'
//   - aggregation - The list of aggregation types (comma separated) to retrieve.
//   - options - PredictiveMetricClientGetOptions contains the optional parameters for the PredictiveMetricClient.Get method.
func (client *PredictiveMetricClient) Get(ctx context.Context, resourceGroupName string, autoscaleSettingName string, timespan string, interval string, metricNamespace string, metricName string, aggregation string, options *PredictiveMetricClientGetOptions) (PredictiveMetricClientGetResponse, error) {
	var err error
	const operationName = "PredictiveMetricClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, autoscaleSettingName, timespan, interval, metricNamespace, metricName, aggregation, options)
	if err != nil {
		return PredictiveMetricClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PredictiveMetricClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PredictiveMetricClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PredictiveMetricClient) getCreateRequest(ctx context.Context, resourceGroupName string, autoscaleSettingName string, timespan string, interval string, metricNamespace string, metricName string, aggregation string, options *PredictiveMetricClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Insights/autoscalesettings/{autoscaleSettingName}/predictiveMetrics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if autoscaleSettingName == "" {
		return nil, errors.New("parameter autoscaleSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{autoscaleSettingName}", url.PathEscape(autoscaleSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("timespan", timespan)
	reqQP.Set("interval", interval)
	reqQP.Set("metricNamespace", metricNamespace)
	reqQP.Set("metricName", metricName)
	reqQP.Set("aggregation", aggregation)
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PredictiveMetricClient) getHandleResponse(resp *http.Response) (PredictiveMetricClientGetResponse, error) {
	result := PredictiveMetricClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PredictiveResponse); err != nil {
		return PredictiveMetricClientGetResponse{}, err
	}
	return result, nil
}