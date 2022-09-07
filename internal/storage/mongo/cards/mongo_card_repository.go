package cards

import (
	"context"
	"fmt"
	"github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCardRepository struct {
	collection *mongo.Collection
}

func (mcr *MongoCardRepository) AddCard(card *cards.Card) (*cards.CardId, error) {
	ctx := context.Background()
	opts := options.InsertOne()
	doc, err := mcr.collection.InsertOne(ctx, card, opts)
	if err != nil {
		return nil, fmt.Errorf("could not save document to mongo: %w", err)
	}
	id := fmt.Sprintf("%v", doc.InsertedID)
	cardId := &cards.CardId{Id: id}
	return cardId, nil
}

func (mcr *MongoCardRepository) GetCard(id string) (*cards.Card, error) {
	var doc cards.Card
	ctx := context.Background()
	if err := mcr.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&doc); err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error finding cart in mongo collection: %w", err)
	} else {
		return &doc, nil
	}
}

func NewMongoCardRepository(coll *mongo.Collection) *MongoCardRepository {
	return &MongoCardRepository{collection: coll}
}
