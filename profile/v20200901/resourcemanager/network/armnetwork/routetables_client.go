//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/profile/v20200901/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RouteTablesClient contains the methods for the RouteTables group.
// Don't use this type directly, use NewRouteTablesClient() instead.
type RouteTablesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRouteTablesClient creates a new instance of RouteTablesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRouteTablesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RouteTablesClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+"/armnetwork.RouteTablesClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RouteTablesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or updates a route table in a specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
//   - resourceGroupName - The name of the resource group.
//   - routeTableName - The name of the route table.
//   - parameters - Parameters supplied to the create or update route table operation.
//   - options - RouteTablesClientBeginCreateOrUpdateOptions contains the optional parameters for the RouteTablesClient.BeginCreateOrUpdate
//     method.
func (client *RouteTablesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, routeTableName string, parameters RouteTable, options *RouteTablesClientBeginCreateOrUpdateOptions) (*runtime.Poller[RouteTablesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, routeTableName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[RouteTablesClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[RouteTablesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or updates a route table in a specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
func (client *RouteTablesClient) createOrUpdate(ctx context.Context, resourceGroupName string, routeTableName string, parameters RouteTable, options *RouteTablesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, routeTableName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RouteTablesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, routeTableName string, parameters RouteTable, options *RouteTablesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeTableName == "" {
		return nil, errors.New("parameter routeTableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified route table.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
//   - resourceGroupName - The name of the resource group.
//   - routeTableName - The name of the route table.
//   - options - RouteTablesClientBeginDeleteOptions contains the optional parameters for the RouteTablesClient.BeginDelete method.
func (client *RouteTablesClient) BeginDelete(ctx context.Context, resourceGroupName string, routeTableName string, options *RouteTablesClientBeginDeleteOptions) (*runtime.Poller[RouteTablesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, routeTableName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[RouteTablesClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[RouteTablesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified route table.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
func (client *RouteTablesClient) deleteOperation(ctx context.Context, resourceGroupName string, routeTableName string, options *RouteTablesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, routeTableName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RouteTablesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, routeTableName string, options *RouteTablesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeTableName == "" {
		return nil, errors.New("parameter routeTableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the specified route table.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
//   - resourceGroupName - The name of the resource group.
//   - routeTableName - The name of the route table.
//   - options - RouteTablesClientGetOptions contains the optional parameters for the RouteTablesClient.Get method.
func (client *RouteTablesClient) Get(ctx context.Context, resourceGroupName string, routeTableName string, options *RouteTablesClientGetOptions) (RouteTablesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, routeTableName, options)
	if err != nil {
		return RouteTablesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RouteTablesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RouteTablesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RouteTablesClient) getCreateRequest(ctx context.Context, resourceGroupName string, routeTableName string, options *RouteTablesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeTableName == "" {
		return nil, errors.New("parameter routeTableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RouteTablesClient) getHandleResponse(resp *http.Response) (RouteTablesClientGetResponse, error) {
	result := RouteTablesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RouteTable); err != nil {
		return RouteTablesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all route tables in a resource group.
//
// Generated from API version 2018-11-01
//   - resourceGroupName - The name of the resource group.
//   - options - RouteTablesClientListOptions contains the optional parameters for the RouteTablesClient.NewListPager method.
func (client *RouteTablesClient) NewListPager(resourceGroupName string, options *RouteTablesClientListOptions) *runtime.Pager[RouteTablesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RouteTablesClientListResponse]{
		More: func(page RouteTablesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RouteTablesClientListResponse) (RouteTablesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RouteTablesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RouteTablesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RouteTablesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *RouteTablesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *RouteTablesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RouteTablesClient) listHandleResponse(resp *http.Response) (RouteTablesClientListResponse, error) {
	result := RouteTablesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RouteTableListResult); err != nil {
		return RouteTablesClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Gets all route tables in a subscription.
//
// Generated from API version 2018-11-01
//   - options - RouteTablesClientListAllOptions contains the optional parameters for the RouteTablesClient.NewListAllPager method.
func (client *RouteTablesClient) NewListAllPager(options *RouteTablesClientListAllOptions) *runtime.Pager[RouteTablesClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[RouteTablesClientListAllResponse]{
		More: func(page RouteTablesClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RouteTablesClientListAllResponse) (RouteTablesClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RouteTablesClientListAllResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RouteTablesClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RouteTablesClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *RouteTablesClient) listAllCreateRequest(ctx context.Context, options *RouteTablesClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/routeTables"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *RouteTablesClient) listAllHandleResponse(resp *http.Response) (RouteTablesClientListAllResponse, error) {
	result := RouteTablesClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RouteTableListResult); err != nil {
		return RouteTablesClientListAllResponse{}, err
	}
	return result, nil
}

// BeginUpdateTags - Updates a route table tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
//   - resourceGroupName - The name of the resource group.
//   - routeTableName - The name of the route table.
//   - parameters - Parameters supplied to update route table tags.
//   - options - RouteTablesClientBeginUpdateTagsOptions contains the optional parameters for the RouteTablesClient.BeginUpdateTags
//     method.
func (client *RouteTablesClient) BeginUpdateTags(ctx context.Context, resourceGroupName string, routeTableName string, parameters TagsObject, options *RouteTablesClientBeginUpdateTagsOptions) (*runtime.Poller[RouteTablesClientUpdateTagsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateTags(ctx, resourceGroupName, routeTableName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[RouteTablesClientUpdateTagsResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[RouteTablesClientUpdateTagsResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UpdateTags - Updates a route table tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
func (client *RouteTablesClient) updateTags(ctx context.Context, resourceGroupName string, routeTableName string, parameters TagsObject, options *RouteTablesClientBeginUpdateTagsOptions) (*http.Response, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, routeTableName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *RouteTablesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, routeTableName string, parameters TagsObject, options *RouteTablesClientBeginUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeTableName == "" {
		return nil, errors.New("parameter routeTableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
