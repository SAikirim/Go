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

package hybridkubernetes

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/hybridkubernetes/mgmt/2020-01-01-preview/hybridkubernetes"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ConnectivityStatus = original.ConnectivityStatus

const (
	Connected  ConnectivityStatus = original.Connected
	Connecting ConnectivityStatus = original.Connecting
	Expired    ConnectivityStatus = original.Expired
	Offline    ConnectivityStatus = original.Offline
)

type ProvisioningState = original.ProvisioningState

const (
	Accepted     ProvisioningState = original.Accepted
	Canceled     ProvisioningState = original.Canceled
	Deleting     ProvisioningState = original.Deleting
	Failed       ProvisioningState = original.Failed
	Provisioning ProvisioningState = original.Provisioning
	Succeeded    ProvisioningState = original.Succeeded
	Updating     ProvisioningState = original.Updating
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	None           ResourceIdentityType = original.None
	SystemAssigned ResourceIdentityType = original.SystemAssigned
)

type AuthenticationDetails = original.AuthenticationDetails
type AuthenticationDetailsValue = original.AuthenticationDetailsValue
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type ConnectedCluster = original.ConnectedCluster
type ConnectedClusterAADProfile = original.ConnectedClusterAADProfile
type ConnectedClusterClient = original.ConnectedClusterClient
type ConnectedClusterCreateFuture = original.ConnectedClusterCreateFuture
type ConnectedClusterDeleteFuture = original.ConnectedClusterDeleteFuture
type ConnectedClusterIdentity = original.ConnectedClusterIdentity
type ConnectedClusterList = original.ConnectedClusterList
type ConnectedClusterListIterator = original.ConnectedClusterListIterator
type ConnectedClusterListPage = original.ConnectedClusterListPage
type ConnectedClusterPatch = original.ConnectedClusterPatch
type ConnectedClusterPatchProperties = original.ConnectedClusterPatchProperties
type ConnectedClusterProperties = original.ConnectedClusterProperties
type CredentialResult = original.CredentialResult
type CredentialResults = original.CredentialResults
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type HybridConnectionConfig = original.HybridConnectionConfig
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationsClient = original.OperationsClient
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewConnectedClusterClient(subscriptionID string) ConnectedClusterClient {
	return original.NewConnectedClusterClient(subscriptionID)
}
func NewConnectedClusterClientWithBaseURI(baseURI string, subscriptionID string) ConnectedClusterClient {
	return original.NewConnectedClusterClientWithBaseURI(baseURI, subscriptionID)
}
func NewConnectedClusterListIterator(page ConnectedClusterListPage) ConnectedClusterListIterator {
	return original.NewConnectedClusterListIterator(page)
}
func NewConnectedClusterListPage(cur ConnectedClusterList, getNextPage func(context.Context, ConnectedClusterList) (ConnectedClusterList, error)) ConnectedClusterListPage {
	return original.NewConnectedClusterListPage(cur, getNextPage)
}
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return original.NewOperationListIterator(page)
}
func NewOperationListPage(cur OperationList, getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return original.NewOperationListPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleConnectivityStatusValues() []ConnectivityStatus {
	return original.PossibleConnectivityStatusValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}