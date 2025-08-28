package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/variations/api/core"
	"github.com/amorindev/go-tmpl/pkg/app/ecomm/variations/domain"
	cShared "github.com/amorindev/go-tmpl/pkg/shared/api/core"
	dShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
)

func (h Handler) CreateVariation(w http.ResponseWriter, r *http.Request) {
	var req core.CreateVariationReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		cShared.RespondError(w, dShared.NewAppError(dShared.ErrCodeInvalidParams, "invalid request body"))
		return
	}

	defer r.Body.Close()

	if err := req.Validate(); err != nil {
		cShared.RespondError(w, err)
		return
	}

	variation := &domain.Variation{
		Name: req.Name,
	}

	if err := h.VariationSrv.CreateVariation(context.Background(), variation); err != nil {
		cShared.RespondError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(variation)
}
