//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armproviderhub

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DefaultRolloutsClient contains the methods for the DefaultRollouts group.
// Don't use this type directly, use NewDefaultRolloutsClient() instead.
type DefaultRolloutsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDefaultRolloutsClient creates a new instance of DefaultRolloutsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDefaultRolloutsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DefaultRolloutsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &DefaultRolloutsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates the rollout details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-20
// providerNamespace - The name of the resource provider hosted within ProviderHub.
// rolloutName - The rollout name.
// properties - The Default rollout properties supplied to the CreateOrUpdate operation.
// options - DefaultRolloutsClientBeginCreateOrUpdateOptions contains the optional parameters for the DefaultRolloutsClient.BeginCreateOrUpdate
// method.
func (client *DefaultRolloutsClient) BeginCreateOrUpdate(ctx context.Context, providerNamespace string, rolloutName string, properties DefaultRollout, options *DefaultRolloutsClientBeginCreateOrUpdateOptions) (*runtime.Poller[DefaultRolloutsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, providerNamespace, rolloutName, properties, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[DefaultRolloutsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DefaultRolloutsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates the rollout details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-20
func (client *DefaultRolloutsClient) createOrUpdate(ctx context.Context, providerNamespace string, rolloutName string, properties DefaultRollout, options *DefaultRolloutsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, providerNamespace, rolloutName, properties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DefaultRolloutsClient) createOrUpdateCreateRequest(ctx context.Context, providerNamespace string, rolloutName string, properties DefaultRollout, options *DefaultRolloutsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/defaultRollouts/{rolloutName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	if rolloutName == "" {
		return nil, errors.New("parameter rolloutName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rolloutName}", url.PathEscape(rolloutName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, properties)
}

// Delete - Deletes the rollout resource. Rollout must be in terminal state.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-20
// providerNamespace - The name of the resource provider hosted within ProviderHub.
// rolloutName - The rollout name.
// options - DefaultRolloutsClientDeleteOptions contains the optional parameters for the DefaultRolloutsClient.Delete method.
func (client *DefaultRolloutsClient) Delete(ctx context.Context, providerNamespace string, rolloutName string, options *DefaultRolloutsClientDeleteOptions) (DefaultRolloutsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, providerNamespace, rolloutName, options)
	if err != nil {
		return DefaultRolloutsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DefaultRolloutsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DefaultRolloutsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DefaultRolloutsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DefaultRolloutsClient) deleteCreateRequest(ctx context.Context, providerNamespace string, rolloutName string, options *DefaultRolloutsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/defaultRollouts/{rolloutName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	if rolloutName == "" {
		return nil, errors.New("parameter rolloutName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rolloutName}", url.PathEscape(rolloutName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the default rollout details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-20
// providerNamespace - The name of the resource provider hosted within ProviderHub.
// rolloutName - The rollout name.
// options - DefaultRolloutsClientGetOptions contains the optional parameters for the DefaultRolloutsClient.Get method.
func (client *DefaultRolloutsClient) Get(ctx context.Context, providerNamespace string, rolloutName string, options *DefaultRolloutsClientGetOptions) (DefaultRolloutsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, providerNamespace, rolloutName, options)
	if err != nil {
		return DefaultRolloutsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DefaultRolloutsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DefaultRolloutsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DefaultRolloutsClient) getCreateRequest(ctx context.Context, providerNamespace string, rolloutName string, options *DefaultRolloutsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/defaultRollouts/{rolloutName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	if rolloutName == "" {
		return nil, errors.New("parameter rolloutName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rolloutName}", url.PathEscape(rolloutName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DefaultRolloutsClient) getHandleResponse(resp *http.Response) (DefaultRolloutsClientGetResponse, error) {
	result := DefaultRolloutsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DefaultRollout); err != nil {
		return DefaultRolloutsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByProviderRegistrationPager - Gets the list of the rollouts for the given provider.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-20
// providerNamespace - The name of the resource provider hosted within ProviderHub.
// options - DefaultRolloutsClientListByProviderRegistrationOptions contains the optional parameters for the DefaultRolloutsClient.ListByProviderRegistration
// method.
func (client *DefaultRolloutsClient) NewListByProviderRegistrationPager(providerNamespace string, options *DefaultRolloutsClientListByProviderRegistrationOptions) *runtime.Pager[DefaultRolloutsClientListByProviderRegistrationResponse] {
	return runtime.NewPager(runtime.PagingHandler[DefaultRolloutsClientListByProviderRegistrationResponse]{
		More: func(page DefaultRolloutsClientListByProviderRegistrationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DefaultRolloutsClientListByProviderRegistrationResponse) (DefaultRolloutsClientListByProviderRegistrationResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByProviderRegistrationCreateRequest(ctx, providerNamespace, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DefaultRolloutsClientListByProviderRegistrationResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DefaultRolloutsClientListByProviderRegistrationResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DefaultRolloutsClientListByProviderRegistrationResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByProviderRegistrationHandleResponse(resp)
		},
	})
}

// listByProviderRegistrationCreateRequest creates the ListByProviderRegistration request.
func (client *DefaultRolloutsClient) listByProviderRegistrationCreateRequest(ctx context.Context, providerNamespace string, options *DefaultRolloutsClientListByProviderRegistrationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/defaultRollouts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByProviderRegistrationHandleResponse handles the ListByProviderRegistration response.
func (client *DefaultRolloutsClient) listByProviderRegistrationHandleResponse(resp *http.Response) (DefaultRolloutsClientListByProviderRegistrationResponse, error) {
	result := DefaultRolloutsClientListByProviderRegistrationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DefaultRolloutArrayResponseWithContinuation); err != nil {
		return DefaultRolloutsClientListByProviderRegistrationResponse{}, err
	}
	return result, nil
}

// Stop - Stops or cancels the rollout, if in progress.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-20
// providerNamespace - The name of the resource provider hosted within ProviderHub.
// rolloutName - The rollout name.
// options - DefaultRolloutsClientStopOptions contains the optional parameters for the DefaultRolloutsClient.Stop method.
func (client *DefaultRolloutsClient) Stop(ctx context.Context, providerNamespace string, rolloutName string, options *DefaultRolloutsClientStopOptions) (DefaultRolloutsClientStopResponse, error) {
	req, err := client.stopCreateRequest(ctx, providerNamespace, rolloutName, options)
	if err != nil {
		return DefaultRolloutsClientStopResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DefaultRolloutsClientStopResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DefaultRolloutsClientStopResponse{}, runtime.NewResponseError(resp)
	}
	return DefaultRolloutsClientStopResponse{}, nil
}

// stopCreateRequest creates the Stop request.
func (client *DefaultRolloutsClient) stopCreateRequest(ctx context.Context, providerNamespace string, rolloutName string, options *DefaultRolloutsClientStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/defaultRollouts/{rolloutName}/stop"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	if rolloutName == "" {
		return nil, errors.New("parameter rolloutName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rolloutName}", url.PathEscape(rolloutName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
