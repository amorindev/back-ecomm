package service

import (
	authMethodP "github.com/amorindev/go-tmpl/pkg/app/auth-methods/port"
	"github.com/amorindev/go-tmpl/pkg/app/users/port"
)

var _ authMethodP.AuthMethodSrv = &Service{}

type Service struct{
    UserRepo port.UserRepo
}

func NewAuthMethodSrv(userRepo port.UserRepo) *Service{
    return &Service{
        UserRepo: userRepo,
    }
}