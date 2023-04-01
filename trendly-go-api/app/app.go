package app

import (
	"log"
	"net/http"
	"os"
	"trendly-go-api/config"

	"trendly-go-api/app/handler"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	Config *config.Config
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	a.Router = mux.NewRouter()
	a.setRouters()
	a.Config = config
	if runtime_api, _ := os.LookupEnv("AWS_LAMBDA_RUNTIME_API"); runtime_api != "" {
		log.Println("Starting up in Lambda Runtime")
		adapter := gorillamux.NewV2(a.Router)
		lambda.Start(adapter.ProxyWithContext)
	} else {
		log.Println("Starting up on own")
		srv := &http.Server{
			Addr:    ":8080",
			Handler: a.Router,
		}
		_ = srv.ListenAndServe()
	}
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/socials", a.handleRequest(handler.GetSocials))
	a.Get("/categories", a.handleRequest(handler.GetCategories))
	a.Get("/discover/{social_id}/{category_id}", a.handleRequest(handler.GetTrendingMedia))
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(conf *config.Config, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.Config, w, r)
	}
}
