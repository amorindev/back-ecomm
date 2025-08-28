package service

import "github.com/amorindev/go-tmpl/pkg/app/ecomm/category/port"

var _ port.CategorySrv = &Service{}

type Service struct{
    CategoryRepo port.CategoryRepo
}

func NewCategorySrv(categoryRepo port.CategoryRepo) *Service{
    return &Service{
        CategoryRepo: categoryRepo,
    }
}