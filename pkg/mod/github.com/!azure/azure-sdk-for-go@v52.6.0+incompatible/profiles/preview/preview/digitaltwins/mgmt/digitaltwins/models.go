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

package digitaltwins

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/digitaltwins/mgmt/2020-03-01-preview/digitaltwins"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type EndpointProvisioningState = original.EndpointProvisioningState

const (
	Canceled     EndpointProvisioningState = original.Canceled
	Deleting     EndpointProvisioningState = original.Deleting
	Failed       EndpointProvisioningState = original.Failed
	Provisioning EndpointProvisioningState = original.Provisioning
	Succeeded    EndpointProvisioningState = original.Succeeded
)

type EndpointType = original.EndpointType

const (
	EndpointTypeDigitalTwinsEndpointResourceProperties EndpointType = original.EndpointTypeDigitalTwinsEndpointResourceProperties
	EndpointTypeEventGrid                              EndpointType = original.EndpointTypeEventGrid
	EndpointTypeEventHub                               EndpointType = original.EndpointTypeEventHub
	EndpointTypeServiceBus                             EndpointType = original.EndpointTypeServiceBus
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateCanceled     ProvisioningState = original.ProvisioningStateCanceled
	ProvisioningStateDeleting     ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed       ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateProvisioning ProvisioningState = original.ProvisioningStateProvisioning
	ProvisioningStateSucceeded    ProvisioningState = original.ProvisioningStateSucceeded
)

type Reason = original.Reason

const (
	AlreadyExists Reason = original.AlreadyExists
	Invalid       Reason = original.Invalid
)

type BaseClient = original.BaseClient
type BasicEndpointResourceProperties = original.BasicEndpointResourceProperties
type CheckNameRequest = original.CheckNameRequest
type CheckNameResult = original.CheckNameResult
type Client = original.Client
type CreateOrUpdateFuture = original.CreateOrUpdateFuture
type DeleteFuture = original.DeleteFuture
type Description = original.Description
type DescriptionListResult = original.DescriptionListResult
type DescriptionListResultIterator = original.DescriptionListResultIterator
type DescriptionListResultPage = original.DescriptionListResultPage
type EndpointClient = original.EndpointClient
type EndpointCreateOrUpdateFuture = original.EndpointCreateOrUpdateFuture
type EndpointDeleteFuture = original.EndpointDeleteFuture
type EndpointResource = original.EndpointResource
type EndpointResourceListResult = original.EndpointResourceListResult
type EndpointResourceListResultIterator = original.EndpointResourceListResultIterator
type EndpointResourceListResultPage = original.EndpointResourceListResultPage
type EndpointResourceProperties = original.EndpointResourceProperties
type ErrorDefinition = original.ErrorDefinition
type ErrorResponse = original.ErrorResponse
type EventGrid = original.EventGrid
type EventHub = original.EventHub
type ExternalResource = original.ExternalResource
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type PatchDescription = original.PatchDescription
type Properties = original.Properties
type Resource = original.Resource
type ServiceBus = original.ServiceBus
type SkuInfo = original.SkuInfo
type UpdateFuture = original.UpdateFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func NewDescriptionListResultIterator(page DescriptionListResultPage) DescriptionListResultIterator {
	return original.NewDescriptionListResultIterator(page)
}
func NewDescriptionListResultPage(cur DescriptionListResult, getNextPage func(context.Context, DescriptionListResult) (DescriptionListResult, error)) DescriptionListResultPage {
	return original.NewDescriptionListResultPage(cur, getNextPage)
}
func NewEndpointClient(subscriptionID string) EndpointClient {
	return original.NewEndpointClient(subscriptionID)
}
func NewEndpointClientWithBaseURI(baseURI string, subscriptionID string) EndpointClient {
	return original.NewEndpointClientWithBaseURI(baseURI, subscriptionID)
}
func NewEndpointResourceListResultIterator(page EndpointResourceListResultPage) EndpointResourceListResultIterator {
	return original.NewEndpointResourceListResultIterator(page)
}
func NewEndpointResourceListResultPage(cur EndpointResourceListResult, getNextPage func(context.Context, EndpointResourceListResult) (EndpointResourceListResult, error)) EndpointResourceListResultPage {
	return original.NewEndpointResourceListResultPage(cur, getNextPage)
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
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleEndpointProvisioningStateValues() []EndpointProvisioningState {
	return original.PossibleEndpointProvisioningStateValues()
}
func PossibleEndpointTypeValues() []EndpointType {
	return original.PossibleEndpointTypeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleReasonValues() []Reason {
	return original.PossibleReasonValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
