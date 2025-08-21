package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/variations/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) FindAllWithOptions(ctx context.Context) ([]*domain.Variation, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "var_options"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "variation_id"},
			{Key: "as", Value: "options"},
		}}},
	}

	cursor, err := r.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error getting variations with options: %w", err)
	}
	defer cursor.Close(ctx)

	var variations []*domain.Variation
	for cursor.Next(ctx) {
		var variation domain.Variation
		err := cursor.Decode(&variation)
		if err != nil {
			return nil, fmt.Errorf("error decoding variation: %w", err)
		}
		variations = append(variations, &variation)
	}

	return variations, nil
}
