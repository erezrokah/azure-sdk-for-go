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

// SecurityGroupsClient contains the methods for the NetworkSecurityGroups group.
// Don't use this type directly, use NewSecurityGroupsClient() instead.
type SecurityGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSecurityGroupsClient creates a new instance of SecurityGroupsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSecurityGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SecurityGroupsClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+"/armnetwork.SecurityGroupsClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SecurityGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a network security group in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
//   - resourceGroupName - The name of the resource group.
//   - networkSecurityGroupName - The name of the network security group.
//   - parameters - Parameters supplied to the create or update network security group operation.
//   - options - SecurityGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the SecurityGroupsClient.BeginCreateOrUpdate
//     method.
func (client *SecurityGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters SecurityGroup, options *SecurityGroupsClientBeginCreateOrUpdateOptions) (*runtime.Poller[SecurityGroupsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, networkSecurityGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SecurityGroupsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[SecurityGroupsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a network security group in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
func (client *SecurityGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters SecurityGroup, options *SecurityGroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, parameters, options)
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
func (client *SecurityGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters SecurityGroup, options *SecurityGroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
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

// BeginDelete - Deletes the specified network security group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
//   - resourceGroupName - The name of the resource group.
//   - networkSecurityGroupName - The name of the network security group.
//   - options - SecurityGroupsClientBeginDeleteOptions contains the optional parameters for the SecurityGroupsClient.BeginDelete
//     method.
func (client *SecurityGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *SecurityGroupsClientBeginDeleteOptions) (*runtime.Poller[SecurityGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, networkSecurityGroupName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SecurityGroupsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[SecurityGroupsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified network security group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
func (client *SecurityGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *SecurityGroupsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, options)
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
func (client *SecurityGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *SecurityGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
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

// Get - Gets the specified network security group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
//   - resourceGroupName - The name of the resource group.
//   - networkSecurityGroupName - The name of the network security group.
//   - options - SecurityGroupsClientGetOptions contains the optional parameters for the SecurityGroupsClient.Get method.
func (client *SecurityGroupsClient) Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *SecurityGroupsClientGetOptions) (SecurityGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, options)
	if err != nil {
		return SecurityGroupsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SecurityGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SecurityGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SecurityGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *SecurityGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
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
func (client *SecurityGroupsClient) getHandleResponse(resp *http.Response) (SecurityGroupsClientGetResponse, error) {
	result := SecurityGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityGroup); err != nil {
		return SecurityGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all network security groups in a resource group.
//
// Generated from API version 2018-11-01
//   - resourceGroupName - The name of the resource group.
//   - options - SecurityGroupsClientListOptions contains the optional parameters for the SecurityGroupsClient.NewListPager method.
func (client *SecurityGroupsClient) NewListPager(resourceGroupName string, options *SecurityGroupsClientListOptions) *runtime.Pager[SecurityGroupsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SecurityGroupsClientListResponse]{
		More: func(page SecurityGroupsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecurityGroupsClientListResponse) (SecurityGroupsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SecurityGroupsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SecurityGroupsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SecurityGroupsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SecurityGroupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *SecurityGroupsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups"
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
func (client *SecurityGroupsClient) listHandleResponse(resp *http.Response) (SecurityGroupsClientListResponse, error) {
	result := SecurityGroupsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityGroupListResult); err != nil {
		return SecurityGroupsClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Gets all network security groups in a subscription.
//
// Generated from API version 2018-11-01
//   - options - SecurityGroupsClientListAllOptions contains the optional parameters for the SecurityGroupsClient.NewListAllPager
//     method.
func (client *SecurityGroupsClient) NewListAllPager(options *SecurityGroupsClientListAllOptions) *runtime.Pager[SecurityGroupsClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[SecurityGroupsClientListAllResponse]{
		More: func(page SecurityGroupsClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecurityGroupsClientListAllResponse) (SecurityGroupsClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SecurityGroupsClientListAllResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SecurityGroupsClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SecurityGroupsClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *SecurityGroupsClient) listAllCreateRequest(ctx context.Context, options *SecurityGroupsClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkSecurityGroups"
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
func (client *SecurityGroupsClient) listAllHandleResponse(resp *http.Response) (SecurityGroupsClientListAllResponse, error) {
	result := SecurityGroupsClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityGroupListResult); err != nil {
		return SecurityGroupsClientListAllResponse{}, err
	}
	return result, nil
}

// BeginUpdateTags - Updates a network security group tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
//   - resourceGroupName - The name of the resource group.
//   - networkSecurityGroupName - The name of the network security group.
//   - parameters - Parameters supplied to update network security group tags.
//   - options - SecurityGroupsClientBeginUpdateTagsOptions contains the optional parameters for the SecurityGroupsClient.BeginUpdateTags
//     method.
func (client *SecurityGroupsClient) BeginUpdateTags(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters TagsObject, options *SecurityGroupsClientBeginUpdateTagsOptions) (*runtime.Poller[SecurityGroupsClientUpdateTagsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateTags(ctx, resourceGroupName, networkSecurityGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SecurityGroupsClientUpdateTagsResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[SecurityGroupsClientUpdateTagsResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UpdateTags - Updates a network security group tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01
func (client *SecurityGroupsClient) updateTags(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters TagsObject, options *SecurityGroupsClientBeginUpdateTagsOptions) (*http.Response, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, parameters, options)
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
func (client *SecurityGroupsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters TagsObject, options *SecurityGroupsClientBeginUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
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
