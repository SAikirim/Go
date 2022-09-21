package support

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

// TicketsClient is the microsoft Azure Support Resource Provider.
type TicketsClient struct {
	BaseClient
}

// NewTicketsClient creates an instance of the TicketsClient client.
func NewTicketsClient(subscriptionID string) TicketsClient {
	return NewTicketsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTicketsClientWithBaseURI creates an instance of the TicketsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewTicketsClientWithBaseURI(baseURI string, subscriptionID string) TicketsClient {
	return TicketsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckNameAvailability check the availability of a resource name. This API should be used to check the uniqueness of
// the name for support ticket creation for the selected subscription.
// Parameters:
// checkNameAvailabilityInput - input to check.
func (client TicketsClient) CheckNameAvailability(ctx context.Context, checkNameAvailabilityInput CheckNameAvailabilityInput) (result CheckNameAvailabilityOutput, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TicketsClient.CheckNameAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: checkNameAvailabilityInput,
			Constraints: []validation.Constraint{{Target: "checkNameAvailabilityInput.Name", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("support.TicketsClient", "CheckNameAvailability", err.Error())
	}

	req, err := client.CheckNameAvailabilityPreparer(ctx, checkNameAvailabilityInput)
	if err != nil {
		err = autorest.NewErrorWithError(err, "support.TicketsClient", "CheckNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "support.TicketsClient", "CheckNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "support.TicketsClient", "CheckNameAvailability", resp, "Failure responding to request")
		return
	}

	return
}

// CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client TicketsClient) CheckNameAvailabilityPreparer(ctx context.Context, checkNameAvailabilityInput CheckNameAvailabilityInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Support/checkNameAvailability", pathParameters),
		autorest.WithJSON(checkNameAvailabilityInput),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client TicketsClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client TicketsClient) CheckNameAvailabilityResponder(resp *http.Response) (result CheckNameAvailabilityOutput, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create creates a new support ticket for Subscription and Service limits (Quota), Technical, Billing, and
// Subscription Management issues for the specified subscription. Learn the [prerequisites](https://aka.ms/supportAPI)
// required to create a support ticket.<br/><br/>Always call the Services and ProblemClassifications API to get the
// most recent set of services and problem categories required for support ticket creation.<br/><br/>Adding attachments
// is not currently supported via the API. To add a file to an existing support ticket, visit the [Manage support
// ticket](https://portal.azure.com/#blade/Microsoft_Azure_Support/HelpAndSupportBlade/managesupportrequest) page in
// the Azure portal, select the support ticket, and use the file upload control to add a new file.<br/><br/>Providing
// consent to share diagnostic information with Azure support is currently not supported via the API. The Azure support
// engineer working on your ticket will reach out to you for consent if your issue requires gathering diagnostic
// information from your Azure resources.<br/><br/>**Creating a support ticket for on-behalf-of**: Include
// _x-ms-authorization-auxiliary_ header to provide an auxiliary token as per
// [documentation](https://docs.microsoft.com/azure/azure-resource-manager/management/authenticate-multi-tenant). The
// primary token will be from the tenant for whom a support ticket is being raised against the subscription, i.e. Cloud
// solution provider (CSP) customer tenant. The auxiliary token will be from the Cloud solution provider (CSP) partner
// tenant.
// Parameters:
// supportTicketName - support ticket name.
// createSupportTicketParameters - support ticket request payload.
func (client TicketsClient) Create(ctx context.Context, supportTicketName string, createSupportTicketParameters TicketDetails) (result TicketsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TicketsClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: createSupportTicketParameters,
			Constraints: []validation.Constraint{{Target: "createSupportTicketParameters.TicketDetailsProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "createSupportTicketParameters.TicketDetailsProperties.Description", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "createSupportTicketParameters.TicketDetailsProperties.ProblemClassificationID", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "createSupportTicketParameters.TicketDetailsProperties.ContactDetails", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "createSupportTicketParameters.TicketDetailsProperties.ContactDetails.FirstName", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "createSupportTicketParameters.TicketDetailsProperties.ContactDetails.LastName", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "createSupportTicketParameters.TicketDetailsProperties.ContactDetails.PrimaryEmailAddress", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "createSupportTicketParameters.TicketDetailsProperties.ContactDetails.PreferredTimeZone", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "createSupportTicketParameters.TicketDetailsProperties.ContactDetails.Country", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "createSupportTicketParameters.TicketDetailsProperties.ContactDetails.PreferredSupportLanguage", Name: validation.Null, Rule: true, Chain: nil},
						}},
					{Target: "createSupportTicketParameters.TicketDetailsProperties.Title", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "createSupportTicketParameters.TicketDetailsProperties.ServiceID", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("support.TicketsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, supportTicketName, createSupportTicketParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "support.TicketsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "support.TicketsClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client TicketsClient) CreatePreparer(ctx context.Context, supportTicketName string, createSupportTicketParameters TicketDetails) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"supportTicketName": autorest.Encode("path", supportTicketName),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	createSupportTicketParameters.ID = nil
	createSupportTicketParameters.Name = nil
	createSupportTicketParameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets/{supportTicketName}", pathParameters),
		autorest.WithJSON(createSupportTicketParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client TicketsClient) CreateSender(req *http.Request) (future TicketsCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client TicketsClient) (td TicketDetails, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "support.TicketsCreateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("support.TicketsCreateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		td.Response.Response, err = future.GetResult(sender)
		if td.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "support.TicketsCreateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && td.Response.Response.StatusCode != http.StatusNoContent {
			td, err = client.CreateResponder(td.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "support.TicketsCreateFuture", "Result", td.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client TicketsClient) CreateResponder(resp *http.Response) (result TicketDetails, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get ticket details for an Azure subscription. Support ticket data is available for 18 months after ticket
// creation. If a ticket was created more than 18 months ago, a request for data might cause an error.
// Parameters:
// supportTicketName - support ticket name.
func (client TicketsClient) Get(ctx context.Context, supportTicketName string) (result TicketDetails, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TicketsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, supportTicketName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "support.TicketsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "support.TicketsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "support.TicketsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client TicketsClient) GetPreparer(ctx context.Context, supportTicketName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"supportTicketName": autorest.Encode("path", supportTicketName),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets/{supportTicketName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client TicketsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client TicketsClient) GetResponder(resp *http.Response) (result TicketDetails, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all the support tickets for an Azure subscription. You can also filter the support tickets by _Status_ or
// _CreatedDate_ using the $filter parameter. Output will be a paged result with _nextLink_, using which you can
// retrieve the next set of support tickets. <br/><br/>Support ticket data is available for 18 months after ticket
// creation. If a ticket was created more than 18 months ago, a request for data might cause an error.
// Parameters:
// top - the number of values to return in the collection. Default is 25 and max is 100.
// filter - the filter to apply on the operation. We support 'odata v4.0' filter semantics. [Learn
// more](https://docs.microsoft.com/odata/concepts/queryoptions-overview). _Status_ filter can only be used
// with Equals ('eq') operator. For _CreatedDate_ filter, the supported operators are Greater Than ('gt') and
// Greater Than or Equals ('ge'). When using both filters, combine them using the logical 'AND'.
func (client TicketsClient) List(ctx context.Context, top *int32, filter string) (result TicketsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TicketsClient.List")
		defer func() {
			sc := -1
			if result.tlr.Response.Response != nil {
				sc = result.tlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, top, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "support.TicketsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.tlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "support.TicketsClient", "List", resp, "Failure sending request")
		return
	}

	result.tlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "support.TicketsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.tlr.hasNextLink() && result.tlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client TicketsClient) ListPreparer(ctx context.Context, top *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client TicketsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client TicketsClient) ListResponder(resp *http.Response) (result TicketsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client TicketsClient) listNextResults(ctx context.Context, lastResults TicketsListResult) (result TicketsListResult, err error) {
	req, err := lastResults.ticketsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "support.TicketsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "support.TicketsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "support.TicketsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client TicketsClient) ListComplete(ctx context.Context, top *int32, filter string) (result TicketsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TicketsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, top, filter)
	return
}

// Update this API allows you to update the severity level, ticket status, and your contact information in the support
// ticket.<br/><br/>Note: The severity levels cannot be changed if a support ticket is actively being worked upon by an
// Azure support engineer. In such a case, contact your support engineer to request severity update by adding a new
// communication using the Communications API.<br/><br/>Changing the ticket status to _closed_ is allowed only on an
// unassigned case. When an engineer is actively working on the ticket, send your ticket closure request by sending a
// note to your engineer.
// Parameters:
// supportTicketName - support ticket name.
// updateSupportTicket - updateSupportTicket object.
func (client TicketsClient) Update(ctx context.Context, supportTicketName string, updateSupportTicket UpdateSupportTicket) (result TicketDetails, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TicketsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, supportTicketName, updateSupportTicket)
	if err != nil {
		err = autorest.NewErrorWithError(err, "support.TicketsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "support.TicketsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "support.TicketsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client TicketsClient) UpdatePreparer(ctx context.Context, supportTicketName string, updateSupportTicket UpdateSupportTicket) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"supportTicketName": autorest.Encode("path", supportTicketName),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets/{supportTicketName}", pathParameters),
		autorest.WithJSON(updateSupportTicket),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client TicketsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client TicketsClient) UpdateResponder(resp *http.Response) (result TicketDetails, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
