package service

import (
	"context"
	"time"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/category/domain"
	dShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
)

func (s *Service) Patch(ctx context.Context, id string, category *domain.Category) (*domain.Category, error) {
	existing, err := s.CategoryRepo.Get(ctx, id)
	if err != nil {
		return nil, dShared.ManageError(err, "")
	}

	// We only update the submitted fields
	if category.Name != "" {
		existing.Name = category.Name
	}

	if category.Desc != nil {
		existing.Desc = category.Desc
	}

	now := time.Now().UTC()
	existing.UpdatedAt = &now

	err = s.CategoryRepo.Update(ctx, id, existing)
	if err != nil {
		return nil, dShared.ManageError(err, "")
	}

	return existing, nil
}
