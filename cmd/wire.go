//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/modern-apis-architecture/coinsure-cards/cmd/graph"
	"github.com/modern-apis-architecture/coinsure-cards/internal/adapter"
	"github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards/repository"
	"github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards/service"
	service2 "github.com/modern-apis-architecture/coinsure-cards/internal/domain/notification/service"
	"github.com/modern-apis-architecture/coinsure-cards/internal/message/cloudevents"
	"github.com/modern-apis-architecture/coinsure-cards/internal/message/kafka"
	"github.com/modern-apis-architecture/coinsure-cards/internal/security/middleware"
	"github.com/modern-apis-architecture/coinsure-cards/internal/security/openid"
	"github.com/modern-apis-architecture/coinsure-cards/internal/storage/banklo"
	"github.com/modern-apis-architecture/coinsure-cards/internal/storage/banklo/account"
	cards2 "github.com/modern-apis-architecture/coinsure-cards/internal/storage/banklo/cards"
	"github.com/modern-apis-architecture/coinsure-cards/internal/storage/banklo/subscription"
	"github.com/modern-apis-architecture/coinsure-cards/internal/storage/mongo"
	"github.com/modern-apis-architecture/coinsure-cards/internal/storage/mongo/cards"
)

func buildAppContainer() (*Application, error) {
	wire.Build(mongo.NewDatabase, mongo.ProvideCollection,
		cards.NewMongoCardRepository,
		banklo.ProvideHttpCli,
		wire.Bind(new(repository.CardRepository), new(*cards.MongoCardRepository)),
		service.NewCardService, graph.NewResolver, openid.NewTokenParser, openid.NewJwksProvider,
		middleware.NewJwtValidator, middleware.NewAuthMiddleware, NewApplication,
		account.NewBankloAccountService, cards2.NewBankloCardService, subscription.NewBankloSubscriptionService,
		wire.Bind(new(service.AccountExternalService), new(*account.BankloAccountService)),
		wire.Bind(new(service.CardExternalService), new(*cards2.BankloCardService)),
		wire.Bind(new(service.SubscriptionExternalService), new(*subscription.BankloSubscriptionService)),
		kafka.ProduceKafkaSender, cloudevents.NewCloudEventsReceiver,
		service2.NewCardUpdateService, adapter.NewWebhookHandler,
	)
	return nil, nil
}
