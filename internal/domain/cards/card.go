package cards

import "time"

type Card struct {
	Id             string    `json:"id" bson:"_id"`
	ValidUntil     time.Time `json:"valid_until" bson:"valid_until"`
	LastFourDigits string    `json:"last_four_digits" bson:"last_four_digits"`
	FirstSixDigits string    `json:"first_six_digits" bson:"first_six_digits"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
	Name           string    `json:"name" bson:"name"`
	Tags           []string  `json:"tags" bson:"tags"`
	User           User      `json:"user" bson:"user"`
}

type User struct {
	Id    string `json:"id" bson:"user_id"`
	Email string `json:"email" bson:"email"`
}
