package service

import (
	"context"
	"time"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/variations/domain"
	dShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
)

func (s *Service) CreateVariation(ctx context.Context, variation *domain.Variation) error {
	now := time.Now().UTC()
	variation.CreatedAt = &now

	err := s.VariationRepo.Insert(ctx, variation)
	if err != nil {
		return dShared.ManageError(err, "")
	}

	return nil
}
