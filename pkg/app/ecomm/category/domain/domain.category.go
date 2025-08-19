package domain

import "time"

type Category struct {
	ID        interface{} `json:"id" bson:"_id"`
	Name      string      `json:"name" bson:"name"`
	Desc      *string     `json:"desc" bson:"desc,omitempty"`
	CreateAt  *time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time  `json:"updated_at" bson:"updated_at"`
}

func FromCore(c Category) *Category {
	return &Category{
		ID:        c.ID,
		Desc:      c.Desc,
		Name:      c.Name,
		CreateAt:  c.CreateAt,
		UpdatedAt: c.UpdatedAt,
	}
}
