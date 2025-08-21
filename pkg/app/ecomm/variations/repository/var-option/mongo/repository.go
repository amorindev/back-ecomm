package mongo

import (
	"github.com/amorindev/go-tmpl/pkg/app/ecomm/variations/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.VarOptionRepo = &Repository{}

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewVarOptionRepo(client *mongo.Client, collection *mongo.Collection) *Repository {
	return &Repository{
		Client:     client,
		Collection: collection,
	}
}
