/*
Copyright The Kubernetes Authors.

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

// Code generated by MockGen. DO NOT EDIT.
// Source: ../loadbalancers.go

// Package mock_loadbalancers is a generated GoMock package.
package mock_loadbalancers

import (
	reflect "reflect"

	autorest "github.com/Azure/go-autorest/autorest"
	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
	v1beta1 "sigs.k8s.io/cluster-api-provider-azure/api/v1beta1"
	azure "sigs.k8s.io/cluster-api-provider-azure/azure"
)

// MockLBScope is a mock of LBScope interface.
type MockLBScope struct {
	ctrl     *gomock.Controller
	recorder *MockLBScopeMockRecorder
}

// MockLBScopeMockRecorder is the mock recorder for MockLBScope.
type MockLBScopeMockRecorder struct {
	mock *MockLBScope
}

// NewMockLBScope creates a new mock instance.
func NewMockLBScope(ctrl *gomock.Controller) *MockLBScope {
	mock := &MockLBScope{ctrl: ctrl}
	mock.recorder = &MockLBScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLBScope) EXPECT() *MockLBScopeMockRecorder {
	return m.recorder
}

// APIServerLBName mocks base method.
func (m *MockLBScope) APIServerLBName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIServerLBName")
	ret0, _ := ret[0].(string)
	return ret0
}

// APIServerLBName indicates an expected call of APIServerLBName.
func (mr *MockLBScopeMockRecorder) APIServerLBName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIServerLBName", reflect.TypeOf((*MockLBScope)(nil).APIServerLBName))
}

// APIServerLBPoolName mocks base method.
func (m *MockLBScope) APIServerLBPoolName(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIServerLBPoolName", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// APIServerLBPoolName indicates an expected call of APIServerLBPoolName.
func (mr *MockLBScopeMockRecorder) APIServerLBPoolName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIServerLBPoolName", reflect.TypeOf((*MockLBScope)(nil).APIServerLBPoolName), arg0)
}

// AdditionalTags mocks base method.
func (m *MockLBScope) AdditionalTags() v1beta1.Tags {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdditionalTags")
	ret0, _ := ret[0].(v1beta1.Tags)
	return ret0
}

// AdditionalTags indicates an expected call of AdditionalTags.
func (mr *MockLBScopeMockRecorder) AdditionalTags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdditionalTags", reflect.TypeOf((*MockLBScope)(nil).AdditionalTags))
}

// Authorizer mocks base method.
func (m *MockLBScope) Authorizer() autorest.Authorizer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorizer")
	ret0, _ := ret[0].(autorest.Authorizer)
	return ret0
}

// Authorizer indicates an expected call of Authorizer.
func (mr *MockLBScopeMockRecorder) Authorizer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorizer", reflect.TypeOf((*MockLBScope)(nil).Authorizer))
}

// AvailabilitySetEnabled mocks base method.
func (m *MockLBScope) AvailabilitySetEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailabilitySetEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AvailabilitySetEnabled indicates an expected call of AvailabilitySetEnabled.
func (mr *MockLBScopeMockRecorder) AvailabilitySetEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailabilitySetEnabled", reflect.TypeOf((*MockLBScope)(nil).AvailabilitySetEnabled))
}

// BaseURI mocks base method.
func (m *MockLBScope) BaseURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// BaseURI indicates an expected call of BaseURI.
func (mr *MockLBScopeMockRecorder) BaseURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseURI", reflect.TypeOf((*MockLBScope)(nil).BaseURI))
}

// ClientID mocks base method.
func (m *MockLBScope) ClientID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientID indicates an expected call of ClientID.
func (mr *MockLBScopeMockRecorder) ClientID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockLBScope)(nil).ClientID))
}

// ClientSecret mocks base method.
func (m *MockLBScope) ClientSecret() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientSecret")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientSecret indicates an expected call of ClientSecret.
func (mr *MockLBScopeMockRecorder) ClientSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientSecret", reflect.TypeOf((*MockLBScope)(nil).ClientSecret))
}

// CloudEnvironment mocks base method.
func (m *MockLBScope) CloudEnvironment() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudEnvironment")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudEnvironment indicates an expected call of CloudEnvironment.
func (mr *MockLBScopeMockRecorder) CloudEnvironment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudEnvironment", reflect.TypeOf((*MockLBScope)(nil).CloudEnvironment))
}

// CloudProviderConfigOverrides mocks base method.
func (m *MockLBScope) CloudProviderConfigOverrides() *v1beta1.CloudProviderConfigOverrides {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudProviderConfigOverrides")
	ret0, _ := ret[0].(*v1beta1.CloudProviderConfigOverrides)
	return ret0
}

// CloudProviderConfigOverrides indicates an expected call of CloudProviderConfigOverrides.
func (mr *MockLBScopeMockRecorder) CloudProviderConfigOverrides() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudProviderConfigOverrides", reflect.TypeOf((*MockLBScope)(nil).CloudProviderConfigOverrides))
}

// ClusterName mocks base method.
func (m *MockLBScope) ClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClusterName indicates an expected call of ClusterName.
func (mr *MockLBScopeMockRecorder) ClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterName", reflect.TypeOf((*MockLBScope)(nil).ClusterName))
}

// ControlPlaneRouteTable mocks base method.
func (m *MockLBScope) ControlPlaneRouteTable() v1beta1.RouteTable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControlPlaneRouteTable")
	ret0, _ := ret[0].(v1beta1.RouteTable)
	return ret0
}

// ControlPlaneRouteTable indicates an expected call of ControlPlaneRouteTable.
func (mr *MockLBScopeMockRecorder) ControlPlaneRouteTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControlPlaneRouteTable", reflect.TypeOf((*MockLBScope)(nil).ControlPlaneRouteTable))
}

// ControlPlaneSubnet mocks base method.
func (m *MockLBScope) ControlPlaneSubnet() v1beta1.SubnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControlPlaneSubnet")
	ret0, _ := ret[0].(v1beta1.SubnetSpec)
	return ret0
}

// ControlPlaneSubnet indicates an expected call of ControlPlaneSubnet.
func (mr *MockLBScopeMockRecorder) ControlPlaneSubnet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControlPlaneSubnet", reflect.TypeOf((*MockLBScope)(nil).ControlPlaneSubnet))
}

// Enabled mocks base method.
func (m *MockLBScope) Enabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Enabled indicates an expected call of Enabled.
func (mr *MockLBScopeMockRecorder) Enabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enabled", reflect.TypeOf((*MockLBScope)(nil).Enabled))
}

// Error mocks base method.
func (m *MockLBScope) Error(err error, msg string, keysAndValues ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{err, msg}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockLBScopeMockRecorder) Error(err, msg interface{}, keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{err, msg}, keysAndValues...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLBScope)(nil).Error), varargs...)
}

// FailureDomains mocks base method.
func (m *MockLBScope) FailureDomains() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FailureDomains")
	ret0, _ := ret[0].([]string)
	return ret0
}

// FailureDomains indicates an expected call of FailureDomains.
func (mr *MockLBScopeMockRecorder) FailureDomains() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailureDomains", reflect.TypeOf((*MockLBScope)(nil).FailureDomains))
}

// GetPrivateDNSZoneName mocks base method.
func (m *MockLBScope) GetPrivateDNSZoneName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivateDNSZoneName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPrivateDNSZoneName indicates an expected call of GetPrivateDNSZoneName.
func (mr *MockLBScopeMockRecorder) GetPrivateDNSZoneName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateDNSZoneName", reflect.TypeOf((*MockLBScope)(nil).GetPrivateDNSZoneName))
}

// HashKey mocks base method.
func (m *MockLBScope) HashKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// HashKey indicates an expected call of HashKey.
func (mr *MockLBScopeMockRecorder) HashKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashKey", reflect.TypeOf((*MockLBScope)(nil).HashKey))
}

// Info mocks base method.
func (m *MockLBScope) Info(msg string, keysAndValues ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{msg}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info.
func (mr *MockLBScopeMockRecorder) Info(msg interface{}, keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{msg}, keysAndValues...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLBScope)(nil).Info), varargs...)
}

// IsAPIServerPrivate mocks base method.
func (m *MockLBScope) IsAPIServerPrivate() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAPIServerPrivate")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAPIServerPrivate indicates an expected call of IsAPIServerPrivate.
func (mr *MockLBScopeMockRecorder) IsAPIServerPrivate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAPIServerPrivate", reflect.TypeOf((*MockLBScope)(nil).IsAPIServerPrivate))
}

// IsIPv6Enabled mocks base method.
func (m *MockLBScope) IsIPv6Enabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsIPv6Enabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsIPv6Enabled indicates an expected call of IsIPv6Enabled.
func (mr *MockLBScopeMockRecorder) IsIPv6Enabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsIPv6Enabled", reflect.TypeOf((*MockLBScope)(nil).IsIPv6Enabled))
}

// IsVnetManaged mocks base method.
func (m *MockLBScope) IsVnetManaged() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsVnetManaged")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsVnetManaged indicates an expected call of IsVnetManaged.
func (mr *MockLBScopeMockRecorder) IsVnetManaged() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsVnetManaged", reflect.TypeOf((*MockLBScope)(nil).IsVnetManaged))
}

// LBSpecs mocks base method.
func (m *MockLBScope) LBSpecs() []azure.LBSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LBSpecs")
	ret0, _ := ret[0].([]azure.LBSpec)
	return ret0
}

// LBSpecs indicates an expected call of LBSpecs.
func (mr *MockLBScopeMockRecorder) LBSpecs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LBSpecs", reflect.TypeOf((*MockLBScope)(nil).LBSpecs))
}

// Location mocks base method.
func (m *MockLBScope) Location() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(string)
	return ret0
}

// Location indicates an expected call of Location.
func (mr *MockLBScopeMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockLBScope)(nil).Location))
}

// NodeSubnets mocks base method.
func (m *MockLBScope) NodeSubnets() []v1beta1.SubnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeSubnets")
	ret0, _ := ret[0].([]v1beta1.SubnetSpec)
	return ret0
}

// NodeSubnets indicates an expected call of NodeSubnets.
func (mr *MockLBScopeMockRecorder) NodeSubnets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeSubnets", reflect.TypeOf((*MockLBScope)(nil).NodeSubnets))
}

// OutboundLBName mocks base method.
func (m *MockLBScope) OutboundLBName(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OutboundLBName", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// OutboundLBName indicates an expected call of OutboundLBName.
func (mr *MockLBScopeMockRecorder) OutboundLBName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OutboundLBName", reflect.TypeOf((*MockLBScope)(nil).OutboundLBName), arg0)
}

// OutboundPoolName mocks base method.
func (m *MockLBScope) OutboundPoolName(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OutboundPoolName", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// OutboundPoolName indicates an expected call of OutboundPoolName.
func (mr *MockLBScopeMockRecorder) OutboundPoolName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OutboundPoolName", reflect.TypeOf((*MockLBScope)(nil).OutboundPoolName), arg0)
}

// ResourceGroup mocks base method.
func (m *MockLBScope) ResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// ResourceGroup indicates an expected call of ResourceGroup.
func (mr *MockLBScopeMockRecorder) ResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceGroup", reflect.TypeOf((*MockLBScope)(nil).ResourceGroup))
}

// SetSubnet mocks base method.
func (m *MockLBScope) SetSubnet(arg0 v1beta1.SubnetSpec) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSubnet", arg0)
}

// SetSubnet indicates an expected call of SetSubnet.
func (mr *MockLBScopeMockRecorder) SetSubnet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSubnet", reflect.TypeOf((*MockLBScope)(nil).SetSubnet), arg0)
}

// Subnet mocks base method.
func (m *MockLBScope) Subnet(arg0 string) v1beta1.SubnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subnet", arg0)
	ret0, _ := ret[0].(v1beta1.SubnetSpec)
	return ret0
}

// Subnet indicates an expected call of Subnet.
func (mr *MockLBScopeMockRecorder) Subnet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subnet", reflect.TypeOf((*MockLBScope)(nil).Subnet), arg0)
}

// Subnets mocks base method.
func (m *MockLBScope) Subnets() v1beta1.Subnets {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subnets")
	ret0, _ := ret[0].(v1beta1.Subnets)
	return ret0
}

// Subnets indicates an expected call of Subnets.
func (mr *MockLBScopeMockRecorder) Subnets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subnets", reflect.TypeOf((*MockLBScope)(nil).Subnets))
}

// SubscriptionID mocks base method.
func (m *MockLBScope) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID.
func (mr *MockLBScopeMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockLBScope)(nil).SubscriptionID))
}

// TenantID mocks base method.
func (m *MockLBScope) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID.
func (mr *MockLBScopeMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockLBScope)(nil).TenantID))
}

// V mocks base method.
func (m *MockLBScope) V(level int) logr.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V", level)
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// V indicates an expected call of V.
func (mr *MockLBScopeMockRecorder) V(level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V", reflect.TypeOf((*MockLBScope)(nil).V), level)
}

// Vnet mocks base method.
func (m *MockLBScope) Vnet() *v1beta1.VnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Vnet")
	ret0, _ := ret[0].(*v1beta1.VnetSpec)
	return ret0
}

// Vnet indicates an expected call of Vnet.
func (mr *MockLBScopeMockRecorder) Vnet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Vnet", reflect.TypeOf((*MockLBScope)(nil).Vnet))
}

// WithName mocks base method.
func (m *MockLBScope) WithName(name string) logr.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithName", name)
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// WithName indicates an expected call of WithName.
func (mr *MockLBScopeMockRecorder) WithName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithName", reflect.TypeOf((*MockLBScope)(nil).WithName), name)
}

// WithValues mocks base method.
func (m *MockLBScope) WithValues(keysAndValues ...interface{}) logr.Logger {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WithValues", varargs...)
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// WithValues indicates an expected call of WithValues.
func (mr *MockLBScopeMockRecorder) WithValues(keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithValues", reflect.TypeOf((*MockLBScope)(nil).WithValues), keysAndValues...)
}