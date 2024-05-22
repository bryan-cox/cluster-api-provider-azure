// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type RouteTables_Route_Spec_ARM struct {
	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name string `json:"name,omitempty"`

	// Properties: Properties of the route.
	Properties *RoutePropertiesFormat_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RouteTables_Route_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (route RouteTables_Route_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (route *RouteTables_Route_Spec_ARM) GetName() string {
	return route.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/routeTables/routes"
func (route *RouteTables_Route_Spec_ARM) GetType() string {
	return "Microsoft.Network/routeTables/routes"
}

// Route resource.
type RoutePropertiesFormat_ARM struct {
	// AddressPrefix: The destination CIDR to which the route applies.
	AddressPrefix *string `json:"addressPrefix,omitempty"`

	// HasBgpOverride: A value indicating whether this route overrides overlapping BGP routes regardless of LPM.
	HasBgpOverride *bool `json:"hasBgpOverride,omitempty"`

	// NextHopIpAddress: The IP address packets should be forwarded to. Next hop values are only allowed in routes where the
	// next hop type is VirtualAppliance.
	NextHopIpAddress *string `json:"nextHopIpAddress,omitempty"`

	// NextHopType: The type of Azure hop the packet should be sent to.
	NextHopType *RouteNextHopType `json:"nextHopType,omitempty"`
}
