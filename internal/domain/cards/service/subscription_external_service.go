package service

import "context"

type SubscriptionExternalService interface {
	Subscribe(ctx context.Context, cardId string) error
}
