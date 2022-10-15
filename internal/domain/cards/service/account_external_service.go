package service

import (
	"context"
	"github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards/service/request"
	cards2 "github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards/service/response"
)

type AccountExternalService interface {
	CreateAccount(ctx context.Context, request *request.CreateAccountRequest) (*cards2.AccountId, error)
	Get(ctx context.Context, id string) (*cards2.AccountId, error)
}
