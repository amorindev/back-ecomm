package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/category/domain"
	domainShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) Get(ctx context.Context, id string) (*domain.Category, error) {
	oID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, domainShared.ErrIncorrectID
	}

	var ctg domain.Category
	err = r.Collection.FindOne(ctx, bson.M{"_id": oID}).Decode(&ctg)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domainShared.ErrNotFound
		}
		return nil, fmt.Errorf("error getting category: %w", err)
	}
	ctg.ID = oID.Hex()

	return &ctg, nil
}
