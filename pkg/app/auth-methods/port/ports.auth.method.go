package port

import (
	"context"

	"github.com/amorindev/go-tmpl/pkg/app/users/domain"
)

type AuthMethodSrv interface {
	SignUp(ctx context.Context, user *domain.User) error
}