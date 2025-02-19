//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armauthorization

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

// RoleAssignmentsClient contains the methods for the RoleAssignments group.
// Don't use this type directly, use NewRoleAssignmentsClient() instead.
type RoleAssignmentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRoleAssignmentsClient creates a new instance of RoleAssignmentsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRoleAssignmentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RoleAssignmentsClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+"/armauthorization.RoleAssignmentsClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RoleAssignmentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Creates a role assignment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-07-01
//   - scope - The scope of the role assignment to create. The scope can be any REST resource instance. For example, use '/subscriptions/{subscription-id}/'
//     for a subscription,
//     '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}' for a resource group, and
//     '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider}/{resource-type}/{resource-name}'
//     for a resource.
//   - roleAssignmentName - A GUID for the role assignment to create. The name must be unique and different for each role assignment.
//   - parameters - Parameters for the role assignment.
//   - options - RoleAssignmentsClientCreateOptions contains the optional parameters for the RoleAssignmentsClient.Create method.
func (client *RoleAssignmentsClient) Create(ctx context.Context, scope string, roleAssignmentName string, parameters RoleAssignmentCreateParameters, options *RoleAssignmentsClientCreateOptions) (RoleAssignmentsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, scope, roleAssignmentName, parameters, options)
	if err != nil {
		return RoleAssignmentsClientCreateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoleAssignmentsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return RoleAssignmentsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *RoleAssignmentsClient) createCreateRequest(ctx context.Context, scope string, roleAssignmentName string, parameters RoleAssignmentCreateParameters, options *RoleAssignmentsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentName == "" {
		return nil, errors.New("parameter roleAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentName}", url.PathEscape(roleAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleResponse handles the Create response.
func (client *RoleAssignmentsClient) createHandleResponse(resp *http.Response) (RoleAssignmentsClientCreateResponse, error) {
	result := RoleAssignmentsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignment); err != nil {
		return RoleAssignmentsClientCreateResponse{}, err
	}
	return result, nil
}

// CreateByID - Creates a role assignment by ID.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-07-01
//   - roleAssignmentID - The fully qualified ID of the role assignment, including the scope, resource name and resource type.
//     Use the format, /{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}.
//     Example: /subscriptions/{subId}/resourcegroups/{rgname}//providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}.
//   - parameters - Parameters for the role assignment.
//   - options - RoleAssignmentsClientCreateByIDOptions contains the optional parameters for the RoleAssignmentsClient.CreateByID
//     method.
func (client *RoleAssignmentsClient) CreateByID(ctx context.Context, roleAssignmentID string, parameters RoleAssignmentCreateParameters, options *RoleAssignmentsClientCreateByIDOptions) (RoleAssignmentsClientCreateByIDResponse, error) {
	req, err := client.createByIDCreateRequest(ctx, roleAssignmentID, parameters, options)
	if err != nil {
		return RoleAssignmentsClientCreateByIDResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoleAssignmentsClientCreateByIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return RoleAssignmentsClientCreateByIDResponse{}, runtime.NewResponseError(resp)
	}
	return client.createByIDHandleResponse(resp)
}

// createByIDCreateRequest creates the CreateByID request.
func (client *RoleAssignmentsClient) createByIDCreateRequest(ctx context.Context, roleAssignmentID string, parameters RoleAssignmentCreateParameters, options *RoleAssignmentsClientCreateByIDOptions) (*policy.Request, error) {
	urlPath := "/{roleAssignmentId}"
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentId}", roleAssignmentID)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createByIDHandleResponse handles the CreateByID response.
func (client *RoleAssignmentsClient) createByIDHandleResponse(resp *http.Response) (RoleAssignmentsClientCreateByIDResponse, error) {
	result := RoleAssignmentsClientCreateByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignment); err != nil {
		return RoleAssignmentsClientCreateByIDResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a role assignment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-07-01
//   - scope - The scope of the role assignment to delete.
//   - roleAssignmentName - The name of the role assignment to delete.
//   - options - RoleAssignmentsClientDeleteOptions contains the optional parameters for the RoleAssignmentsClient.Delete method.
func (client *RoleAssignmentsClient) Delete(ctx context.Context, scope string, roleAssignmentName string, options *RoleAssignmentsClientDeleteOptions) (RoleAssignmentsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, scope, roleAssignmentName, options)
	if err != nil {
		return RoleAssignmentsClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoleAssignmentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return RoleAssignmentsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *RoleAssignmentsClient) deleteCreateRequest(ctx context.Context, scope string, roleAssignmentName string, options *RoleAssignmentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentName == "" {
		return nil, errors.New("parameter roleAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentName}", url.PathEscape(roleAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *RoleAssignmentsClient) deleteHandleResponse(resp *http.Response) (RoleAssignmentsClientDeleteResponse, error) {
	result := RoleAssignmentsClientDeleteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignment); err != nil {
		return RoleAssignmentsClientDeleteResponse{}, err
	}
	return result, nil
}

// DeleteByID - Deletes a role assignment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-07-01
//   - roleAssignmentID - The fully qualified ID of the role assignment, including the scope, resource name and resource type.
//     Use the format, /{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}.
//     Example: /subscriptions/{subId}/resourcegroups/{rgname}//providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}.
//   - options - RoleAssignmentsClientDeleteByIDOptions contains the optional parameters for the RoleAssignmentsClient.DeleteByID
//     method.
func (client *RoleAssignmentsClient) DeleteByID(ctx context.Context, roleAssignmentID string, options *RoleAssignmentsClientDeleteByIDOptions) (RoleAssignmentsClientDeleteByIDResponse, error) {
	req, err := client.deleteByIDCreateRequest(ctx, roleAssignmentID, options)
	if err != nil {
		return RoleAssignmentsClientDeleteByIDResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoleAssignmentsClientDeleteByIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return RoleAssignmentsClientDeleteByIDResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteByIDHandleResponse(resp)
}

// deleteByIDCreateRequest creates the DeleteByID request.
func (client *RoleAssignmentsClient) deleteByIDCreateRequest(ctx context.Context, roleAssignmentID string, options *RoleAssignmentsClientDeleteByIDOptions) (*policy.Request, error) {
	urlPath := "/{roleAssignmentId}"
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentId}", roleAssignmentID)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteByIDHandleResponse handles the DeleteByID response.
func (client *RoleAssignmentsClient) deleteByIDHandleResponse(resp *http.Response) (RoleAssignmentsClientDeleteByIDResponse, error) {
	result := RoleAssignmentsClientDeleteByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignment); err != nil {
		return RoleAssignmentsClientDeleteByIDResponse{}, err
	}
	return result, nil
}

// Get - Get the specified role assignment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-07-01
//   - scope - The scope of the role assignment.
//   - roleAssignmentName - The name of the role assignment to get.
//   - options - RoleAssignmentsClientGetOptions contains the optional parameters for the RoleAssignmentsClient.Get method.
func (client *RoleAssignmentsClient) Get(ctx context.Context, scope string, roleAssignmentName string, options *RoleAssignmentsClientGetOptions) (RoleAssignmentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, roleAssignmentName, options)
	if err != nil {
		return RoleAssignmentsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoleAssignmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleAssignmentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RoleAssignmentsClient) getCreateRequest(ctx context.Context, scope string, roleAssignmentName string, options *RoleAssignmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentName == "" {
		return nil, errors.New("parameter roleAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentName}", url.PathEscape(roleAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoleAssignmentsClient) getHandleResponse(resp *http.Response) (RoleAssignmentsClientGetResponse, error) {
	result := RoleAssignmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignment); err != nil {
		return RoleAssignmentsClientGetResponse{}, err
	}
	return result, nil
}

// GetByID - Gets a role assignment by ID.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-07-01
//   - roleAssignmentID - The fully qualified ID of the role assignment, including the scope, resource name and resource type.
//     Use the format, /{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}.
//     Example: /subscriptions/{subId}/resourcegroups/{rgname}//providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}.
//   - options - RoleAssignmentsClientGetByIDOptions contains the optional parameters for the RoleAssignmentsClient.GetByID method.
func (client *RoleAssignmentsClient) GetByID(ctx context.Context, roleAssignmentID string, options *RoleAssignmentsClientGetByIDOptions) (RoleAssignmentsClientGetByIDResponse, error) {
	req, err := client.getByIDCreateRequest(ctx, roleAssignmentID, options)
	if err != nil {
		return RoleAssignmentsClientGetByIDResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoleAssignmentsClientGetByIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleAssignmentsClientGetByIDResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByIDHandleResponse(resp)
}

// getByIDCreateRequest creates the GetByID request.
func (client *RoleAssignmentsClient) getByIDCreateRequest(ctx context.Context, roleAssignmentID string, options *RoleAssignmentsClientGetByIDOptions) (*policy.Request, error) {
	urlPath := "/{roleAssignmentId}"
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentId}", roleAssignmentID)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByIDHandleResponse handles the GetByID response.
func (client *RoleAssignmentsClient) getByIDHandleResponse(resp *http.Response) (RoleAssignmentsClientGetByIDResponse, error) {
	result := RoleAssignmentsClientGetByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignment); err != nil {
		return RoleAssignmentsClientGetByIDResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all role assignments for the subscription.
//
// Generated from API version 2015-07-01
//   - options - RoleAssignmentsClientListOptions contains the optional parameters for the RoleAssignmentsClient.NewListPager
//     method.
func (client *RoleAssignmentsClient) NewListPager(options *RoleAssignmentsClientListOptions) *runtime.Pager[RoleAssignmentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RoleAssignmentsClientListResponse]{
		More: func(page RoleAssignmentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RoleAssignmentsClientListResponse) (RoleAssignmentsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RoleAssignmentsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RoleAssignmentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RoleAssignmentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *RoleAssignmentsClient) listCreateRequest(ctx context.Context, options *RoleAssignmentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/roleAssignments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2015-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RoleAssignmentsClient) listHandleResponse(resp *http.Response) (RoleAssignmentsClientListResponse, error) {
	result := RoleAssignmentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentListResult); err != nil {
		return RoleAssignmentsClientListResponse{}, err
	}
	return result, nil
}

// NewListForResourcePager - Gets role assignments for a resource.
//
// Generated from API version 2015-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceProviderNamespace - The namespace of the resource provider.
//   - parentResourcePath - The parent resource identity.
//   - resourceType - The resource type of the resource.
//   - resourceName - The name of the resource to get role assignments for.
//   - options - RoleAssignmentsClientListForResourceOptions contains the optional parameters for the RoleAssignmentsClient.NewListForResourcePager
//     method.
func (client *RoleAssignmentsClient) NewListForResourcePager(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, options *RoleAssignmentsClientListForResourceOptions) *runtime.Pager[RoleAssignmentsClientListForResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[RoleAssignmentsClientListForResourceResponse]{
		More: func(page RoleAssignmentsClientListForResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RoleAssignmentsClientListForResourceResponse) (RoleAssignmentsClientListForResourceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listForResourceCreateRequest(ctx, resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RoleAssignmentsClientListForResourceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RoleAssignmentsClientListForResourceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RoleAssignmentsClientListForResourceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listForResourceHandleResponse(resp)
		},
	})
}

// listForResourceCreateRequest creates the ListForResource request.
func (client *RoleAssignmentsClient) listForResourceCreateRequest(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, options *RoleAssignmentsClientListForResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/roleAssignments"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
	urlPath = strings.ReplaceAll(urlPath, "{parentResourcePath}", parentResourcePath)
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", resourceType)
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2015-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForResourceHandleResponse handles the ListForResource response.
func (client *RoleAssignmentsClient) listForResourceHandleResponse(resp *http.Response) (RoleAssignmentsClientListForResourceResponse, error) {
	result := RoleAssignmentsClientListForResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentListResult); err != nil {
		return RoleAssignmentsClientListForResourceResponse{}, err
	}
	return result, nil
}

// NewListForResourceGroupPager - Gets role assignments for a resource group.
//
// Generated from API version 2015-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - RoleAssignmentsClientListForResourceGroupOptions contains the optional parameters for the RoleAssignmentsClient.NewListForResourceGroupPager
//     method.
func (client *RoleAssignmentsClient) NewListForResourceGroupPager(resourceGroupName string, options *RoleAssignmentsClientListForResourceGroupOptions) *runtime.Pager[RoleAssignmentsClientListForResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[RoleAssignmentsClientListForResourceGroupResponse]{
		More: func(page RoleAssignmentsClientListForResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RoleAssignmentsClientListForResourceGroupResponse) (RoleAssignmentsClientListForResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listForResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RoleAssignmentsClientListForResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RoleAssignmentsClientListForResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RoleAssignmentsClientListForResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listForResourceGroupHandleResponse(resp)
		},
	})
}

// listForResourceGroupCreateRequest creates the ListForResourceGroup request.
func (client *RoleAssignmentsClient) listForResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *RoleAssignmentsClientListForResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/roleAssignments"
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
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2015-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForResourceGroupHandleResponse handles the ListForResourceGroup response.
func (client *RoleAssignmentsClient) listForResourceGroupHandleResponse(resp *http.Response) (RoleAssignmentsClientListForResourceGroupResponse, error) {
	result := RoleAssignmentsClientListForResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentListResult); err != nil {
		return RoleAssignmentsClientListForResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListForScopePager - Gets role assignments for a scope.
//
// Generated from API version 2015-07-01
//   - scope - The scope of the role assignments.
//   - options - RoleAssignmentsClientListForScopeOptions contains the optional parameters for the RoleAssignmentsClient.NewListForScopePager
//     method.
func (client *RoleAssignmentsClient) NewListForScopePager(scope string, options *RoleAssignmentsClientListForScopeOptions) *runtime.Pager[RoleAssignmentsClientListForScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[RoleAssignmentsClientListForScopeResponse]{
		More: func(page RoleAssignmentsClientListForScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RoleAssignmentsClientListForScopeResponse) (RoleAssignmentsClientListForScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listForScopeCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RoleAssignmentsClientListForScopeResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RoleAssignmentsClientListForScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RoleAssignmentsClientListForScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listForScopeHandleResponse(resp)
		},
	})
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *RoleAssignmentsClient) listForScopeCreateRequest(ctx context.Context, scope string, options *RoleAssignmentsClientListForScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignments"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2015-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForScopeHandleResponse handles the ListForScope response.
func (client *RoleAssignmentsClient) listForScopeHandleResponse(resp *http.Response) (RoleAssignmentsClientListForScopeResponse, error) {
	result := RoleAssignmentsClientListForScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentListResult); err != nil {
		return RoleAssignmentsClientListForScopeResponse{}, err
	}
	return result, nil
}
