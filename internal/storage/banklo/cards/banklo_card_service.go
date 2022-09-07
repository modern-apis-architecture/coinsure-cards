package cards

import (
	"bytes"
	"encoding/json"
	"github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards/service/request"
	cards "github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards/service/response"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type BankloCardService struct {
	cli *http.Client
}

func (bcs *BankloCardService) Create(accountId string, request *request.CreateCardRequest) (*cards.CardCreated, error) {
	rootUrl := os.Getenv("CARDS_ISSUER_ROOT_URL")
	body, _ := json.Marshal(request)
	req, err := http.NewRequest(http.MethodPost, rootUrl+"/accounts/"+accountId+"/cards", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	resp, err := bcs.cli.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)
	rb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	cc := &cards.CardCreated{}
	jsonErr := json.Unmarshal(rb, cc)
	if jsonErr != nil {
		return nil, err
	}
	return cc, nil
}

func NewBankloCardService(cli *http.Client) *BankloCardService {
	return &BankloCardService{cli: cli}
}
