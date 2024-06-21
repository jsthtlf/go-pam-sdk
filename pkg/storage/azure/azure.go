package azure

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

type ReplayStorage struct {
	AccountName    string
	AccountKey     string
	ContainerName  string
	EndpointSuffix string
}

func NewReplayStorage(accountName, accountKey, containerName, endpointSuffix string) ReplayStorage {
	return ReplayStorage{
		AccountName:    accountName,
		AccountKey:     accountKey,
		ContainerName:  containerName,
		EndpointSuffix: endpointSuffix,
	}
}

func (a ReplayStorage) Upload(gZipFilePath, target string) error {
	file, err := os.Open(gZipFilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	credential, err := azblob.NewSharedKeyCredential(a.AccountName, a.AccountKey)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf("https://%s.blob.%s/", a.AccountName, a.EndpointSuffix)
	client, err := azblob.NewClientWithSharedKeyCredential(endpoint, credential, nil)
	if err != nil {
		return err
	}

	// TODO Проверить: создание контейнера | контейнер создан
	_, err = client.CreateContainer(context.TODO(), a.ContainerName, nil)
	if err != nil {
		// TODO обработать, если контейнер уже создан
		return err
	}

	// TODO Проверить: загрузку файла | загрузку файлов
	_, err = client.UploadStream(context.TODO(),
		a.ContainerName,
		target,
		file,
		&azblob.UploadStreamOptions{
			BlockSize:   4 * 1024 * 1024,
			Concurrency: 16,
		})
	if err != nil {
		return err
	}

	return nil
}

func (a ReplayStorage) TypeName() string {
	return "azure"
}
