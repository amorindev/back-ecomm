package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/variations/domain"
	dShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) Insert(ctx context.Context, varOption *domain.VarOption) error {
	id := bson.NewObjectID()
	varOption.ID = id

	variationOID, err := bson.ObjectIDFromHex(varOption.VariationID.(string))
	if err != nil {
		return dShared.ErrIncorrectID
	}
	varOption.VariationID = variationOID

	_, err = r.Collection.InsertOne(ctx, varOption)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("%w: error inserting varOption: %s", dShared.ErrDuplicateKey, err.Error())
		}
		return fmt.Errorf("error inserting varOption: %s", err.Error())
	}

	varOption.ID = id.Hex()
	varOption.VariationID = variationOID.Hex()

	return nil
}
