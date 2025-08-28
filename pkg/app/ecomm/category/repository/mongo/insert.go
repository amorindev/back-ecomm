package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/category/domain"
	domainShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) Insert(ctx context.Context, category *domain.Category) error {
	id := bson.NewObjectID()
	category.ID = id

	_, err := r.Collection.InsertOne(ctx, category)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("%w: error inserting category: %s", domainShared.ErrDuplicateKey, err.Error())
		}
		return fmt.Errorf("error inserting category: %s", err.Error())
	}
	category.ID = id.Hex()
	return nil
}
