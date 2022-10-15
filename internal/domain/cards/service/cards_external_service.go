package service

import (
	"context"
	"github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards/service/request"
	cards "github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards/service/response"
)

type CardExternalService interface {
	Create(ctx context.Context, accountId string, request *request.CreateCardRequest) (*cards.CardCreated, error)
}
