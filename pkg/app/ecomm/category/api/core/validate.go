package core

import (
	"errors"
	"strings"

	"github.com/amorindev/go-tmpl/pkg/shared/domain"
)

// ValidateCreate checks the fields required for creating a category
func validateCategoryFields(name string, desc *string) error {
	if strings.TrimSpace(name) == "" {
		return domain.NewAppError(domain.ErrCodeInvalidParams, "name is required")
	}
	if len(name) > 100 {
		return errors.New("name cannot exceed 100 characters")
	}
	if desc != nil && len(*desc) > 255 {
		return domain.NewAppError(domain.ErrCodeInvalidParams, "desc cannot exceed 255 characters")
	}
	return nil
}

