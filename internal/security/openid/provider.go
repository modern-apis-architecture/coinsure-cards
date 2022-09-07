package openid

import (
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"log"
	"net/url"
	"os"
	"time"
)

func NewJwksProvider() *jwks.CachingProvider {
	issuerURL, err := url.Parse(os.Getenv("OPENID_ISSUER"))
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}
	return jwks.NewCachingProvider(issuerURL, 5*time.Minute)
}
