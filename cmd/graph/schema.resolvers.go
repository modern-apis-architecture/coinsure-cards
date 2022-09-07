package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/modern-apis-architecture/coinsure-cards/cmd/graph/generated"
	"github.com/modern-apis-architecture/coinsure-cards/cmd/graph/model"
	"github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards"
	"github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards/service/request"
	"github.com/modern-apis-architecture/coinsure-cards/internal/security/middleware"
)

// CreateCard is the resolver for the createCard field.
func (r *mutationResolver) CreateCard(ctx context.Context, input model.CreateCardInput) (*model.Card, error) {
	user := middleware.ForContext(ctx)
	id, _ := uuid.NewUUID()
	card := &cards.Card{
		Id:             id.String(),
		ValidUntil:     time.Now().Add(time.Hour * 60000),
		LastFourDigits: "1234",
		FirstSixDigits: "123456",
		CreatedAt:      time.Now(),
		Name:           input.Name,
		Tags:           input.Tags,
		User: cards.User{
			Id: user.Id,
		},
		External: cards.External{},
		Status:   "pending",
	}
	pd := &request.PersonalData{
		Address: request.Address{
			Number:  input.PersonalData.Address.Number,
			ZipCode: input.PersonalData.Address.ZipCode,
		},
		BirthDate:  input.PersonalData.BirthDate,
		Document:   input.PersonalData.Document,
		LastName:   input.PersonalData.LastName,
		MotherName: input.PersonalData.MotherName,
		Name:       input.PersonalData.Name,
		Phone: request.Phone{
			Code:   input.PersonalData.Phone.Code,
			Number: input.PersonalData.Phone.Number,
		},
	}
	_, err := r.cardSvc.Store(pd, card)
	if err != nil {
		return nil, err
	}
	cr := &model.Card{
		ID:         id.String(),
		ValidUntil: card.ValidUntil.String(),
		Tags:       card.Tags,
		Name:       card.Name,
		User: &model.User{
			ID: user.Id,
		},
	}
	return cr, nil
}

// FindCard is the resolver for the findCard field.
func (r *queryResolver) FindCard(ctx context.Context, id string) (*model.Card, error) {
	card, err := r.cardSvc.Find(id)
	if err != nil {
		return nil, err
	}
	c := &model.Card{
		ID:         card.Id,
		ValidUntil: card.ValidUntil.String(),
		Tags:       card.Tags,
		Name:       card.Name,
		User: &model.User{
			ID: card.User.Id,
		},
		External: &model.External{CardID: card.External.Id},
	}
	return c, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
