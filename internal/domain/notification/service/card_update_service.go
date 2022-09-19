package service

import (
	"context"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/client"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/google/uuid"
	"github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards"
	"github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards/repository"
	"github.com/modern-apis-architecture/coinsure-cards/internal/domain/notification"
	"log"
)

type CardUpdateService struct {
	repo repository.CardRepository
	cc   client.Client
}

func (cus *CardUpdateService) ReceiveNotification(cupd *notification.CardUpdate) error {
	card, err := cus.repo.GetCard(cupd.CardId)
	if err != nil {
		return err
	}
	e := createEvent(card)
	if result := cus.cc.Send(context.Background(), e); cloudevents.IsUndelivered(result) {
		log.Printf("failed to send: %v", result)
	} else {
		log.Printf("sent: %s, accepted: %t", e.ID(), cloudevents.IsACK(result))
	}

	err = cus.repo.UpdateCard(cupd.CardId, cupd.Status, cupd.AccountId)
	if err != nil {
		return err
	}
	return nil
}

func createEvent(card *cards.Card) event.Event {
	e := cloudevents.NewEvent()
	e.SetID(uuid.New().String())
	e.SetType("coinsure.card.ready")
	e.SetSource("coinsure.cards")
	evt := &notification.CardCreatedEvent{
		Id:          card.Id,
		HolderName:  card.User.Name,
		PhoneNumber: card.User.PhoneNumber,
	}

	_ = e.SetData(cloudevents.ApplicationJSON, evt)
	return e
}

func NewCardUpdateService(repo repository.CardRepository, cc client.Client) *CardUpdateService {
	return &CardUpdateService{repo: repo, cc: cc}
}
