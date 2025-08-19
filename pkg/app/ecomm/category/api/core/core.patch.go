package core

import (
	"strings"

	"github.com/amorindev/go-tmpl/pkg/shared/domain"
)

// Note: We use pointers for both fields to distinguish between "not provided" and "explicitly empty".
//   - If the pointer is nil → the field was not sent, so we don't update it.
//   - If the pointer is non-nil but empty (e.g. ""), it means the client wants to clear the value.
//
// This allows partial updates without accidentally overwriting fields with zero values.
type PatchCategoryReq struct {
	Name *string `json:"name"`
	Desc *string `json:"desc"`
}

func (req PatchCategoryReq) Validate() error {
	if req.Name != nil {
		name := strings.TrimSpace(*req.Name)
		if name == "" {
			return domain.NewAppError(domain.ErrCodeInvalidParams, "name is required")
		}

		if len(name) > 100 {
			return domain.NewAppError(domain.ErrCodeInvalidParams, "name cannot exceed 100 characters")
		}
	}

	if req.Desc != nil {
		desc := strings.TrimSpace(*req.Desc)
		if len(desc) > 255 {
			return domain.NewAppError(domain.ErrCodeInvalidParams, "desc cannot exceed 255 characters")
		}
	}
	return nil
}
