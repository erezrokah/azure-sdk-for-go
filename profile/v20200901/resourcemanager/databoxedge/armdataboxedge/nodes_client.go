//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdataboxedge

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

// NodesClient contains the methods for the Nodes group.
// Don't use this type directly, use NewNodesClient() instead.
type NodesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNodesClient creates a new instance of NodesClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNodesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NodesClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+"/armdataboxedge.NodesClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NodesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByDataBoxEdgeDevicePager - Gets all the nodes currently configured under this Data Box Edge device
//
// Generated from API version 2019-08-01
//   - deviceName - The device name.
//   - resourceGroupName - The resource group name.
//   - options - NodesClientListByDataBoxEdgeDeviceOptions contains the optional parameters for the NodesClient.NewListByDataBoxEdgeDevicePager
//     method.
func (client *NodesClient) NewListByDataBoxEdgeDevicePager(deviceName string, resourceGroupName string, options *NodesClientListByDataBoxEdgeDeviceOptions) *runtime.Pager[NodesClientListByDataBoxEdgeDeviceResponse] {
	return runtime.NewPager(runtime.PagingHandler[NodesClientListByDataBoxEdgeDeviceResponse]{
		More: func(page NodesClientListByDataBoxEdgeDeviceResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *NodesClientListByDataBoxEdgeDeviceResponse) (NodesClientListByDataBoxEdgeDeviceResponse, error) {
			req, err := client.listByDataBoxEdgeDeviceCreateRequest(ctx, deviceName, resourceGroupName, options)
			if err != nil {
				return NodesClientListByDataBoxEdgeDeviceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return NodesClientListByDataBoxEdgeDeviceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NodesClientListByDataBoxEdgeDeviceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDataBoxEdgeDeviceHandleResponse(resp)
		},
	})
}

// listByDataBoxEdgeDeviceCreateRequest creates the ListByDataBoxEdgeDevice request.
func (client *NodesClient) listByDataBoxEdgeDeviceCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *NodesClientListByDataBoxEdgeDeviceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/nodes"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDataBoxEdgeDeviceHandleResponse handles the ListByDataBoxEdgeDevice response.
func (client *NodesClient) listByDataBoxEdgeDeviceHandleResponse(resp *http.Response) (NodesClientListByDataBoxEdgeDeviceResponse, error) {
	result := NodesClientListByDataBoxEdgeDeviceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NodeList); err != nil {
		return NodesClientListByDataBoxEdgeDeviceResponse{}, err
	}
	return result, nil
}
