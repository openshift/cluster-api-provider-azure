//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcehealth

import "time"

// AvailabilityStatus - availabilityStatus of a resource.
type AvailabilityStatus struct {
	// Azure Resource Manager Identity for the availabilityStatuses resource.
	ID *string

	// Azure Resource Manager geo location of the resource.
	Location *string

	// current.
	Name *string

	// Properties of availability state.
	Properties *AvailabilityStatusProperties

	// Microsoft.ResourceHealth/AvailabilityStatuses.
	Type *string
}

// AvailabilityStatusListResult - The List availabilityStatus operation response.
type AvailabilityStatusListResult struct {
	// REQUIRED; The list of availabilityStatuses.
	Value []*AvailabilityStatus

	// The URI to fetch the next page of availabilityStatuses. Call ListNext() with this URI to fetch the next page of availabilityStatuses.
	NextLink *string
}

// AvailabilityStatusProperties - Properties of availability state.
type AvailabilityStatusProperties struct {
	// The Article Id
	ArticleID *string

	// Availability status of the resource. When it is null, this availabilityStatus object represents an availability impacting
	// event
	AvailabilityState *AvailabilityStateValues

	// When a context field is set to Platform, this field will reflect if the event was planned or unplanned. If the context
	// field does not have a value of Platform, then this field will be ignored.
	Category *string

	// When an event is created, it can either be triggered by a customer or the platform of the resource and this field will
	// illustrate that. This field is connected to the category field in this object.
	Context *string

	// Details of the availability status.
	DetailedStatus *string

	// In case of an availability impacting event, it describes the category of a PlatformInitiated health impacting event. Examples
	// are Planned, Unplanned etc.
	HealthEventCategory *string

	// In case of an availability impacting event, it describes where the health impacting event was originated. Examples are
	// PlatformInitiated, UserInitiated etc.
	HealthEventCause *string

	// It is a unique Id that identifies the event
	HealthEventID *string

	// In case of an availability impacting event, it describes when the health impacting event was originated. Examples are Lifecycle,
	// Downtime, Fault Analysis etc.
	HealthEventType *string

	// Timestamp for when last change in health status occurred.
	OccurredTime *time.Time

	// Chronicity of the availability transition.
	ReasonChronicity *ReasonChronicityTypes

	// When the resource's availabilityState is Unavailable, it describes where the health impacting event was originated. Examples
	// are planned, unplanned, user initiated or an outage etc.
	ReasonType *string

	// An annotation describing a change in the availabilityState to Available from Unavailable with a reasonType of type Unplanned
	RecentlyResolved *AvailabilityStatusPropertiesRecentlyResolved

	// Lists actions the user can take based on the current availabilityState of the resource.
	RecommendedActions []*RecommendedAction

	// Timestamp for when the health was last checked.
	ReportedTime *time.Time

	// When the resource's availabilityState is Unavailable and the reasonType is not User Initiated, it provides the date and
	// time for when the issue is expected to be resolved.
	ResolutionETA *time.Time

	// When the resource's availabilityState is Unavailable, it provides the Timestamp for when the health impacting event was
	// received.
	RootCauseAttributionTime *time.Time

	// Lists the service impacting events that may be affecting the health of the resource.
	ServiceImpactingEvents []*ServiceImpactingEvent

	// Summary description of the availability status.
	Summary *string

	// Title description of the availability status.
	Title *string
}

// AvailabilityStatusPropertiesRecentlyResolved - An annotation describing a change in the availabilityState to Available
// from Unavailable with a reasonType of type Unplanned
type AvailabilityStatusPropertiesRecentlyResolved struct {
	// Timestamp when the availabilityState changes to Available.
	ResolvedTime *time.Time

	// Brief description of cause of the resource becoming unavailable.
	UnavailabilitySummary *string

	// Timestamp for when the availabilityState changed to Unavailable
	UnavailableOccurredTime *time.Time
}

// EmergingIssue - On-going emerging issue from azure status.
type EmergingIssue struct {
	// Timestamp for when last time refreshed for ongoing emerging issue.
	RefreshTimestamp *time.Time

	// The list of emerging issues of active event type.
	StatusActiveEvents []*StatusActiveEvent

	// The list of emerging issues of banner type.
	StatusBanners []*StatusBanner
}

// EmergingIssueImpact - Object of the emerging issue impact on services and regions.
type EmergingIssueImpact struct {
	// The impacted service id.
	ID *string

	// The impacted service name.
	Name *string

	// The list of impacted regions for corresponding emerging issues.
	Regions []*ImpactedRegion
}

// EmergingIssueListResult - The list of emerging issues.
type EmergingIssueListResult struct {
	// The link used to get the next page of emerging issues.
	NextLink *string

	// The list of emerging issues.
	Value []*EmergingIssuesGetResult
}

// EmergingIssuesGetResult - The Get EmergingIssues operation response.
type EmergingIssuesGetResult struct {
	// The emerging issue entity properties.
	Properties *EmergingIssue

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// Event - Service health event
type Event struct {
	// Properties of event.
	Properties *EventProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// EventImpactedResource - Impacted resource for an event.
type EventImpactedResource struct {
	// Properties of impacted resource.
	Properties *EventImpactedResourceProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// EventImpactedResourceListResult - The List of eventImpactedResources operation response.
type EventImpactedResourceListResult struct {
	// REQUIRED; The list of eventImpactedResources.
	Value []*EventImpactedResource

	// The URI to fetch the next page of events. Call ListNext() with this URI to fetch the next page of impacted resource.
	NextLink *string
}

// EventImpactedResourceProperties - Properties of impacted resource.
type EventImpactedResourceProperties struct {
	// Additional information.
	Info []*KeyValueItem

	// READ-ONLY; Impacted resource region name.
	TargetRegion *string

	// READ-ONLY; Identity for resource within Microsoft cloud.
	TargetResourceID *string

	// READ-ONLY; Resource type within Microsoft cloud.
	TargetResourceType *string
}

// EventProperties - Properties of event.
type EventProperties struct {
	// Additional information
	AdditionalInformation *EventPropertiesAdditionalInformation

	// Article of event.
	Article *EventPropertiesArticle

	// Contains the communication message for the event, that could include summary, root cause and other details.
	Description *string

	// duration in seconds
	Duration *int32

	// Tells if we want to enable or disable Microsoft Support for this event.
	EnableChatWithUs *bool

	// Tells if we want to enable or disable Microsoft Support for this event.
	EnableMicrosoftSupport *bool

	// Level of event.
	EventLevel *EventLevelValues

	// Source of event.
	EventSource *EventSourceValues

	// Type of event.
	EventType *EventTypeValues

	// The id of the Incident
	ExternalIncidentID *string

	// Frequently asked questions for the service health event.
	Faqs []*Faq

	// Header text of event.
	Header *string

	// Stage for HIR Document
	HirStage *string

	// List services impacted by the service health event.
	Impact []*Impact

	// It provides the Timestamp for when the health impacting event resolved.
	ImpactMitigationTime *time.Time

	// It provides the Timestamp for when the health impacting event started.
	ImpactStartTime *time.Time

	// The type of the impact
	ImpactType *string

	// It provides information if the event is High incident rate event or not.
	IsHIR *bool

	// It provides the Timestamp for when the health impacting event was last updated.
	LastUpdateTime *time.Time

	// Level of insight.
	Level *LevelValues

	// Useful links of event.
	Links []*Link

	// Is true if the event is platform initiated.
	PlatformInitiated *bool

	// Priority level of the event. Has value from 0 to 23. 0 is the highest priority. Service issue events have higher priority
	// followed by planned maintenance and health advisory. Critical events have
	// higher priority followed by error, warning and informational. Furthermore, active events have higher priority than resolved.
	Priority *int32

	// The reason for the Incident
	Reason *string

	// Recommended actions of event.
	RecommendedActions *EventPropertiesRecommendedActions

	// Current status of event.
	Status *EventStatusValues

	// Summary text of event.
	Summary *string

	// Title text of event.
	Title *string
}

// EventPropertiesAdditionalInformation - Additional information
type EventPropertiesAdditionalInformation struct {
	// Additional information Message
	Message *string
}

// EventPropertiesArticle - Article of event.
type EventPropertiesArticle struct {
	// Article content of event.
	ArticleContent *string

	// Article Id
	ArticleID *string

	// It provides a map of parameter name and value
	Parameters any
}

// EventPropertiesRecommendedActions - Recommended actions of event.
type EventPropertiesRecommendedActions struct {
	// Recommended actions for the service health event.
	Actions []*EventPropertiesRecommendedActionsItem

	// Recommended action locale for the service health event.
	LocaleCode *string

	// Recommended action title for the service health event.
	Message *string
}

// EventPropertiesRecommendedActionsItem - Recommended action for the service health event.
type EventPropertiesRecommendedActionsItem struct {
	// Recommended action text
	ActionText *string

	// Recommended action group Id for the service health event.
	GroupID *int32
}

// Events - The List events operation response.
type Events struct {
	// REQUIRED; The list of event.
	Value []*Event

	// The URI to fetch the next page of events. Call ListNext() with this URI to fetch the next page of events.
	NextLink *string
}

// Faq - Frequently asked question for the service health event
type Faq struct {
	// FAQ answer for the service health event.
	Answer *string

	// FAQ locale for the service health event.
	LocaleCode *string

	// FAQ question for the service health event.
	Question *string
}

// Impact - Azure service impacted by the service health event.
type Impact struct {
	// List regions impacted by the service health event.
	ImpactedRegions []*ImpactedServiceRegion

	// Impacted service name.
	ImpactedService *string
}

// ImpactedRegion - Object of impacted region.
type ImpactedRegion struct {
	// The impacted region id.
	ID *string

	// The impacted region name.
	Name *string
}

// ImpactedServiceRegion - Azure region impacted by the service health event.
type ImpactedServiceRegion struct {
	// Impacted region name.
	ImpactedRegion *string

	// List subscription impacted by the service health event.
	ImpactedSubscriptions []*string

	// List tenant impacted by the service health event.
	ImpactedTenants []*string

	// It provides the Timestamp for when the last update for the service health event.
	LastUpdateTime *time.Time

	// Current status of event in the region.
	Status *EventStatusValues

	// List of updates for given service health event.
	Updates []*Update
}

// KeyValueItem - Key value tuple.
type KeyValueItem struct {
	// READ-ONLY; Key of tuple.
	Key *string

	// READ-ONLY; Value of tuple.
	Value *string
}

// Link - Useful links for service health event.
type Link struct {
	// It provides the name of portal extension blade to produce link for given service health event.
	BladeName *string

	// Display text of link.
	DisplayText *LinkDisplayText

	// It provides the name of portal extension to produce link for given service health event.
	ExtensionName *string

	// It provides a map of parameter name and value for portal extension blade to produce lik for given service health event.
	Parameters any

	// Type of link.
	Type *LinkTypeValues
}

// LinkDisplayText - Display text of link.
type LinkDisplayText struct {
	// Localized display text of link.
	LocalizedValue *string

	// Display text of link.
	Value *string
}

// MetadataEntity - The metadata entity contract.
type MetadataEntity struct {
	// The metadata entity properties.
	Properties *MetadataEntityProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// MetadataEntityListResult - The list of metadata entities
type MetadataEntityListResult struct {
	// The link used to get the next page of metadata.
	NextLink *string

	// The list of metadata entities.
	Value []*MetadataEntity
}

// MetadataEntityProperties - The metadata entity properties
type MetadataEntityProperties struct {
	// The list of scenarios applicable to this metadata entity.
	ApplicableScenarios []*Scenario

	// The list of keys on which this entity depends on.
	DependsOn []*string

	// The display name.
	DisplayName *string

	// The list of supported values.
	SupportedValues []*MetadataSupportedValueDetail
}

// MetadataSupportedValueDetail - The metadata supported value detail.
type MetadataSupportedValueDetail struct {
	// The display name.
	DisplayName *string

	// The id.
	ID *string

	// The list of associated resource types.
	ResourceTypes []*string
}

// Operation available in the Microsoft.ResourceHealth resource provider.
type Operation struct {
	// Properties of the operation.
	Display *OperationDisplay

	// Name of the operation.
	Name *string
}

// OperationDisplay - Properties of the operation.
type OperationDisplay struct {
	// Description of the operation.
	Description *string

	// Operation name.
	Operation *string

	// Provider name.
	Provider *string

	// Resource name.
	Resource *string
}

// OperationListResult - Lists the operations response.
type OperationListResult struct {
	// REQUIRED; List of operations available in the Microsoft.ResourceHealth resource provider.
	Value []*Operation
}

// RecommendedAction - Lists actions the user can take based on the current availabilityState of the resource.
type RecommendedAction struct {
	// Recommended action.
	Action *string

	// Link to the action
	ActionURL *string

	// the comment for the Action
	ActionURLComment *string

	// Substring of action, it describes which text should host the action URL.
	ActionURLText *string
}

// ServiceImpactingEvent - Lists the service impacting events that may be affecting the health of the resource.
type ServiceImpactingEvent struct {
	// Correlation id for the event
	CorrelationID *string

	// Timestamp for when the event started.
	EventStartTime *time.Time

	// Timestamp for when event was submitted/detected.
	EventStatusLastModifiedTime *time.Time

	// Properties of the service impacting event.
	IncidentProperties *ServiceImpactingEventIncidentProperties

	// Status of the service impacting event.
	Status *ServiceImpactingEventStatus
}

// ServiceImpactingEventIncidentProperties - Properties of the service impacting event.
type ServiceImpactingEventIncidentProperties struct {
	// Type of Event.
	IncidentType *string

	// Region impacted by the event.
	Region *string

	// Service impacted by the event.
	Service *string

	// Title of the incident.
	Title *string
}

// ServiceImpactingEventStatus - Status of the service impacting event.
type ServiceImpactingEventStatus struct {
	// Current status of the event
	Value *string
}

// StatusActiveEvent - Active event type of emerging issue.
type StatusActiveEvent struct {
	// The cloud type of this active event.
	Cloud *string

	// The details of active event.
	Description *string

	// The list of emerging issues impacts.
	Impacts []*EmergingIssueImpact

	// The last time modified on this banner.
	LastModifiedTime *time.Time

	// The boolean value of this active event if published or not.
	Published *bool

	// The severity level of this active event.
	Severity *SeverityValues

	// The stage of this active event.
	Stage *StageValues

	// The impact start time on this active event.
	StartTime *time.Time

	// The active event title.
	Title *string

	// The tracking id of this active event.
	TrackingID *string
}

// StatusBanner - Banner type of emerging issue.
type StatusBanner struct {
	// The cloud type of this banner.
	Cloud *string

	// The last time modified on this banner.
	LastModifiedTime *time.Time

	// The details of banner.
	Message *string

	// The banner title.
	Title *string
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// Update for service health event.
type Update struct {
	// Summary text for the given update for the service health event.
	Summary *string

	// It provides the Timestamp for the given update for the service health event.
	UpdateDateTime *time.Time
}
