package blockchain

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

// MemberProvisioningState enumerates the values for member provisioning state.
type MemberProvisioningState string

const (
	// Deleting ...
	Deleting MemberProvisioningState = "Deleting"
	// Failed ...
	Failed MemberProvisioningState = "Failed"
	// NotSpecified ...
	NotSpecified MemberProvisioningState = "NotSpecified"
	// Stale ...
	Stale MemberProvisioningState = "Stale"
	// Succeeded ...
	Succeeded MemberProvisioningState = "Succeeded"
	// Updating ...
	Updating MemberProvisioningState = "Updating"
)

// PossibleMemberProvisioningStateValues returns an array of possible values for the MemberProvisioningState const type.
func PossibleMemberProvisioningStateValues() []MemberProvisioningState {
	return []MemberProvisioningState{Deleting, Failed, NotSpecified, Stale, Succeeded, Updating}
}

// NameAvailabilityReason enumerates the values for name availability reason.
type NameAvailabilityReason string

const (
	// NameAvailabilityReasonAlreadyExists ...
	NameAvailabilityReasonAlreadyExists NameAvailabilityReason = "AlreadyExists"
	// NameAvailabilityReasonInvalid ...
	NameAvailabilityReasonInvalid NameAvailabilityReason = "Invalid"
	// NameAvailabilityReasonNotSpecified ...
	NameAvailabilityReasonNotSpecified NameAvailabilityReason = "NotSpecified"
)

// PossibleNameAvailabilityReasonValues returns an array of possible values for the NameAvailabilityReason const type.
func PossibleNameAvailabilityReasonValues() []NameAvailabilityReason {
	return []NameAvailabilityReason{NameAvailabilityReasonAlreadyExists, NameAvailabilityReasonInvalid, NameAvailabilityReasonNotSpecified}
}

// NodeProvisioningState enumerates the values for node provisioning state.
type NodeProvisioningState string

const (
	// NodeProvisioningStateDeleting ...
	NodeProvisioningStateDeleting NodeProvisioningState = "Deleting"
	// NodeProvisioningStateFailed ...
	NodeProvisioningStateFailed NodeProvisioningState = "Failed"
	// NodeProvisioningStateNotSpecified ...
	NodeProvisioningStateNotSpecified NodeProvisioningState = "NotSpecified"
	// NodeProvisioningStateSucceeded ...
	NodeProvisioningStateSucceeded NodeProvisioningState = "Succeeded"
	// NodeProvisioningStateUpdating ...
	NodeProvisioningStateUpdating NodeProvisioningState = "Updating"
)

// PossibleNodeProvisioningStateValues returns an array of possible values for the NodeProvisioningState const type.
func PossibleNodeProvisioningStateValues() []NodeProvisioningState {
	return []NodeProvisioningState{NodeProvisioningStateDeleting, NodeProvisioningStateFailed, NodeProvisioningStateNotSpecified, NodeProvisioningStateSucceeded, NodeProvisioningStateUpdating}
}

// Protocol enumerates the values for protocol.
type Protocol string

const (
	// ProtocolCorda ...
	ProtocolCorda Protocol = "Corda"
	// ProtocolNotSpecified ...
	ProtocolNotSpecified Protocol = "NotSpecified"
	// ProtocolParity ...
	ProtocolParity Protocol = "Parity"
	// ProtocolQuorum ...
	ProtocolQuorum Protocol = "Quorum"
)

// PossibleProtocolValues returns an array of possible values for the Protocol const type.
func PossibleProtocolValues() []Protocol {
	return []Protocol{ProtocolCorda, ProtocolNotSpecified, ProtocolParity, ProtocolQuorum}
}
