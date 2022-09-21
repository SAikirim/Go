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

// TenantAccessClient is the apiManagement Client
type TenantAccessClient struct {
	BaseClient
}

// NewTenantAccessClient creates an instance of the TenantAccessClient client.
func NewTenantAccessClient(subscriptionID string) TenantAccessClient {
	return NewTenantAccessClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTenantAccessClientWithBaseURI creates an instance of the TenantAccessClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewTenantAccessClientWithBaseURI(baseURI string, subscriptionID string) TenantAccessClient {
	return TenantAccessClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create update tenant access information details.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// parameters - parameters supplied to retrieve the Tenant Access Information.
// accessName - the identifier of the Access configuration.
// ifMatch - eTag of the Entity. ETag should match the current entity state from the header response of the GET
// request or it should be * for unconditional update.
func (client TenantAccessClient) Create(ctx context.Context, resourceGroupName string, serviceName string, parameters AccessInformationCreateParameters, accessName AccessIDName, ifMatch string) (result AccessInformationContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TenantAccessClient.Create")
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
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.TenantAccessClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, serviceName, parameters, accessName, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client TenantAccessClient) CreatePreparer(ctx context.Context, resourceGroupName string, serviceName string, parameters AccessInformationCreateParameters, accessName AccessIDName, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessName":        autorest.Encode("path", accessName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client TenantAccessClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client TenantAccessClient) CreateResponder(resp *http.Response) (result AccessInformationContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get tenant access information details without secrets.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// accessName - the identifier of the Access configuration.
func (client TenantAccessClient) Get(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName) (result AccessInformationContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TenantAccessClient.Get")
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
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.TenantAccessClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, accessName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client TenantAccessClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessName":        autorest.Encode("path", accessName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client TenantAccessClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client TenantAccessClient) GetResponder(resp *http.Response) (result AccessInformationContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetEntityTag tenant access metadata
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// accessName - the identifier of the Access configuration.
func (client TenantAccessClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TenantAccessClient.GetEntityTag")
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
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.TenantAccessClient", "GetEntityTag", err.Error())
	}

	req, err := client.GetEntityTagPreparer(ctx, resourceGroupName, serviceName, accessName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "GetEntityTag", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEntityTagSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "GetEntityTag", resp, "Failure sending request")
		return
	}

	result, err = client.GetEntityTagResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "GetEntityTag", resp, "Failure responding to request")
		return
	}

	return
}

// GetEntityTagPreparer prepares the GetEntityTag request.
func (client TenantAccessClient) GetEntityTagPreparer(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessName":        autorest.Encode("path", accessName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsHead(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetEntityTagSender sends the GetEntityTag request. The method will close the
// http.Response Body if it receives an error.
func (client TenantAccessClient) GetEntityTagSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetEntityTagResponder handles the response to the GetEntityTag request. The method always
// closes the http.Response Body.
func (client TenantAccessClient) GetEntityTagResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByService returns list of access infos - for Git and Management endpoints.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// filter - not used
func (client TenantAccessClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string, filter string) (result AccessInformationCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TenantAccessClient.ListByService")
		defer func() {
			sc := -1
			if result.aic.Response.Response != nil {
				sc = result.aic.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.TenantAccessClient", "ListByService", err.Error())
	}

	result.fn = client.listByServiceNextResults
	req, err := client.ListByServicePreparer(ctx, resourceGroupName, serviceName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "ListByService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.aic.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "ListByService", resp, "Failure sending request")
		return
	}

	result.aic, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "ListByService", resp, "Failure responding to request")
		return
	}
	if result.aic.hasNextLink() && result.aic.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByServicePreparer prepares the ListByService request.
func (client TenantAccessClient) ListByServicePreparer(ctx context.Context, resourceGroupName string, serviceName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServiceSender sends the ListByService request. The method will close the
// http.Response Body if it receives an error.
func (client TenantAccessClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client TenantAccessClient) ListByServiceResponder(resp *http.Response) (result AccessInformationCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByServiceNextResults retrieves the next set of results, if any.
func (client TenantAccessClient) listByServiceNextResults(ctx context.Context, lastResults AccessInformationCollection) (result AccessInformationCollection, err error) {
	req, err := lastResults.accessInformationCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "listByServiceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "listByServiceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "listByServiceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByServiceComplete enumerates all values, automatically crossing page boundaries as required.
func (client TenantAccessClient) ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string) (result AccessInformationCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TenantAccessClient.ListByService")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByService(ctx, resourceGroupName, serviceName, filter)
	return
}

// ListSecrets get tenant access information details.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// accessName - the identifier of the Access configuration.
func (client TenantAccessClient) ListSecrets(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName) (result AccessInformationSecretsContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TenantAccessClient.ListSecrets")
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
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.TenantAccessClient", "ListSecrets", err.Error())
	}

	req, err := client.ListSecretsPreparer(ctx, resourceGroupName, serviceName, accessName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "ListSecrets", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSecretsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "ListSecrets", resp, "Failure sending request")
		return
	}

	result, err = client.ListSecretsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "ListSecrets", resp, "Failure responding to request")
		return
	}

	return
}

// ListSecretsPreparer prepares the ListSecrets request.
func (client TenantAccessClient) ListSecretsPreparer(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessName":        autorest.Encode("path", accessName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}/listSecrets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSecretsSender sends the ListSecrets request. The method will close the
// http.Response Body if it receives an error.
func (client TenantAccessClient) ListSecretsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListSecretsResponder handles the response to the ListSecrets request. The method always
// closes the http.Response Body.
func (client TenantAccessClient) ListSecretsResponder(resp *http.Response) (result AccessInformationSecretsContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RegeneratePrimaryKey regenerate primary access key
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// accessName - the identifier of the Access configuration.
func (client TenantAccessClient) RegeneratePrimaryKey(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TenantAccessClient.RegeneratePrimaryKey")
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
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.TenantAccessClient", "RegeneratePrimaryKey", err.Error())
	}

	req, err := client.RegeneratePrimaryKeyPreparer(ctx, resourceGroupName, serviceName, accessName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "RegeneratePrimaryKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegeneratePrimaryKeySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "RegeneratePrimaryKey", resp, "Failure sending request")
		return
	}

	result, err = client.RegeneratePrimaryKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "RegeneratePrimaryKey", resp, "Failure responding to request")
		return
	}

	return
}

// RegeneratePrimaryKeyPreparer prepares the RegeneratePrimaryKey request.
func (client TenantAccessClient) RegeneratePrimaryKeyPreparer(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessName":        autorest.Encode("path", accessName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}/regeneratePrimaryKey", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegeneratePrimaryKeySender sends the RegeneratePrimaryKey request. The method will close the
// http.Response Body if it receives an error.
func (client TenantAccessClient) RegeneratePrimaryKeySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RegeneratePrimaryKeyResponder handles the response to the RegeneratePrimaryKey request. The method always
// closes the http.Response Body.
func (client TenantAccessClient) RegeneratePrimaryKeyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// RegenerateSecondaryKey regenerate secondary access key
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// accessName - the identifier of the Access configuration.
func (client TenantAccessClient) RegenerateSecondaryKey(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TenantAccessClient.RegenerateSecondaryKey")
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
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.TenantAccessClient", "RegenerateSecondaryKey", err.Error())
	}

	req, err := client.RegenerateSecondaryKeyPreparer(ctx, resourceGroupName, serviceName, accessName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "RegenerateSecondaryKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegenerateSecondaryKeySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "RegenerateSecondaryKey", resp, "Failure sending request")
		return
	}

	result, err = client.RegenerateSecondaryKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "RegenerateSecondaryKey", resp, "Failure responding to request")
		return
	}

	return
}

// RegenerateSecondaryKeyPreparer prepares the RegenerateSecondaryKey request.
func (client TenantAccessClient) RegenerateSecondaryKeyPreparer(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessName":        autorest.Encode("path", accessName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}/regenerateSecondaryKey", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegenerateSecondaryKeySender sends the RegenerateSecondaryKey request. The method will close the
// http.Response Body if it receives an error.
func (client TenantAccessClient) RegenerateSecondaryKeySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RegenerateSecondaryKeyResponder handles the response to the RegenerateSecondaryKey request. The method always
// closes the http.Response Body.
func (client TenantAccessClient) RegenerateSecondaryKeyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update update tenant access information details.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// parameters - parameters supplied to retrieve the Tenant Access Information.
// accessName - the identifier of the Access configuration.
// ifMatch - eTag of the Entity. ETag should match the current entity state from the header response of the GET
// request or it should be * for unconditional update.
func (client TenantAccessClient) Update(ctx context.Context, resourceGroupName string, serviceName string, parameters AccessInformationUpdateParameters, accessName AccessIDName, ifMatch string) (result AccessInformationContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TenantAccessClient.Update")
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
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.TenantAccessClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, serviceName, parameters, accessName, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client TenantAccessClient) UpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, parameters AccessInformationUpdateParameters, accessName AccessIDName, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessName":        autorest.Encode("path", accessName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client TenantAccessClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client TenantAccessClient) UpdateResponder(resp *http.Response) (result AccessInformationContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
