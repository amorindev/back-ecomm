package service

import (
	"context"

	"github.com/amorindev/go-tmpl/pkg/shared/domain"
)

func (s *Service) DeleteVariation(ctx context.Context, id string) error {
	err := s.VariationRepo.Delete(ctx, id)
	if err != nil {
		return domain.ManageError(err, "")
	}
	return nil
}
