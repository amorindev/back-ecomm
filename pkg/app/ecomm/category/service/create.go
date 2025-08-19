package service

import (
	"context"
	"time"

	dShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
	"github.com/amorindev/go-tmpl/pkg/app/ecomm/category/domain"
)

func (s *Service) Create(ctx context.Context, category *domain.Category) error {
    now := time.Now().UTC()
    category.CreateAt = &now

    err := s.CategoryRepo.Insert(ctx,category)
    if err != nil {
      return dShared.ManageError(err, "")
    }

    return nil
}
