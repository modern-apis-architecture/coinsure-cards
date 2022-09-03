package repository

import "github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards"

type CardRepository interface {
	AddCard(card cards.Card) error
	GetCard(id string) (cards.Card, error)
}
