package service

import (
	"context"
	"time"

	dShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
	"github.com/amorindev/go-tmpl/pkg/app/ecomm/variations/domain"
)

func (s *Service) CreateVarOption(ctx context.Context, varOption *domain.VarOption) error {
	now := time.Now().UTC()
	varOption.CreatedAt = &now

	err := s.VarOptionRepo.Insert(ctx,varOption)
    if err != nil {
      return dShared.ManageError(err, "")
    }

	return nil
}

