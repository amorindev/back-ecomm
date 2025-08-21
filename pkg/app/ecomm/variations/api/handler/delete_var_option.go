package handler

import (
	"context"
	"net/http"

	cShared "github.com/amorindev/go-tmpl/pkg/shared/api/core"
	dShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
)

func (h Handler) DeleteVarOption(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	variationID := r.PathValue("variationId")

	if id == "" {
		cShared.RespondError(w, dShared.NewAppError(dShared.ErrCodeInvalidParams, "missing varOption id"))
		return
	}

	if variationID == "" {
		cShared.RespondError(w, dShared.NewAppError(dShared.ErrCodeInvalidParams, "missing variationId id"))
		return
	}

	if err := h.VariationSrv.DeleteVarOption(context.Background(), id, variationID); err != nil {
		cShared.RespondError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
