//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

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

// IntegrationServiceEnvironmentManagedApisClient contains the methods for the IntegrationServiceEnvironmentManagedApis group.
// Don't use this type directly, use NewIntegrationServiceEnvironmentManagedApisClient() instead.
type IntegrationServiceEnvironmentManagedApisClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIntegrationServiceEnvironmentManagedApisClient creates a new instance of IntegrationServiceEnvironmentManagedApisClient with the specified values.
// subscriptionID - The subscription id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIntegrationServiceEnvironmentManagedApisClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IntegrationServiceEnvironmentManagedApisClient, error) {
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
	client := &IntegrationServiceEnvironmentManagedApisClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginDelete - Deletes the integration service environment managed Api.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroup - The resource group.
// integrationServiceEnvironmentName - The integration service environment name.
// apiName - The api name.
// options - IntegrationServiceEnvironmentManagedApisClientBeginDeleteOptions contains the optional parameters for the IntegrationServiceEnvironmentManagedApisClient.BeginDelete
// method.
func (client *IntegrationServiceEnvironmentManagedApisClient) BeginDelete(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, options *IntegrationServiceEnvironmentManagedApisClientBeginDeleteOptions) (*runtime.Poller[IntegrationServiceEnvironmentManagedApisClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroup, integrationServiceEnvironmentName, apiName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[IntegrationServiceEnvironmentManagedApisClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[IntegrationServiceEnvironmentManagedApisClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the integration service environment managed Api.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
func (client *IntegrationServiceEnvironmentManagedApisClient) deleteOperation(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, options *IntegrationServiceEnvironmentManagedApisClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, apiName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IntegrationServiceEnvironmentManagedApisClient) deleteCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, options *IntegrationServiceEnvironmentManagedApisClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis/{apiName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	if apiName == "" {
		return nil, errors.New("parameter apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(apiName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the integration service environment managed Api.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroup - The resource group name.
// integrationServiceEnvironmentName - The integration service environment name.
// apiName - The api name.
// options - IntegrationServiceEnvironmentManagedApisClientGetOptions contains the optional parameters for the IntegrationServiceEnvironmentManagedApisClient.Get
// method.
func (client *IntegrationServiceEnvironmentManagedApisClient) Get(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, options *IntegrationServiceEnvironmentManagedApisClientGetOptions) (IntegrationServiceEnvironmentManagedApisClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, apiName, options)
	if err != nil {
		return IntegrationServiceEnvironmentManagedApisClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationServiceEnvironmentManagedApisClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationServiceEnvironmentManagedApisClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IntegrationServiceEnvironmentManagedApisClient) getCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, options *IntegrationServiceEnvironmentManagedApisClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis/{apiName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	if apiName == "" {
		return nil, errors.New("parameter apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(apiName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntegrationServiceEnvironmentManagedApisClient) getHandleResponse(resp *http.Response) (IntegrationServiceEnvironmentManagedApisClientGetResponse, error) {
	result := IntegrationServiceEnvironmentManagedApisClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationServiceEnvironmentManagedAPI); err != nil {
		return IntegrationServiceEnvironmentManagedApisClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the integration service environment managed Apis.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroup - The resource group.
// integrationServiceEnvironmentName - The integration service environment name.
// options - IntegrationServiceEnvironmentManagedApisClientListOptions contains the optional parameters for the IntegrationServiceEnvironmentManagedApisClient.List
// method.
func (client *IntegrationServiceEnvironmentManagedApisClient) NewListPager(resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentManagedApisClientListOptions) *runtime.Pager[IntegrationServiceEnvironmentManagedApisClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[IntegrationServiceEnvironmentManagedApisClientListResponse]{
		More: func(page IntegrationServiceEnvironmentManagedApisClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IntegrationServiceEnvironmentManagedApisClientListResponse) (IntegrationServiceEnvironmentManagedApisClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return IntegrationServiceEnvironmentManagedApisClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return IntegrationServiceEnvironmentManagedApisClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IntegrationServiceEnvironmentManagedApisClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *IntegrationServiceEnvironmentManagedApisClient) listCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentManagedApisClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IntegrationServiceEnvironmentManagedApisClient) listHandleResponse(resp *http.Response) (IntegrationServiceEnvironmentManagedApisClientListResponse, error) {
	result := IntegrationServiceEnvironmentManagedApisClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationServiceEnvironmentManagedAPIListResult); err != nil {
		return IntegrationServiceEnvironmentManagedApisClientListResponse{}, err
	}
	return result, nil
}

// BeginPut - Puts the integration service environment managed Api.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroup - The resource group name.
// integrationServiceEnvironmentName - The integration service environment name.
// apiName - The api name.
// integrationServiceEnvironmentManagedAPI - The integration service environment managed api.
// options - IntegrationServiceEnvironmentManagedApisClientBeginPutOptions contains the optional parameters for the IntegrationServiceEnvironmentManagedApisClient.BeginPut
// method.
func (client *IntegrationServiceEnvironmentManagedApisClient) BeginPut(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, integrationServiceEnvironmentManagedAPI IntegrationServiceEnvironmentManagedAPI, options *IntegrationServiceEnvironmentManagedApisClientBeginPutOptions) (*runtime.Poller[IntegrationServiceEnvironmentManagedApisClientPutResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.put(ctx, resourceGroup, integrationServiceEnvironmentName, apiName, integrationServiceEnvironmentManagedAPI, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[IntegrationServiceEnvironmentManagedApisClientPutResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[IntegrationServiceEnvironmentManagedApisClientPutResponse](options.ResumeToken, client.pl, nil)
	}
}

// Put - Puts the integration service environment managed Api.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
func (client *IntegrationServiceEnvironmentManagedApisClient) put(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, integrationServiceEnvironmentManagedAPI IntegrationServiceEnvironmentManagedAPI, options *IntegrationServiceEnvironmentManagedApisClientBeginPutOptions) (*http.Response, error) {
	req, err := client.putCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, apiName, integrationServiceEnvironmentManagedAPI, options)
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

// putCreateRequest creates the Put request.
func (client *IntegrationServiceEnvironmentManagedApisClient) putCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, integrationServiceEnvironmentManagedAPI IntegrationServiceEnvironmentManagedAPI, options *IntegrationServiceEnvironmentManagedApisClientBeginPutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis/{apiName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	if apiName == "" {
		return nil, errors.New("parameter apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(apiName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, integrationServiceEnvironmentManagedAPI)
}
