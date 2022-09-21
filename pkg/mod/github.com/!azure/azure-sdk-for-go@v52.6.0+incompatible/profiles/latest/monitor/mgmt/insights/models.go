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

package insights

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/monitor/mgmt/2020-10-01/insights"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ActionGroup = original.ActionGroup
type ActionList = original.ActionList
type ActivityLogAlertResource = original.ActivityLogAlertResource
type ActivityLogAlertsClient = original.ActivityLogAlertsClient
type AlertRuleAllOfCondition = original.AlertRuleAllOfCondition
type AlertRuleAnyOfOrLeafCondition = original.AlertRuleAnyOfOrLeafCondition
type AlertRuleLeafCondition = original.AlertRuleLeafCondition
type AlertRuleList = original.AlertRuleList
type AlertRuleListIterator = original.AlertRuleListIterator
type AlertRuleListPage = original.AlertRuleListPage
type AlertRulePatchObject = original.AlertRulePatchObject
type AlertRulePatchProperties = original.AlertRulePatchProperties
type AlertRuleProperties = original.AlertRuleProperties
type AzureResource = original.AzureResource
type BaseClient = original.BaseClient
type ErrorResponse = original.ErrorResponse

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewActivityLogAlertsClient(subscriptionID string) ActivityLogAlertsClient {
	return original.NewActivityLogAlertsClient(subscriptionID)
}
func NewActivityLogAlertsClientWithBaseURI(baseURI string, subscriptionID string) ActivityLogAlertsClient {
	return original.NewActivityLogAlertsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAlertRuleListIterator(page AlertRuleListPage) AlertRuleListIterator {
	return original.NewAlertRuleListIterator(page)
}
func NewAlertRuleListPage(cur AlertRuleList, getNextPage func(context.Context, AlertRuleList) (AlertRuleList, error)) AlertRuleListPage {
	return original.NewAlertRuleListPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
