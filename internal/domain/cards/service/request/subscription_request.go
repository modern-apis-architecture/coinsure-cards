package request

type CreateSubscriptionRequest struct {
	Token string `json:"token"`
	Url   string `json:"url"`
}
