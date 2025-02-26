//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcehealth

// AvailabilityStatusesClientGetByResourceOptions contains the optional parameters for the AvailabilityStatusesClient.GetByResource
// method.
type AvailabilityStatusesClientGetByResourceOptions struct {
	// Setting $expand=recommendedactions in url query expands the recommendedactions in the response.
	Expand *string

	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}

// AvailabilityStatusesClientListByResourceGroupOptions contains the optional parameters for the AvailabilityStatusesClient.NewListByResourceGroupPager
// method.
type AvailabilityStatusesClientListByResourceGroupOptions struct {
	// Setting $expand=recommendedactions in url query expands the recommendedactions in the response.
	Expand *string

	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}

// AvailabilityStatusesClientListBySubscriptionIDOptions contains the optional parameters for the AvailabilityStatusesClient.NewListBySubscriptionIDPager
// method.
type AvailabilityStatusesClientListBySubscriptionIDOptions struct {
	// Setting $expand=recommendedactions in url query expands the recommendedactions in the response.
	Expand *string

	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}

// AvailabilityStatusesClientListOptions contains the optional parameters for the AvailabilityStatusesClient.NewListPager
// method.
type AvailabilityStatusesClientListOptions struct {
	// Setting $expand=recommendedactions in url query expands the recommendedactions in the response.
	Expand *string

	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}

// ChildAvailabilityStatusesClientGetByResourceOptions contains the optional parameters for the ChildAvailabilityStatusesClient.GetByResource
// method.
type ChildAvailabilityStatusesClientGetByResourceOptions struct {
	// Setting $expand=recommendedactions in url query expands the recommendedactions in the response.
	Expand *string

	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}

// ChildAvailabilityStatusesClientListOptions contains the optional parameters for the ChildAvailabilityStatusesClient.NewListPager
// method.
type ChildAvailabilityStatusesClientListOptions struct {
	// Setting $expand=recommendedactions in url query expands the recommendedactions in the response.
	Expand *string

	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}

// ChildResourcesClientListOptions contains the optional parameters for the ChildResourcesClient.NewListPager method.
type ChildResourcesClientListOptions struct {
	// Setting $expand=recommendedactions in url query expands the recommendedactions in the response.
	Expand *string

	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}

// EmergingIssuesClientGetOptions contains the optional parameters for the EmergingIssuesClient.Get method.
type EmergingIssuesClientGetOptions struct {
	// placeholder for future optional parameters
}

// EmergingIssuesClientListOptions contains the optional parameters for the EmergingIssuesClient.NewListPager method.
type EmergingIssuesClientListOptions struct {
	// placeholder for future optional parameters
}

// EventClientFetchDetailsBySubscriptionIDAndTrackingIDOptions contains the optional parameters for the EventClient.FetchDetailsBySubscriptionIDAndTrackingID
// method.
type EventClientFetchDetailsBySubscriptionIDAndTrackingIDOptions struct {
	// placeholder for future optional parameters
}

// EventClientFetchDetailsByTenantIDAndTrackingIDOptions contains the optional parameters for the EventClient.FetchDetailsByTenantIDAndTrackingID
// method.
type EventClientFetchDetailsByTenantIDAndTrackingIDOptions struct {
	// placeholder for future optional parameters
}

// EventClientGetBySubscriptionIDAndTrackingIDOptions contains the optional parameters for the EventClient.GetBySubscriptionIDAndTrackingID
// method.
type EventClientGetBySubscriptionIDAndTrackingIDOptions struct {
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string

	// Specifies from when to return events, based on the lastUpdateTime property. For example, queryStartTime = 7/24/2020 OR
	// queryStartTime=7%2F24%2F2020
	QueryStartTime *string
}

// EventClientGetByTenantIDAndTrackingIDOptions contains the optional parameters for the EventClient.GetByTenantIDAndTrackingID
// method.
type EventClientGetByTenantIDAndTrackingIDOptions struct {
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string

	// Specifies from when to return events, based on the lastUpdateTime property. For example, queryStartTime = 7/24/2020 OR
	// queryStartTime=7%2F24%2F2020
	QueryStartTime *string
}

// EventsClientListBySingleResourceOptions contains the optional parameters for the EventsClient.NewListBySingleResourcePager
// method.
type EventsClientListBySingleResourceOptions struct {
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}

// EventsClientListBySubscriptionIDOptions contains the optional parameters for the EventsClient.NewListBySubscriptionIDPager
// method.
type EventsClientListBySubscriptionIDOptions struct {
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string

	// Specifies from when to return events, based on the lastUpdateTime property. For example, queryStartTime = 7/24/2020 OR
	// queryStartTime=7%2F24%2F2020
	QueryStartTime *string
}

// EventsClientListByTenantIDOptions contains the optional parameters for the EventsClient.NewListByTenantIDPager method.
type EventsClientListByTenantIDOptions struct {
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string

	// Specifies from when to return events, based on the lastUpdateTime property. For example, queryStartTime = 7/24/2020 OR
	// queryStartTime=7%2F24%2F2020
	QueryStartTime *string
}

// ImpactedResourcesClientGetByTenantIDOptions contains the optional parameters for the ImpactedResourcesClient.GetByTenantID
// method.
type ImpactedResourcesClientGetByTenantIDOptions struct {
	// placeholder for future optional parameters
}

// ImpactedResourcesClientGetOptions contains the optional parameters for the ImpactedResourcesClient.Get method.
type ImpactedResourcesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ImpactedResourcesClientListBySubscriptionIDAndEventIDOptions contains the optional parameters for the ImpactedResourcesClient.NewListBySubscriptionIDAndEventIDPager
// method.
type ImpactedResourcesClientListBySubscriptionIDAndEventIDOptions struct {
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}

// ImpactedResourcesClientListByTenantIDAndEventIDOptions contains the optional parameters for the ImpactedResourcesClient.NewListByTenantIDAndEventIDPager
// method.
type ImpactedResourcesClientListByTenantIDAndEventIDOptions struct {
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}

// MetadataClientGetEntityOptions contains the optional parameters for the MetadataClient.GetEntity method.
type MetadataClientGetEntityOptions struct {
	// placeholder for future optional parameters
}

// MetadataClientListOptions contains the optional parameters for the MetadataClient.NewListPager method.
type MetadataClientListOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// SecurityAdvisoryImpactedResourcesClientListBySubscriptionIDAndEventIDOptions contains the optional parameters for the SecurityAdvisoryImpactedResourcesClient.NewListBySubscriptionIDAndEventIDPager
// method.
type SecurityAdvisoryImpactedResourcesClientListBySubscriptionIDAndEventIDOptions struct {
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}

// SecurityAdvisoryImpactedResourcesClientListByTenantIDAndEventIDOptions contains the optional parameters for the SecurityAdvisoryImpactedResourcesClient.NewListByTenantIDAndEventIDPager
// method.
type SecurityAdvisoryImpactedResourcesClientListByTenantIDAndEventIDOptions struct {
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}
