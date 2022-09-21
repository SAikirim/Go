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

package aad

import original "github.com/Azure/azure-sdk-for-go/services/aad/mgmt/2017-04-01/aad"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type Category = original.Category

const (
	AuditLogs  Category = original.AuditLogs
	SignInLogs Category = original.SignInLogs
)

type CategoryType = original.CategoryType

const (
	Logs CategoryType = original.Logs
)

type BaseClient = original.BaseClient
type DiagnosticSettings = original.DiagnosticSettings
type DiagnosticSettingsCategory = original.DiagnosticSettingsCategory
type DiagnosticSettingsCategoryClient = original.DiagnosticSettingsCategoryClient
type DiagnosticSettingsCategoryResource = original.DiagnosticSettingsCategoryResource
type DiagnosticSettingsCategoryResourceCollection = original.DiagnosticSettingsCategoryResourceCollection
type DiagnosticSettingsClient = original.DiagnosticSettingsClient
type DiagnosticSettingsResource = original.DiagnosticSettingsResource
type DiagnosticSettingsResourceCollection = original.DiagnosticSettingsResourceCollection
type Display = original.Display
type ErrorDefinition = original.ErrorDefinition
type ErrorResponse = original.ErrorResponse
type LogSettings = original.LogSettings
type OperationsClient = original.OperationsClient
type OperationsDiscovery = original.OperationsDiscovery
type OperationsDiscoveryCollection = original.OperationsDiscoveryCollection
type ProxyOnlyResource = original.ProxyOnlyResource
type RetentionPolicy = original.RetentionPolicy

func New() BaseClient {
	return original.New()
}
func NewDiagnosticSettingsCategoryClient() DiagnosticSettingsCategoryClient {
	return original.NewDiagnosticSettingsCategoryClient()
}
func NewDiagnosticSettingsCategoryClientWithBaseURI(baseURI string) DiagnosticSettingsCategoryClient {
	return original.NewDiagnosticSettingsCategoryClientWithBaseURI(baseURI)
}
func NewDiagnosticSettingsClient() DiagnosticSettingsClient {
	return original.NewDiagnosticSettingsClient()
}
func NewDiagnosticSettingsClientWithBaseURI(baseURI string) DiagnosticSettingsClient {
	return original.NewDiagnosticSettingsClientWithBaseURI(baseURI)
}
func NewOperationsClient() OperationsClient {
	return original.NewOperationsClient()
}
func NewOperationsClientWithBaseURI(baseURI string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleCategoryTypeValues() []CategoryType {
	return original.PossibleCategoryTypeValues()
}
func PossibleCategoryValues() []Category {
	return original.PossibleCategoryValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
