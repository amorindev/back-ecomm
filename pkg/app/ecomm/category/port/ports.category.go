package port

import (
	"context"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/category/domain"
)

type CategoryRepo interface {
	Insert(ctx context.Context, category *domain.Category) error
	Get(ctx context.Context, id string) (*domain.Category, error)
	FindAll(ctx context.Context) ([]*domain.Category, error)
	Update(ctx context.Context, id string, category *domain.Category) error
	Delete(ctx context.Context, id string) error
}

type CategorySrv interface {
	Create(ctx context.Context, category *domain.Category) error
	GetAll(ctx context.Context) ([]*domain.Category, error)
	Update(ctx context.Context, id string, category *domain.Category) error
	Delete(ctx context.Context, id string) error
	Patch(ctx context.Context, id string, category *domain.Category) (*domain.Category, error)
}
