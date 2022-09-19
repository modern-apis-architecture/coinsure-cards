package cards

import (
	cards "github.com/modern-apis-architecture/coinsure-cards/internal/domain/cards/service/response"
	"time"
)

type Card struct {
	Id             string    `json:"id" bson:"_id"`
	ValidUntil     time.Time `json:"valid_until" bson:"valid_until"`
	LastFourDigits string    `json:"last_four_digits" bson:"last_four_digits"`
	FirstSixDigits string    `json:"first_six_digits" bson:"first_six_digits"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
	Name           string    `json:"name" bson:"name"`
	Tags           []string  `json:"tags" bson:"tags"`
	User           User      `json:"user" bson:"user"`
	External       External  `json:"external" bson:"external"`
	Status         string    `json:"status" bson:"status"`
}

type User struct {
	Id          string `json:"id" bson:"user_id"`
	Name        string `json:"name" bson:"name"`
	PhoneNumber string `json:"phone_number"`
}

type External struct {
	Card    CardExternal    `json:"card" bson:"card"`
	Account cards.AccountId `json:"account" bson:"account"`
}

type CardExternal struct {
	Id string `json:"id" bson:"id"`
}
