package service

import "github.com/amorindev/go-tmpl/pkg/app/ecomm/variations/port"

var _ port.VariationSrv = &Service{}

type Service struct {
	VariationRepo port.VariationRepo
    VarOptionRepo port.VarOptionRepo
}

func NewVariationSrv(variationRepo port.VariationRepo, varOptionRepo port.VarOptionRepo) *Service {
	return &Service{
        VariationRepo: variationRepo,
        VarOptionRepo: varOptionRepo,
	}
}