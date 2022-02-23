package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func NewContainerClient(accountName string, containerName string) (azblob.ContainerClient, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return azblob.ContainerClient{}, fmt.Errorf("invalid credentials: %w", err)
	}

	serviceClient, err := azblob.NewServiceClient(fmt.Sprintf("https://%s.blob.core.windows.net/", accountName), cred, nil)
	if err != nil {
		return azblob.ContainerClient{}, fmt.Errorf("invalid credentials: %w", err)
	}

	return serviceClient.NewContainerClient(containerName), nil
}

func upload(accountName string, containerName string, blobName string, data []byte, metadata map[string]string) error {
	containerClient, err := NewContainerClient(accountName, containerName)
	if err != nil {
		return err
	}

	blob := containerClient.NewBlockBlobClient(blobName)

	options := azblob.HighLevelUploadToBlockBlobOption{
		Metadata: metadata,
	}

	_, err = blob.UploadBufferToBlockBlob(context.Background(), data, options)
	if err != nil {
		return err
	}

	return nil
}

func waitForBlobToBeRemoved(accountName string, containerName string, blobToFind string) error {
	containerClient, err := NewContainerClient(accountName, containerName)
	if err != nil {
		return err
	}

	ctx := context.Background()
	retries := 500
	for retries > 0 {
		pager := containerClient.ListBlobsFlat(&azblob.ContainerListBlobFlatSegmentOptions{})
		for pager.NextPage(ctx) {

			resp := pager.PageResponse()

			blobExist := false
			for _, blob := range resp.ContainerListBlobFlatSegmentResult.Segment.BlobItems {
				if *blob.Name == blobToFind {
					blobExist = true
					break
				}
			}

			if blobExist {
				time.Sleep(500 * time.Millisecond)
				retries--
				break
			} else {
				return nil
			}
		}
	}

	return fmt.Errorf("max attempts reached: blob '%s/%s/%s' is still present", accountName, containerName, blobToFind)
}
