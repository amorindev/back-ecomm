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

func (h Handler) CreateVarOption(w http.ResponseWriter, r *http.Request) {
	variationID := r.PathValue("variationId")

	if variationID == "" {
		cShared.RespondError(w, dShared.NewAppError(dShared.ErrCodeInvalidParams, "missing variation id"))
		return
	}

	var req core.CreateVarOptionReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		cShared.RespondError(w, dShared.NewAppError(dShared.ErrCodeInvalidParams, "invalid request body"))
		return
	}
	defer r.Body.Close()

	if err := req.Validate(); err != nil {
		cShared.RespondError(w, err)
		return
	}

	varOption := &domain.VarOption{
		Label:       req.Label,
		Value:       req.Value,
		VariationID: variationID,
	}

	if err := h.VariationSrv.CreateVarOption(context.Background(), varOption); err != nil {
		cShared.RespondError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(varOption)

}
