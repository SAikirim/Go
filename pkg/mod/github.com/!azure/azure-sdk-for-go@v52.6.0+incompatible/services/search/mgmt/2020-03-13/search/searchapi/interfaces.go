package searchapi

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
	"github.com/Azure/azure-sdk-for-go/services/search/mgmt/2020-03-13/search"
	"github.com/Azure/go-autorest/autorest"
	"github.com/gofrs/uuid"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result search.OperationListResult, err error)
}

var _ OperationsClientAPI = (*search.OperationsClient)(nil)

// AdminKeysClientAPI contains the set of methods on the AdminKeysClient type.
type AdminKeysClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, searchServiceName string, clientRequestID *uuid.UUID) (result search.AdminKeyResult, err error)
	Regenerate(ctx context.Context, resourceGroupName string, searchServiceName string, keyKind search.AdminKeyKind, clientRequestID *uuid.UUID) (result search.AdminKeyResult, err error)
}

var _ AdminKeysClientAPI = (*search.AdminKeysClient)(nil)

// QueryKeysClientAPI contains the set of methods on the QueryKeysClient type.
type QueryKeysClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, searchServiceName string, name string, clientRequestID *uuid.UUID) (result search.QueryKey, err error)
	Delete(ctx context.Context, resourceGroupName string, searchServiceName string, key string, clientRequestID *uuid.UUID) (result autorest.Response, err error)
	ListBySearchService(ctx context.Context, resourceGroupName string, searchServiceName string, clientRequestID *uuid.UUID) (result search.ListQueryKeysResultPage, err error)
	ListBySearchServiceComplete(ctx context.Context, resourceGroupName string, searchServiceName string, clientRequestID *uuid.UUID) (result search.ListQueryKeysResultIterator, err error)
}

var _ QueryKeysClientAPI = (*search.QueryKeysClient)(nil)

// ServicesClientAPI contains the set of methods on the ServicesClient type.
type ServicesClientAPI interface {
	CheckNameAvailability(ctx context.Context, checkNameAvailabilityInput search.CheckNameAvailabilityInput, clientRequestID *uuid.UUID) (result search.CheckNameAvailabilityOutput, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, searchServiceName string, service search.Service, clientRequestID *uuid.UUID) (result search.ServicesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, searchServiceName string, clientRequestID *uuid.UUID) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, searchServiceName string, clientRequestID *uuid.UUID) (result search.Service, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, clientRequestID *uuid.UUID) (result search.ServiceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, clientRequestID *uuid.UUID) (result search.ServiceListResultIterator, err error)
	ListBySubscription(ctx context.Context, clientRequestID *uuid.UUID) (result search.ServiceListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context, clientRequestID *uuid.UUID) (result search.ServiceListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, searchServiceName string, service search.Service, clientRequestID *uuid.UUID) (result search.Service, err error)
}

var _ ServicesClientAPI = (*search.ServicesClient)(nil)

// PrivateLinkResourcesClientAPI contains the set of methods on the PrivateLinkResourcesClient type.
type PrivateLinkResourcesClientAPI interface {
	ListSupported(ctx context.Context, resourceGroupName string, searchServiceName string, clientRequestID *uuid.UUID) (result search.PrivateLinkResourcesResult, err error)
}

var _ PrivateLinkResourcesClientAPI = (*search.PrivateLinkResourcesClient)(nil)

// PrivateEndpointConnectionsClientAPI contains the set of methods on the PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, searchServiceName string, privateEndpointConnectionName string, clientRequestID *uuid.UUID) (result search.PrivateEndpointConnection, err error)
	Get(ctx context.Context, resourceGroupName string, searchServiceName string, privateEndpointConnectionName string, clientRequestID *uuid.UUID) (result search.PrivateEndpointConnection, err error)
	ListByService(ctx context.Context, resourceGroupName string, searchServiceName string, clientRequestID *uuid.UUID) (result search.PrivateEndpointConnectionListResultPage, err error)
	ListByServiceComplete(ctx context.Context, resourceGroupName string, searchServiceName string, clientRequestID *uuid.UUID) (result search.PrivateEndpointConnectionListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, searchServiceName string, privateEndpointConnectionName string, privateEndpointConnection search.PrivateEndpointConnection, clientRequestID *uuid.UUID) (result search.PrivateEndpointConnection, err error)
}

var _ PrivateEndpointConnectionsClientAPI = (*search.PrivateEndpointConnectionsClient)(nil)
