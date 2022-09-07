package service

import (
	"github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards/service/request"
	cards2 "github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards/service/response"
)

type AccountExternalService interface {
	CreateAccount(request *request.CreateAccountRequest) (*cards2.AccountId, error)
}
