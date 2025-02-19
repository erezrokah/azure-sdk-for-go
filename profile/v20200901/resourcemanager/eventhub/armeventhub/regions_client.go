//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armeventhub

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

// RegionsClient contains the methods for the Regions group.
// Don't use this type directly, use NewRegionsClient() instead.
type RegionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRegionsClient creates a new instance of RegionsClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRegionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RegionsClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+"/armeventhub.RegionsClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RegionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListBySKUPager - Gets the available Regions for a given sku
//
// Generated from API version 2017-04-01
//   - sku - The sku type.
//   - options - RegionsClientListBySKUOptions contains the optional parameters for the RegionsClient.NewListBySKUPager method.
func (client *RegionsClient) NewListBySKUPager(sku string, options *RegionsClientListBySKUOptions) *runtime.Pager[RegionsClientListBySKUResponse] {
	return runtime.NewPager(runtime.PagingHandler[RegionsClientListBySKUResponse]{
		More: func(page RegionsClientListBySKUResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RegionsClientListBySKUResponse) (RegionsClientListBySKUResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySKUCreateRequest(ctx, sku, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RegionsClientListBySKUResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RegionsClientListBySKUResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RegionsClientListBySKUResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySKUHandleResponse(resp)
		},
	})
}

// listBySKUCreateRequest creates the ListBySKU request.
func (client *RegionsClient) listBySKUCreateRequest(ctx context.Context, sku string, options *RegionsClientListBySKUOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EventHub/sku/{sku}/regions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if sku == "" {
		return nil, errors.New("parameter sku cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sku}", url.PathEscape(sku))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySKUHandleResponse handles the ListBySKU response.
func (client *RegionsClient) listBySKUHandleResponse(resp *http.Response) (RegionsClientListBySKUResponse, error) {
	result := RegionsClientListBySKUResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MessagingRegionsListResult); err != nil {
		return RegionsClientListBySKUResponse{}, err
	}
	return result, nil
}
