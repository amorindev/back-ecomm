package port

import "context"

type FileStorageAdt interface {
	UploadImage(ctx context.Context, imgPath string, file []byte, contentType string) error
	GetImage(ctx context.Context, imgPath string) (string, error)
}

type FileStorageSrv interface {
	UploadImage(ctx context.Context, imgPath string, file []byte, contentType string) (string, error)
	GetImage(ctx context.Context, imgPath string) (string, error)
}
