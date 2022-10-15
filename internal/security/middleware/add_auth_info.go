package middleware

import (
	"context"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/modern-apis-architecture/coinsure-cards/internal/security/openid"
	"gopkg.in/square/go-jose.v2/jwt"
	"net/http"
)

type AuthMiddleware struct {
	tp       *openid.TokenParser
	provider *jwks.CachingProvider
}

type UserData struct {
	Id string
}

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	data string
}

func (am *AuthMiddleware) AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := jwtmiddleware.AuthHeaderTokenExtractor(r)
			if err != nil || len(token) == 0 {
				//http.Error(w, "Token not found", http.StatusUnauthorized)
				next.ServeHTTP(w, r)
				return
			}
			jwtt, errT := am.tp.ParseToken(token)
			if errT != nil {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			key, err := am.provider.KeyFunc(context.TODO())
			if err != nil {
				http.Error(w, "invalid key", http.StatusUnauthorized)
				return
			}
			claimDest := &jwt.Claims{}
			err2 := jwtt.Claims(key, claimDest)
			if err2 != nil {
				http.Error(w, "invalid claims", http.StatusUnauthorized)
				return
			}
			user := &UserData{
				Id: claimDest.ID,
			}
			ctx := context.WithValue(r.Context(), userCtxKey, user)
			nc := context.WithValue(ctx, "external-auth", token)
			r = r.WithContext(nc)
			next.ServeHTTP(w, r)
		})
	}
}

func NewAuthMiddleware(tp *openid.TokenParser, provider *jwks.CachingProvider) *AuthMiddleware {
	return &AuthMiddleware{tp: tp, provider: provider}
}

func ForContext(ctx context.Context) *UserData {
	raw, _ := ctx.Value(userCtxKey).(*UserData)
	return raw
}
