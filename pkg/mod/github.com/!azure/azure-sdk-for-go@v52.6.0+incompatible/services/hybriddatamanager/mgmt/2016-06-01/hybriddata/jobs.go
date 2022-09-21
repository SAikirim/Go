package hybriddata

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

// JobsClient is the client for the Jobs methods of the Hybriddata service.
type JobsClient struct {
	BaseClient
}

// NewJobsClient creates an instance of the JobsClient client.
func NewJobsClient(subscriptionID string) JobsClient {
	return NewJobsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewJobsClientWithBaseURI creates an instance of the JobsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewJobsClientWithBaseURI(baseURI string, subscriptionID string) JobsClient {
	return JobsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Cancel cancels the given job.
// Parameters:
// dataServiceName - the name of the data service of the job definition.
// jobDefinitionName - the name of the job definition of the job.
// jobID - the job id of the job queried.
// resourceGroupName - the Resource Group Name
// dataManagerName - the name of the DataManager Resource within the specified resource group. DataManager
// names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
func (client JobsClient) Cancel(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string) (result JobsCancelFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.Cancel")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dataManagerName,
			Constraints: []validation.Constraint{{Target: "dataManagerName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "dataManagerName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "dataManagerName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybriddata.JobsClient", "Cancel", err.Error())
	}

	req, err := client.CancelPreparer(ctx, dataServiceName, jobDefinitionName, jobID, resourceGroupName, dataManagerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobsClient", "Cancel", nil, "Failure preparing request")
		return
	}

	result, err = client.CancelSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobsClient", "Cancel", nil, "Failure sending request")
		return
	}

	return
}

// CancelPreparer prepares the Cancel request.
func (client JobsClient) CancelPreparer(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataManagerName":   autorest.Encode("path", dataManagerName),
		"dataServiceName":   autorest.Encode("path", dataServiceName),
		"jobDefinitionName": autorest.Encode("path", jobDefinitionName),
		"jobId":             autorest.Encode("path", jobID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}/jobs/{jobId}/cancel", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) CancelSender(req *http.Request) (future JobsCancelFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client JobsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "hybriddata.JobsCancelFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("hybriddata.JobsCancelFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client JobsClient) CancelResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get this method gets a data manager job given the jobId.
// Parameters:
// dataServiceName - the name of the data service of the job definition.
// jobDefinitionName - the name of the job definition of the job.
// jobID - the job id of the job queried.
// resourceGroupName - the Resource Group Name
// dataManagerName - the name of the DataManager Resource within the specified resource group. DataManager
// names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
// expand - $expand is supported on details parameter for job, which provides details on the job stages.
func (client JobsClient) Get(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string, expand string) (result Job, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dataManagerName,
			Constraints: []validation.Constraint{{Target: "dataManagerName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "dataManagerName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "dataManagerName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybriddata.JobsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, dataServiceName, jobDefinitionName, jobID, resourceGroupName, dataManagerName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hybriddata.JobsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client JobsClient) GetPreparer(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataManagerName":   autorest.Encode("path", dataManagerName),
		"dataServiceName":   autorest.Encode("path", dataServiceName),
		"jobDefinitionName": autorest.Encode("path", jobDefinitionName),
		"jobId":             autorest.Encode("path", jobID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}/jobs/{jobId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client JobsClient) GetResponder(resp *http.Response) (result Job, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDataManager this method gets all the jobs at the data manager resource level.
// Parameters:
// resourceGroupName - the Resource Group Name
// dataManagerName - the name of the DataManager Resource within the specified resource group. DataManager
// names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
// filter - oData Filter options
func (client JobsClient) ListByDataManager(ctx context.Context, resourceGroupName string, dataManagerName string, filter string) (result JobListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.ListByDataManager")
		defer func() {
			sc := -1
			if result.jl.Response.Response != nil {
				sc = result.jl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dataManagerName,
			Constraints: []validation.Constraint{{Target: "dataManagerName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "dataManagerName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "dataManagerName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybriddata.JobsClient", "ListByDataManager", err.Error())
	}

	result.fn = client.listByDataManagerNextResults
	req, err := client.ListByDataManagerPreparer(ctx, resourceGroupName, dataManagerName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobsClient", "ListByDataManager", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDataManagerSender(req)
	if err != nil {
		result.jl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hybriddata.JobsClient", "ListByDataManager", resp, "Failure sending request")
		return
	}

	result.jl, err = client.ListByDataManagerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobsClient", "ListByDataManager", resp, "Failure responding to request")
		return
	}
	if result.jl.hasNextLink() && result.jl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByDataManagerPreparer prepares the ListByDataManager request.
func (client JobsClient) ListByDataManagerPreparer(ctx context.Context, resourceGroupName string, dataManagerName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataManagerName":   autorest.Encode("path", dataManagerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/jobs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDataManagerSender sends the ListByDataManager request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) ListByDataManagerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByDataManagerResponder handles the response to the ListByDataManager request. The method always
// closes the http.Response Body.
func (client JobsClient) ListByDataManagerResponder(resp *http.Response) (result JobList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByDataManagerNextResults retrieves the next set of results, if any.
func (client JobsClient) listByDataManagerNextResults(ctx context.Context, lastResults JobList) (result JobList, err error) {
	req, err := lastResults.jobListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "hybriddata.JobsClient", "listByDataManagerNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByDataManagerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "hybriddata.JobsClient", "listByDataManagerNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByDataManagerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobsClient", "listByDataManagerNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByDataManagerComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobsClient) ListByDataManagerComplete(ctx context.Context, resourceGroupName string, dataManagerName string, filter string) (result JobListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.ListByDataManager")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByDataManager(ctx, resourceGroupName, dataManagerName, filter)
	return
}

// ListByDataService this method gets all the jobs of a data service type in a given resource.
// Parameters:
// dataServiceName - the name of the data service of interest.
// resourceGroupName - the Resource Group Name
// dataManagerName - the name of the DataManager Resource within the specified resource group. DataManager
// names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
// filter - oData Filter options
func (client JobsClient) ListByDataService(ctx context.Context, dataServiceName string, resourceGroupName string, dataManagerName string, filter string) (result JobListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.ListByDataService")
		defer func() {
			sc := -1
			if result.jl.Response.Response != nil {
				sc = result.jl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dataManagerName,
			Constraints: []validation.Constraint{{Target: "dataManagerName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "dataManagerName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "dataManagerName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybriddata.JobsClient", "ListByDataService", err.Error())
	}

	result.fn = client.listByDataServiceNextResults
	req, err := client.ListByDataServicePreparer(ctx, dataServiceName, resourceGroupName, dataManagerName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobsClient", "ListByDataService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDataServiceSender(req)
	if err != nil {
		result.jl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hybriddata.JobsClient", "ListByDataService", resp, "Failure sending request")
		return
	}

	result.jl, err = client.ListByDataServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobsClient", "ListByDataService", resp, "Failure responding to request")
		return
	}
	if result.jl.hasNextLink() && result.jl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByDataServicePreparer prepares the ListByDataService request.
func (client JobsClient) ListByDataServicePreparer(ctx context.Context, dataServiceName string, resourceGroupName string, dataManagerName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataManagerName":   autorest.Encode("path", dataManagerName),
		"dataServiceName":   autorest.Encode("path", dataServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDataServiceSender sends the ListByDataService request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) ListByDataServiceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByDataServiceResponder handles the response to the ListByDataService request. The method always
// closes the http.Response Body.
func (client JobsClient) ListByDataServiceResponder(resp *http.Response) (result JobList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByDataServiceNextResults retrieves the next set of results, if any.
func (client JobsClient) listByDataServiceNextResults(ctx context.Context, lastResults JobList) (result JobList, err error) {
	req, err := lastResults.jobListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "hybriddata.JobsClient", "listByDataServiceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByDataServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "hybriddata.JobsClient", "listByDataServiceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByDataServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobsClient", "listByDataServiceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByDataServiceComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobsClient) ListByDataServiceComplete(ctx context.Context, dataServiceName string, resourceGroupName string, dataManagerName string, filter string) (result JobListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.ListByDataService")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByDataService(ctx, dataServiceName, resourceGroupName, dataManagerName, filter)
	return
}

// ListByJobDefinition this method gets all the jobs of a given job definition.
// Parameters:
// dataServiceName - the name of the data service of the job definition.
// jobDefinitionName - the name of the job definition for which jobs are needed.
// resourceGroupName - the Resource Group Name
// dataManagerName - the name of the DataManager Resource within the specified resource group. DataManager
// names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
// filter - oData Filter options
func (client JobsClient) ListByJobDefinition(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string, filter string) (result JobListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.ListByJobDefinition")
		defer func() {
			sc := -1
			if result.jl.Response.Response != nil {
				sc = result.jl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dataManagerName,
			Constraints: []validation.Constraint{{Target: "dataManagerName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "dataManagerName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "dataManagerName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybriddata.JobsClient", "ListByJobDefinition", err.Error())
	}

	result.fn = client.listByJobDefinitionNextResults
	req, err := client.ListByJobDefinitionPreparer(ctx, dataServiceName, jobDefinitionName, resourceGroupName, dataManagerName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobsClient", "ListByJobDefinition", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByJobDefinitionSender(req)
	if err != nil {
		result.jl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hybriddata.JobsClient", "ListByJobDefinition", resp, "Failure sending request")
		return
	}

	result.jl, err = client.ListByJobDefinitionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobsClient", "ListByJobDefinition", resp, "Failure responding to request")
		return
	}
	if result.jl.hasNextLink() && result.jl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByJobDefinitionPreparer prepares the ListByJobDefinition request.
func (client JobsClient) ListByJobDefinitionPreparer(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataManagerName":   autorest.Encode("path", dataManagerName),
		"dataServiceName":   autorest.Encode("path", dataServiceName),
		"jobDefinitionName": autorest.Encode("path", jobDefinitionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}/jobs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByJobDefinitionSender sends the ListByJobDefinition request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) ListByJobDefinitionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByJobDefinitionResponder handles the response to the ListByJobDefinition request. The method always
// closes the http.Response Body.
func (client JobsClient) ListByJobDefinitionResponder(resp *http.Response) (result JobList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByJobDefinitionNextResults retrieves the next set of results, if any.
func (client JobsClient) listByJobDefinitionNextResults(ctx context.Context, lastResults JobList) (result JobList, err error) {
	req, err := lastResults.jobListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "hybriddata.JobsClient", "listByJobDefinitionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByJobDefinitionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "hybriddata.JobsClient", "listByJobDefinitionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByJobDefinitionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobsClient", "listByJobDefinitionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByJobDefinitionComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobsClient) ListByJobDefinitionComplete(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string, filter string) (result JobListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.ListByJobDefinition")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByJobDefinition(ctx, dataServiceName, jobDefinitionName, resourceGroupName, dataManagerName, filter)
	return
}

// Resume resumes the given job.
// Parameters:
// dataServiceName - the name of the data service of the job definition.
// jobDefinitionName - the name of the job definition of the job.
// jobID - the job id of the job queried.
// resourceGroupName - the Resource Group Name
// dataManagerName - the name of the DataManager Resource within the specified resource group. DataManager
// names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
func (client JobsClient) Resume(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string) (result JobsResumeFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.Resume")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dataManagerName,
			Constraints: []validation.Constraint{{Target: "dataManagerName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "dataManagerName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "dataManagerName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybriddata.JobsClient", "Resume", err.Error())
	}

	req, err := client.ResumePreparer(ctx, dataServiceName, jobDefinitionName, jobID, resourceGroupName, dataManagerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobsClient", "Resume", nil, "Failure preparing request")
		return
	}

	result, err = client.ResumeSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.JobsClient", "Resume", nil, "Failure sending request")
		return
	}

	return
}

// ResumePreparer prepares the Resume request.
func (client JobsClient) ResumePreparer(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataManagerName":   autorest.Encode("path", dataManagerName),
		"dataServiceName":   autorest.Encode("path", dataServiceName),
		"jobDefinitionName": autorest.Encode("path", jobDefinitionName),
		"jobId":             autorest.Encode("path", jobID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}/jobs/{jobId}/resume", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResumeSender sends the Resume request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) ResumeSender(req *http.Request) (future JobsResumeFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client JobsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "hybriddata.JobsResumeFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("hybriddata.JobsResumeFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// ResumeResponder handles the response to the Resume request. The method always
// closes the http.Response Body.
func (client JobsClient) ResumeResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}