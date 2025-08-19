package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/category/api/core"
	"github.com/amorindev/go-tmpl/pkg/app/ecomm/category/domain"
	coreShared "github.com/amorindev/go-tmpl/pkg/shared/api/core"
	domainShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
)

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req core.CreateCategoryReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		coreShared.RespondError(w, domainShared.NewAppError(domainShared.ErrCodeInvalidParams, "invalid request body"))
		return
	}

	defer r.Body.Close()

	if err := req.Validate(); err != nil {
		coreShared.RespondError(w, err)
		return
	}

	category := &domain.Category{
		Name: req.Name,
		Desc: req.Desc,
	}

	if err := h.CategorySrv.Create(context.Background(), category); err != nil {
		coreShared.RespondError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}
