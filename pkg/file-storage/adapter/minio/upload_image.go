package minio

import (
	"bytes"
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

func (a *Adapter) UploadImage(ctx context.Context, imgPath string, file []byte, contentType string) error {
	fileReader := bytes.NewReader(file)

	options := minio.PutObjectOptions{
		ContentType: contentType,
	}

	fileSize := int64(len(file))
	_, err := a.MinioClient.PutObject(ctx, a.BucketName, imgPath, fileReader, fileSize, options)
	if err != nil {
		return fmt.Errorf("failed to upload image %q to bucket %q: %w", imgPath, a.BucketName, err)
	}

	return nil
}
