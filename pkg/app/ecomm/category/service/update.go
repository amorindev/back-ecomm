package service

import (
	"context"
	"time"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/category/domain"
	dShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
)

func (h *Service) Update(ctx context.Context, id string, category *domain.Category) error {
	now := time.Now().UTC()
	category.UpdatedAt = &now

	err := h.CategoryRepo.Update(ctx, id, category)
	if err != nil {
		return dShared.ManageError(err, "")
	}
	return nil
}
