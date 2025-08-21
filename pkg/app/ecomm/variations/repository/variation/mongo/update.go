package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/variations/domain"
	dShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (r *Repository) Update(ctx context.Context, variation *domain.Variation) error {
	oID, err := bson.ObjectIDFromHex(variation.ID.(string))
	if err != nil {
		return dShared.ErrIncorrectID
	}

	update := bson.M{
		"$set": bson.M{
			"name":       variation.Name,
			"updated_at": variation.UpdatedAt,
		},
	}

	filter := bson.M{
		"_id": oID,
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err = r.Collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(variation)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return dShared.ErrNotFound
		}
		return fmt.Errorf("failed to update variation: %w", err)
	}

	variation.ID = oID.Hex()

	return nil
}
