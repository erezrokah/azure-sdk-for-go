package aad

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/aad/mgmt/2020-07-01-preview/aad"

// ARMProxyResource the resource model definition for a ARM proxy resource. It will have everything other
// than required location and tags
type ARMProxyResource struct {
	// ID - READ-ONLY; The unique resource identifier of the Azure AD PrivateLink Policy.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the Azure AD PrivateLink Policy.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of Azure resource.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for ARMProxyResource.
func (apr ARMProxyResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// AzureADMetricsConfig azureADMetrics resource.
type AzureADMetricsConfig struct {
	autorest.Response `json:"-"`
	Properties        *AzureADMetricsPropertiesFormat `json:"properties,omitempty"`
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

// MarshalJSON is the custom marshaler for AzureADMetricsConfig.
func (amc AzureADMetricsConfig) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if amc.Properties != nil {
		objectMap["properties"] = amc.Properties
	}
	if amc.Tags != nil {
		objectMap["tags"] = amc.Tags
	}
	if amc.Location != nil {
		objectMap["location"] = amc.Location
	}
	return json.Marshal(objectMap)
}

// AzureADMetricsCreateOrUpdateFuture an abstraction for monitoring and retrieving the results of a
// long-running operation.
type AzureADMetricsCreateOrUpdateFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(AzureADMetricsClient) (AzureADMetricsConfig, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *AzureADMetricsCreateOrUpdateFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for AzureADMetricsCreateOrUpdateFuture.Result.
func (future *AzureADMetricsCreateOrUpdateFuture) result(client AzureADMetricsClient) (aamc AzureADMetricsConfig, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.AzureADMetricsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		aamc.Response.Response = future.Response()
		err = azure.NewAsyncOpIncompleteError("aad.AzureADMetricsCreateOrUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if aamc.Response.Response, err = future.GetResult(sender); err == nil && aamc.Response.Response.StatusCode != http.StatusNoContent {
		aamc, err = client.CreateOrUpdateResponder(aamc.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "aad.AzureADMetricsCreateOrUpdateFuture", "Result", aamc.Response.Response, "Failure responding to request")
		}
	}
	return
}

// AzureADMetricsListResult a list of AzureADMetrics resources
type AzureADMetricsListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; Array of AzureADMetrics resources.
	Value *[]AzureADMetricsConfig `json:"value,omitempty"`
	// NextLink - READ-ONLY; The link used to get the next page of operations.
	NextLink *string `json:"nextLink,omitempty"`
}

// MarshalJSON is the custom marshaler for AzureADMetricsListResult.
func (amlr AzureADMetricsListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// AzureADMetricsListResultIterator provides access to a complete listing of AzureADMetricsConfig values.
type AzureADMetricsListResultIterator struct {
	i    int
	page AzureADMetricsListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *AzureADMetricsListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AzureADMetricsListResultIterator.NextWithContext")
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
func (iter *AzureADMetricsListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter AzureADMetricsListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter AzureADMetricsListResultIterator) Response() AzureADMetricsListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter AzureADMetricsListResultIterator) Value() AzureADMetricsConfig {
	if !iter.page.NotDone() {
		return AzureADMetricsConfig{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the AzureADMetricsListResultIterator type.
func NewAzureADMetricsListResultIterator(page AzureADMetricsListResultPage) AzureADMetricsListResultIterator {
	return AzureADMetricsListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (amlr AzureADMetricsListResult) IsEmpty() bool {
	return amlr.Value == nil || len(*amlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (amlr AzureADMetricsListResult) hasNextLink() bool {
	return amlr.NextLink != nil && len(*amlr.NextLink) != 0
}

// azureADMetricsListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (amlr AzureADMetricsListResult) azureADMetricsListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !amlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(amlr.NextLink)))
}

// AzureADMetricsListResultPage contains a page of AzureADMetricsConfig values.
type AzureADMetricsListResultPage struct {
	fn    func(context.Context, AzureADMetricsListResult) (AzureADMetricsListResult, error)
	aamlr AzureADMetricsListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *AzureADMetricsListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AzureADMetricsListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.aamlr)
		if err != nil {
			return err
		}
		page.aamlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *AzureADMetricsListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page AzureADMetricsListResultPage) NotDone() bool {
	return !page.aamlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page AzureADMetricsListResultPage) Response() AzureADMetricsListResult {
	return page.aamlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page AzureADMetricsListResultPage) Values() []AzureADMetricsConfig {
	if page.aamlr.IsEmpty() {
		return nil
	}
	return *page.aamlr.Value
}

// Creates a new instance of the AzureADMetricsListResultPage type.
func NewAzureADMetricsListResultPage(cur AzureADMetricsListResult, getNextPage func(context.Context, AzureADMetricsListResult) (AzureADMetricsListResult, error)) AzureADMetricsListResultPage {
	return AzureADMetricsListResultPage{
		fn:    getNextPage,
		aamlr: cur,
	}
}

// AzureADMetricsPropertiesFormat ...
type AzureADMetricsPropertiesFormat struct {
	// ProvisioningState - READ-ONLY; The provisioning state of the resource. Possible values include: 'ProvisioningStateSucceeded', 'ProvisioningStateCreated', 'ProvisioningStateFailed'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
}

// MarshalJSON is the custom marshaler for AzureADMetricsPropertiesFormat.
func (aampf AzureADMetricsPropertiesFormat) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// AzureADMetricsUpdateParameter azureADMetrics parameters to be updated.
type AzureADMetricsUpdateParameter struct {
	// Tags - Resource tags to be updated.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for AzureADMetricsUpdateParameter.
func (amup AzureADMetricsUpdateParameter) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if amup.Tags != nil {
		objectMap["tags"] = amup.Tags
	}
	return json.Marshal(objectMap)
}

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

// MarshalJSON is the custom marshaler for AzureEntityResource.
func (aer AzureEntityResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// AzureResourceBase common properties for all Azure resources.
type AzureResourceBase struct {
	// ID - READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty"`
	// Name - Name of this resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for AzureResourceBase.
func (arb AzureResourceBase) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if arb.Name != nil {
		objectMap["name"] = arb.Name
	}
	return json.Marshal(objectMap)
}

// ErrorDefinition error definition.
type ErrorDefinition struct {
	// Code - READ-ONLY; Service specific error code which serves as the substatus for the HTTP error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; Description of the error.
	Message *string `json:"message,omitempty"`
	// Details - READ-ONLY; Internal error details.
	Details *[]ErrorDefinition `json:"details,omitempty"`
}

// MarshalJSON is the custom marshaler for ErrorDefinition.
func (ed ErrorDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ErrorResponse error response.
type ErrorResponse struct {
	// Error - The error details.
	Error *ErrorDefinition `json:"error,omitempty"`
}

// PrivateLinkForAzureAdCreateFuture an abstraction for monitoring and retrieving the results of a
// long-running operation.
type PrivateLinkForAzureAdCreateFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(PrivateLinkForAzureAdClient) (PrivateLinkPolicy, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *PrivateLinkForAzureAdCreateFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for PrivateLinkForAzureAdCreateFuture.Result.
func (future *PrivateLinkForAzureAdCreateFuture) result(client PrivateLinkForAzureAdClient) (plp PrivateLinkPolicy, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.PrivateLinkForAzureAdCreateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		plp.Response.Response = future.Response()
		err = azure.NewAsyncOpIncompleteError("aad.PrivateLinkForAzureAdCreateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if plp.Response.Response, err = future.GetResult(sender); err == nil && plp.Response.Response.StatusCode != http.StatusNoContent {
		plp, err = client.CreateResponder(plp.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "aad.PrivateLinkForAzureAdCreateFuture", "Result", plp.Response.Response, "Failure responding to request")
		}
	}
	return
}

// PrivateLinkPolicy privateLink Policy configuration object.
type PrivateLinkPolicy struct {
	autorest.Response `json:"-"`
	// OwnerTenantID - Guid of the owner tenant
	OwnerTenantID *string `json:"ownerTenantId,omitempty"`
	// AllTenants - Flag indicating whether all tenants are allowed
	AllTenants *bool `json:"allTenants,omitempty"`
	// Tenants - The list of tenantIds.
	Tenants *[]string `json:"tenants,omitempty"`
	// ResourceName - Name of the private link policy resource
	ResourceName *string `json:"resourceName,omitempty"`
	// SubscriptionID - Subscription Identifier
	SubscriptionID *string `json:"subscriptionId,omitempty"`
	// ResourceGroup - Name of the resource group
	ResourceGroup *string `json:"resourceGroup,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// ID - READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty"`
	// Name - Name of this resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for PrivateLinkPolicy.
func (plp PrivateLinkPolicy) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if plp.OwnerTenantID != nil {
		objectMap["ownerTenantId"] = plp.OwnerTenantID
	}
	if plp.AllTenants != nil {
		objectMap["allTenants"] = plp.AllTenants
	}
	if plp.Tenants != nil {
		objectMap["tenants"] = plp.Tenants
	}
	if plp.ResourceName != nil {
		objectMap["resourceName"] = plp.ResourceName
	}
	if plp.SubscriptionID != nil {
		objectMap["subscriptionId"] = plp.SubscriptionID
	}
	if plp.ResourceGroup != nil {
		objectMap["resourceGroup"] = plp.ResourceGroup
	}
	if plp.Tags != nil {
		objectMap["tags"] = plp.Tags
	}
	if plp.Name != nil {
		objectMap["name"] = plp.Name
	}
	return json.Marshal(objectMap)
}

// PrivateLinkPolicyListResult a list of private link policies
type PrivateLinkPolicyListResult struct {
	autorest.Response `json:"-"`
	// Value - Array of private link policies
	Value *[]PrivateLinkPolicy `json:"value,omitempty"`
	// NextLink - The link used to get the next page of operations.
	NextLink *string `json:"nextLink,omitempty"`
}

// PrivateLinkPolicyListResultIterator provides access to a complete listing of PrivateLinkPolicy values.
type PrivateLinkPolicyListResultIterator struct {
	i    int
	page PrivateLinkPolicyListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *PrivateLinkPolicyListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkPolicyListResultIterator.NextWithContext")
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
func (iter *PrivateLinkPolicyListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter PrivateLinkPolicyListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter PrivateLinkPolicyListResultIterator) Response() PrivateLinkPolicyListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter PrivateLinkPolicyListResultIterator) Value() PrivateLinkPolicy {
	if !iter.page.NotDone() {
		return PrivateLinkPolicy{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the PrivateLinkPolicyListResultIterator type.
func NewPrivateLinkPolicyListResultIterator(page PrivateLinkPolicyListResultPage) PrivateLinkPolicyListResultIterator {
	return PrivateLinkPolicyListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (plplr PrivateLinkPolicyListResult) IsEmpty() bool {
	return plplr.Value == nil || len(*plplr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (plplr PrivateLinkPolicyListResult) hasNextLink() bool {
	return plplr.NextLink != nil && len(*plplr.NextLink) != 0
}

// privateLinkPolicyListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (plplr PrivateLinkPolicyListResult) privateLinkPolicyListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !plplr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(plplr.NextLink)))
}

// PrivateLinkPolicyListResultPage contains a page of PrivateLinkPolicy values.
type PrivateLinkPolicyListResultPage struct {
	fn    func(context.Context, PrivateLinkPolicyListResult) (PrivateLinkPolicyListResult, error)
	plplr PrivateLinkPolicyListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *PrivateLinkPolicyListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkPolicyListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.plplr)
		if err != nil {
			return err
		}
		page.plplr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *PrivateLinkPolicyListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page PrivateLinkPolicyListResultPage) NotDone() bool {
	return !page.plplr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page PrivateLinkPolicyListResultPage) Response() PrivateLinkPolicyListResult {
	return page.plplr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page PrivateLinkPolicyListResultPage) Values() []PrivateLinkPolicy {
	if page.plplr.IsEmpty() {
		return nil
	}
	return *page.plplr.Value
}

// Creates a new instance of the PrivateLinkPolicyListResultPage type.
func NewPrivateLinkPolicyListResultPage(cur PrivateLinkPolicyListResult, getNextPage func(context.Context, PrivateLinkPolicyListResult) (PrivateLinkPolicyListResult, error)) PrivateLinkPolicyListResultPage {
	return PrivateLinkPolicyListResultPage{
		fn:    getNextPage,
		plplr: cur,
	}
}

// PrivateLinkPolicyUpdateParameter private Link policy parameters to be updated.
type PrivateLinkPolicyUpdateParameter struct {
	// Tags - Resource tags to be updated.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for PrivateLinkPolicyUpdateParameter.
func (plpup PrivateLinkPolicyUpdateParameter) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if plpup.Tags != nil {
		objectMap["tags"] = plpup.Tags
	}
	return json.Marshal(objectMap)
}

// PrivateLinkResource a private link resource
type PrivateLinkResource struct {
	autorest.Response `json:"-"`
	// PrivateLinkResourceProperties - Resource properties.
	*PrivateLinkResourceProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; The unique resource identifier of the Azure AD PrivateLink Policy.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the Azure AD PrivateLink Policy.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of Azure resource.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for PrivateLinkResource.
func (plr PrivateLinkResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if plr.PrivateLinkResourceProperties != nil {
		objectMap["properties"] = plr.PrivateLinkResourceProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for PrivateLinkResource struct.
func (plr *PrivateLinkResource) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var privateLinkResourceProperties PrivateLinkResourceProperties
				err = json.Unmarshal(*v, &privateLinkResourceProperties)
				if err != nil {
					return err
				}
				plr.PrivateLinkResourceProperties = &privateLinkResourceProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				plr.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				plr.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				plr.Type = &typeVar
			}
		}
	}

	return nil
}

// PrivateLinkResourceListResult a list of private link resources
type PrivateLinkResourceListResult struct {
	autorest.Response `json:"-"`
	// Value - Array of private link resources
	Value *[]PrivateLinkResource `json:"value,omitempty"`
	// NextLink - The link used to get the next page of operations.
	NextLink *string `json:"nextLink,omitempty"`
}

// PrivateLinkResourceListResultIterator provides access to a complete listing of PrivateLinkResource
// values.
type PrivateLinkResourceListResultIterator struct {
	i    int
	page PrivateLinkResourceListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *PrivateLinkResourceListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkResourceListResultIterator.NextWithContext")
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
func (iter *PrivateLinkResourceListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter PrivateLinkResourceListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter PrivateLinkResourceListResultIterator) Response() PrivateLinkResourceListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter PrivateLinkResourceListResultIterator) Value() PrivateLinkResource {
	if !iter.page.NotDone() {
		return PrivateLinkResource{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the PrivateLinkResourceListResultIterator type.
func NewPrivateLinkResourceListResultIterator(page PrivateLinkResourceListResultPage) PrivateLinkResourceListResultIterator {
	return PrivateLinkResourceListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (plrlr PrivateLinkResourceListResult) IsEmpty() bool {
	return plrlr.Value == nil || len(*plrlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (plrlr PrivateLinkResourceListResult) hasNextLink() bool {
	return plrlr.NextLink != nil && len(*plrlr.NextLink) != 0
}

// privateLinkResourceListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (plrlr PrivateLinkResourceListResult) privateLinkResourceListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !plrlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(plrlr.NextLink)))
}

// PrivateLinkResourceListResultPage contains a page of PrivateLinkResource values.
type PrivateLinkResourceListResultPage struct {
	fn    func(context.Context, PrivateLinkResourceListResult) (PrivateLinkResourceListResult, error)
	plrlr PrivateLinkResourceListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *PrivateLinkResourceListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkResourceListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.plrlr)
		if err != nil {
			return err
		}
		page.plrlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *PrivateLinkResourceListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page PrivateLinkResourceListResultPage) NotDone() bool {
	return !page.plrlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page PrivateLinkResourceListResultPage) Response() PrivateLinkResourceListResult {
	return page.plrlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page PrivateLinkResourceListResultPage) Values() []PrivateLinkResource {
	if page.plrlr.IsEmpty() {
		return nil
	}
	return *page.plrlr.Value
}

// Creates a new instance of the PrivateLinkResourceListResultPage type.
func NewPrivateLinkResourceListResultPage(cur PrivateLinkResourceListResult, getNextPage func(context.Context, PrivateLinkResourceListResult) (PrivateLinkResourceListResult, error)) PrivateLinkResourceListResultPage {
	return PrivateLinkResourceListResultPage{
		fn:    getNextPage,
		plrlr: cur,
	}
}

// PrivateLinkResourceProperties properties of a private link resource.
type PrivateLinkResourceProperties struct {
	// GroupID - READ-ONLY; The private link resource group id.
	GroupID *string `json:"groupId,omitempty"`
	// RequiredMembers - READ-ONLY; The private link resource required member names.
	RequiredMembers *[]string `json:"requiredMembers,omitempty"`
}

// MarshalJSON is the custom marshaler for PrivateLinkResourceProperties.
func (plrp PrivateLinkResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
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

// MarshalJSON is the custom marshaler for ProxyResource.
func (pr ProxyResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
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

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
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
