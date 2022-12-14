package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func NewDatabase() (*mongo.Database, error) {
	dsn := os.Getenv("MONGO_DSN")
	if dsn == "" {
		dsn = "mongodb://localhost:27017"
	}
	ctx := context.Background()
	if m, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn)); err != nil {
		return nil, err
	} else {
		return m.Database("coinsure-cards"), nil
	}
}
