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

package converters

import (
	asonetworkv1 "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	"k8s.io/utils/ptr"
	infrav1 "sigs.k8s.io/cluster-api-provider-azure/api/v1beta1"
)

// IpTagsToASO converts a CAPZ IP tag to an asonetworkv1.IpTag.
func IpTagsToASO(ipTags []infrav1.IPTag) []asonetworkv1.IpTag {
	if len(ipTags) == 0 {
		return nil
	}
	skdIPTags := make([]asonetworkv1.IpTag, len(ipTags))
	for i, ipTag := range ipTags {
		skdIPTags[i] = asonetworkv1.IpTag{
			IpTagType: ptr.To(ipTag.Type),
			Tag:       ptr.To(ipTag.Tag),
		}
	}
	return skdIPTags
}

func StringsPtrsToSlice(strings []*string) []string {
	result := make([]string, len(strings))
	for i, s := range strings {
		if s != nil {
			result[i] = *s
		}
	}
	return result
}
