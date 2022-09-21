package attestationapi

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
	"github.com/Azure/azure-sdk-for-go/services/attestation/2018-09-01/attestation"
)

// PolicyClientAPI contains the set of methods on the PolicyClient type.
type PolicyClientAPI interface {
	Get(ctx context.Context, tenantBaseURL string, tee attestation.TeeKind) (result attestation.SetObject, err error)
	PrepareToSet(ctx context.Context, tenantBaseURL string, tee attestation.TeeKind, policyJws string) (result attestation.SetObject, err error)
	Reset(ctx context.Context, tenantBaseURL string, tee attestation.TeeKind, policyJws string) (result attestation.SetObject, err error)
	Set(ctx context.Context, tenantBaseURL string, tee attestation.TeeKind, newAttestationPolicy string) (result attestation.SetObject, err error)
}

var _ PolicyClientAPI = (*attestation.PolicyClient)(nil)

// PolicyCertificatesClientAPI contains the set of methods on the PolicyCertificatesClient type.
type PolicyCertificatesClientAPI interface {
	Add(ctx context.Context, tenantBaseURL string, policyCertificateToAdd string) (result attestation.SetObject, err error)
	Get(ctx context.Context, tenantBaseURL string) (result attestation.SetObject, err error)
	Remove(ctx context.Context, tenantBaseURL string, policyCertificateToRemove string) (result attestation.SetObject, err error)
}

var _ PolicyCertificatesClientAPI = (*attestation.PolicyCertificatesClient)(nil)

// SigningCertificatesClientAPI contains the set of methods on the SigningCertificatesClient type.
type SigningCertificatesClientAPI interface {
	Get(ctx context.Context, tenantBaseURL string) (result attestation.SetObject, err error)
}

var _ SigningCertificatesClientAPI = (*attestation.SigningCertificatesClient)(nil)

// MetadataConfigurationClientAPI contains the set of methods on the MetadataConfigurationClient type.
type MetadataConfigurationClientAPI interface {
	Get(ctx context.Context, tenantBaseURL string) (result attestation.SetObject, err error)
}

var _ MetadataConfigurationClientAPI = (*attestation.MetadataConfigurationClient)(nil)
