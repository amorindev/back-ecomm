package handler

import (
	"encoding/json"
	"net/http"

	"github.com/amorindev/go-tmpl/pkg/shared/api/core"
)

func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.CategorySrv.GetAll(r.Context())
	if err != nil {
		core.RespondError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}
