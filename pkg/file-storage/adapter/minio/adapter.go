package minio

import (
	"github.com/amorindev/go-tmpl/pkg/file-storage/port"
	"github.com/minio/minio-go/v7"
)

var _ port.FileStorageAdt = &Adapter{}

type Adapter struct {
	MinioClient *minio.Client
	BucketName  string
}

func NewMinioAdt(client *minio.Client, bucketName string) *Adapter {
	return &Adapter{
		MinioClient: client,
		BucketName: bucketName,
	}
}
