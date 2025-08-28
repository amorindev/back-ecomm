package service

import (
	"context"

	dShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
	"github.com/amorindev/go-tmpl/pkg/app/ecomm/category/domain"
)

func (s *Service) GetAll(ctx context.Context) ([]*domain.Category, error) {
	ctgs, err := s.CategoryRepo.FindAll(ctx)
	if err != nil {
		return nil, dShared.ManageError(err, "")
	}
	return ctgs, nil
}
