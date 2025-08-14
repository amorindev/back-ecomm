package port

import (
	"context"

	"github.com/amorindev/go-tmpl/pkg/app/users/domain"
)

type UserRepo interface {
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	Insert(ctx context.Context, user *domain.User) error
}
