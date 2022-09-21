package subscription

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

// Client is the the subscription client
type Client struct {
	BaseClient
}

// NewClient creates an instance of the Client client.
func NewClient() Client {
	return NewClientWithBaseURI(DefaultBaseURI)
}

// NewClientWithBaseURI creates an instance of the Client client using a custom endpoint.  Use this when interacting
// with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewClientWithBaseURI(baseURI string) Client {
	return Client{NewWithBaseURI(baseURI)}
}

// Cancel the operation to cancel a subscription
// Parameters:
// subscriptionID - subscription Id.
func (client Client) Cancel(ctx context.Context, subscriptionID string) (result CanceledSubscriptionID, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Cancel")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CancelPreparer(ctx, subscriptionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "Cancel", nil, "Failure preparing request")
		return
	}

	resp, err := client.CancelSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "subscription.Client", "Cancel", resp, "Failure sending request")
		return
	}

	result, err = client.CancelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "Cancel", resp, "Failure responding to request")
		return
	}

	return
}

// CancelPreparer prepares the Cancel request.
func (client Client) CancelPreparer(ctx context.Context, subscriptionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2019-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Subscription/cancel", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client Client) CancelSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client Client) CancelResponder(resp *http.Response) (result CanceledSubscriptionID, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateAlias create Alias Subscription.
// Parameters:
// aliasName - alias Name
func (client Client) CreateAlias(ctx context.Context, aliasName string, body PutAliasRequest) (result CreateAliasFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.CreateAlias")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Properties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "body.Properties.DisplayName", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "body.Properties.BillingScope", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("subscription.Client", "CreateAlias", err.Error())
	}

	req, err := client.CreateAliasPreparer(ctx, aliasName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "CreateAlias", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateAliasSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "CreateAlias", nil, "Failure sending request")
		return
	}

	return
}

// CreateAliasPreparer prepares the CreateAlias request.
func (client Client) CreateAliasPreparer(ctx context.Context, aliasName string, body PutAliasRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"aliasName": autorest.Encode("path", aliasName),
	}

	const APIVersion = "2019-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Subscription/aliases/{aliasName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateAliasSender sends the CreateAlias request. The method will close the
// http.Response Body if it receives an error.
func (client Client) CreateAliasSender(req *http.Request) (future CreateAliasFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client Client) (par PutAliasResponse, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "subscription.CreateAliasFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("subscription.CreateAliasFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		par.Response.Response, err = future.GetResult(sender)
		if par.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "subscription.CreateAliasFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && par.Response.Response.StatusCode != http.StatusNoContent {
			par, err = client.CreateAliasResponder(par.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "subscription.CreateAliasFuture", "Result", par.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateAliasResponder handles the response to the CreateAlias request. The method always
// closes the http.Response Body.
func (client Client) CreateAliasResponder(resp *http.Response) (result PutAliasResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateCspSubscription the operation to create a new CSP subscription.
// Parameters:
// billingAccountName - the name of the Microsoft Customer Agreement billing account for which you want to
// create the subscription.
// customerName - the name of the customer.
// body - the subscription creation parameters.
func (client Client) CreateCspSubscription(ctx context.Context, billingAccountName string, customerName string, body ModernCspSubscriptionCreationParameters) (result CreateCspSubscriptionFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.CreateCspSubscription")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.DisplayName", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "body.SkuID", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("subscription.Client", "CreateCspSubscription", err.Error())
	}

	req, err := client.CreateCspSubscriptionPreparer(ctx, billingAccountName, customerName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "CreateCspSubscription", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateCspSubscriptionSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "CreateCspSubscription", nil, "Failure sending request")
		return
	}

	return
}

// CreateCspSubscriptionPreparer prepares the CreateCspSubscription request.
func (client Client) CreateCspSubscriptionPreparer(ctx context.Context, billingAccountName string, customerName string, body ModernCspSubscriptionCreationParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"customerName":       autorest.Encode("path", customerName),
	}

	const APIVersion = "2019-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}/providers/Microsoft.Subscription/createSubscription", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateCspSubscriptionSender sends the CreateCspSubscription request. The method will close the
// http.Response Body if it receives an error.
func (client Client) CreateCspSubscriptionSender(req *http.Request) (future CreateCspSubscriptionFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client Client) (cr CreationResult, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "subscription.CreateCspSubscriptionFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("subscription.CreateCspSubscriptionFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		cr.Response.Response, err = future.GetResult(sender)
		if cr.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "subscription.CreateCspSubscriptionFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && cr.Response.Response.StatusCode != http.StatusNoContent {
			cr, err = client.CreateCspSubscriptionResponder(cr.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "subscription.CreateCspSubscriptionFuture", "Result", cr.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateCspSubscriptionResponder handles the response to the CreateCspSubscription request. The method always
// closes the http.Response Body.
func (client Client) CreateCspSubscriptionResponder(resp *http.Response) (result CreationResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateSubscription the operation to create a new WebDirect or EA Azure subscription.
// Parameters:
// billingAccountName - the name of the Microsoft Customer Agreement billing account for which you want to
// create the subscription.
// billingProfileName - the name of the billing profile in the billing account for which you want to create the
// subscription.
// invoiceSectionName - the name of the invoice section in the billing account for which you want to create the
// subscription.
// body - the subscription creation parameters.
func (client Client) CreateSubscription(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, body ModernSubscriptionCreationParameters) (result CreateSubscriptionFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.CreateSubscription")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.DisplayName", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "body.SkuID", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "body.Owner", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.Owner.ObjectID", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("subscription.Client", "CreateSubscription", err.Error())
	}

	req, err := client.CreateSubscriptionPreparer(ctx, billingAccountName, billingProfileName, invoiceSectionName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "CreateSubscription", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSubscriptionSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "CreateSubscription", nil, "Failure sending request")
		return
	}

	return
}

// CreateSubscriptionPreparer prepares the CreateSubscription request.
func (client Client) CreateSubscriptionPreparer(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, body ModernSubscriptionCreationParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
		"invoiceSectionName": autorest.Encode("path", invoiceSectionName),
	}

	const APIVersion = "2019-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/providers/Microsoft.Subscription/createSubscription", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSubscriptionSender sends the CreateSubscription request. The method will close the
// http.Response Body if it receives an error.
func (client Client) CreateSubscriptionSender(req *http.Request) (future CreateSubscriptionFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client Client) (cr CreationResult, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "subscription.CreateSubscriptionFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("subscription.CreateSubscriptionFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		cr.Response.Response, err = future.GetResult(sender)
		if cr.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "subscription.CreateSubscriptionFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && cr.Response.Response.StatusCode != http.StatusNoContent {
			cr, err = client.CreateSubscriptionResponder(cr.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "subscription.CreateSubscriptionFuture", "Result", cr.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateSubscriptionResponder handles the response to the CreateSubscription request. The method always
// closes the http.Response Body.
func (client Client) CreateSubscriptionResponder(resp *http.Response) (result CreationResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateSubscriptionInEnrollmentAccount creates an Azure subscription
// Parameters:
// enrollmentAccountName - the name of the enrollment account to which the subscription will be billed.
// body - the subscription creation parameters.
func (client Client) CreateSubscriptionInEnrollmentAccount(ctx context.Context, enrollmentAccountName string, body CreationParameters) (result CreateSubscriptionInEnrollmentAccountFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.CreateSubscriptionInEnrollmentAccount")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateSubscriptionInEnrollmentAccountPreparer(ctx, enrollmentAccountName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "CreateSubscriptionInEnrollmentAccount", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSubscriptionInEnrollmentAccountSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "CreateSubscriptionInEnrollmentAccount", nil, "Failure sending request")
		return
	}

	return
}

// CreateSubscriptionInEnrollmentAccountPreparer prepares the CreateSubscriptionInEnrollmentAccount request.
func (client Client) CreateSubscriptionInEnrollmentAccountPreparer(ctx context.Context, enrollmentAccountName string, body CreationParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"enrollmentAccountName": autorest.Encode("path", enrollmentAccountName),
	}

	const APIVersion = "2019-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/enrollmentAccounts/{enrollmentAccountName}/providers/Microsoft.Subscription/createSubscription", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSubscriptionInEnrollmentAccountSender sends the CreateSubscriptionInEnrollmentAccount request. The method will close the
// http.Response Body if it receives an error.
func (client Client) CreateSubscriptionInEnrollmentAccountSender(req *http.Request) (future CreateSubscriptionInEnrollmentAccountFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client Client) (cr CreationResult, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "subscription.CreateSubscriptionInEnrollmentAccountFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("subscription.CreateSubscriptionInEnrollmentAccountFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		cr.Response.Response, err = future.GetResult(sender)
		if cr.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "subscription.CreateSubscriptionInEnrollmentAccountFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && cr.Response.Response.StatusCode != http.StatusNoContent {
			cr, err = client.CreateSubscriptionInEnrollmentAccountResponder(cr.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "subscription.CreateSubscriptionInEnrollmentAccountFuture", "Result", cr.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateSubscriptionInEnrollmentAccountResponder handles the response to the CreateSubscriptionInEnrollmentAccount request. The method always
// closes the http.Response Body.
func (client Client) CreateSubscriptionInEnrollmentAccountResponder(resp *http.Response) (result CreationResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteAlias delete Alias.
// Parameters:
// aliasName - alias Name
func (client Client) DeleteAlias(ctx context.Context, aliasName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.DeleteAlias")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteAliasPreparer(ctx, aliasName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "DeleteAlias", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteAliasSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "subscription.Client", "DeleteAlias", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteAliasResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "DeleteAlias", resp, "Failure responding to request")
		return
	}

	return
}

// DeleteAliasPreparer prepares the DeleteAlias request.
func (client Client) DeleteAliasPreparer(ctx context.Context, aliasName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"aliasName": autorest.Encode("path", aliasName),
	}

	const APIVersion = "2019-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Subscription/aliases/{aliasName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteAliasSender sends the DeleteAlias request. The method will close the
// http.Response Body if it receives an error.
func (client Client) DeleteAliasSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteAliasResponder handles the response to the DeleteAlias request. The method always
// closes the http.Response Body.
func (client Client) DeleteAliasResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Enable the operation to enable a subscription
// Parameters:
// subscriptionID - subscription Id.
func (client Client) Enable(ctx context.Context, subscriptionID string) (result EnabledSubscriptionID, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Enable")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.EnablePreparer(ctx, subscriptionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "Enable", nil, "Failure preparing request")
		return
	}

	resp, err := client.EnableSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "subscription.Client", "Enable", resp, "Failure sending request")
		return
	}

	result, err = client.EnableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "Enable", resp, "Failure responding to request")
		return
	}

	return
}

// EnablePreparer prepares the Enable request.
func (client Client) EnablePreparer(ctx context.Context, subscriptionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2019-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Subscription/enable", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// EnableSender sends the Enable request. The method will close the
// http.Response Body if it receives an error.
func (client Client) EnableSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// EnableResponder handles the response to the Enable request. The method always
// closes the http.Response Body.
func (client Client) EnableResponder(resp *http.Response) (result EnabledSubscriptionID, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAlias get Alias Subscription.
// Parameters:
// aliasName - alias Name
func (client Client) GetAlias(ctx context.Context, aliasName string) (result PutAliasResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.GetAlias")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetAliasPreparer(ctx, aliasName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "GetAlias", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAliasSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "subscription.Client", "GetAlias", resp, "Failure sending request")
		return
	}

	result, err = client.GetAliasResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "GetAlias", resp, "Failure responding to request")
		return
	}

	return
}

// GetAliasPreparer prepares the GetAlias request.
func (client Client) GetAliasPreparer(ctx context.Context, aliasName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"aliasName": autorest.Encode("path", aliasName),
	}

	const APIVersion = "2019-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Subscription/aliases/{aliasName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAliasSender sends the GetAlias request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetAliasSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetAliasResponder handles the response to the GetAlias request. The method always
// closes the http.Response Body.
func (client Client) GetAliasResponder(resp *http.Response) (result PutAliasResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAlias get Alias Subscription.
func (client Client) ListAlias(ctx context.Context) (result PutAliasListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.ListAlias")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListAliasPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "ListAlias", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAliasSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "subscription.Client", "ListAlias", resp, "Failure sending request")
		return
	}

	result, err = client.ListAliasResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "ListAlias", resp, "Failure responding to request")
		return
	}

	return
}

// ListAliasPreparer prepares the ListAlias request.
func (client Client) ListAliasPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2019-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Subscription/aliases"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAliasSender sends the ListAlias request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ListAliasSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListAliasResponder handles the response to the ListAlias request. The method always
// closes the http.Response Body.
func (client Client) ListAliasResponder(resp *http.Response) (result PutAliasListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Rename the operation to rename a subscription
// Parameters:
// subscriptionID - subscription Id.
// body - subscription Name
func (client Client) Rename(ctx context.Context, subscriptionID string, body Name) (result RenamedSubscriptionID, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Rename")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RenamePreparer(ctx, subscriptionID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "Rename", nil, "Failure preparing request")
		return
	}

	resp, err := client.RenameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "subscription.Client", "Rename", resp, "Failure sending request")
		return
	}

	result, err = client.RenameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "Rename", resp, "Failure responding to request")
		return
	}

	return
}

// RenamePreparer prepares the Rename request.
func (client Client) RenamePreparer(ctx context.Context, subscriptionID string, body Name) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2019-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Subscription/rename", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RenameSender sends the Rename request. The method will close the
// http.Response Body if it receives an error.
func (client Client) RenameSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RenameResponder handles the response to the Rename request. The method always
// closes the http.Response Body.
func (client Client) RenameResponder(resp *http.Response) (result RenamedSubscriptionID, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}