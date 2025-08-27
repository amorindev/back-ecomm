package helpers

import (
	"sort"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/products/domain"
)

// CalculateVariations builds and assigns the list of unique product variations
// (e.g., Color, Size, Material) based on the options defined across all ProductItems.
//
// Process:
//  1. Iterate over each ProductItem and its options.
//  2. Group option values by variation name using a map (acting as a set to avoid duplicates).
//  3. Convert each set into a slice and sort values alphabetically for consistency.
//  4. Assign the resulting list to p.Variations.
//
// Example:
//  - ProductItems with options:
//      Color=Red, Size=M
//      Color=Blue, Size=L
//  - Result in p.Variations:
//      [
//        {Name: "Color", Values: ["Blue", "Red"]},
//        {Name: "Size", Values: ["L", "M"]}
//      ]
func CalculateVariations(p *domain.Product) {
	variationMap := make(map[string]map[string]struct{})

	for _, item := range p.ProductItems {
		for _, opt := range item.Options {
			if _, exists := variationMap[opt.Name]; !exists {
				variationMap[opt.Name] = make(map[string]struct{})
			}
			variationMap[opt.Name][opt.VarOptName] = struct{}{}
		}
	}

	var variations []*domain.Variation
	for name, valuesSet := range variationMap {
		var values []string
		for val := range valuesSet {
			values = append(values, val)
		}
		sort.Strings(values)
		variations = append(variations, &domain.Variation{
			Name:   name,
			Values: values,
		})
	}

	p.Variations = variations
}
