package storage

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

type AzureReplayStorage struct {
	AccountName    string
	AccountKey     string
	ContainerName  string
	EndpointSuffix string
}

func NewAzureReplayStorage(accountName, accountKey, containerName, endpointSuffix string) AzureReplayStorage {
	return AzureReplayStorage{
		AccountName:    accountName,
		AccountKey:     accountKey,
		ContainerName:  containerName,
		EndpointSuffix: endpointSuffix,
	}
}

func (a AzureReplayStorage) Upload(gZipFilePath, target string) error {
	file, err := os.Open(gZipFilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	credential, err := azblob.NewSharedKeyCredential(a.AccountName, a.AccountKey)
	if err != nil {
		return err
	}
	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})
	endpoint := fmt.Sprintf("https://%s.blob.%s/%s", a.AccountName, a.EndpointSuffix, a.ContainerName)
	URL, _ := url.Parse(endpoint)
	containerURL := azblob.NewContainerURL(*URL, p)
	blobURL := containerURL.NewBlockBlobURL(target)

	commonResp, err := azblob.UploadFileToBlockBlob(context.TODO(), file, blobURL, azblob.UploadToBlockBlobOptions{
		BlockSize:   4 * 1024 * 1024,
		Parallelism: 16})
	if err != nil {
		return err
	}
	if httpResp := commonResp.Response(); httpResp != nil {
		defer httpResp.Body.Close()
	}
	return err
}

func (a AzureReplayStorage) TypeName() string {
	return TypeAzure
}
