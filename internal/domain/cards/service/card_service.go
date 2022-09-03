package service

import "github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards/repository"

type CardService struct {
	repo repository.CardRepository
}
