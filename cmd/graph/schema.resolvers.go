package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/modern-apis-architecture/coinsure-cards/cmd/graph/generated"
	"github.com/modern-apis-architecture/coinsure-cards/cmd/graph/model"
)

// CreateCard is the resolver for the createCard field.
func (r *mutationResolver) CreateCard(ctx context.Context, input model.CreateCardInput) (*model.Card, error) {
	panic(fmt.Errorf("not implemented: CreateCard - createCard"))
}

// FindCard is the resolver for the findCard field.
func (r *queryResolver) FindCard(ctx context.Context, id string) (*model.Card, error) {
	panic(fmt.Errorf("not implemented: FindCard - findCard"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
