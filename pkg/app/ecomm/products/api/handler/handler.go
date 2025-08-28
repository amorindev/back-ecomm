package handler

import (
	"net/http"

	productP "github.com/amorindev/go-tmpl/pkg/app/ecomm/products/port"
)

type Handler struct {
	ProductSrv productP.ProductSrv
}

func NewProductHandler(server *http.ServeMux, productSrv productP.ProductSrv) *Handler {
	h := &Handler{
		ProductSrv: productSrv,
	}

	server.HandleFunc("GET /products", h.GetAll)
	server.HandleFunc("POST /products", h.Create)
	server.HandleFunc("DELETE /products/{id}", h.Delete)

	return h
}
