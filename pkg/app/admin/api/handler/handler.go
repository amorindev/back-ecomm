package handler

import "net/http"

type Handler struct {
	ApiBaseUrl string
}

func NewAdminHandler(mux *http.ServeMux, apiBaseUrl string) *Handler {
	h := &Handler{
		ApiBaseUrl: apiBaseUrl,
	}


	// Templates
	//mux.HandleFunc("/admin/home", h.HomePage)
	//mux.HandleFunc("/admin/other", h.OtherPage)
	mux.HandleFunc("/admin/categories", h.CategoriesPage)
	mux.HandleFunc("/admin/variations", h.VariationsPage)
	mux.HandleFunc("/admin/products", h.ProductsPage)

	return h
}
