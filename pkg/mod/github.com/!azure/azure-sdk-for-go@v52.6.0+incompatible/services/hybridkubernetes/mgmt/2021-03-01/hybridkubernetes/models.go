package hybridkubernetes

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
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/hybridkubernetes/mgmt/2021-03-01/hybridkubernetes"

// AzureEntityResource the resource model definition for an Azure Resource Manager resource with an etag.
type AzureEntityResource struct {
	// Etag - READ-ONLY; Resource Etag.
	Etag *string `json:"etag,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// ConnectedCluster represents a connected cluster.
type ConnectedCluster struct {
	autorest.Response `json:"-"`
	// Identity - The identity of the connected cluster.
	Identity *ConnectedClusterIdentity `json:"identity,omitempty"`
	// ConnectedClusterProperties - Describes the connected cluster resource properties.
	*ConnectedClusterProperties `json:"properties,omitempty"`
	// SystemData - READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData `json:"systemData,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for ConnectedCluster.
func (cc ConnectedCluster) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cc.Identity != nil {
		objectMap["identity"] = cc.Identity
	}
	if cc.ConnectedClusterProperties != nil {
		objectMap["properties"] = cc.ConnectedClusterProperties
	}
	if cc.Tags != nil {
		objectMap["tags"] = cc.Tags
	}
	if cc.Location != nil {
		objectMap["location"] = cc.Location
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ConnectedCluster struct.
func (cc *ConnectedCluster) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "identity":
			if v != nil {
				var identity ConnectedClusterIdentity
				err = json.Unmarshal(*v, &identity)
				if err != nil {
					return err
				}
				cc.Identity = &identity
			}
		case "properties":
			if v != nil {
				var connectedClusterProperties ConnectedClusterProperties
				err = json.Unmarshal(*v, &connectedClusterProperties)
				if err != nil {
					return err
				}
				cc.ConnectedClusterProperties = &connectedClusterProperties
			}
		case "systemData":
			if v != nil {
				var systemData SystemData
				err = json.Unmarshal(*v, &systemData)
				if err != nil {
					return err
				}
				cc.SystemData = &systemData
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				cc.Tags = tags
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				cc.Location = &location
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				cc.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				cc.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				cc.Type = &typeVar
			}
		}
	}

	return nil
}

// ConnectedClusterCreateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type ConnectedClusterCreateFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(ConnectedClusterClient) (ConnectedCluster, error)
}

// ConnectedClusterDeleteFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type ConnectedClusterDeleteFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(ConnectedClusterClient) (autorest.Response, error)
}

// ConnectedClusterIdentity identity for the connected cluster.
type ConnectedClusterIdentity struct {
	// PrincipalID - READ-ONLY; The principal id of connected cluster identity. This property will only be provided for a system assigned identity.
	PrincipalID *string `json:"principalId,omitempty"`
	// TenantID - READ-ONLY; The tenant id associated with the connected cluster. This property will only be provided for a system assigned identity.
	TenantID *string `json:"tenantId,omitempty"`
	// Type - The type of identity used for the connected cluster. The type 'SystemAssigned, includes a system created identity. The type 'None' means no identity is assigned to the connected cluster. Possible values include: 'None', 'SystemAssigned'
	Type ResourceIdentityType `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for ConnectedClusterIdentity.
func (cci ConnectedClusterIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cci.Type != "" {
		objectMap["type"] = cci.Type
	}
	return json.Marshal(objectMap)
}

// ConnectedClusterList the paginated list of connected Clusters
type ConnectedClusterList struct {
	autorest.Response `json:"-"`
	// Value - The list of connected clusters
	Value *[]ConnectedCluster `json:"value,omitempty"`
	// NextLink - The link to fetch the next page of connected cluster
	NextLink *string `json:"nextLink,omitempty"`
}

// ConnectedClusterListIterator provides access to a complete listing of ConnectedCluster values.
type ConnectedClusterListIterator struct {
	i    int
	page ConnectedClusterListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ConnectedClusterListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConnectedClusterListIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *ConnectedClusterListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ConnectedClusterListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ConnectedClusterListIterator) Response() ConnectedClusterList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ConnectedClusterListIterator) Value() ConnectedCluster {
	if !iter.page.NotDone() {
		return ConnectedCluster{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ConnectedClusterListIterator type.
func NewConnectedClusterListIterator(page ConnectedClusterListPage) ConnectedClusterListIterator {
	return ConnectedClusterListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (ccl ConnectedClusterList) IsEmpty() bool {
	return ccl.Value == nil || len(*ccl.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (ccl ConnectedClusterList) hasNextLink() bool {
	return ccl.NextLink != nil && len(*ccl.NextLink) != 0
}

// connectedClusterListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ccl ConnectedClusterList) connectedClusterListPreparer(ctx context.Context) (*http.Request, error) {
	if !ccl.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ccl.NextLink)))
}

// ConnectedClusterListPage contains a page of ConnectedCluster values.
type ConnectedClusterListPage struct {
	fn  func(context.Context, ConnectedClusterList) (ConnectedClusterList, error)
	ccl ConnectedClusterList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ConnectedClusterListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConnectedClusterListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.ccl)
		if err != nil {
			return err
		}
		page.ccl = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ConnectedClusterListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ConnectedClusterListPage) NotDone() bool {
	return !page.ccl.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ConnectedClusterListPage) Response() ConnectedClusterList {
	return page.ccl
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ConnectedClusterListPage) Values() []ConnectedCluster {
	if page.ccl.IsEmpty() {
		return nil
	}
	return *page.ccl.Value
}

// Creates a new instance of the ConnectedClusterListPage type.
func NewConnectedClusterListPage(cur ConnectedClusterList, getNextPage func(context.Context, ConnectedClusterList) (ConnectedClusterList, error)) ConnectedClusterListPage {
	return ConnectedClusterListPage{
		fn:  getNextPage,
		ccl: cur,
	}
}

// ConnectedClusterPatch object containing updates for patch operations.
type ConnectedClusterPatch struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Properties - Describes the connected cluster resource properties that can be updated during PATCH operation.
	Properties interface{} `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for ConnectedClusterPatch.
func (ccp ConnectedClusterPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ccp.Tags != nil {
		objectMap["tags"] = ccp.Tags
	}
	if ccp.Properties != nil {
		objectMap["properties"] = ccp.Properties
	}
	return json.Marshal(objectMap)
}

// ConnectedClusterProperties properties of the connected cluster.
type ConnectedClusterProperties struct {
	// AgentPublicKeyCertificate - Base64 encoded public certificate used by the agent to do the initial handshake to the backend services in Azure.
	AgentPublicKeyCertificate *string `json:"agentPublicKeyCertificate,omitempty"`
	// KubernetesVersion - READ-ONLY; The Kubernetes version of the connected cluster resource
	KubernetesVersion *string `json:"kubernetesVersion,omitempty"`
	// TotalNodeCount - READ-ONLY; Number of nodes present in the connected cluster resource
	TotalNodeCount *int32 `json:"totalNodeCount,omitempty"`
	// TotalCoreCount - READ-ONLY; Number of CPU cores present in the connected cluster resource
	TotalCoreCount *int32 `json:"totalCoreCount,omitempty"`
	// AgentVersion - READ-ONLY; Version of the agent running on the connected cluster resource
	AgentVersion *string `json:"agentVersion,omitempty"`
	// ProvisioningState - Provisioning state of the connected cluster resource. Possible values include: 'Succeeded', 'Failed', 'Canceled', 'Provisioning', 'Updating', 'Deleting', 'Accepted'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// Distribution - The Kubernetes distribution running on this connected cluster.
	Distribution *string `json:"distribution,omitempty"`
	// Infrastructure - The infrastructure on which the Kubernetes cluster represented by this connected cluster is running on.
	Infrastructure *string `json:"infrastructure,omitempty"`
	// Offering - READ-ONLY; Connected cluster offering
	Offering *string `json:"offering,omitempty"`
	// ManagedIdentityCertificateExpirationTime - READ-ONLY; Expiration time of the managed identity certificate
	ManagedIdentityCertificateExpirationTime *date.Time `json:"managedIdentityCertificateExpirationTime,omitempty"`
	// LastConnectivityTime - READ-ONLY; Time representing the last instance when heart beat was received from the cluster
	LastConnectivityTime *date.Time `json:"lastConnectivityTime,omitempty"`
	// ConnectivityStatus - READ-ONLY; Represents the connectivity status of the connected cluster. Possible values include: 'Connecting', 'Connected', 'Offline', 'Expired'
	ConnectivityStatus ConnectivityStatus `json:"connectivityStatus,omitempty"`
}

// MarshalJSON is the custom marshaler for ConnectedClusterProperties.
func (ccp ConnectedClusterProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ccp.AgentPublicKeyCertificate != nil {
		objectMap["agentPublicKeyCertificate"] = ccp.AgentPublicKeyCertificate
	}
	if ccp.ProvisioningState != "" {
		objectMap["provisioningState"] = ccp.ProvisioningState
	}
	if ccp.Distribution != nil {
		objectMap["distribution"] = ccp.Distribution
	}
	if ccp.Infrastructure != nil {
		objectMap["infrastructure"] = ccp.Infrastructure
	}
	return json.Marshal(objectMap)
}

// ErrorAdditionalInfo the resource management error additional info.
type ErrorAdditionalInfo struct {
	// Type - READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty"`
	// Info - READ-ONLY; The additional info.
	Info interface{} `json:"info,omitempty"`
}

// ErrorDetail the error detail.
type ErrorDetail struct {
	// Code - READ-ONLY; The error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; The error message.
	Message *string `json:"message,omitempty"`
	// Target - READ-ONLY; The error target.
	Target *string `json:"target,omitempty"`
	// Details - READ-ONLY; The error details.
	Details *[]ErrorDetail `json:"details,omitempty"`
	// AdditionalInfo - READ-ONLY; The error additional info.
	AdditionalInfo *[]ErrorAdditionalInfo `json:"additionalInfo,omitempty"`
}

// ErrorResponse common error response for all Azure Resource Manager APIs to return error details for
// failed operations. (This also follows the OData error response format.).
type ErrorResponse struct {
	// Error - The error object.
	Error *ErrorDetail `json:"error,omitempty"`
}

// Operation the Connected cluster API operation
type Operation struct {
	// Name - READ-ONLY; Operation name: {Microsoft.Kubernetes}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - READ-ONLY; The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.connectedClusters
	Provider *string `json:"provider,omitempty"`
	// Resource - Connected Cluster Resource on which the operation is performed
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
	// Description - Description of the operation.
	Description *string `json:"description,omitempty"`
}

// OperationList the paginated list of connected cluster API operations.
type OperationList struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; The list of connected cluster API operations.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - The link to fetch the next page of connected cluster API operations.
	NextLink *string `json:"nextLink,omitempty"`
}

// MarshalJSON is the custom marshaler for OperationList.
func (ol OperationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ol.NextLink != nil {
		objectMap["nextLink"] = ol.NextLink
	}
	return json.Marshal(objectMap)
}

// OperationListIterator provides access to a complete listing of Operation values.
type OperationListIterator struct {
	i    int
	page OperationListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListIterator) Response() OperationList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationListIterator type.
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return OperationListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (ol OperationList) IsEmpty() bool {
	return ol.Value == nil || len(*ol.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (ol OperationList) hasNextLink() bool {
	return ol.NextLink != nil && len(*ol.NextLink) != 0
}

// operationListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ol OperationList) operationListPreparer(ctx context.Context) (*http.Request, error) {
	if !ol.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ol.NextLink)))
}

// OperationListPage contains a page of Operation values.
type OperationListPage struct {
	fn func(context.Context, OperationList) (OperationList, error)
	ol OperationList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.ol)
		if err != nil {
			return err
		}
		page.ol = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListPage) NotDone() bool {
	return !page.ol.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListPage) Response() OperationList {
	return page.ol
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListPage) Values() []Operation {
	if page.ol.IsEmpty() {
		return nil
	}
	return *page.ol.Value
}

// Creates a new instance of the OperationListPage type.
func NewOperationListPage(cur OperationList, getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return OperationListPage{
		fn: getNextPage,
		ol: cur,
	}
}

// ProxyResource the resource model definition for a Azure Resource Manager proxy resource. It will not
// have tags and a location
type ProxyResource struct {
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// Resource common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// SystemData metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// CreatedBy - The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`
	// CreatedByType - The type of identity that created the resource. Possible values include: 'User', 'Application', 'ManagedIdentity', 'Key'
	CreatedByType CreatedByType `json:"createdByType,omitempty"`
	// CreatedAt - The timestamp of resource creation (UTC).
	CreatedAt *date.Time `json:"createdAt,omitempty"`
	// LastModifiedBy - The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// LastModifiedByType - The type of identity that last modified the resource. Possible values include: 'LastModifiedByTypeUser', 'LastModifiedByTypeApplication', 'LastModifiedByTypeManagedIdentity', 'LastModifiedByTypeKey'
	LastModifiedByType LastModifiedByType `json:"lastModifiedByType,omitempty"`
	// LastModifiedAt - The timestamp of resource modification (UTC).
	LastModifiedAt *date.Time `json:"lastModifiedAt,omitempty"`
}

// TrackedResource the resource model definition for an Azure Resource Manager tracked top level resource
// which has 'tags' and a 'location'
type TrackedResource struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	if tr.Location != nil {
		objectMap["location"] = tr.Location
	}
	return json.Marshal(objectMap)
}
