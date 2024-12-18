// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

// Nat Gateway resource.
type NatGateway_STATUS_ARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Nat Gateway properties.
	Properties *NatGatewayPropertiesFormat_STATUS_ARM `json:"properties,omitempty"`

	// Sku: The nat gateway SKU.
	Sku *NatGatewaySku_STATUS_ARM `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`

	// Zones: A list of availability zones denoting the zone in which Nat Gateway should be deployed.
	Zones []string `json:"zones,omitempty"`
}

// Nat Gateway properties.
type NatGatewayPropertiesFormat_STATUS_ARM struct {
	// IdleTimeoutInMinutes: The idle timeout of the nat gateway.
	IdleTimeoutInMinutes *int `json:"idleTimeoutInMinutes,omitempty"`

	// ProvisioningState: The provisioning state of the NAT gateway resource.
	ProvisioningState *ApplicationGatewayProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// PublicIpAddresses: An array of public ip addresses associated with the nat gateway resource.
	PublicIpAddresses []ApplicationGatewaySubResource_STATUS_ARM `json:"publicIpAddresses,omitempty"`

	// PublicIpPrefixes: An array of public ip prefixes associated with the nat gateway resource.
	PublicIpPrefixes []ApplicationGatewaySubResource_STATUS_ARM `json:"publicIpPrefixes,omitempty"`

	// ResourceGuid: The resource GUID property of the NAT gateway resource.
	ResourceGuid *string `json:"resourceGuid,omitempty"`

	// Subnets: An array of references to the subnets using this nat gateway resource.
	Subnets []ApplicationGatewaySubResource_STATUS_ARM `json:"subnets,omitempty"`
}

// SKU of nat gateway.
type NatGatewaySku_STATUS_ARM struct {
	// Name: Name of Nat Gateway SKU.
	Name *NatGatewaySku_Name_STATUS `json:"name,omitempty"`
}

type NatGatewaySku_Name_STATUS string

const NatGatewaySku_Name_STATUS_Standard = NatGatewaySku_Name_STATUS("Standard")

// Mapping from string to NatGatewaySku_Name_STATUS
var natGatewaySku_Name_STATUS_Values = map[string]NatGatewaySku_Name_STATUS{
	"standard": NatGatewaySku_Name_STATUS_Standard,
}