package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/category/domain"
	dShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (r *Repository) Update(ctx context.Context, id string, category *domain.Category) error {
	oID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return dShared.ErrIncorrectID
	}

	update := bson.M{
		"$set": bson.M{
			"name":       category.Name,
			"desc":       category.Desc,
			"updated_at": category.UpdatedAt,
		},
	}

	filter := bson.M{"_id": oID}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	
	err = r.Collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&category)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return dShared.ErrNotFound
		}
		return fmt.Errorf("failed to update category: %w", err)
	}

	return nil
}
