package domain

// ProductStatus defines the possible states of a product
type ProductStatus string

const (
	ProductStatusActive   ProductStatus = "active"   // The product is available for sale
	ProductStatusInactive ProductStatus = "inactive" // The product is not available for customers
)
