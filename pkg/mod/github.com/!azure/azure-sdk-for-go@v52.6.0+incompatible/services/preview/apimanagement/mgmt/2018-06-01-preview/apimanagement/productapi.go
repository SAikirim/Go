package apimanagement

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ProductAPIClient is the apiManagement Client
type ProductAPIClient struct {
	BaseClient
}

// NewProductAPIClient creates an instance of the ProductAPIClient client.
func NewProductAPIClient(subscriptionID string) ProductAPIClient {
	return NewProductAPIClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProductAPIClientWithBaseURI creates an instance of the ProductAPIClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewProductAPIClientWithBaseURI(baseURI string, subscriptionID string) ProductAPIClient {
	return ProductAPIClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckEntityExists checks that API entity specified by identifier is associated with the Product entity.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// productID - product identifier. Must be unique in the current API Management service instance.
// apiid - API revision identifier. Must be unique in the current API Management service instance. Non-current
// revision has ;rev=n as a suffix where n is the revision number.
func (client ProductAPIClient) CheckEntityExists(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiid string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductAPIClient.CheckEntityExists")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: productID,
			Constraints: []validation.Constraint{{Target: "productID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "productID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "productID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.ProductAPIClient", "CheckEntityExists", err.Error())
	}

	req, err := client.CheckEntityExistsPreparer(ctx, resourceGroupName, serviceName, productID, apiid)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ProductAPIClient", "CheckEntityExists", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckEntityExistsSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.ProductAPIClient", "CheckEntityExists", resp, "Failure sending request")
		return
	}

	result, err = client.CheckEntityExistsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ProductAPIClient", "CheckEntityExists", resp, "Failure responding to request")
		return
	}

	return
}

// CheckEntityExistsPreparer prepares the CheckEntityExists request.
func (client ProductAPIClient) CheckEntityExistsPreparer(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiid string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"productId":         autorest.Encode("path", productID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsHead(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/apis/{apiId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckEntityExistsSender sends the CheckEntityExists request. The method will close the
// http.Response Body if it receives an error.
func (client ProductAPIClient) CheckEntityExistsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckEntityExistsResponder handles the response to the CheckEntityExists request. The method always
// closes the http.Response Body.
func (client ProductAPIClient) CheckEntityExistsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdate adds an API to the specified product.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// productID - product identifier. Must be unique in the current API Management service instance.
// apiid - API revision identifier. Must be unique in the current API Management service instance. Non-current
// revision has ;rev=n as a suffix where n is the revision number.
func (client ProductAPIClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiid string) (result APIContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductAPIClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: productID,
			Constraints: []validation.Constraint{{Target: "productID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "productID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "productID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.ProductAPIClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, productID, apiid)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ProductAPIClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.ProductAPIClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ProductAPIClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ProductAPIClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiid string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"productId":         autorest.Encode("path", productID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/apis/{apiId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ProductAPIClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ProductAPIClient) CreateOrUpdateResponder(resp *http.Response) (result APIContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified API from the specified product.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// productID - product identifier. Must be unique in the current API Management service instance.
// apiid - API revision identifier. Must be unique in the current API Management service instance. Non-current
// revision has ;rev=n as a suffix where n is the revision number.
func (client ProductAPIClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiid string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductAPIClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: productID,
			Constraints: []validation.Constraint{{Target: "productID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "productID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "productID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.ProductAPIClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, productID, apiid)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ProductAPIClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.ProductAPIClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ProductAPIClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ProductAPIClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiid string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"productId":         autorest.Encode("path", productID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/apis/{apiId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ProductAPIClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ProductAPIClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByProduct lists a collection of the APIs associated with a product.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// productID - product identifier. Must be unique in the current API Management service instance.
// filter - | Field       | Supported operators    | Supported functions               |
// |-------------|------------------------|-----------------------------------|
//
// |name | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith|
// |displayName | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith|
// |description | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith|
// |serviceUrl | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith|
// |path | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith|
// top - number of records to return.
// skip - number of records to skip.
func (client ProductAPIClient) ListByProduct(ctx context.Context, resourceGroupName string, serviceName string, productID string, filter string, top *int32, skip *int32) (result APICollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductAPIClient.ListByProduct")
		defer func() {
			sc := -1
			if result.ac.Response.Response != nil {
				sc = result.ac.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: productID,
			Constraints: []validation.Constraint{{Target: "productID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "productID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "productID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.ProductAPIClient", "ListByProduct", err.Error())
	}

	result.fn = client.listByProductNextResults
	req, err := client.ListByProductPreparer(ctx, resourceGroupName, serviceName, productID, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ProductAPIClient", "ListByProduct", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByProductSender(req)
	if err != nil {
		result.ac.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.ProductAPIClient", "ListByProduct", resp, "Failure sending request")
		return
	}

	result.ac, err = client.ListByProductResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ProductAPIClient", "ListByProduct", resp, "Failure responding to request")
		return
	}
	if result.ac.hasNextLink() && result.ac.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByProductPreparer prepares the ListByProduct request.
func (client ProductAPIClient) ListByProductPreparer(ctx context.Context, resourceGroupName string, serviceName string, productID string, filter string, top *int32, skip *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"productId":         autorest.Encode("path", productID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/apis", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByProductSender sends the ListByProduct request. The method will close the
// http.Response Body if it receives an error.
func (client ProductAPIClient) ListByProductSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByProductResponder handles the response to the ListByProduct request. The method always
// closes the http.Response Body.
func (client ProductAPIClient) ListByProductResponder(resp *http.Response) (result APICollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByProductNextResults retrieves the next set of results, if any.
func (client ProductAPIClient) listByProductNextResults(ctx context.Context, lastResults APICollection) (result APICollection, err error) {
	req, err := lastResults.aPICollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.ProductAPIClient", "listByProductNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByProductSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.ProductAPIClient", "listByProductNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByProductResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ProductAPIClient", "listByProductNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByProductComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProductAPIClient) ListByProductComplete(ctx context.Context, resourceGroupName string, serviceName string, productID string, filter string, top *int32, skip *int32) (result APICollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductAPIClient.ListByProduct")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByProduct(ctx, resourceGroupName, serviceName, productID, filter, top, skip)
	return
}
