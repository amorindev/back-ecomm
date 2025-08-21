package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/variations/domain"
	dShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) Insert(ctx context.Context, variation *domain.Variation) error {
	id := bson.NewObjectID()
	variation.ID = id

	_, err := r.Collection.InsertOne(ctx, variation)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("%w: error inserting variation: %s", dShared.ErrDuplicateKey, err.Error())
		}
		return fmt.Errorf("error inserting variation: %s", err.Error())
	}

	variation.ID = id.Hex()

	return nil
}
