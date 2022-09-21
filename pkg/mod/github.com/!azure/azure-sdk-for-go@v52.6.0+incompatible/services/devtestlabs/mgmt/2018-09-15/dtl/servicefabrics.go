package dtl

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

// ServiceFabricsClient is the the DevTest Labs Client.
type ServiceFabricsClient struct {
	BaseClient
}

// NewServiceFabricsClient creates an instance of the ServiceFabricsClient client.
func NewServiceFabricsClient(subscriptionID string) ServiceFabricsClient {
	return NewServiceFabricsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServiceFabricsClientWithBaseURI creates an instance of the ServiceFabricsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewServiceFabricsClientWithBaseURI(baseURI string, subscriptionID string) ServiceFabricsClient {
	return ServiceFabricsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or replace an existing service fabric. This operation can take a while to complete.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// userName - the name of the user profile.
// name - the name of the service fabric.
// serviceFabric - a Service Fabric.
func (client ServiceFabricsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, userName string, name string, serviceFabric ServiceFabric) (result ServiceFabricsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceFabricsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceFabric,
			Constraints: []validation.Constraint{{Target: "serviceFabric.ServiceFabricProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "serviceFabric.ServiceFabricProperties.ApplicableSchedule", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "serviceFabric.ServiceFabricProperties.ApplicableSchedule.ApplicableScheduleProperties", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "serviceFabric.ServiceFabricProperties.ApplicableSchedule.ApplicableScheduleProperties.LabVmsShutdown", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "serviceFabric.ServiceFabricProperties.ApplicableSchedule.ApplicableScheduleProperties.LabVmsShutdown.ScheduleProperties", Name: validation.Null, Rule: true, Chain: nil}}},
							{Target: "serviceFabric.ServiceFabricProperties.ApplicableSchedule.ApplicableScheduleProperties.LabVmsStartup", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "serviceFabric.ServiceFabricProperties.ApplicableSchedule.ApplicableScheduleProperties.LabVmsStartup.ScheduleProperties", Name: validation.Null, Rule: true, Chain: nil}}},
						}},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("dtl.ServiceFabricsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, labName, userName, name, serviceFabric)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ServiceFabricsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, labName string, userName string, name string, serviceFabric ServiceFabric) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userName":          autorest.Encode("path", userName),
	}

	const APIVersion = "2018-09-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}", pathParameters),
		autorest.WithJSON(serviceFabric),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceFabricsClient) CreateOrUpdateSender(req *http.Request) (future ServiceFabricsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ServiceFabricsClient) (sf ServiceFabric, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("dtl.ServiceFabricsCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		sf.Response.Response, err = future.GetResult(sender)
		if sf.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsCreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && sf.Response.Response.StatusCode != http.StatusNoContent {
			sf, err = client.CreateOrUpdateResponder(sf.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsCreateOrUpdateFuture", "Result", sf.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ServiceFabricsClient) CreateOrUpdateResponder(resp *http.Response) (result ServiceFabric, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete service fabric. This operation can take a while to complete.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// userName - the name of the user profile.
// name - the name of the service fabric.
func (client ServiceFabricsClient) Delete(ctx context.Context, resourceGroupName string, labName string, userName string, name string) (result ServiceFabricsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceFabricsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, labName, userName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ServiceFabricsClient) DeletePreparer(ctx context.Context, resourceGroupName string, labName string, userName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userName":          autorest.Encode("path", userName),
	}

	const APIVersion = "2018-09-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceFabricsClient) DeleteSender(req *http.Request) (future ServiceFabricsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ServiceFabricsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("dtl.ServiceFabricsDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ServiceFabricsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get service fabric.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// userName - the name of the user profile.
// name - the name of the service fabric.
// expand - specify the $expand query. Example: 'properties($expand=applicableSchedule)'
func (client ServiceFabricsClient) Get(ctx context.Context, resourceGroupName string, labName string, userName string, name string, expand string) (result ServiceFabric, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceFabricsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, labName, userName, name, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ServiceFabricsClient) GetPreparer(ctx context.Context, resourceGroupName string, labName string, userName string, name string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userName":          autorest.Encode("path", userName),
	}

	const APIVersion = "2018-09-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceFabricsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServiceFabricsClient) GetResponder(resp *http.Response) (result ServiceFabric, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list service fabrics in a given user profile.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// userName - the name of the user profile.
// expand - specify the $expand query. Example: 'properties($expand=applicableSchedule)'
// filter - the filter to apply to the operation. Example: '$filter=contains(name,'myName')
// top - the maximum number of resources to return from the operation. Example: '$top=10'
// orderby - the ordering expression for the results, using OData notation. Example: '$orderby=name desc'
func (client ServiceFabricsClient) List(ctx context.Context, resourceGroupName string, labName string, userName string, expand string, filter string, top *int32, orderby string) (result ServiceFabricListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceFabricsClient.List")
		defer func() {
			sc := -1
			if result.sfl.Response.Response != nil {
				sc = result.sfl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, labName, userName, expand, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.sfl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsClient", "List", resp, "Failure sending request")
		return
	}

	result.sfl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.sfl.hasNextLink() && result.sfl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ServiceFabricsClient) ListPreparer(ctx context.Context, resourceGroupName string, labName string, userName string, expand string, filter string, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userName":          autorest.Encode("path", userName),
	}

	const APIVersion = "2018-09-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceFabricsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ServiceFabricsClient) ListResponder(resp *http.Response) (result ServiceFabricList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ServiceFabricsClient) listNextResults(ctx context.Context, lastResults ServiceFabricList) (result ServiceFabricList, err error) {
	req, err := lastResults.serviceFabricListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dtl.ServiceFabricsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dtl.ServiceFabricsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ServiceFabricsClient) ListComplete(ctx context.Context, resourceGroupName string, labName string, userName string, expand string, filter string, top *int32, orderby string) (result ServiceFabricListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceFabricsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, labName, userName, expand, filter, top, orderby)
	return
}

// ListApplicableSchedules lists the applicable start/stop schedules, if any.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// userName - the name of the user profile.
// name - the name of the service fabric.
func (client ServiceFabricsClient) ListApplicableSchedules(ctx context.Context, resourceGroupName string, labName string, userName string, name string) (result ApplicableSchedule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceFabricsClient.ListApplicableSchedules")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListApplicableSchedulesPreparer(ctx, resourceGroupName, labName, userName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsClient", "ListApplicableSchedules", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListApplicableSchedulesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsClient", "ListApplicableSchedules", resp, "Failure sending request")
		return
	}

	result, err = client.ListApplicableSchedulesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsClient", "ListApplicableSchedules", resp, "Failure responding to request")
		return
	}

	return
}

// ListApplicableSchedulesPreparer prepares the ListApplicableSchedules request.
func (client ServiceFabricsClient) ListApplicableSchedulesPreparer(ctx context.Context, resourceGroupName string, labName string, userName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userName":          autorest.Encode("path", userName),
	}

	const APIVersion = "2018-09-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}/listApplicableSchedules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListApplicableSchedulesSender sends the ListApplicableSchedules request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceFabricsClient) ListApplicableSchedulesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListApplicableSchedulesResponder handles the response to the ListApplicableSchedules request. The method always
// closes the http.Response Body.
func (client ServiceFabricsClient) ListApplicableSchedulesResponder(resp *http.Response) (result ApplicableSchedule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Start start a service fabric. This operation can take a while to complete.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// userName - the name of the user profile.
// name - the name of the service fabric.
func (client ServiceFabricsClient) Start(ctx context.Context, resourceGroupName string, labName string, userName string, name string) (result ServiceFabricsStartFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceFabricsClient.Start")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StartPreparer(ctx, resourceGroupName, labName, userName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsClient", "Start", nil, "Failure preparing request")
		return
	}

	result, err = client.StartSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsClient", "Start", nil, "Failure sending request")
		return
	}

	return
}

// StartPreparer prepares the Start request.
func (client ServiceFabricsClient) StartPreparer(ctx context.Context, resourceGroupName string, labName string, userName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userName":          autorest.Encode("path", userName),
	}

	const APIVersion = "2018-09-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceFabricsClient) StartSender(req *http.Request) (future ServiceFabricsStartFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ServiceFabricsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsStartFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("dtl.ServiceFabricsStartFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client ServiceFabricsClient) StartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Stop stop a service fabric This operation can take a while to complete.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// userName - the name of the user profile.
// name - the name of the service fabric.
func (client ServiceFabricsClient) Stop(ctx context.Context, resourceGroupName string, labName string, userName string, name string) (result ServiceFabricsStopFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceFabricsClient.Stop")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StopPreparer(ctx, resourceGroupName, labName, userName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsClient", "Stop", nil, "Failure preparing request")
		return
	}

	result, err = client.StopSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsClient", "Stop", nil, "Failure sending request")
		return
	}

	return
}

// StopPreparer prepares the Stop request.
func (client ServiceFabricsClient) StopPreparer(ctx context.Context, resourceGroupName string, labName string, userName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userName":          autorest.Encode("path", userName),
	}

	const APIVersion = "2018-09-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceFabricsClient) StopSender(req *http.Request) (future ServiceFabricsStopFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ServiceFabricsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsStopFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("dtl.ServiceFabricsStopFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client ServiceFabricsClient) StopResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update allows modifying tags of service fabrics. All other properties will be ignored.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// userName - the name of the user profile.
// name - the name of the service fabric.
// serviceFabric - a Service Fabric.
func (client ServiceFabricsClient) Update(ctx context.Context, resourceGroupName string, labName string, userName string, name string, serviceFabric ServiceFabricFragment) (result ServiceFabric, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceFabricsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, labName, userName, name, serviceFabric)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ServiceFabricsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ServiceFabricsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, labName string, userName string, name string, serviceFabric ServiceFabricFragment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userName":          autorest.Encode("path", userName),
	}

	const APIVersion = "2018-09-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}", pathParameters),
		autorest.WithJSON(serviceFabric),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceFabricsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ServiceFabricsClient) UpdateResponder(resp *http.Response) (result ServiceFabric, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
