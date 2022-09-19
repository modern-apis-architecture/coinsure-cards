package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/labstack/echo/v4"
	"github.com/modern-apis-architecture/coinsure-cards/cmd/graph"
	"github.com/modern-apis-architecture/coinsure-cards/cmd/graph/generated"
	"github.com/modern-apis-architecture/coinsure-cards/internal/adapter"
	"github.com/modern-apis-architecture/coinsure-cards/internal/security/middleware"
	"log"
	"net/http"
	"os"
	"sync"
)

type Application struct {
	authMid  *middleware.AuthMiddleware
	jv       *middleware.JwtValidator
	resolver *graph.Resolver
	wh       *adapter.WebhookHandler
}

func (app *Application) Run() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		app.grqphqlServer()
		wg.Done()
	}()
	go func() {
		app.restServer()
		wg.Done()
	}()
	wg.Wait()
}

func (app *Application) grqphqlServer() {
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

func (app *Application) restServer() {
	e := echo.New()
	e.POST("/cards-webhook", app.wh.Handle)
	e.Logger.Fatal(e.Start(":6000"))
}

func NewApplication(authMid *middleware.AuthMiddleware, jv *middleware.JwtValidator, resolver *graph.Resolver, wh *adapter.WebhookHandler) *Application {
	return &Application{
		authMid:  authMid,
		jv:       jv,
		resolver: resolver,
		wh:       wh,
	}
}
