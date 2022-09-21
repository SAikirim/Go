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

package storagesync

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/storagesync/mgmt/2020-03-01/storagesync"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ChangeDetectionMode = original.ChangeDetectionMode

const (
	Default   ChangeDetectionMode = original.Default
	Recursive ChangeDetectionMode = original.Recursive
)

type FeatureStatus = original.FeatureStatus

const (
	Off FeatureStatus = original.Off
	On  FeatureStatus = original.On
)

type IncomingTrafficPolicy = original.IncomingTrafficPolicy

const (
	AllowAllTraffic          IncomingTrafficPolicy = original.AllowAllTraffic
	AllowVirtualNetworksOnly IncomingTrafficPolicy = original.AllowVirtualNetworksOnly
)

type InitialDownloadPolicy = original.InitialDownloadPolicy

const (
	AvoidTieredFiles           InitialDownloadPolicy = original.AvoidTieredFiles
	NamespaceOnly              InitialDownloadPolicy = original.NamespaceOnly
	NamespaceThenModifiedFiles InitialDownloadPolicy = original.NamespaceThenModifiedFiles
)

type LocalCacheMode = original.LocalCacheMode

const (
	DownloadNewAndModifiedFiles LocalCacheMode = original.DownloadNewAndModifiedFiles
	UpdateLocallyCachedFiles    LocalCacheMode = original.UpdateLocallyCachedFiles
)

type NameAvailabilityReason = original.NameAvailabilityReason

const (
	AlreadyExists NameAvailabilityReason = original.AlreadyExists
	Invalid       NameAvailabilityReason = original.Invalid
)

type OperationDirection = original.OperationDirection

const (
	Cancel OperationDirection = original.Cancel
	Do     OperationDirection = original.Do
	Undo   OperationDirection = original.Undo
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

type ProgressType = original.ProgressType

const (
	Download   ProgressType = original.Download
	Initialize ProgressType = original.Initialize
	None       ProgressType = original.None
	Recall     ProgressType = original.Recall
	Upload     ProgressType = original.Upload
)

type Reason = original.Reason

const (
	Deleted      Reason = original.Deleted
	Registered   Reason = original.Registered
	Suspended    Reason = original.Suspended
	Unregistered Reason = original.Unregistered
	Warned       Reason = original.Warned
)

type ServerEndpointCloudTieringHealthState = original.ServerEndpointCloudTieringHealthState

const (
	ServerEndpointCloudTieringHealthStateError   ServerEndpointCloudTieringHealthState = original.ServerEndpointCloudTieringHealthStateError
	ServerEndpointCloudTieringHealthStateHealthy ServerEndpointCloudTieringHealthState = original.ServerEndpointCloudTieringHealthStateHealthy
)

type ServerEndpointOfflineDataTransferState = original.ServerEndpointOfflineDataTransferState

const (
	Complete   ServerEndpointOfflineDataTransferState = original.Complete
	InProgress ServerEndpointOfflineDataTransferState = original.InProgress
	NotRunning ServerEndpointOfflineDataTransferState = original.NotRunning
	Stopping   ServerEndpointOfflineDataTransferState = original.Stopping
)

type ServerEndpointSyncActivityState = original.ServerEndpointSyncActivityState

const (
	ServerEndpointSyncActivityStateDownload          ServerEndpointSyncActivityState = original.ServerEndpointSyncActivityStateDownload
	ServerEndpointSyncActivityStateUpload            ServerEndpointSyncActivityState = original.ServerEndpointSyncActivityStateUpload
	ServerEndpointSyncActivityStateUploadAndDownload ServerEndpointSyncActivityState = original.ServerEndpointSyncActivityStateUploadAndDownload
)

type ServerEndpointSyncHealthState = original.ServerEndpointSyncHealthState

const (
	ServerEndpointSyncHealthStateError                                    ServerEndpointSyncHealthState = original.ServerEndpointSyncHealthStateError
	ServerEndpointSyncHealthStateHealthy                                  ServerEndpointSyncHealthState = original.ServerEndpointSyncHealthStateHealthy
	ServerEndpointSyncHealthStateNoActivity                               ServerEndpointSyncHealthState = original.ServerEndpointSyncHealthStateNoActivity
	ServerEndpointSyncHealthStateSyncBlockedForChangeDetectionPostRestore ServerEndpointSyncHealthState = original.ServerEndpointSyncHealthStateSyncBlockedForChangeDetectionPostRestore
	ServerEndpointSyncHealthStateSyncBlockedForRestore                    ServerEndpointSyncHealthState = original.ServerEndpointSyncHealthStateSyncBlockedForRestore
)

type WorkflowStatus = original.WorkflowStatus

const (
	WorkflowStatusAborted   WorkflowStatus = original.WorkflowStatusAborted
	WorkflowStatusActive    WorkflowStatus = original.WorkflowStatusActive
	WorkflowStatusExpired   WorkflowStatus = original.WorkflowStatusExpired
	WorkflowStatusFailed    WorkflowStatus = original.WorkflowStatusFailed
	WorkflowStatusSucceeded WorkflowStatus = original.WorkflowStatusSucceeded
)

type APIError = original.APIError
type AzureEntityResource = original.AzureEntityResource
type BackupRequest = original.BackupRequest
type BaseClient = original.BaseClient
type CheckNameAvailabilityParameters = original.CheckNameAvailabilityParameters
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type CloudEndpoint = original.CloudEndpoint
type CloudEndpointArray = original.CloudEndpointArray
type CloudEndpointCreateParameters = original.CloudEndpointCreateParameters
type CloudEndpointCreateParametersProperties = original.CloudEndpointCreateParametersProperties
type CloudEndpointProperties = original.CloudEndpointProperties
type CloudEndpointsClient = original.CloudEndpointsClient
type CloudEndpointsCreateFuture = original.CloudEndpointsCreateFuture
type CloudEndpointsDeleteFuture = original.CloudEndpointsDeleteFuture
type CloudEndpointsPostBackupFuture = original.CloudEndpointsPostBackupFuture
type CloudEndpointsPostRestoreFuture = original.CloudEndpointsPostRestoreFuture
type CloudEndpointsPreBackupFuture = original.CloudEndpointsPreBackupFuture
type CloudEndpointsPreRestoreFuture = original.CloudEndpointsPreRestoreFuture
type CloudEndpointsTriggerChangeDetectionFuture = original.CloudEndpointsTriggerChangeDetectionFuture
type CloudTieringCachePerformance = original.CloudTieringCachePerformance
type CloudTieringDatePolicyStatus = original.CloudTieringDatePolicyStatus
type CloudTieringFilesNotTiering = original.CloudTieringFilesNotTiering
type CloudTieringSpaceSavings = original.CloudTieringSpaceSavings
type CloudTieringVolumeFreeSpacePolicyStatus = original.CloudTieringVolumeFreeSpacePolicyStatus
type Error = original.Error
type ErrorDetails = original.ErrorDetails
type FilesNotTieringError = original.FilesNotTieringError
type OperationDisplayInfo = original.OperationDisplayInfo
type OperationDisplayResource = original.OperationDisplayResource
type OperationEntity = original.OperationEntity
type OperationEntityListResult = original.OperationEntityListResult
type OperationEntityListResultIterator = original.OperationEntityListResultIterator
type OperationEntityListResultPage = original.OperationEntityListResultPage
type OperationStatus = original.OperationStatus
type OperationStatusClient = original.OperationStatusClient
type OperationsClient = original.OperationsClient
type PostBackupResponse = original.PostBackupResponse
type PostBackupResponseProperties = original.PostBackupResponseProperties
type PostRestoreRequest = original.PostRestoreRequest
type PreRestoreRequest = original.PreRestoreRequest
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateEndpointConnectionsCreateFuture = original.PrivateEndpointConnectionsCreateFuture
type PrivateEndpointConnectionsDeleteFuture = original.PrivateEndpointConnectionsDeleteFuture
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceListResult = original.PrivateLinkResourceListResult
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type ProxyResource = original.ProxyResource
type RecallActionParameters = original.RecallActionParameters
type RegisteredServer = original.RegisteredServer
type RegisteredServerArray = original.RegisteredServerArray
type RegisteredServerCreateParameters = original.RegisteredServerCreateParameters
type RegisteredServerCreateParametersProperties = original.RegisteredServerCreateParametersProperties
type RegisteredServerProperties = original.RegisteredServerProperties
type RegisteredServersClient = original.RegisteredServersClient
type RegisteredServersCreateFuture = original.RegisteredServersCreateFuture
type RegisteredServersDeleteFuture = original.RegisteredServersDeleteFuture
type RegisteredServersTriggerRolloverFuture = original.RegisteredServersTriggerRolloverFuture
type Resource = original.Resource
type ResourcesMoveInfo = original.ResourcesMoveInfo
type RestoreFileSpec = original.RestoreFileSpec
type ServerEndpoint = original.ServerEndpoint
type ServerEndpointArray = original.ServerEndpointArray
type ServerEndpointCloudTieringStatus = original.ServerEndpointCloudTieringStatus
type ServerEndpointCreateParameters = original.ServerEndpointCreateParameters
type ServerEndpointCreateParametersProperties = original.ServerEndpointCreateParametersProperties
type ServerEndpointFilesNotSyncingError = original.ServerEndpointFilesNotSyncingError
type ServerEndpointProperties = original.ServerEndpointProperties
type ServerEndpointRecallError = original.ServerEndpointRecallError
type ServerEndpointRecallStatus = original.ServerEndpointRecallStatus
type ServerEndpointSyncActivityStatus = original.ServerEndpointSyncActivityStatus
type ServerEndpointSyncSessionStatus = original.ServerEndpointSyncSessionStatus
type ServerEndpointSyncStatus = original.ServerEndpointSyncStatus
type ServerEndpointUpdateParameters = original.ServerEndpointUpdateParameters
type ServerEndpointUpdateProperties = original.ServerEndpointUpdateProperties
type ServerEndpointsClient = original.ServerEndpointsClient
type ServerEndpointsCreateFuture = original.ServerEndpointsCreateFuture
type ServerEndpointsDeleteFuture = original.ServerEndpointsDeleteFuture
type ServerEndpointsRecallActionFuture = original.ServerEndpointsRecallActionFuture
type ServerEndpointsUpdateFuture = original.ServerEndpointsUpdateFuture
type Service = original.Service
type ServiceArray = original.ServiceArray
type ServiceCreateParameters = original.ServiceCreateParameters
type ServiceCreateParametersProperties = original.ServiceCreateParametersProperties
type ServiceProperties = original.ServiceProperties
type ServiceUpdateParameters = original.ServiceUpdateParameters
type ServiceUpdateProperties = original.ServiceUpdateProperties
type ServicesClient = original.ServicesClient
type ServicesCreateFuture = original.ServicesCreateFuture
type ServicesDeleteFuture = original.ServicesDeleteFuture
type ServicesUpdateFuture = original.ServicesUpdateFuture
type SubscriptionState = original.SubscriptionState
type SyncGroup = original.SyncGroup
type SyncGroupArray = original.SyncGroupArray
type SyncGroupCreateParameters = original.SyncGroupCreateParameters
type SyncGroupProperties = original.SyncGroupProperties
type SyncGroupsClient = original.SyncGroupsClient
type TrackedResource = original.TrackedResource
type TriggerChangeDetectionParameters = original.TriggerChangeDetectionParameters
type TriggerRolloverRequest = original.TriggerRolloverRequest
type Workflow = original.Workflow
type WorkflowArray = original.WorkflowArray
type WorkflowProperties = original.WorkflowProperties
type WorkflowsClient = original.WorkflowsClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewCloudEndpointsClient(subscriptionID string) CloudEndpointsClient {
	return original.NewCloudEndpointsClient(subscriptionID)
}
func NewCloudEndpointsClientWithBaseURI(baseURI string, subscriptionID string) CloudEndpointsClient {
	return original.NewCloudEndpointsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationEntityListResultIterator(page OperationEntityListResultPage) OperationEntityListResultIterator {
	return original.NewOperationEntityListResultIterator(page)
}
func NewOperationEntityListResultPage(cur OperationEntityListResult, getNextPage func(context.Context, OperationEntityListResult) (OperationEntityListResult, error)) OperationEntityListResultPage {
	return original.NewOperationEntityListResultPage(cur, getNextPage)
}
func NewOperationStatusClient(subscriptionID string) OperationStatusClient {
	return original.NewOperationStatusClient(subscriptionID)
}
func NewOperationStatusClientWithBaseURI(baseURI string, subscriptionID string) OperationStatusClient {
	return original.NewOperationStatusClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
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
func NewRegisteredServersClient(subscriptionID string) RegisteredServersClient {
	return original.NewRegisteredServersClient(subscriptionID)
}
func NewRegisteredServersClientWithBaseURI(baseURI string, subscriptionID string) RegisteredServersClient {
	return original.NewRegisteredServersClientWithBaseURI(baseURI, subscriptionID)
}
func NewServerEndpointsClient(subscriptionID string) ServerEndpointsClient {
	return original.NewServerEndpointsClient(subscriptionID)
}
func NewServerEndpointsClientWithBaseURI(baseURI string, subscriptionID string) ServerEndpointsClient {
	return original.NewServerEndpointsClientWithBaseURI(baseURI, subscriptionID)
}
func NewServicesClient(subscriptionID string) ServicesClient {
	return original.NewServicesClient(subscriptionID)
}
func NewServicesClientWithBaseURI(baseURI string, subscriptionID string) ServicesClient {
	return original.NewServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewSyncGroupsClient(subscriptionID string) SyncGroupsClient {
	return original.NewSyncGroupsClient(subscriptionID)
}
func NewSyncGroupsClientWithBaseURI(baseURI string, subscriptionID string) SyncGroupsClient {
	return original.NewSyncGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewWorkflowsClient(subscriptionID string) WorkflowsClient {
	return original.NewWorkflowsClient(subscriptionID)
}
func NewWorkflowsClientWithBaseURI(baseURI string, subscriptionID string) WorkflowsClient {
	return original.NewWorkflowsClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleChangeDetectionModeValues() []ChangeDetectionMode {
	return original.PossibleChangeDetectionModeValues()
}
func PossibleFeatureStatusValues() []FeatureStatus {
	return original.PossibleFeatureStatusValues()
}
func PossibleIncomingTrafficPolicyValues() []IncomingTrafficPolicy {
	return original.PossibleIncomingTrafficPolicyValues()
}
func PossibleInitialDownloadPolicyValues() []InitialDownloadPolicy {
	return original.PossibleInitialDownloadPolicyValues()
}
func PossibleLocalCacheModeValues() []LocalCacheMode {
	return original.PossibleLocalCacheModeValues()
}
func PossibleNameAvailabilityReasonValues() []NameAvailabilityReason {
	return original.PossibleNameAvailabilityReasonValues()
}
func PossibleOperationDirectionValues() []OperationDirection {
	return original.PossibleOperationDirectionValues()
}
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return original.PossiblePrivateEndpointConnectionProvisioningStateValues()
}
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return original.PossiblePrivateEndpointServiceConnectionStatusValues()
}
func PossibleProgressTypeValues() []ProgressType {
	return original.PossibleProgressTypeValues()
}
func PossibleReasonValues() []Reason {
	return original.PossibleReasonValues()
}
func PossibleServerEndpointCloudTieringHealthStateValues() []ServerEndpointCloudTieringHealthState {
	return original.PossibleServerEndpointCloudTieringHealthStateValues()
}
func PossibleServerEndpointOfflineDataTransferStateValues() []ServerEndpointOfflineDataTransferState {
	return original.PossibleServerEndpointOfflineDataTransferStateValues()
}
func PossibleServerEndpointSyncActivityStateValues() []ServerEndpointSyncActivityState {
	return original.PossibleServerEndpointSyncActivityStateValues()
}
func PossibleServerEndpointSyncHealthStateValues() []ServerEndpointSyncHealthState {
	return original.PossibleServerEndpointSyncHealthStateValues()
}
func PossibleWorkflowStatusValues() []WorkflowStatus {
	return original.PossibleWorkflowStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
