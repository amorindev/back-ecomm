package core

import "github.com/amorindev/go-tmpl/pkg/app/ecomm/products/domain"

type ProductResp struct {
	Count    int64             `json:"count"`
	Pages    int64             `json:"pages"`
	Next     *string           `json:"next"`
	Previous *string           `json:"previous"`
	Products []*domain.Product `json:"products"`
}
