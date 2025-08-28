package core

import (
	"strings"

	"github.com/amorindev/go-tmpl/pkg/shared/domain"
)

type CreateProductReq struct {
	CategoryID string  `json:"category_id,omitempty"`
	Name       string  `json:"name"`
	Desc       *string `json:"desc"`
	FilePath   string  `json:"file_path"`
	ProductItems []*CreateProductItem `json:"product_items"`
}

type CreateProductItem struct {
	QtyInStock   int           `json:"qty_in_stock"`
	Price        float64       `json:"price"`
	FilePath     string        `json:"file_path"`
	ContentType  string        `json:"content_type"`
	VarOptionIDs []interface{} `json:"var_option_ids,omitempty"`
}

func (c CreateProductReq) Validate() error {
	if strings.TrimSpace(c.CategoryID) == "" {
		return domain.NewAppError(domain.ErrCodeInvalidParams, "category_id is required")
	}
	if strings.TrimSpace(c.Name) == "" {
		return domain.NewAppError(domain.ErrCodeInvalidParams, "name is required")
	}
	if c.Desc != nil && strings.TrimSpace(*c.Desc) == "" {
		return domain.NewAppError(domain.ErrCodeInvalidParams, "desc is required")
	}
	if strings.TrimSpace(c.FilePath) == "" {
		return domain.NewAppError(domain.ErrCodeInvalidParams, "file_path is required")
	}

	// Invalid product (no items at all)
	if len(c.ProductItems) == 0 {
		return domain.NewAppError(domain.ErrCodeInvalidParams, "at least one product item is required")
	}

	// Product with variants (must have at least 2 items)
	if len(c.ProductItems) > 1 {
		for _, item := range c.ProductItems {
			if len(item.VarOptionIDs) == 0 {
				return domain.NewAppError(domain.ErrCodeInvalidParams, "variant options are required for product items")
			}
			if strings.TrimSpace(item.FilePath) == "" {
				return domain.NewAppError(domain.ErrCodeInvalidParams, "file_path is required for each product item")
			}
			if item.Price <= 0 {
				return domain.NewAppError(domain.ErrCodeInvalidParams, "price must be greater than zero")
			}
			if item.QtyInStock < 0 {
				return domain.NewAppError(domain.ErrCodeInvalidParams, "qty_in_stock cannot be negative")
			}
		}
	}

	// Product without variants (only 1 item, no variant options)
	// Note: FilePath is not validated here, because it is assigned in the handler from main_image
	if len(c.ProductItems) == 1 {
		item := c.ProductItems[0]
		if len(item.VarOptionIDs) > 0 {
			return domain.NewAppError(domain.ErrCodeInvalidParams, "single product must not contain variant options")
		}
		if item.Price <= 0 {
			return domain.NewAppError(domain.ErrCodeInvalidParams, "price must be greater than zero")
		}
		if item.QtyInStock < 0 {
			return domain.NewAppError(domain.ErrCodeInvalidParams, "qty_in_stock cannot be negative")
		}
	}

	return nil
}
