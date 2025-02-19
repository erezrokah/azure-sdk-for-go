//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsynapse

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

// KustoPoolDataConnectionsClient contains the methods for the KustoPoolDataConnections group.
// Don't use this type directly, use NewKustoPoolDataConnectionsClient() instead.
type KustoPoolDataConnectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewKustoPoolDataConnectionsClient creates a new instance of KustoPoolDataConnectionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewKustoPoolDataConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*KustoPoolDataConnectionsClient, error) {
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
	client := &KustoPoolDataConnectionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CheckNameAvailability - Checks that the data connection name is valid and is not already in use.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - kustoPoolName - The name of the Kusto pool.
//   - databaseName - The name of the database in the Kusto pool.
//   - dataConnectionName - The name of the data connection.
//   - options - KustoPoolDataConnectionsClientCheckNameAvailabilityOptions contains the optional parameters for the KustoPoolDataConnectionsClient.CheckNameAvailability
//     method.
func (client *KustoPoolDataConnectionsClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName DataConnectionCheckNameRequest, options *KustoPoolDataConnectionsClientCheckNameAvailabilityOptions) (KustoPoolDataConnectionsClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, resourceGroupName, workspaceName, kustoPoolName, databaseName, dataConnectionName, options)
	if err != nil {
		return KustoPoolDataConnectionsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return KustoPoolDataConnectionsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return KustoPoolDataConnectionsClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *KustoPoolDataConnectionsClient) checkNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName DataConnectionCheckNameRequest, options *KustoPoolDataConnectionsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/databases/{databaseName}/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if kustoPoolName == "" {
		return nil, errors.New("parameter kustoPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kustoPoolName}", url.PathEscape(kustoPoolName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, dataConnectionName)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *KustoPoolDataConnectionsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (KustoPoolDataConnectionsClientCheckNameAvailabilityResponse, error) {
	result := KustoPoolDataConnectionsClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameResult); err != nil {
		return KustoPoolDataConnectionsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Creates or updates a data connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - kustoPoolName - The name of the Kusto pool.
//   - databaseName - The name of the database in the Kusto pool.
//   - dataConnectionName - The name of the data connection.
//   - parameters - The data connection parameters supplied to the CreateOrUpdate operation.
//   - options - KustoPoolDataConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the KustoPoolDataConnectionsClient.BeginCreateOrUpdate
//     method.
func (client *KustoPoolDataConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName string, parameters DataConnectionClassification, options *KustoPoolDataConnectionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[KustoPoolDataConnectionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, kustoPoolName, databaseName, dataConnectionName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[KustoPoolDataConnectionsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[KustoPoolDataConnectionsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a data connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01-preview
func (client *KustoPoolDataConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName string, parameters DataConnectionClassification, options *KustoPoolDataConnectionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, kustoPoolName, databaseName, dataConnectionName, parameters, options)
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
func (client *KustoPoolDataConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName string, parameters DataConnectionClassification, options *KustoPoolDataConnectionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/databases/{databaseName}/dataConnections/{dataConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if kustoPoolName == "" {
		return nil, errors.New("parameter kustoPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kustoPoolName}", url.PathEscape(kustoPoolName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if dataConnectionName == "" {
		return nil, errors.New("parameter dataConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataConnectionName}", url.PathEscape(dataConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDataConnectionValidation - Checks that the data connection parameters are valid.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - kustoPoolName - The name of the Kusto pool.
//   - databaseName - The name of the database in the Kusto pool.
//   - parameters - The data connection parameters supplied to the CreateOrUpdate operation.
//   - options - KustoPoolDataConnectionsClientBeginDataConnectionValidationOptions contains the optional parameters for the KustoPoolDataConnectionsClient.BeginDataConnectionValidation
//     method.
func (client *KustoPoolDataConnectionsClient) BeginDataConnectionValidation(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, parameters DataConnectionValidation, options *KustoPoolDataConnectionsClientBeginDataConnectionValidationOptions) (*runtime.Poller[KustoPoolDataConnectionsClientDataConnectionValidationResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.dataConnectionValidation(ctx, resourceGroupName, workspaceName, kustoPoolName, databaseName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[KustoPoolDataConnectionsClientDataConnectionValidationResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[KustoPoolDataConnectionsClientDataConnectionValidationResponse](options.ResumeToken, client.pl, nil)
	}
}

// DataConnectionValidation - Checks that the data connection parameters are valid.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01-preview
func (client *KustoPoolDataConnectionsClient) dataConnectionValidation(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, parameters DataConnectionValidation, options *KustoPoolDataConnectionsClientBeginDataConnectionValidationOptions) (*http.Response, error) {
	req, err := client.dataConnectionValidationCreateRequest(ctx, resourceGroupName, workspaceName, kustoPoolName, databaseName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// dataConnectionValidationCreateRequest creates the DataConnectionValidation request.
func (client *KustoPoolDataConnectionsClient) dataConnectionValidationCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, parameters DataConnectionValidation, options *KustoPoolDataConnectionsClientBeginDataConnectionValidationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/databases/{databaseName}/dataConnectionValidation"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if kustoPoolName == "" {
		return nil, errors.New("parameter kustoPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kustoPoolName}", url.PathEscape(kustoPoolName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the data connection with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - kustoPoolName - The name of the Kusto pool.
//   - databaseName - The name of the database in the Kusto pool.
//   - dataConnectionName - The name of the data connection.
//   - options - KustoPoolDataConnectionsClientBeginDeleteOptions contains the optional parameters for the KustoPoolDataConnectionsClient.BeginDelete
//     method.
func (client *KustoPoolDataConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName string, options *KustoPoolDataConnectionsClientBeginDeleteOptions) (*runtime.Poller[KustoPoolDataConnectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, kustoPoolName, databaseName, dataConnectionName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[KustoPoolDataConnectionsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[KustoPoolDataConnectionsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the data connection with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01-preview
func (client *KustoPoolDataConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName string, options *KustoPoolDataConnectionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, kustoPoolName, databaseName, dataConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *KustoPoolDataConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName string, options *KustoPoolDataConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/databases/{databaseName}/dataConnections/{dataConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if kustoPoolName == "" {
		return nil, errors.New("parameter kustoPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kustoPoolName}", url.PathEscape(kustoPoolName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if dataConnectionName == "" {
		return nil, errors.New("parameter dataConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataConnectionName}", url.PathEscape(dataConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns a data connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - kustoPoolName - The name of the Kusto pool.
//   - databaseName - The name of the database in the Kusto pool.
//   - dataConnectionName - The name of the data connection.
//   - options - KustoPoolDataConnectionsClientGetOptions contains the optional parameters for the KustoPoolDataConnectionsClient.Get
//     method.
func (client *KustoPoolDataConnectionsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName string, options *KustoPoolDataConnectionsClientGetOptions) (KustoPoolDataConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, kustoPoolName, databaseName, dataConnectionName, options)
	if err != nil {
		return KustoPoolDataConnectionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return KustoPoolDataConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return KustoPoolDataConnectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *KustoPoolDataConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName string, options *KustoPoolDataConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/databases/{databaseName}/dataConnections/{dataConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if kustoPoolName == "" {
		return nil, errors.New("parameter kustoPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kustoPoolName}", url.PathEscape(kustoPoolName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if dataConnectionName == "" {
		return nil, errors.New("parameter dataConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataConnectionName}", url.PathEscape(dataConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *KustoPoolDataConnectionsClient) getHandleResponse(resp *http.Response) (KustoPoolDataConnectionsClientGetResponse, error) {
	result := KustoPoolDataConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return KustoPoolDataConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDatabasePager - Returns the list of data connections of the given Kusto pool database.
//
// Generated from API version 2021-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - kustoPoolName - The name of the Kusto pool.
//   - databaseName - The name of the database in the Kusto pool.
//   - options - KustoPoolDataConnectionsClientListByDatabaseOptions contains the optional parameters for the KustoPoolDataConnectionsClient.NewListByDatabasePager
//     method.
func (client *KustoPoolDataConnectionsClient) NewListByDatabasePager(resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, options *KustoPoolDataConnectionsClientListByDatabaseOptions) *runtime.Pager[KustoPoolDataConnectionsClientListByDatabaseResponse] {
	return runtime.NewPager(runtime.PagingHandler[KustoPoolDataConnectionsClientListByDatabaseResponse]{
		More: func(page KustoPoolDataConnectionsClientListByDatabaseResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *KustoPoolDataConnectionsClientListByDatabaseResponse) (KustoPoolDataConnectionsClientListByDatabaseResponse, error) {
			req, err := client.listByDatabaseCreateRequest(ctx, resourceGroupName, workspaceName, kustoPoolName, databaseName, options)
			if err != nil {
				return KustoPoolDataConnectionsClientListByDatabaseResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return KustoPoolDataConnectionsClientListByDatabaseResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return KustoPoolDataConnectionsClientListByDatabaseResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDatabaseHandleResponse(resp)
		},
	})
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *KustoPoolDataConnectionsClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, options *KustoPoolDataConnectionsClientListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/databases/{databaseName}/dataConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if kustoPoolName == "" {
		return nil, errors.New("parameter kustoPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kustoPoolName}", url.PathEscape(kustoPoolName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *KustoPoolDataConnectionsClient) listByDatabaseHandleResponse(resp *http.Response) (KustoPoolDataConnectionsClientListByDatabaseResponse, error) {
	result := KustoPoolDataConnectionsClientListByDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataConnectionListResult); err != nil {
		return KustoPoolDataConnectionsClientListByDatabaseResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a data connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - kustoPoolName - The name of the Kusto pool.
//   - databaseName - The name of the database in the Kusto pool.
//   - dataConnectionName - The name of the data connection.
//   - parameters - The data connection parameters supplied to the Update operation.
//   - options - KustoPoolDataConnectionsClientBeginUpdateOptions contains the optional parameters for the KustoPoolDataConnectionsClient.BeginUpdate
//     method.
func (client *KustoPoolDataConnectionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName string, parameters DataConnectionClassification, options *KustoPoolDataConnectionsClientBeginUpdateOptions) (*runtime.Poller[KustoPoolDataConnectionsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, workspaceName, kustoPoolName, databaseName, dataConnectionName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[KustoPoolDataConnectionsClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[KustoPoolDataConnectionsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates a data connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01-preview
func (client *KustoPoolDataConnectionsClient) update(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName string, parameters DataConnectionClassification, options *KustoPoolDataConnectionsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceName, kustoPoolName, databaseName, dataConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *KustoPoolDataConnectionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, dataConnectionName string, parameters DataConnectionClassification, options *KustoPoolDataConnectionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/databases/{databaseName}/dataConnections/{dataConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if kustoPoolName == "" {
		return nil, errors.New("parameter kustoPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kustoPoolName}", url.PathEscape(kustoPoolName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if dataConnectionName == "" {
		return nil, errors.New("parameter dataConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataConnectionName}", url.PathEscape(dataConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
