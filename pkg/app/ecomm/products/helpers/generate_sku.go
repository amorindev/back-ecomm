package helpers

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/products/domain"
)

func GenerateItemSKU(p *domain.Product, item *domain.ProductItem) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Category prefix (default "CAT", otherwise first 3 chars of CategoryName)
	categoryPart := "CAT"
	if p.CategoryName != "" {
		categoryPart = strings.ToUpper(p.CategoryName[:min(3, len(p.CategoryName))])
	}

	// Product name prefix (first 3 chars of product name, ignoring spaces)
	namePart := strings.ToUpper(strings.ReplaceAll(p.Name, " ", ""))[:min(3, len(p.Name))]

	// Variation options (e.g., color "RED", size "L"), take first 2 chars of each option
	var optionPart string
	if len(item.Options) > 0 {
		for _, opt := range item.Options {
			optionPart += strings.ToUpper(opt.VarOptName[:min(2, len(opt.VarOptName))])
		}
	}

	// Random number (4 digits)
	randomPart := fmt.Sprintf("%04d", r.Intn(10000))

	// Final SKU format: CATEGORY-NAME-OPTIONS-RANDOM
	return fmt.Sprintf("%s-%s-%s-%s", categoryPart, namePart, optionPart, randomPart)
}

// Helper function: returns the smaller of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
