package middlewares

import "github.com/amorindev/go-tmpl/internal/auth"

type AuthMiddleware struct {
	AuthSrv *auth.TokenSrv
}

func NewAuthMdw(authSrv *auth.TokenSrv) *AuthMiddleware {
	return &AuthMiddleware{
		AuthSrv: authSrv,
	}
}
