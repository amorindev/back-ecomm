package domain

import "time"

// Variation represents a variation of a product (e.g. Color, Size).
// bson omitempty we don’t want to insert it, but it will be useful for building from the database
type Variation struct {
	ID        interface{}  `json:"id" bson:"_id" `
	Name      string       `json:"name" bson:"name"`
	CreatedAt *time.Time   `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time   `json:"updated_at" bson:"updated_at"`
	Options   []*VarOption `json:"options,omitempty" bson:"options,omitempty"`
}

// Option represents each option inside a variation
// e.g. "Black" "#000000"
// e.g. "XL" nil
type VarOption struct {
	ID          interface{} `json:"id" bson:"_id"`
	VariationID interface{} `json:"variation_id" bson:"variation_id"`
	Label       string      `json:"label" bson:"label"`
	Value       *string     `json:"value" bson:"value"`
	CreatedAt   *time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt   *time.Time  `json:"updated_at" bson:"updated_at"`
}
