package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/category/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) FindAll(ctx context.Context) ([]*domain.Category, error) {
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error getting categories: %w", err)
	}
	defer cursor.Close(ctx)

	var categories []*domain.Category
	for cursor.Next(ctx) {
		var category domain.Category
		err := cursor.Decode(&category)
		if err != nil {
			return nil, fmt.Errorf("error decoding category: %w", err)
		}
		categories = append(categories, &category)
	}

	return categories, nil
}
