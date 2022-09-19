package adapter

import (
	"github.com/labstack/echo/v4"
	"github.com/modern-apis-architecture/coinsure-cards/internal/domain/notification"
	service2 "github.com/modern-apis-architecture/coinsure-cards/internal/domain/notification/service"
	"net/http"
)

type WebhookHandler struct {
	cardSvc *service2.CardUpdateService
}

func (wh *WebhookHandler) Handle(ctx echo.Context) error {
	ct := &notification.CardUpdate{}
	if err := ctx.Bind(ct); err != nil {
		return ctx.JSON(http.StatusBadRequest, &echo.Map{"data": err.Error()})
	}
	err := wh.cardSvc.ReceiveNotification(ct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &echo.Map{"data": err.Error()})
	}
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &echo.Map{"data": err.Error()})
	}
	return ctx.JSON(200, &echo.Map{"message": "success"})
}

func NewWebhookHandler(cardSvc *service2.CardUpdateService) *WebhookHandler {
	return &WebhookHandler{cardSvc: cardSvc}
}
