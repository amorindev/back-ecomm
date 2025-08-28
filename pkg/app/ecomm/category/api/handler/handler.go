package handler

import (
	"net/http"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/category/port"
)

type Handler struct {
	CategorySrv port.CategorySrv
}

func NewCategoryHandler(mux *http.ServeMux, categorySrv port.CategorySrv) *Handler {
	h := &Handler{
		CategorySrv: categorySrv,
	}

	mux.HandleFunc("POST /categories", h.Create)
	mux.HandleFunc("GET /categories", h.GetAll)
	mux.HandleFunc("PUT /categories/{id}", h.Update)
	mux.HandleFunc("PATCH /categories/{id}", h.Patch)
	mux.HandleFunc("DELETE /categories/{id}", h.Delete)

	return h
}
