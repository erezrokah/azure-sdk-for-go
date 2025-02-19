//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2023-01-01/examples/ListPrivateLinkResources.json
func ExamplePrivateLinkResourcesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservices.NewPrivateLinkResourcesClient("6c48fa17-39c7-45f1-90ac-47a587128ace", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("petesting", "pemsi-ecy-rsv2", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PrivateLinkResources = armrecoveryservices.PrivateLinkResources{
		// 	Value: []*armrecoveryservices.PrivateLinkResource{
		// 		{
		// 			Name: to.Ptr("backupResource"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/Vaults/privateLinkResources"),
		// 			ID: to.Ptr("/subscriptions/6c48fa17-39c7-45f1-90ac-47a587128ace/resourceGroups/petesting/providers/Microsoft.RecoveryServices/Vaults/pemsi-ecy-rsv2/privateLinkResources/backupResource"),
		// 			Properties: &armrecoveryservices.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("AzureBackup"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("backup-fab1"),
		// 					to.Ptr("backup-rec2"),
		// 					to.Ptr("backup-prot1"),
		// 					to.Ptr("backup-ecs1"),
		// 					to.Ptr("backup-tel1"),
		// 					to.Ptr("backup-wbcm1"),
		// 					to.Ptr("backup-fc1"),
		// 					to.Ptr("backup-id1")},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.ecy.backup.windowsazure.com"),
		// 						to.Ptr("privatelink.queue.core.windows.net"),
		// 						to.Ptr("privatelink.blob.core.windows.net")},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("siteRecoveryResource"),
		// 					Type: to.Ptr("Microsoft.RecoveryServices/vaults/privateLinkResources"),
		// 					ID: to.Ptr("/subscriptions/6c48fa17-39c7-45f1-90ac-47a587128ace/resourceGroups/gaallarg/providers/Microsoft.RecoveryServices/vaults/amchandnTest2702A/privateLinkResources/siteRecoveryResource"),
		// 					Properties: &armrecoveryservices.PrivateLinkResourceProperties{
		// 						GroupID: to.Ptr("AzureSiteRecovery"),
		// 						RequiredMembers: []*string{
		// 							to.Ptr("siteRecovery-rcm1"),
		// 							to.Ptr("siteRecovery-prot2"),
		// 							to.Ptr("siteRecovery-tel1"),
		// 							to.Ptr("siteRecovery-srs1"),
		// 							to.Ptr("siteRecovery-prot2b"),
		// 							to.Ptr("siteRecovery-id1")},
		// 							RequiredZoneNames: []*string{
		// 								to.Ptr("privatelink.ecy.siterecovery.windowsazure.com"),
		// 								to.Ptr("privatelink.queue.core.windows.net"),
		// 								to.Ptr("privatelink.blob.core.windows.net")},
		// 							},
		// 					}},
		// 				}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2023-01-01/examples/GetPrivateLinkResources.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservices.NewPrivateLinkResourcesClient("6c48fa17-39c7-45f1-90ac-47a587128ace", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "petesting", "pemsi-ecy-rsv2", "backupResource", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResource = armrecoveryservices.PrivateLinkResource{
	// 	Name: to.Ptr("backupResource"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/Vaults/privateLinkResources"),
	// 	ID: to.Ptr("/subscriptions/6c48fa17-39c7-45f1-90ac-47a587128ace/resourceGroups/petesting/providers/Microsoft.RecoveryServices/Vaults/pemsi-ecy-rsv2/privateLinkResources/backupResource"),
	// 	Properties: &armrecoveryservices.PrivateLinkResourceProperties{
	// 		GroupID: to.Ptr("AzureBackup"),
	// 		RequiredMembers: []*string{
	// 			to.Ptr("backup-fab1"),
	// 			to.Ptr("backup-rec2"),
	// 			to.Ptr("backup-prot1"),
	// 			to.Ptr("backup-ecs1"),
	// 			to.Ptr("backup-tel1"),
	// 			to.Ptr("backup-wbcm1"),
	// 			to.Ptr("backup-fc1"),
	// 			to.Ptr("backup-id1")},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("privatelink.ecy.backup.windowsazure.com"),
	// 				to.Ptr("privatelink.queue.core.windows.net"),
	// 				to.Ptr("privatelink.blob.core.windows.net")},
	// 			},
	// 		}
}
