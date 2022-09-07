package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func ProvideCollection(db *mongo.Database) (*mongo.Collection, error) {
	return db.Collection("cards"), nil
}
