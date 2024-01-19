// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=virtualnetworks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={virtualnetworks/status,virtualnetworks/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20201101.VirtualNetwork
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2020-11-01/virtualNetwork.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}
type VirtualNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetwork_Spec   `json:"spec,omitempty"`
	Status            VirtualNetwork_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &VirtualNetwork{}

// GetConditions returns the conditions of the resource
func (network *VirtualNetwork) GetConditions() conditions.Conditions {
	return network.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (network *VirtualNetwork) SetConditions(conditions conditions.Conditions) {
	network.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &VirtualNetwork{}

// AzureName returns the Azure name of the resource
func (network *VirtualNetwork) AzureName() string {
	return network.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (network VirtualNetwork) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (network *VirtualNetwork) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (network *VirtualNetwork) GetSpec() genruntime.ConvertibleSpec {
	return &network.Spec
}

// GetStatus returns the status of this resource
func (network *VirtualNetwork) GetStatus() genruntime.ConvertibleStatus {
	return &network.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (network *VirtualNetwork) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/virtualNetworks"
func (network *VirtualNetwork) GetType() string {
	return "Microsoft.Network/virtualNetworks"
}

// NewEmptyStatus returns a new empty (blank) status
func (network *VirtualNetwork) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &VirtualNetwork_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (network *VirtualNetwork) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(network.Spec)
	return network.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (network *VirtualNetwork) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*VirtualNetwork_STATUS); ok {
		network.Status = *st
		return nil
	}

	// Convert status to required version
	var st VirtualNetwork_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	network.Status = st
	return nil
}

// Hub marks that this VirtualNetwork is the hub type for conversion
func (network *VirtualNetwork) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (network *VirtualNetwork) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: network.Spec.OriginalVersion,
		Kind:    "VirtualNetwork",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20201101.VirtualNetwork
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2020-11-01/virtualNetwork.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}
type VirtualNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetwork `json:"items"`
}

// Storage version of v1api20201101.VirtualNetwork_Spec
type VirtualNetwork_Spec struct {
	AddressSpace *AddressSpace `json:"addressSpace,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName            string                        `json:"azureName,omitempty"`
	BgpCommunities       *VirtualNetworkBgpCommunities `json:"bgpCommunities,omitempty"`
	DdosProtectionPlan   *SubResource                  `json:"ddosProtectionPlan,omitempty"`
	DhcpOptions          *DhcpOptions                  `json:"dhcpOptions,omitempty"`
	EnableDdosProtection *bool                         `json:"enableDdosProtection,omitempty"`
	EnableVmProtection   *bool                         `json:"enableVmProtection,omitempty"`
	ExtendedLocation     *ExtendedLocation             `json:"extendedLocation,omitempty"`
	IpAllocations        []SubResource                 `json:"ipAllocations,omitempty"`
	Location             *string                       `json:"location,omitempty"`
	OriginalVersion      string                        `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &VirtualNetwork_Spec{}

// ConvertSpecFrom populates our VirtualNetwork_Spec from the provided source
func (network *VirtualNetwork_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == network {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(network)
}

// ConvertSpecTo populates the provided destination from our VirtualNetwork_Spec
func (network *VirtualNetwork_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == network {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(network)
}

// Storage version of v1api20201101.VirtualNetwork_STATUS
// Virtual Network resource.
type VirtualNetwork_STATUS struct {
	AddressSpace         *AddressSpace_STATUS                 `json:"addressSpace,omitempty"`
	BgpCommunities       *VirtualNetworkBgpCommunities_STATUS `json:"bgpCommunities,omitempty"`
	Conditions           []conditions.Condition               `json:"conditions,omitempty"`
	DdosProtectionPlan   *SubResource_STATUS                  `json:"ddosProtectionPlan,omitempty"`
	DhcpOptions          *DhcpOptions_STATUS                  `json:"dhcpOptions,omitempty"`
	EnableDdosProtection *bool                                `json:"enableDdosProtection,omitempty"`
	EnableVmProtection   *bool                                `json:"enableVmProtection,omitempty"`
	Etag                 *string                              `json:"etag,omitempty"`
	ExtendedLocation     *ExtendedLocation_STATUS             `json:"extendedLocation,omitempty"`
	Id                   *string                              `json:"id,omitempty"`
	IpAllocations        []SubResource_STATUS                 `json:"ipAllocations,omitempty"`
	Location             *string                              `json:"location,omitempty"`
	Name                 *string                              `json:"name,omitempty"`
	PropertyBag          genruntime.PropertyBag               `json:"$propertyBag,omitempty"`
	ProvisioningState    *string                              `json:"provisioningState,omitempty"`
	ResourceGuid         *string                              `json:"resourceGuid,omitempty"`
	Tags                 map[string]string                    `json:"tags,omitempty"`
	Type                 *string                              `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &VirtualNetwork_STATUS{}

// ConvertStatusFrom populates our VirtualNetwork_STATUS from the provided source
func (network *VirtualNetwork_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == network {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(network)
}

// ConvertStatusTo populates the provided destination from our VirtualNetwork_STATUS
func (network *VirtualNetwork_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == network {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(network)
}

// Storage version of v1api20201101.AddressSpace
// AddressSpace contains an array of IP address ranges that can be used by subnets of the virtual network.
type AddressSpace struct {
	AddressPrefixes []string               `json:"addressPrefixes,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20201101.AddressSpace_STATUS
// AddressSpace contains an array of IP address ranges that can be used by subnets of the virtual network.
type AddressSpace_STATUS struct {
	AddressPrefixes []string               `json:"addressPrefixes,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20201101.DhcpOptions
// DhcpOptions contains an array of DNS servers available to VMs deployed in the virtual network. Standard DHCP option for
// a subnet overrides VNET DHCP options.
type DhcpOptions struct {
	DnsServers  []string               `json:"dnsServers,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20201101.DhcpOptions_STATUS
// DhcpOptions contains an array of DNS servers available to VMs deployed in the virtual network. Standard DHCP option for
// a subnet overrides VNET DHCP options.
type DhcpOptions_STATUS struct {
	DnsServers  []string               `json:"dnsServers,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20201101.VirtualNetworkBgpCommunities
// Bgp Communities sent over ExpressRoute with each route corresponding to a prefix in this VNET.
type VirtualNetworkBgpCommunities struct {
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	VirtualNetworkCommunity *string                `json:"virtualNetworkCommunity,omitempty"`
}

// Storage version of v1api20201101.VirtualNetworkBgpCommunities_STATUS
// Bgp Communities sent over ExpressRoute with each route corresponding to a prefix in this VNET.
type VirtualNetworkBgpCommunities_STATUS struct {
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RegionalCommunity       *string                `json:"regionalCommunity,omitempty"`
	VirtualNetworkCommunity *string                `json:"virtualNetworkCommunity,omitempty"`
}

func init() {
	SchemeBuilder.Register(&VirtualNetwork{}, &VirtualNetworkList{})
}
