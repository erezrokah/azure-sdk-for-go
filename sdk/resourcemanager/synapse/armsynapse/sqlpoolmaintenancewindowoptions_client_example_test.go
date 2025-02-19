//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetMaintenanceWindowOptions.json
func ExampleSQLPoolMaintenanceWindowOptionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewSQLPoolMaintenanceWindowOptionsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "samplerg", "testworkspace", "testsp", "current", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MaintenanceWindowOptions = armsynapse.MaintenanceWindowOptions{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools/maintenanceWindowOptions"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/samplerg/providers/Microsoft.Synapse/workspaces/testworkspace/sqlPools/testsp/maintenanceWindowOptions/current"),
	// 	Properties: &armsynapse.MaintenanceWindowOptionsProperties{
	// 		AllowMultipleMaintenanceWindowsPerCycle: to.Ptr(true),
	// 		DefaultDurationInMinutes: to.Ptr[int32](120),
	// 		IsEnabled: to.Ptr(true),
	// 		MaintenanceWindowCycles: []*armsynapse.MaintenanceWindowTimeRange{
	// 			{
	// 				DayOfWeek: to.Ptr(armsynapse.DayOfWeekSaturday),
	// 				Duration: to.Ptr("PT60M"),
	// 				StartTime: to.Ptr("00:00:00"),
	// 		}},
	// 		MinCycles: to.Ptr[int32](2),
	// 		MinDurationInMinutes: to.Ptr[int32](60),
	// 		TimeGranularityInMinutes: to.Ptr[int32](5),
	// 	},
	// }
}
