package service

import (
	"context"
	"time"

	dShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
	"github.com/amorindev/go-tmpl/pkg/app/ecomm/variations/domain"
)

func (s *Service) UpdateVariation(ctx context.Context, variation *domain.Variation) error {
	now := time.Now().UTC()
	variation.UpdatedAt = &now

	err := s.VariationRepo.Update(ctx, variation)
	if err != nil {
		return dShared.ManageError(err, "")
	}
    return nil
}

