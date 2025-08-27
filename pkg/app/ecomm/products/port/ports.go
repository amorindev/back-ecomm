package port

import (
	"context"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/products/domain"
)

type ProductRepo interface {
	Insert(ctx context.Context, product *domain.Product) error
	FindAll(ctx context.Context, limit int64, page int64) ([]*domain.Product, error)
	Count(ctx context.Context) (int64, error)
	Delete(ctx context.Context, id string) error
}

type ProductSrv interface {
	GetAll(ctx context.Context, limit int64, page int64) ([]*domain.Product, int64, int64, error)
	Create(ctx context.Context, product *domain.Product) error
	Delete(ctx context.Context, id string) error
}
