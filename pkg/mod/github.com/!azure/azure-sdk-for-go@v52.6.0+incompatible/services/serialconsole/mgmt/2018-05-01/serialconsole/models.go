package serialconsole

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
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/serialconsole/mgmt/2018-05-01/serialconsole"

// DisableSerialConsoleResult returns whether or not Serial Console is disabled.
type DisableSerialConsoleResult struct {
	// Disabled - Whether or not Serial Console is disabled.
	Disabled *bool `json:"disabled,omitempty"`
}

// EnableSerialConsoleResult returns whether or not Serial Console is disabled (enabled).
type EnableSerialConsoleResult struct {
	// Disabled - Whether or not Serial Console is disabled (enabled).
	Disabled *bool `json:"disabled,omitempty"`
}

// GetSerialConsoleSubscriptionNotFound error saying that the provided subscription could not be found
type GetSerialConsoleSubscriptionNotFound struct {
	// Code - Error code
	Code *string `json:"code,omitempty"`
	// Message - Subscription not found message
	Message *string `json:"message,omitempty"`
}

// Operations serial Console operations
type Operations struct {
	autorest.Response `json:"-"`
	// Value - A list of Serial Console operations
	Value *[]OperationsValueItem `json:"value,omitempty"`
}

// OperationsValueItem ...
type OperationsValueItem struct {
	Name         *string                     `json:"name,omitempty"`
	IsDataAction *string                     `json:"isDataAction,omitempty"`
	Display      *OperationsValueItemDisplay `json:"display,omitempty"`
}

// OperationsValueItemDisplay ...
type OperationsValueItemDisplay struct {
	Provider    *string `json:"provider,omitempty"`
	Resource    *string `json:"resource,omitempty"`
	Operation   *string `json:"operation,omitempty"`
	Description *string `json:"description,omitempty"`
}

// SetObject ...
type SetObject struct {
	autorest.Response `json:"-"`
	Value             interface{} `json:"value,omitempty"`
}

// Status returns whether or not Serial Console is disabled.
type Status struct {
	// Disabled - Whether or not Serial Console is disabled.
	Disabled *bool `json:"disabled,omitempty"`
}