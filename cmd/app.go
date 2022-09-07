package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/modern-apis-architecture/coinsure-cards/cmd/graph"
	"github.com/modern-apis-architecture/coinsure-cards/cmd/graph/generated"
	"github.com/modern-apis-architecture/coinsure-cards/internal/security/middleware"
	"log"
	"net/http"
	"os"
)

type Application struct {
	authMid  *middleware.AuthMiddleware
	jv       *middleware.JwtValidator
	resolver *graph.Resolver
}

func (app *Application) Run() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := chi.NewRouter()

	//	router.Use(app.jv.EnsureValidToken())
	router.Use(app.authMid.AuthMiddleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: app.resolver}))

	router.Handle("/", playground.Handler("Coinsure Cards", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func NewApplication(authMid *middleware.AuthMiddleware, jv *middleware.JwtValidator, resolver *graph.Resolver) *Application {
	return &Application{
		authMid:  authMid,
		jv:       jv,
		resolver: resolver,
	}
}
