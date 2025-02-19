//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/NameSpaces/PrivateLinkResourcesGet.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewPrivateLinkResourcesClient("subID", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "ArunMonocle", "sdk-Namespace-2924", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResourcesListResult = armservicebus.PrivateLinkResourcesListResult{
	// 	Value: []*armservicebus.PrivateLinkResource{
	// 		{
	// 			Name: to.Ptr("namespace"),
	// 			Type: to.Ptr("Microsoft.ServiceBus/namespaces/privateLinkResources"),
	// 			ID: to.Ptr("subscriptions/subID/resourceGroups/SDK-ServiceBus-4794/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-5828/privateLinkResources/namespace"),
	// 			Properties: &armservicebus.PrivateLinkResourceProperties{
	// 				GroupID: to.Ptr("namespace"),
	// 				RequiredMembers: []*string{
	// 					to.Ptr("namespace")},
	// 					RequiredZoneNames: []*string{
	// 						to.Ptr("privatelink.servicebus.windows.net")},
	// 					},
	// 			}},
	// 		}
}
