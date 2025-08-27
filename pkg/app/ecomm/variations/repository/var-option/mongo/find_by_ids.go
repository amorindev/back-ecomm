package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/variations/domain"
	dShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) FindByIDs(ctx context.Context, ids []string) ([]*domain.VarOption, error) {

	objectIDs := make([]bson.ObjectID, 0, len(ids))
	for _, idStr := range ids {
		oid, err := bson.ObjectIDFromHex(idStr)
		if err != nil {
			return nil, fmt.Errorf("%w: invalid id %s: %w", dShared.ErrIncorrectID, idStr, err)
		}
		objectIDs = append(objectIDs, oid)
	}

	filter := bson.M{"_id": bson.M{"$in": objectIDs}}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find VarOptions by IDs: %w", err)
	}
	defer cursor.Close(ctx)

	var results []*domain.VarOption
	if err := cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("failed to decode VarOptions: %w", err)
	}

	return results, nil
}
