package service

import (
	categoryP "github.com/amorindev/go-tmpl/pkg/app/ecomm/category/port"
	productP "github.com/amorindev/go-tmpl/pkg/app/ecomm/products/port"
	varOptP "github.com/amorindev/go-tmpl/pkg/app/ecomm/variations/port"
	fileStgP "github.com/amorindev/go-tmpl/pkg/file-storage/port"
)

var _ productP.ProductSrv = &Service{}

type Service struct {
	ProductRepo    productP.ProductRepo
	CategoryRepo   categoryP.CategoryRepo
	FileStorageSrv fileStgP.FileStorageSrv
	VarOptionRepo  varOptP.VarOptionRepo
}

func NewProductSrv(
	productRepo productP.ProductRepo,
	varOptionRepo varOptP.VarOptionRepo,
	categoryRepo categoryP.CategoryRepo,
	fileStorageSrv fileStgP.FileStorageSrv,
) *Service {
	return &Service{
		ProductRepo:    productRepo,
		VarOptionRepo:  varOptionRepo,
		CategoryRepo:   categoryRepo,
		FileStorageSrv: fileStorageSrv,
	}
}
