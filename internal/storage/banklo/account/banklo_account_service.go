package account

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards/service/request"
	cards2 "github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards/service/response"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type BankloAccountService struct {
	cli *http.Client
}

func (bas *BankloAccountService) CreateAccount(ctx context.Context, request *request.CreateAccountRequest) (*cards2.AccountId, error) {
	rootUrl := os.Getenv("CARDS_ISSUER_ROOT_URL")
	body, _ := json.Marshal(request)
	req, err := http.NewRequest(http.MethodPost, rootUrl+"/accounts", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	log.Info(ctx.Value("external-auth").(string))
	req.Header.Set("Authorization", "Bearer "+ctx.Value("external-auth").(string))
	if err != nil {
		return nil, err
	}
	resp, err := bas.cli.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)
	if err != nil {
		log.Errorf("error to create card HTTP %v", err)
		return nil, err
	}
	rb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	acc := &cards2.AccountId{}
	jsonErr := json.Unmarshal(rb, acc)
	if jsonErr != nil {
		return nil, err
	}
	log.Info("Account created!!!")
	return acc, nil
}

func (bas *BankloAccountService) Get(ctx context.Context, id string) (*cards2.AccountId, error) {
	rootUrl := os.Getenv("CARDS_ISSUER_ROOT_URL")
	req, err := http.NewRequest(http.MethodGet, rootUrl+"/accounts/"+id, nil)
	req.Header.Set("Content-Type", "application/json")
	log.Info(ctx.Value("external-auth").(string))
	req.Header.Set("Authorization", "Bearer "+ctx.Value("external-auth").(string))
	if err != nil {
		return nil, err
	}
	resp, err := bas.cli.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)
	if err != nil {
		return nil, err
	}
	rb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	acc := &cards2.AccountId{}
	jsonErr := json.Unmarshal(rb, acc)
	if jsonErr != nil {
		return nil, err
	}
	return acc, nil
}

func NewBankloAccountService(cli *http.Client) *BankloAccountService {
	return &BankloAccountService{cli: cli}
}
