package graph

import "github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	cardSvc *service.CardService
}

func NewResolver(cardSvc *service.CardService) *Resolver {
	return &Resolver{cardSvc: cardSvc}
}
