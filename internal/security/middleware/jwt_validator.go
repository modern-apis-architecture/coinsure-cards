package middleware

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"log"
	"net/http"
	"os"
	"time"
)

type JwtValidator struct {
	provider *jwks.CachingProvider
}

func (jv *JwtValidator) EnsureValidToken() func(next http.Handler) http.Handler {
	jwtValidator, err := validator.New(
		jv.provider.KeyFunc,
		validator.RS256,
		jv.provider.IssuerURL.String(),
		[]string{os.Getenv("OPENID_CARDS_AUDIENCE")},
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Encountered error while validating JWT: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"Failed to validate JWT."}`))
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return func(next http.Handler) http.Handler {
		return middleware.CheckJWT(next)
	}
}

func NewJwtValidator(provider *jwks.CachingProvider) *JwtValidator {
	return &JwtValidator{provider: provider}
}
