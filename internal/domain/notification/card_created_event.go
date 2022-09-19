package notification

type CardCreatedEvent struct {
	Id          string `json:"id"`
	HolderName  string `json:"holder_name"`
	PhoneNumber string `json:"phone_number"`
}
