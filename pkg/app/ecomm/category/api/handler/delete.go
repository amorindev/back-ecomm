package handler

import (
	"context"
	"net/http"

	cShared "github.com/amorindev/go-tmpl/pkg/shared/api/core"
	dShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
)

func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		cShared.RespondError(w, dShared.NewAppError(dShared.ErrCodeInvalidParams, "missing category id"))
		return
	}

	if err := h.CategorySrv.Delete(context.Background(), id); err != nil {
		cShared.RespondError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
