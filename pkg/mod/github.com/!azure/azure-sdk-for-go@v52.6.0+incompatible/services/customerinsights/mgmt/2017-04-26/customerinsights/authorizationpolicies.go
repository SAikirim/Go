package customerinsights

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

// AuthorizationPoliciesClient is the the Azure Customer Insights management API provides a RESTful set of web services
// that interact with Azure Customer Insights service to manage your resources. The API has entities that capture the
// relationship between an end user and the Azure Customer Insights service.
type AuthorizationPoliciesClient struct {
	BaseClient
}

// NewAuthorizationPoliciesClient creates an instance of the AuthorizationPoliciesClient client.
func NewAuthorizationPoliciesClient(subscriptionID string) AuthorizationPoliciesClient {
	return NewAuthorizationPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAuthorizationPoliciesClientWithBaseURI creates an instance of the AuthorizationPoliciesClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewAuthorizationPoliciesClientWithBaseURI(baseURI string, subscriptionID string) AuthorizationPoliciesClient {
	return AuthorizationPoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates an authorization policy or updates an existing authorization policy.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// authorizationPolicyName - the name of the policy.
// parameters - parameters supplied to the CreateOrUpdate authorization policy operation.
func (client AuthorizationPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, authorizationPolicyName string, parameters AuthorizationPolicyResourceFormat) (result AuthorizationPolicyResourceFormat, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AuthorizationPoliciesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: authorizationPolicyName,
			Constraints: []validation.Constraint{{Target: "authorizationPolicyName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "authorizationPolicyName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "authorizationPolicyName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]$|^[A-Za-z0-9][\w-\.]*[A-Za-z0-9]$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.AuthorizationPolicy", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.AuthorizationPolicy.Permissions", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.AuthorizationPolicy.Permissions", Name: validation.UniqueItems, Rule: true, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("customerinsights.AuthorizationPoliciesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, hubName, authorizationPolicyName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.AuthorizationPoliciesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.AuthorizationPoliciesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.AuthorizationPoliciesClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client AuthorizationPoliciesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, hubName string, authorizationPolicyName string, parameters AuthorizationPolicyResourceFormat) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"authorizationPolicyName": autorest.Encode("path", authorizationPolicyName),
		"hubName":                 autorest.Encode("path", hubName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/authorizationPolicies/{authorizationPolicyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client AuthorizationPoliciesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client AuthorizationPoliciesClient) CreateOrUpdateResponder(resp *http.Response) (result AuthorizationPolicyResourceFormat, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets an authorization policy in the hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// authorizationPolicyName - the name of the policy.
func (client AuthorizationPoliciesClient) Get(ctx context.Context, resourceGroupName string, hubName string, authorizationPolicyName string) (result AuthorizationPolicyResourceFormat, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AuthorizationPoliciesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, hubName, authorizationPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.AuthorizationPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.AuthorizationPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.AuthorizationPoliciesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client AuthorizationPoliciesClient) GetPreparer(ctx context.Context, resourceGroupName string, hubName string, authorizationPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"authorizationPolicyName": autorest.Encode("path", authorizationPolicyName),
		"hubName":                 autorest.Encode("path", hubName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/authorizationPolicies/{authorizationPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AuthorizationPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AuthorizationPoliciesClient) GetResponder(resp *http.Response) (result AuthorizationPolicyResourceFormat, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByHub gets all the authorization policies in a specified hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
func (client AuthorizationPoliciesClient) ListByHub(ctx context.Context, resourceGroupName string, hubName string) (result AuthorizationPolicyListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AuthorizationPoliciesClient.ListByHub")
		defer func() {
			sc := -1
			if result.aplr.Response.Response != nil {
				sc = result.aplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByHubNextResults
	req, err := client.ListByHubPreparer(ctx, resourceGroupName, hubName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.AuthorizationPoliciesClient", "ListByHub", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByHubSender(req)
	if err != nil {
		result.aplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.AuthorizationPoliciesClient", "ListByHub", resp, "Failure sending request")
		return
	}

	result.aplr, err = client.ListByHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.AuthorizationPoliciesClient", "ListByHub", resp, "Failure responding to request")
		return
	}
	if result.aplr.hasNextLink() && result.aplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByHubPreparer prepares the ListByHub request.
func (client AuthorizationPoliciesClient) ListByHubPreparer(ctx context.Context, resourceGroupName string, hubName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/authorizationPolicies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByHubSender sends the ListByHub request. The method will close the
// http.Response Body if it receives an error.
func (client AuthorizationPoliciesClient) ListByHubSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByHubResponder handles the response to the ListByHub request. The method always
// closes the http.Response Body.
func (client AuthorizationPoliciesClient) ListByHubResponder(resp *http.Response) (result AuthorizationPolicyListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByHubNextResults retrieves the next set of results, if any.
func (client AuthorizationPoliciesClient) listByHubNextResults(ctx context.Context, lastResults AuthorizationPolicyListResult) (result AuthorizationPolicyListResult, err error) {
	req, err := lastResults.authorizationPolicyListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "customerinsights.AuthorizationPoliciesClient", "listByHubNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByHubSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "customerinsights.AuthorizationPoliciesClient", "listByHubNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.AuthorizationPoliciesClient", "listByHubNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByHubComplete enumerates all values, automatically crossing page boundaries as required.
func (client AuthorizationPoliciesClient) ListByHubComplete(ctx context.Context, resourceGroupName string, hubName string) (result AuthorizationPolicyListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AuthorizationPoliciesClient.ListByHub")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByHub(ctx, resourceGroupName, hubName)
	return
}

// RegeneratePrimaryKey regenerates the primary policy key of the specified authorization policy.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// authorizationPolicyName - the name of the policy.
func (client AuthorizationPoliciesClient) RegeneratePrimaryKey(ctx context.Context, resourceGroupName string, hubName string, authorizationPolicyName string) (result AuthorizationPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AuthorizationPoliciesClient.RegeneratePrimaryKey")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RegeneratePrimaryKeyPreparer(ctx, resourceGroupName, hubName, authorizationPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.AuthorizationPoliciesClient", "RegeneratePrimaryKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegeneratePrimaryKeySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.AuthorizationPoliciesClient", "RegeneratePrimaryKey", resp, "Failure sending request")
		return
	}

	result, err = client.RegeneratePrimaryKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.AuthorizationPoliciesClient", "RegeneratePrimaryKey", resp, "Failure responding to request")
		return
	}

	return
}

// RegeneratePrimaryKeyPreparer prepares the RegeneratePrimaryKey request.
func (client AuthorizationPoliciesClient) RegeneratePrimaryKeyPreparer(ctx context.Context, resourceGroupName string, hubName string, authorizationPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"authorizationPolicyName": autorest.Encode("path", authorizationPolicyName),
		"hubName":                 autorest.Encode("path", hubName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/authorizationPolicies/{authorizationPolicyName}/regeneratePrimaryKey", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegeneratePrimaryKeySender sends the RegeneratePrimaryKey request. The method will close the
// http.Response Body if it receives an error.
func (client AuthorizationPoliciesClient) RegeneratePrimaryKeySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RegeneratePrimaryKeyResponder handles the response to the RegeneratePrimaryKey request. The method always
// closes the http.Response Body.
func (client AuthorizationPoliciesClient) RegeneratePrimaryKeyResponder(resp *http.Response) (result AuthorizationPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RegenerateSecondaryKey regenerates the secondary policy key of the specified authorization policy.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// authorizationPolicyName - the name of the policy.
func (client AuthorizationPoliciesClient) RegenerateSecondaryKey(ctx context.Context, resourceGroupName string, hubName string, authorizationPolicyName string) (result AuthorizationPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AuthorizationPoliciesClient.RegenerateSecondaryKey")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RegenerateSecondaryKeyPreparer(ctx, resourceGroupName, hubName, authorizationPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.AuthorizationPoliciesClient", "RegenerateSecondaryKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegenerateSecondaryKeySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.AuthorizationPoliciesClient", "RegenerateSecondaryKey", resp, "Failure sending request")
		return
	}

	result, err = client.RegenerateSecondaryKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.AuthorizationPoliciesClient", "RegenerateSecondaryKey", resp, "Failure responding to request")
		return
	}

	return
}

// RegenerateSecondaryKeyPreparer prepares the RegenerateSecondaryKey request.
func (client AuthorizationPoliciesClient) RegenerateSecondaryKeyPreparer(ctx context.Context, resourceGroupName string, hubName string, authorizationPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"authorizationPolicyName": autorest.Encode("path", authorizationPolicyName),
		"hubName":                 autorest.Encode("path", hubName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/authorizationPolicies/{authorizationPolicyName}/regenerateSecondaryKey", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegenerateSecondaryKeySender sends the RegenerateSecondaryKey request. The method will close the
// http.Response Body if it receives an error.
func (client AuthorizationPoliciesClient) RegenerateSecondaryKeySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RegenerateSecondaryKeyResponder handles the response to the RegenerateSecondaryKey request. The method always
// closes the http.Response Body.
func (client AuthorizationPoliciesClient) RegenerateSecondaryKeyResponder(resp *http.Response) (result AuthorizationPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
