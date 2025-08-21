package service

import (
	"context"

	dShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
	"github.com/amorindev/go-tmpl/pkg/app/ecomm/variations/domain"
)

func (s *Service) GetAllVariationsWithOptions(ctx context.Context) ([]*domain.Variation, error) {
	variations, err := s.VariationRepo.FindAllWithOptions(ctx)
	if err != nil {
		return nil, dShared.ManageError(err, "")
	}
    return variations, nil
}
