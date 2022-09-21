package confluentapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/confluent/mgmt/2020-03-01/confluent"
)

// MarketplaceAgreementsClientAPI contains the set of methods on the MarketplaceAgreementsClient type.
type MarketplaceAgreementsClientAPI interface {
	Create(ctx context.Context, body *confluent.AgreementResource) (result confluent.AgreementResource, err error)
	List(ctx context.Context) (result confluent.AgreementResourceListResponsePage, err error)
	ListComplete(ctx context.Context) (result confluent.AgreementResourceListResponseIterator, err error)
}

var _ MarketplaceAgreementsClientAPI = (*confluent.MarketplaceAgreementsClient)(nil)

// OrganizationOperationsClientAPI contains the set of methods on the OrganizationOperationsClient type.
type OrganizationOperationsClientAPI interface {
	List(ctx context.Context) (result confluent.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result confluent.OperationListResultIterator, err error)
}

var _ OrganizationOperationsClientAPI = (*confluent.OrganizationOperationsClient)(nil)

// OrganizationClientAPI contains the set of methods on the OrganizationClient type.
type OrganizationClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, organizationName string, body *confluent.OrganizationResource) (result confluent.OrganizationCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, organizationName string) (result confluent.OrganizationDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, organizationName string) (result confluent.OrganizationResource, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result confluent.OrganizationResourceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result confluent.OrganizationResourceListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result confluent.OrganizationResourceListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result confluent.OrganizationResourceListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, organizationName string, body *confluent.OrganizationResourceUpdate) (result confluent.OrganizationResource, err error)
}

var _ OrganizationClientAPI = (*confluent.OrganizationClient)(nil)
