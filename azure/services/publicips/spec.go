/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package publicips

import (
	"context"
	"strings"

	asonetworkv1 "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
	infrav1 "sigs.k8s.io/cluster-api-provider-azure/api/v1beta1"
	"sigs.k8s.io/cluster-api-provider-azure/azure/converters"
)

// PublicIPSpec defines the specification for a Public IP.
type PublicIPSpec struct {
	Name             string
	Namespace        string
	ResourceGroup    string
	ClusterName      string
	DNSName          string
	IsIPv6           bool
	Location         string
	ExtendedLocation *infrav1.ExtendedLocationSpec
	FailureDomains   []*string
	AdditionalTags   infrav1.Tags
	IPTags           []infrav1.IPTag
}

// ResourceRef implements azure.ASOResourceSpecGetter.
func (s *PublicIPSpec) ResourceRef() *asonetworkv1.PublicIPAddress {
	return &asonetworkv1.PublicIPAddress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      s.Name,
			Namespace: s.Namespace,
		},
	}
}

// Parameters implements azure.ASOResourceSpecGetter.
func (s *PublicIPSpec) Parameters(_ context.Context, existingPublicIPAddress *asonetworkv1.PublicIPAddress) (params *asonetworkv1.PublicIPAddress, err error) {
	publicIPAddress := &asonetworkv1.PublicIPAddress{}
	publicIPAddress.Spec = asonetworkv1.PublicIPAddress_Spec{}

	if existingPublicIPAddress != nil {
		publicIPAddress = existingPublicIPAddress
	}

	publicIPAddress.Spec.PublicIPAddressVersion = ptr.To(asonetworkv1.IPVersion_IPv4)
	if s.IsIPv6 {
		publicIPAddress.Spec.PublicIPAddressVersion = ptr.To(asonetworkv1.IPVersion_IPv6)
	}

	// only set DNS properties if there is a DNS name specified
	var dnsSettings *asonetworkv1.PublicIPAddressDnsSettings
	if s.DNSName != "" {
		dnsSettings = &asonetworkv1.PublicIPAddressDnsSettings{
			DomainNameLabel: ptr.To(strings.Split(s.DNSName, ".")[0]),
			Fqdn:            ptr.To(s.DNSName),
		}
		publicIPAddress.Spec.DnsSettings = dnsSettings
	}

	publicIPAddress.Spec.Tags = map[string]string{
		"ClusterName": s.ClusterName,
		"Lifecycle":   string(infrav1.ResourceLifecycleOwned),
		"Name":        s.Name,
	}
	for k, v := range s.AdditionalTags {
		publicIPAddress.Spec.Tags[k] = v
	}

	publicIPAddress.Spec.Sku = &asonetworkv1.PublicIPAddressSku{Name: ptr.To(asonetworkv1.PublicIPAddressSku_Name_Standard)}
	publicIPAddress.Spec.AzureName = s.Name
	publicIPAddress.Spec.Location = ptr.To(s.Location)
	publicIPAddress.Spec.ExtendedLocation = converters.ExtendedLocationToNetworkASO(s.ExtendedLocation)
	publicIPAddress.Spec.PublicIPAllocationMethod = ptr.To(asonetworkv1.IPAllocationMethod_Dynamic)
	publicIPAddress.Spec.IpTags = converters.IpTagsToASO(s.IPTags)
	publicIPAddress.Spec.Zones = converters.StringsPtrsToSlice(s.FailureDomains)

	return publicIPAddress, nil
}

// WasManaged implements azure.ASOResourceSpecGetter.
func (s *PublicIPSpec) WasManaged(_ *asonetworkv1.PublicIPAddress) bool {
	return false
}
