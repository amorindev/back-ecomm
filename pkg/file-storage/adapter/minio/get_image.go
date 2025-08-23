package minio

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

func (a *Adapter) GetImage(ctx context.Context, imgPath string) (string, error) {
	// * Set request parameters for content-disposition.
	reqParams := make(url.Values)

	// * Generates a presigned url which expires in a day.
	expiration := time.Hour * 24 * 7
	presignedURL, err := a.MinioClient.PresignedGetObject(ctx, a.BucketName, imgPath, expiration, reqParams)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL for image %q in bucket %q: %w", imgPath, a.BucketName, err)
	}

	return presignedURL.String(), nil
}
