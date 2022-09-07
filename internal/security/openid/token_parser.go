package openid

import (
	"fmt"
	"gopkg.in/square/go-jose.v2/jwt"
)

type TokenParser struct {
}

func (tp *TokenParser) ParseToken(tokenStr string) (*jwt.JSONWebToken, error) {
	token, err := jwt.ParseSigned(tokenStr)
	if err != nil {
		return nil, fmt.Errorf("could not parse the token: %w", err)
	}
	return token, nil
}

func NewTokenParser() *TokenParser {
	return &TokenParser{}
}
