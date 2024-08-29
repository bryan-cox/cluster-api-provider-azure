/*
Copyright 2019 The Kubernetes Authors.

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
	asonetworkv1 "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	infrav1 "sigs.k8s.io/cluster-api-provider-azure/api/v1beta1"
	"sigs.k8s.io/cluster-api-provider-azure/azure"
	"sigs.k8s.io/cluster-api-provider-azure/azure/services/aso"
	"sigs.k8s.io/cluster-api-provider-azure/azure/services/async"
)

const serviceName = "publicips"

// PublicIPScope defines the scope interface for a public IP service.
type PublicIPScope interface {
	azure.Authorizer
	aso.Scope
	azure.ClusterDescriber
	PublicIPSpecs() []azure.ASOResourceSpecGetter[*asonetworkv1.PublicIPAddress]
}

// Service provides operations on Azure resources.
type Service struct {
	Scope PublicIPScope
	*aso.Service[*asonetworkv1.PublicIPAddress, PublicIPScope]
	async.Getter
	async.TagsGetter
}

// New creates a new service.
func New(scope PublicIPScope) *Service {
	svc := aso.NewService[*asonetworkv1.PublicIPAddress, PublicIPScope](serviceName, scope)
	svc.Specs = scope.PublicIPSpecs()
	svc.ConditionType = infrav1.PublicIPsReadyCondition

	return &Service{
		Scope:   scope,
		Service: svc,
	}
}
