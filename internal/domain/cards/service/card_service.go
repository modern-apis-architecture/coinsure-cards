package service

import (
	"fmt"
	"github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards"
	"github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards/repository"
	"github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards/service/request"
)

const PhoneNumber = "+55%s%s"

type CardService struct {
	repo   repository.CardRepository
	acSvc  AccountExternalService
	ccSvc  CardExternalService
	subSvc SubscriptionExternalService
}

func NewCardService(repo repository.CardRepository, acSvc AccountExternalService, ccSvc CardExternalService, subSvc SubscriptionExternalService) *CardService {
	return &CardService{
		repo:   repo,
		acSvc:  acSvc,
		ccSvc:  ccSvc,
		subSvc: subSvc,
	}
}

func (cs *CardService) Find(id string) (*cards.Card, error) {
	return cs.repo.GetCard(id)
}

func (cs *CardService) Store(personalData *request.PersonalData, card *cards.Card) (*cards.CardId, error) {
	accReq := &request.CreateAccountRequest{
		PersonalData: *personalData,
	}
	account, err := cs.acSvc.CreateAccount(accReq)
	if err != nil {
		return nil, err
	}
	ccReq := &request.CreateCardRequest{Name: card.Name}
	cardCreated, err := cs.ccSvc.Create(account.Id, ccReq)
	if err != nil {
		return nil, err
	}
	err = cs.subSvc.Subscribe(cardCreated.Id)
	if err != nil {
		return nil, err
	}
	card.User.PhoneNumber = fmt.Sprintf(PhoneNumber, personalData.Phone.Code, personalData.Phone.Number)
	return cs.repo.AddCard(card)
}
