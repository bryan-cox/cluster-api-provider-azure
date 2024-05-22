// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

// Bastion Host resource.
type BastionHost_STATUS_ARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Represents the bastion host resource.
	Properties *BastionHostPropertiesFormat_STATUS_ARM `json:"properties,omitempty"`

	// Sku: The sku of this Bastion Host.
	Sku *Sku_STATUS_ARM `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// Properties of the Bastion Host.
type BastionHostPropertiesFormat_STATUS_ARM struct {
	// DisableCopyPaste: Enable/Disable Copy/Paste feature of the Bastion Host resource.
	DisableCopyPaste *bool `json:"disableCopyPaste,omitempty"`

	// DnsName: FQDN for the endpoint on which bastion host is accessible.
	DnsName *string `json:"dnsName,omitempty"`

	// EnableFileCopy: Enable/Disable File Copy feature of the Bastion Host resource.
	EnableFileCopy *bool `json:"enableFileCopy,omitempty"`

	// EnableIpConnect: Enable/Disable IP Connect feature of the Bastion Host resource.
	EnableIpConnect *bool `json:"enableIpConnect,omitempty"`

	// EnableShareableLink: Enable/Disable Shareable Link of the Bastion Host resource.
	EnableShareableLink *bool `json:"enableShareableLink,omitempty"`

	// EnableTunneling: Enable/Disable Tunneling feature of the Bastion Host resource.
	EnableTunneling *bool `json:"enableTunneling,omitempty"`

	// IpConfigurations: IP configuration of the Bastion Host resource.
	IpConfigurations []BastionHostIPConfiguration_STATUS_ARM `json:"ipConfigurations,omitempty"`

	// ProvisioningState: The provisioning state of the bastion host resource.
	ProvisioningState *BastionHostProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// ScaleUnits: The scale units for the Bastion Host resource.
	ScaleUnits *int `json:"scaleUnits,omitempty"`
}

// The sku of this Bastion Host.
type Sku_STATUS_ARM struct {
	// Name: The name of this Bastion Host.
	Name *Sku_Name_STATUS `json:"name,omitempty"`
}

// IP configuration of an Bastion Host.
type BastionHostIPConfiguration_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type Sku_Name_STATUS string

const (
	Sku_Name_STATUS_Basic    = Sku_Name_STATUS("Basic")
	Sku_Name_STATUS_Standard = Sku_Name_STATUS("Standard")
)
