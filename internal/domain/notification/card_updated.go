package notification

type CardUpdate struct {
	CardId    string `json:"card_id"`
	Status    string `json:"status"`
	AccountId string `json:"account_id"`
}
