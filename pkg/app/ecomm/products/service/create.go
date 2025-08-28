package service

import (
	"context"
	"fmt"
	"time"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/products/domain"
	"github.com/amorindev/go-tmpl/pkg/app/ecomm/products/helpers"
	dShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
)

func (s *Service) Create(ctx context.Context, product *domain.Product) error {
	now := time.Now().UTC()
	product.CreatedAt = &now
	product.UpdatedAt = &now
	product.Status = domain.ProductStatusActive

	bucketFolderStruct := "products/"
	product.FilePath = fmt.Sprintf("%s%s", bucketFolderStruct, product.FilePath)

	uniqueFilePath, err := s.FileStorageSrv.UploadImage(ctx, product.FilePath, product.File, product.ContentType)
	if err != nil {
		return dShared.ManageError(err, "")
	}
	product.FilePath = uniqueFilePath

	// to create the sku
	ctg, err := s.CategoryRepo.Get(ctx, product.CategoryID.(string))
	if err != nil {
		return dShared.ManageError(err, "")
	}
	product.CategoryName = ctg.Name

	for _, pItem := range product.ProductItems {
		pItem.CreatedAt = &now
		pItem.UpdatedAt = &now
		pItem.FilePath = fmt.Sprintf("%s%s", bucketFolderStruct, pItem.FilePath)

		// Create Sku
		ids := make([]string, 0, len(pItem.VarOptionIDs))
		for _, id := range pItem.VarOptionIDs {
			ids = append(ids, id.(string))
		}

		varOptions, err := s.VarOptionRepo.FindByIDs(ctx, ids)
		if err != nil {
			return dShared.ManageError(err, "")
		}

		options := make([]*domain.Option, 0, len(pItem.VarOptionIDs))
		for _, varOpt := range varOptions {
			option := &domain.Option{
				VarOptName: varOpt.Label,
			}
			options = append(options, option)
		}

		pItem.Options = options

		pItem.Sku = helpers.GenerateItemSKU(product, pItem)
		// we clean so as not to save in the database
		pItem.Options = nil
		product.CategoryName = ""

		uniqueFP, err := s.FileStorageSrv.UploadImage(ctx, pItem.FilePath, pItem.File, pItem.ContentType)
		if err != nil {
			return dShared.ManageError(err, "")
		}
		pItem.FilePath = uniqueFP
	}

	err = s.ProductRepo.Insert(ctx, product)
	if err != nil {
		return dShared.ManageError(err, "")
	}

	return nil
}
