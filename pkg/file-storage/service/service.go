package service

import "github.com/amorindev/go-tmpl/pkg/file-storage/port"

var _ port.FileStorageSrv = &Service{}

type Service struct {
	FileStgAdt port.FileStorageAdt
}

func NewFileStgSrv(fileStgAdt port.FileStorageAdt) *Service {
	return &Service{
        FileStgAdt: fileStgAdt,
	}
}