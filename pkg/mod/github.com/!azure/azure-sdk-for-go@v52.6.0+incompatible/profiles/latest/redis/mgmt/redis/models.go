// +build go1.9

// Copyright 2021 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package redis

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2020-06-01/redis"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type DayOfWeek = original.DayOfWeek

const (
	Everyday  DayOfWeek = original.Everyday
	Friday    DayOfWeek = original.Friday
	Monday    DayOfWeek = original.Monday
	Saturday  DayOfWeek = original.Saturday
	Sunday    DayOfWeek = original.Sunday
	Thursday  DayOfWeek = original.Thursday
	Tuesday   DayOfWeek = original.Tuesday
	Wednesday DayOfWeek = original.Wednesday
	Weekend   DayOfWeek = original.Weekend
)

type KeyType = original.KeyType

const (
	Primary   KeyType = original.Primary
	Secondary KeyType = original.Secondary
)

type PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningState

const (
	Creating  PrivateEndpointConnectionProvisioningState = original.Creating
	Deleting  PrivateEndpointConnectionProvisioningState = original.Deleting
	Failed    PrivateEndpointConnectionProvisioningState = original.Failed
	Succeeded PrivateEndpointConnectionProvisioningState = original.Succeeded
)

type PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatus

const (
	Approved PrivateEndpointServiceConnectionStatus = original.Approved
	Pending  PrivateEndpointServiceConnectionStatus = original.Pending
	Rejected PrivateEndpointServiceConnectionStatus = original.Rejected
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateCreating               ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleting               ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateDisabled               ProvisioningState = original.ProvisioningStateDisabled
	ProvisioningStateFailed                 ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateLinking                ProvisioningState = original.ProvisioningStateLinking
	ProvisioningStateProvisioning           ProvisioningState = original.ProvisioningStateProvisioning
	ProvisioningStateRecoveringScaleFailure ProvisioningState = original.ProvisioningStateRecoveringScaleFailure
	ProvisioningStateScaling                ProvisioningState = original.ProvisioningStateScaling
	ProvisioningStateSucceeded              ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUnlinking              ProvisioningState = original.ProvisioningStateUnlinking
	ProvisioningStateUnprovisioning         ProvisioningState = original.ProvisioningStateUnprovisioning
	ProvisioningStateUpdating               ProvisioningState = original.ProvisioningStateUpdating
)

type PublicNetworkAccess = original.PublicNetworkAccess

const (
	Disabled PublicNetworkAccess = original.Disabled
	Enabled  PublicNetworkAccess = original.Enabled
)

type RebootType = original.RebootType

const (
	AllNodes      RebootType = original.AllNodes
	PrimaryNode   RebootType = original.PrimaryNode
	SecondaryNode RebootType = original.SecondaryNode
)

type ReplicationRole = original.ReplicationRole

const (
	ReplicationRolePrimary   ReplicationRole = original.ReplicationRolePrimary
	ReplicationRoleSecondary ReplicationRole = original.ReplicationRoleSecondary
)

type SkuFamily = original.SkuFamily

const (
	C SkuFamily = original.C
	P SkuFamily = original.P
)

type SkuName = original.SkuName

const (
	Basic    SkuName = original.Basic
	Premium  SkuName = original.Premium
	Standard SkuName = original.Standard
)

type TLSVersion = original.TLSVersion

const (
	OneFullStopOne  TLSVersion = original.OneFullStopOne
	OneFullStopTwo  TLSVersion = original.OneFullStopTwo
	OneFullStopZero TLSVersion = original.OneFullStopZero
)

type AccessKeys = original.AccessKeys
type BaseClient = original.BaseClient
type CheckNameAvailabilityParameters = original.CheckNameAvailabilityParameters
type Client = original.Client
type CommonProperties = original.CommonProperties
type CreateFuture = original.CreateFuture
type CreateParameters = original.CreateParameters
type CreateProperties = original.CreateProperties
type DeleteFuture = original.DeleteFuture
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type ExportDataFuture = original.ExportDataFuture
type ExportRDBParameters = original.ExportRDBParameters
type FirewallRule = original.FirewallRule
type FirewallRuleCreateParameters = original.FirewallRuleCreateParameters
type FirewallRuleListResult = original.FirewallRuleListResult
type FirewallRuleListResultIterator = original.FirewallRuleListResultIterator
type FirewallRuleListResultPage = original.FirewallRuleListResultPage
type FirewallRuleProperties = original.FirewallRuleProperties
type FirewallRulesClient = original.FirewallRulesClient
type ForceRebootResponse = original.ForceRebootResponse
type ImportDataFuture = original.ImportDataFuture
type ImportRDBParameters = original.ImportRDBParameters
type InstanceDetails = original.InstanceDetails
type LinkedServer = original.LinkedServer
type LinkedServerClient = original.LinkedServerClient
type LinkedServerCreateFuture = original.LinkedServerCreateFuture
type LinkedServerCreateParameters = original.LinkedServerCreateParameters
type LinkedServerCreateProperties = original.LinkedServerCreateProperties
type LinkedServerProperties = original.LinkedServerProperties
type LinkedServerWithProperties = original.LinkedServerWithProperties
type LinkedServerWithPropertiesList = original.LinkedServerWithPropertiesList
type LinkedServerWithPropertiesListIterator = original.LinkedServerWithPropertiesListIterator
type LinkedServerWithPropertiesListPage = original.LinkedServerWithPropertiesListPage
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type NotificationListResponse = original.NotificationListResponse
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type PatchSchedule = original.PatchSchedule
type PatchScheduleListResult = original.PatchScheduleListResult
type PatchScheduleListResultIterator = original.PatchScheduleListResultIterator
type PatchScheduleListResultPage = original.PatchScheduleListResultPage
type PatchSchedulesClient = original.PatchSchedulesClient
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateEndpointConnectionsPutFuture = original.PrivateEndpointConnectionsPutFuture
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceListResult = original.PrivateLinkResourceListResult
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type Properties = original.Properties
type ProxyResource = original.ProxyResource
type RebootParameters = original.RebootParameters
type RegenerateKeyParameters = original.RegenerateKeyParameters
type Resource = original.Resource
type ResourceType = original.ResourceType
type ScheduleEntries = original.ScheduleEntries
type ScheduleEntry = original.ScheduleEntry
type Sku = original.Sku
type TrackedResource = original.TrackedResource
type UpdateParameters = original.UpdateParameters
type UpdateProperties = original.UpdateProperties
type UpgradeNotification = original.UpgradeNotification

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func NewFirewallRuleListResultIterator(page FirewallRuleListResultPage) FirewallRuleListResultIterator {
	return original.NewFirewallRuleListResultIterator(page)
}
func NewFirewallRuleListResultPage(cur FirewallRuleListResult, getNextPage func(context.Context, FirewallRuleListResult) (FirewallRuleListResult, error)) FirewallRuleListResultPage {
	return original.NewFirewallRuleListResultPage(cur, getNextPage)
}
func NewFirewallRulesClient(subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClient(subscriptionID)
}
func NewFirewallRulesClientWithBaseURI(baseURI string, subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewLinkedServerClient(subscriptionID string) LinkedServerClient {
	return original.NewLinkedServerClient(subscriptionID)
}
func NewLinkedServerClientWithBaseURI(baseURI string, subscriptionID string) LinkedServerClient {
	return original.NewLinkedServerClientWithBaseURI(baseURI, subscriptionID)
}
func NewLinkedServerWithPropertiesListIterator(page LinkedServerWithPropertiesListPage) LinkedServerWithPropertiesListIterator {
	return original.NewLinkedServerWithPropertiesListIterator(page)
}
func NewLinkedServerWithPropertiesListPage(cur LinkedServerWithPropertiesList, getNextPage func(context.Context, LinkedServerWithPropertiesList) (LinkedServerWithPropertiesList, error)) LinkedServerWithPropertiesListPage {
	return original.NewLinkedServerWithPropertiesListPage(cur, getNextPage)
}
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return original.NewListResultIterator(page)
}
func NewListResultPage(cur ListResult, getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return original.NewListResultPage(cur, getNextPage)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPatchScheduleListResultIterator(page PatchScheduleListResultPage) PatchScheduleListResultIterator {
	return original.NewPatchScheduleListResultIterator(page)
}
func NewPatchScheduleListResultPage(cur PatchScheduleListResult, getNextPage func(context.Context, PatchScheduleListResult) (PatchScheduleListResult, error)) PatchScheduleListResultPage {
	return original.NewPatchScheduleListResultPage(cur, getNextPage)
}
func NewPatchSchedulesClient(subscriptionID string) PatchSchedulesClient {
	return original.NewPatchSchedulesClient(subscriptionID)
}
func NewPatchSchedulesClientWithBaseURI(baseURI string, subscriptionID string) PatchSchedulesClient {
	return original.NewPatchSchedulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClient(subscriptionID)
}
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleDayOfWeekValues() []DayOfWeek {
	return original.PossibleDayOfWeekValues()
}
func PossibleKeyTypeValues() []KeyType {
	return original.PossibleKeyTypeValues()
}
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return original.PossiblePrivateEndpointConnectionProvisioningStateValues()
}
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return original.PossiblePrivateEndpointServiceConnectionStatusValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return original.PossiblePublicNetworkAccessValues()
}
func PossibleRebootTypeValues() []RebootType {
	return original.PossibleRebootTypeValues()
}
func PossibleReplicationRoleValues() []ReplicationRole {
	return original.PossibleReplicationRoleValues()
}
func PossibleSkuFamilyValues() []SkuFamily {
	return original.PossibleSkuFamilyValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleTLSVersionValues() []TLSVersion {
	return original.PossibleTLSVersionValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
