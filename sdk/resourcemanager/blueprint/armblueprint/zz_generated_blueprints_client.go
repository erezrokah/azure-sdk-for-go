//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblueprint

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

// BlueprintsClient contains the methods for the Blueprints group.
// Don't use this type directly, use NewBlueprintsClient() instead.
type BlueprintsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewBlueprintsClient creates a new instance of BlueprintsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBlueprintsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*BlueprintsClient, error) {
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
	client := &BlueprintsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a blueprint definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01-preview
// resourceScope - The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'),
// subscription (format: '/subscriptions/{subscriptionId}').
// blueprintName - Name of the blueprint definition.
// blueprint - Blueprint definition.
// options - BlueprintsClientCreateOrUpdateOptions contains the optional parameters for the BlueprintsClient.CreateOrUpdate
// method.
func (client *BlueprintsClient) CreateOrUpdate(ctx context.Context, resourceScope string, blueprintName string, blueprint Blueprint, options *BlueprintsClientCreateOrUpdateOptions) (BlueprintsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceScope, blueprintName, blueprint, options)
	if err != nil {
		return BlueprintsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BlueprintsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return BlueprintsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *BlueprintsClient) createOrUpdateCreateRequest(ctx context.Context, resourceScope string, blueprintName string, blueprint Blueprint, options *BlueprintsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{resourceScope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceScope}", resourceScope)
	if blueprintName == "" {
		return nil, errors.New("parameter blueprintName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blueprintName}", url.PathEscape(blueprintName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, blueprint)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *BlueprintsClient) createOrUpdateHandleResponse(resp *http.Response) (BlueprintsClientCreateOrUpdateResponse, error) {
	result := BlueprintsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Blueprint); err != nil {
		return BlueprintsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a blueprint definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01-preview
// resourceScope - The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'),
// subscription (format: '/subscriptions/{subscriptionId}').
// blueprintName - Name of the blueprint definition.
// options - BlueprintsClientDeleteOptions contains the optional parameters for the BlueprintsClient.Delete method.
func (client *BlueprintsClient) Delete(ctx context.Context, resourceScope string, blueprintName string, options *BlueprintsClientDeleteOptions) (BlueprintsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceScope, blueprintName, options)
	if err != nil {
		return BlueprintsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BlueprintsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return BlueprintsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *BlueprintsClient) deleteCreateRequest(ctx context.Context, resourceScope string, blueprintName string, options *BlueprintsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{resourceScope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceScope}", resourceScope)
	if blueprintName == "" {
		return nil, errors.New("parameter blueprintName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blueprintName}", url.PathEscape(blueprintName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *BlueprintsClient) deleteHandleResponse(resp *http.Response) (BlueprintsClientDeleteResponse, error) {
	result := BlueprintsClientDeleteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Blueprint); err != nil {
		return BlueprintsClientDeleteResponse{}, err
	}
	return result, nil
}

// Get - Get a blueprint definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01-preview
// resourceScope - The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'),
// subscription (format: '/subscriptions/{subscriptionId}').
// blueprintName - Name of the blueprint definition.
// options - BlueprintsClientGetOptions contains the optional parameters for the BlueprintsClient.Get method.
func (client *BlueprintsClient) Get(ctx context.Context, resourceScope string, blueprintName string, options *BlueprintsClientGetOptions) (BlueprintsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceScope, blueprintName, options)
	if err != nil {
		return BlueprintsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BlueprintsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BlueprintsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BlueprintsClient) getCreateRequest(ctx context.Context, resourceScope string, blueprintName string, options *BlueprintsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceScope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceScope}", resourceScope)
	if blueprintName == "" {
		return nil, errors.New("parameter blueprintName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blueprintName}", url.PathEscape(blueprintName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BlueprintsClient) getHandleResponse(resp *http.Response) (BlueprintsClientGetResponse, error) {
	result := BlueprintsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Blueprint); err != nil {
		return BlueprintsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List blueprint definitions.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01-preview
// resourceScope - The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'),
// subscription (format: '/subscriptions/{subscriptionId}').
// options - BlueprintsClientListOptions contains the optional parameters for the BlueprintsClient.List method.
func (client *BlueprintsClient) NewListPager(resourceScope string, options *BlueprintsClientListOptions) *runtime.Pager[BlueprintsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BlueprintsClientListResponse]{
		More: func(page BlueprintsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BlueprintsClientListResponse) (BlueprintsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceScope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BlueprintsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return BlueprintsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BlueprintsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *BlueprintsClient) listCreateRequest(ctx context.Context, resourceScope string, options *BlueprintsClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceScope}/providers/Microsoft.Blueprint/blueprints"
	urlPath = strings.ReplaceAll(urlPath, "{resourceScope}", resourceScope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BlueprintsClient) listHandleResponse(resp *http.Response) (BlueprintsClientListResponse, error) {
	result := BlueprintsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.List); err != nil {
		return BlueprintsClientListResponse{}, err
	}
	return result, nil
}
