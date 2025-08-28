package service

import (
	"context"
	"math"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/products/domain"
	"github.com/amorindev/go-tmpl/pkg/app/ecomm/products/helpers"
	dShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
)

func (s *Service) GetAll(ctx context.Context, limit int64, page int64) ([]*domain.Product, int64, int64, error) {
	products, err := s.ProductRepo.FindAll(ctx, limit, page)
	if err != nil {
		return nil, 0, 0, dShared.ManageError(err, "")
	}

	for _, product := range products {
		productUrl, err := s.FileStorageSrv.GetImage(ctx, product.FilePath)
		if err != nil {
			return nil, 0, 0, dShared.ManageError(err, "")
		}
		product.ImgUrl = productUrl

		for _, pItem := range product.ProductItems {
			pItemUrl, err := s.FileStorageSrv.GetImage(ctx, pItem.FilePath)
			if err != nil {
				return nil, 0, 0, dShared.ManageError(err, "")

			}
			pItem.ImgUrl = pItemUrl
			pItem.FilePath = ""
			pItem.VarOptionIDs = nil
		}
		product.FilePath = ""
		helpers.CalculateVariations(product)
	}

	count, err := s.ProductRepo.Count(ctx)
	if err != nil {
		return nil, 0, 0, dShared.ManageError(err, "")
	}
	totalPages := int64(math.Ceil(float64(count) / float64(limit)))

	return products, count, totalPages, nil
}
