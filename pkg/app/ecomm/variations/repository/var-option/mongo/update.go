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

func (r *Repository) Update(ctx context.Context, varOption *domain.VarOption) error {
	oID, err := bson.ObjectIDFromHex(varOption.ID.(string))
	if err != nil {
		return dShared.ErrIncorrectID
	}

	variationOID, err := bson.ObjectIDFromHex(varOption.VariationID.(string))
	if err != nil {
		return dShared.ErrIncorrectID
	}

	update := bson.M{
		"$set": bson.M{
			"label":      varOption.Label,
			"value":      varOption.Value,
			"updated_at": varOption.UpdatedAt,
		},
	}

	filter := bson.M{
		"_id":          oID,
		"variation_id": variationOID,
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err = r.Collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&varOption)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return dShared.ErrNotFound
		}
		return fmt.Errorf("failed to update varOption: %w", err)
	}
	varOption.ID = oID.Hex()
	varOption.VariationID = variationOID.Hex()

	return nil
}
