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

package azurestack

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/azurestack/mgmt/2017-06-01/azurestack"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type Category = original.Category

const (
	ADFS    Category = original.ADFS
	AzureAD Category = original.AzureAD
)

type CompatibilityIssue = original.CompatibilityIssue

const (
	ADFSIdentitySystemRequired      CompatibilityIssue = original.ADFSIdentitySystemRequired
	AzureADIdentitySystemRequired   CompatibilityIssue = original.AzureADIdentitySystemRequired
	CapacityBillingModelRequired    CompatibilityIssue = original.CapacityBillingModelRequired
	ConnectionToAzureRequired       CompatibilityIssue = original.ConnectionToAzureRequired
	ConnectionToInternetRequired    CompatibilityIssue = original.ConnectionToInternetRequired
	DevelopmentBillingModelRequired CompatibilityIssue = original.DevelopmentBillingModelRequired
	DisconnectedEnvironmentRequired CompatibilityIssue = original.DisconnectedEnvironmentRequired
	HigherDeviceVersionRequired     CompatibilityIssue = original.HigherDeviceVersionRequired
	LowerDeviceVersionRequired      CompatibilityIssue = original.LowerDeviceVersionRequired
	PayAsYouGoBillingModelRequired  CompatibilityIssue = original.PayAsYouGoBillingModelRequired
)

type ComputeRole = original.ComputeRole

const (
	IaaS ComputeRole = original.IaaS
	None ComputeRole = original.None
	PaaS ComputeRole = original.PaaS
)

type OperatingSystem = original.OperatingSystem

const (
	OperatingSystemLinux   OperatingSystem = original.OperatingSystemLinux
	OperatingSystemNone    OperatingSystem = original.OperatingSystemNone
	OperatingSystemWindows OperatingSystem = original.OperatingSystemWindows
)

type ProvisioningState = original.ProvisioningState

const (
	Canceled  ProvisioningState = original.Canceled
	Creating  ProvisioningState = original.Creating
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
)

type ActivationKeyResult = original.ActivationKeyResult
type BaseClient = original.BaseClient
type CloudManifestFileClient = original.CloudManifestFileClient
type CloudManifestFileDeploymentData = original.CloudManifestFileDeploymentData
type CloudManifestFileEnvironmentEndpoints = original.CloudManifestFileEnvironmentEndpoints
type CloudManifestFileProperties = original.CloudManifestFileProperties
type CloudManifestFileResponse = original.CloudManifestFileResponse
type Compatibility = original.Compatibility
type CustomerSubscription = original.CustomerSubscription
type CustomerSubscriptionList = original.CustomerSubscriptionList
type CustomerSubscriptionListIterator = original.CustomerSubscriptionListIterator
type CustomerSubscriptionListPage = original.CustomerSubscriptionListPage
type CustomerSubscriptionProperties = original.CustomerSubscriptionProperties
type CustomerSubscriptionsClient = original.CustomerSubscriptionsClient
type DataDiskImage = original.DataDiskImage
type DeviceConfiguration = original.DeviceConfiguration
type Display = original.Display
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type ExtendedProduct = original.ExtendedProduct
type ExtendedProductProperties = original.ExtendedProductProperties
type IconUris = original.IconUris
type MarketplaceProductLogUpdate = original.MarketplaceProductLogUpdate
type Operation = original.Operation
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationsClient = original.OperationsClient
type OsDiskImage = original.OsDiskImage
type Product = original.Product
type ProductLink = original.ProductLink
type ProductList = original.ProductList
type ProductListIterator = original.ProductListIterator
type ProductListPage = original.ProductListPage
type ProductLog = original.ProductLog
type ProductNestedProperties = original.ProductNestedProperties
type ProductProperties = original.ProductProperties
type ProductsClient = original.ProductsClient
type Registration = original.Registration
type RegistrationList = original.RegistrationList
type RegistrationListIterator = original.RegistrationListIterator
type RegistrationListPage = original.RegistrationListPage
type RegistrationParameter = original.RegistrationParameter
type RegistrationParameterProperties = original.RegistrationParameterProperties
type RegistrationProperties = original.RegistrationProperties
type RegistrationsClient = original.RegistrationsClient
type Resource = original.Resource
type TrackedResource = original.TrackedResource
type URI = original.URI
type VirtualMachineExtensionProductProperties = original.VirtualMachineExtensionProductProperties
type VirtualMachineProductProperties = original.VirtualMachineProductProperties

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewCloudManifestFileClient(subscriptionID string) CloudManifestFileClient {
	return original.NewCloudManifestFileClient(subscriptionID)
}
func NewCloudManifestFileClientWithBaseURI(baseURI string, subscriptionID string) CloudManifestFileClient {
	return original.NewCloudManifestFileClientWithBaseURI(baseURI, subscriptionID)
}
func NewCustomerSubscriptionListIterator(page CustomerSubscriptionListPage) CustomerSubscriptionListIterator {
	return original.NewCustomerSubscriptionListIterator(page)
}
func NewCustomerSubscriptionListPage(cur CustomerSubscriptionList, getNextPage func(context.Context, CustomerSubscriptionList) (CustomerSubscriptionList, error)) CustomerSubscriptionListPage {
	return original.NewCustomerSubscriptionListPage(cur, getNextPage)
}
func NewCustomerSubscriptionsClient(subscriptionID string) CustomerSubscriptionsClient {
	return original.NewCustomerSubscriptionsClient(subscriptionID)
}
func NewCustomerSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) CustomerSubscriptionsClient {
	return original.NewCustomerSubscriptionsClientWithBaseURI(baseURI, subscriptionID)
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
func NewProductListIterator(page ProductListPage) ProductListIterator {
	return original.NewProductListIterator(page)
}
func NewProductListPage(cur ProductList, getNextPage func(context.Context, ProductList) (ProductList, error)) ProductListPage {
	return original.NewProductListPage(cur, getNextPage)
}
func NewProductsClient(subscriptionID string) ProductsClient {
	return original.NewProductsClient(subscriptionID)
}
func NewProductsClientWithBaseURI(baseURI string, subscriptionID string) ProductsClient {
	return original.NewProductsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRegistrationListIterator(page RegistrationListPage) RegistrationListIterator {
	return original.NewRegistrationListIterator(page)
}
func NewRegistrationListPage(cur RegistrationList, getNextPage func(context.Context, RegistrationList) (RegistrationList, error)) RegistrationListPage {
	return original.NewRegistrationListPage(cur, getNextPage)
}
func NewRegistrationsClient(subscriptionID string) RegistrationsClient {
	return original.NewRegistrationsClient(subscriptionID)
}
func NewRegistrationsClientWithBaseURI(baseURI string, subscriptionID string) RegistrationsClient {
	return original.NewRegistrationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCategoryValues() []Category {
	return original.PossibleCategoryValues()
}
func PossibleCompatibilityIssueValues() []CompatibilityIssue {
	return original.PossibleCompatibilityIssueValues()
}
func PossibleComputeRoleValues() []ComputeRole {
	return original.PossibleComputeRoleValues()
}
func PossibleOperatingSystemValues() []OperatingSystem {
	return original.PossibleOperatingSystemValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}