package domain

import "time"

type ProductItem struct {
	ID           interface{}   `json:"id" bson:"_id"`
	Sku          string        `json:"sku" bson:"sku"`
	QtyInStock   int           `json:"qty_in_stock" bson:"qty_in_stock"`
	Price        float64       `json:"price" bson:"price"`
	FilePath     string        `json:"-" bson:"file_path"`
	File         []byte        `json:"-" bson:"-"`
	ContentType  string        `json:"-" bson:"-"`
	ImgUrl       string        `json:"img_url" bson:"-"`
	CreatedAt    *time.Time    `json:"created_at" bson:"created_at"`
	UpdatedAt    *time.Time    `json:"updated_at" bson:"updated_at"`
	VarOptionIDs []interface{} `json:"var_option_ids,omitempty" bson:"var_option_ids"`
	Options      []*Option     `json:"options" bson:"options,omitempty"`
}

// Option represents a specific variation choice for a product item
// For the response, the data is built directly from the DB, that's why it includes `bson` tags.
// Name: Color, VarOptName Blue, VarOptValue: #0000FF
// Name: Size, VarOptName S, VarOptValue: nil
type Option struct {
	Name        string `json:"name" bson:"name"`
	VarOptName  string `json:"var_opt_name" bson:"var_opt_name"`
	VarOptValue string `json:"var_opt_value" bson:"var_opt_value"`
}
