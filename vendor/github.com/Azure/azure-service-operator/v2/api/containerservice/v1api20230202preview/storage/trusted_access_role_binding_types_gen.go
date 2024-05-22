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

// +kubebuilder:rbac:groups=containerservice.azure.com,resources=trustedaccessrolebindings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=containerservice.azure.com,resources={trustedaccessrolebindings/status,trustedaccessrolebindings/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230202preview.TrustedAccessRoleBinding
// Generator information:
// - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-02-02-preview/managedClusters.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/trustedAccessRoleBindings/{trustedAccessRoleBindingName}
type TrustedAccessRoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedClusters_TrustedAccessRoleBinding_Spec   `json:"spec,omitempty"`
	Status            ManagedClusters_TrustedAccessRoleBinding_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &TrustedAccessRoleBinding{}

// GetConditions returns the conditions of the resource
func (binding *TrustedAccessRoleBinding) GetConditions() conditions.Conditions {
	return binding.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (binding *TrustedAccessRoleBinding) SetConditions(conditions conditions.Conditions) {
	binding.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &TrustedAccessRoleBinding{}

// AzureName returns the Azure name of the resource
func (binding *TrustedAccessRoleBinding) AzureName() string {
	return binding.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-02-02-preview"
func (binding TrustedAccessRoleBinding) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (binding *TrustedAccessRoleBinding) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (binding *TrustedAccessRoleBinding) GetSpec() genruntime.ConvertibleSpec {
	return &binding.Spec
}

// GetStatus returns the status of this resource
func (binding *TrustedAccessRoleBinding) GetStatus() genruntime.ConvertibleStatus {
	return &binding.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (binding *TrustedAccessRoleBinding) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerService/managedClusters/trustedAccessRoleBindings"
func (binding *TrustedAccessRoleBinding) GetType() string {
	return "Microsoft.ContainerService/managedClusters/trustedAccessRoleBindings"
}

// NewEmptyStatus returns a new empty (blank) status
func (binding *TrustedAccessRoleBinding) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ManagedClusters_TrustedAccessRoleBinding_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (binding *TrustedAccessRoleBinding) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(binding.Spec)
	return binding.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (binding *TrustedAccessRoleBinding) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ManagedClusters_TrustedAccessRoleBinding_STATUS); ok {
		binding.Status = *st
		return nil
	}

	// Convert status to required version
	var st ManagedClusters_TrustedAccessRoleBinding_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	binding.Status = st
	return nil
}

// Hub marks that this TrustedAccessRoleBinding is the hub type for conversion
func (binding *TrustedAccessRoleBinding) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (binding *TrustedAccessRoleBinding) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: binding.Spec.OriginalVersion,
		Kind:    "TrustedAccessRoleBinding",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230202preview.TrustedAccessRoleBinding
// Generator information:
// - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-02-02-preview/managedClusters.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/trustedAccessRoleBindings/{trustedAccessRoleBindingName}
type TrustedAccessRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrustedAccessRoleBinding `json:"items"`
}

// Storage version of v1api20230202preview.ManagedClusters_TrustedAccessRoleBinding_Spec
type ManagedClusters_TrustedAccessRoleBinding_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string `json:"azureName,omitempty"`
	OriginalVersion string `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a containerservice.azure.com/ManagedCluster resource
	Owner       *genruntime.KnownResourceReference `group:"containerservice.azure.com" json:"owner,omitempty" kind:"ManagedCluster"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Roles       []string                           `json:"roles,omitempty"`

	// +kubebuilder:validation:Required
	// SourceResourceReference: The ARM resource ID of source resource that trusted access is configured for.
	SourceResourceReference *genruntime.ResourceReference `armReference:"SourceResourceId" json:"sourceResourceReference,omitempty"`
}

var _ genruntime.ConvertibleSpec = &ManagedClusters_TrustedAccessRoleBinding_Spec{}

// ConvertSpecFrom populates our ManagedClusters_TrustedAccessRoleBinding_Spec from the provided source
func (binding *ManagedClusters_TrustedAccessRoleBinding_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == binding {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(binding)
}

// ConvertSpecTo populates the provided destination from our ManagedClusters_TrustedAccessRoleBinding_Spec
func (binding *ManagedClusters_TrustedAccessRoleBinding_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == binding {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(binding)
}

// Storage version of v1api20230202preview.ManagedClusters_TrustedAccessRoleBinding_STATUS
type ManagedClusters_TrustedAccessRoleBinding_STATUS struct {
	Conditions        []conditions.Condition `json:"conditions,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Name              *string                `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	Roles             []string               `json:"roles,omitempty"`
	SourceResourceId  *string                `json:"sourceResourceId,omitempty"`
	SystemData        *SystemData_STATUS     `json:"systemData,omitempty"`
	Type              *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &ManagedClusters_TrustedAccessRoleBinding_STATUS{}

// ConvertStatusFrom populates our ManagedClusters_TrustedAccessRoleBinding_STATUS from the provided source
func (binding *ManagedClusters_TrustedAccessRoleBinding_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == binding {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(binding)
}

// ConvertStatusTo populates the provided destination from our ManagedClusters_TrustedAccessRoleBinding_STATUS
func (binding *ManagedClusters_TrustedAccessRoleBinding_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == binding {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(binding)
}

func init() {
	SchemeBuilder.Register(&TrustedAccessRoleBinding{}, &TrustedAccessRoleBindingList{})
}
