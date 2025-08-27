package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/products/domain"
	dShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) Insert(ctx context.Context, product *domain.Product) error {
	id := bson.NewObjectID()
	product.ID = id

	for _, pItem := range product.ProductItems {
		pItem.ID = bson.NewObjectID()
		var oIDs []interface{}
		for _, varOptID := range pItem.VarOptionIDs {
			oID, err := bson.ObjectIDFromHex(varOptID.(string))
			if err != nil {
				return dShared.ErrIncorrectID
			}
			oIDs = append(oIDs, oID)
		}
		pItem.VarOptionIDs = oIDs
	}

	// Assign ID category
	ctgOID, err := bson.ObjectIDFromHex(product.CategoryID.(string))
	if err != nil {
		return dShared.ErrIncorrectID
	}
	product.CategoryID = ctgOID

	_, err = r.Collection.InsertOne(ctx, product)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("%w: error inserting product: %w", dShared.ErrDuplicateKey, err)
		}
		return fmt.Errorf("error inserting product: %w", err)
	}

	product.ID = id.Hex()
	product.CategoryID = ctgOID.Hex()

	return nil
}
