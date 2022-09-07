package service

import (
	"github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards/service/request"
	cards "github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards/service/response"
)

type CardExternalService interface {
	Create(accountId string, request *request.CreateCardRequest) (*cards.CardCreated, error)
}
