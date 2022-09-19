package repository

import "github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards"

type CardRepository interface {
	AddCard(card *cards.Card) (*cards.CardId, error)
	GetCard(id string) (*cards.Card, error)
	UpdateCard(id string, status string, accountId string) error
}
