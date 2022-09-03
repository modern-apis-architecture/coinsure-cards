// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Card struct {
	ID         string    `json:"id"`
	ValidUntil string    `json:"valid_until"`
	Tags       []string  `json:"tags"`
	Name       *string   `json:"name"`
	User       *User     `json:"user"`
	External   *External `json:"external"`
}

type CreateCardInput struct {
	ValidUntil string   `json:"valid_until"`
	Tags       []string `json:"tags"`
	Name       *string  `json:"name"`
}

type External struct {
	CardID string `json:"card_id"`
}

type User struct {
	Email string  `json:"email"`
	ID    *string `json:"id"`
}
