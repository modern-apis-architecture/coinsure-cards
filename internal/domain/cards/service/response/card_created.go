package cards

type CardCreated struct {
	Id             string `json:"id" bson:"_id"`
	LastFourDigits string `json:"last_four_digits" bson:"last_four_digits"`
	FirstSixDigits string `json:"first_six_digits" bson:"first_six_digits"`
}
