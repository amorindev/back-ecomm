package handler

import (
	"context"
	"net/http"

	cShared "github.com/amorindev/go-tmpl/pkg/shared/api/core"
	dShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
)

func (h Handler) DeleteVariation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		cShared.RespondError(w, dShared.NewAppError(dShared.ErrCodeInvalidParams, "missing variation id"))
		return
	}

	if err := h.VariationSrv.DeleteVariation(context.Background(), id); err != nil {
		cShared.RespondError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
