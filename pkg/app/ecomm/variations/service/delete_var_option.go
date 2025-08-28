package service

import (
	"context"

	"github.com/amorindev/go-tmpl/pkg/shared/domain"
)

func (s *Service) DeleteVarOption(ctx context.Context, id string, variationID string) error {
	err := s.VarOptionRepo.Delete(ctx, id, variationID)
	if err != nil {
		return domain.ManageError(err, "")
	}
	return nil
}
