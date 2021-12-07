//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiotsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotsecurity/armiotsecurity"
)

// x-ms-original-file: specification/iotsecurity/resource-manager/Microsoft.IoTSecurity/preview/2021-07-01-preview/examples/Alerts/GetAlertList.json
func ExampleIotAlertsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotsecurity.NewIotAlertsClient("<subscription-id>",
		"<iot-defender-location>", cred, nil)
	pager := client.List("<device-group-name>",
		&armiotsecurity.IotAlertsListOptions{SkipToken: nil})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("AlertModel.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/iotsecurity/resource-manager/Microsoft.IoTSecurity/preview/2021-07-01-preview/examples/Alerts/GetAlert.json
func ExampleIotAlertsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotsecurity.NewIotAlertsClient("<subscription-id>",
		"<iot-defender-location>", cred, nil)
	res, err := client.Get(ctx,
		"<device-group-name>",
		"<alert-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("AlertModel.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/iotsecurity/resource-manager/Microsoft.IoTSecurity/preview/2021-07-01-preview/examples/Alerts/PatchAlert.json
func ExampleIotAlertsClient_Patch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotsecurity.NewIotAlertsClient("<subscription-id>",
		"<iot-defender-location>", cred, nil)
	res, err := client.Patch(ctx,
		"<device-group-name>",
		"<alert-id>",
		armiotsecurity.AlertPatchPropertiesModel{
			Properties: &armiotsecurity.AlertPatchPropertiesModelProperties{
				Severity: armiotsecurity.AlertSeverityMedium.ToPtr(),
				Status:   armiotsecurity.AlertStatusInProgress.ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("AlertModel.ID: %s\n", *res.ID)
}