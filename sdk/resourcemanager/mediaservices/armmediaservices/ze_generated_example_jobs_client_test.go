//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices"
)

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/jobs-list-all-filter-by-created.json
func ExampleJobsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewJobsClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<account-name>",
		"<transform-name>",
		&armmediaservices.JobsClientListOptions{Filter: to.StringPtr("<filter>"),
			Orderby: to.StringPtr("<orderby>"),
		})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/jobs-get-by-name.json
func ExampleJobsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewJobsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<transform-name>",
		"<job-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.JobsClientGetResult)
}

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/jobs-create.json
func ExampleJobsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewJobsClient("<subscription-id>", cred, nil)
	_, err = client.Create(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<transform-name>",
		"<job-name>",
		armmediaservices.Job{
			Properties: &armmediaservices.JobProperties{
				CorrelationData: map[string]*string{
					"Key 2": to.StringPtr("Value 2"),
					"key1":  to.StringPtr("value1"),
				},
				Input: &armmediaservices.JobInputAsset{
					ODataType: to.StringPtr("<odata-type>"),
					AssetName: to.StringPtr("<asset-name>"),
				},
				Outputs: []armmediaservices.JobOutputClassification{
					&armmediaservices.JobOutputAsset{
						ODataType: to.StringPtr("<odata-type>"),
						AssetName: to.StringPtr("<asset-name>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/jobs-delete.json
func ExampleJobsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewJobsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<transform-name>",
		"<job-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/jobs-update.json
func ExampleJobsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewJobsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<transform-name>",
		"<job-name>",
		armmediaservices.Job{
			Properties: &armmediaservices.JobProperties{
				Description: to.StringPtr("<description>"),
				Input: &armmediaservices.JobInputAsset{
					ODataType: to.StringPtr("<odata-type>"),
					AssetName: to.StringPtr("<asset-name>"),
				},
				Outputs: []armmediaservices.JobOutputClassification{
					&armmediaservices.JobOutputAsset{
						ODataType: to.StringPtr("<odata-type>"),
						AssetName: to.StringPtr("<asset-name>"),
					}},
				Priority: armmediaservices.Priority("High").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.JobsClientUpdateResult)
}

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/jobs-cancel.json
func ExampleJobsClient_CancelJob() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewJobsClient("<subscription-id>", cred, nil)
	_, err = client.CancelJob(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<transform-name>",
		"<job-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}