package service

type SubscriptionExternalService interface {
	Subscribe(cardId string) error
}
