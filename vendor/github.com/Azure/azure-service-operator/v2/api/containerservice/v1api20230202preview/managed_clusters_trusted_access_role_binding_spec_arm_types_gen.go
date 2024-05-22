// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230202preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type ManagedClusters_TrustedAccessRoleBinding_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: Properties for trusted access role binding
	Properties *TrustedAccessRoleBindingProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &ManagedClusters_TrustedAccessRoleBinding_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-02-02-preview"
func (binding ManagedClusters_TrustedAccessRoleBinding_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (binding *ManagedClusters_TrustedAccessRoleBinding_Spec_ARM) GetName() string {
	return binding.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerService/managedClusters/trustedAccessRoleBindings"
func (binding *ManagedClusters_TrustedAccessRoleBinding_Spec_ARM) GetType() string {
	return "Microsoft.ContainerService/managedClusters/trustedAccessRoleBindings"
}

// Properties for trusted access role binding
type TrustedAccessRoleBindingProperties_ARM struct {
	// Roles: A list of roles to bind, each item is a resource type qualified role name. For example:
	// 'Microsoft.MachineLearningServices/workspaces/reader'.
	Roles            []string `json:"roles"`
	SourceResourceId *string  `json:"sourceResourceId,omitempty"`
}
