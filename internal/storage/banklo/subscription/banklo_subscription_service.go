package subscription

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards/service/request"
	"io"
	"net/http"
	"os"
)

type BankloSubscriptionService struct {
	cli *http.Client
}

func (bsur *BankloSubscriptionService) Subscribe(ctx context.Context, cardId string) error {
	sr := &request.CreateSubscriptionRequest{
		Token: "AAA",
		Url:   "http://localhost:9999/cards/" + cardId,
	}
	body, _ := json.Marshal(sr)
	rootUrl := os.Getenv("CARDS_ISSUER_ROOT_URL")
	req, err := http.NewRequest(http.MethodPost, rootUrl+"/cards/"+cardId+"/subscribe", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", ctx.Value("external-auth").(string))
	if err != nil {
		return err
	}
	resp, err := bsur.cli.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)
	return nil
}

func NewBankloSubscriptionService(cli *http.Client) *BankloSubscriptionService {
	return &BankloSubscriptionService{cli: cli}
}
