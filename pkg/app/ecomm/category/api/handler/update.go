package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/category/api/core"
	"github.com/amorindev/go-tmpl/pkg/app/ecomm/category/domain"
	cShared "github.com/amorindev/go-tmpl/pkg/shared/api/core"
	dShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
)

func (h Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		cShared.RespondError(w, dShared.NewAppError(dShared.ErrCodeInvalidParams, "missing category id"))
		return
	}

	var req core.UpdateCategoryReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		cShared.RespondError(w, dShared.NewAppError(dShared.ErrCodeInvalidParams, "invalid request body"))
		return 
	}

	defer r.Body.Close()

	if err := req.Validate(); err != nil {
		cShared.RespondError(w, err)
		return
	}

	category := &domain.Category{
		Name: req.Name,
		Desc: req.Desc,
	}

	if err := h.CategorySrv.Update(context.Background(), id, category); err != nil {
		cShared.RespondError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)
}
