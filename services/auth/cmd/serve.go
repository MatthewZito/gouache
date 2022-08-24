package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/exbotanical/gouache/cache"
	"github.com/exbotanical/gouache/controllers"
	"github.com/exbotanical/gouache/models"
	"github.com/exbotanical/gouache/repositories"
	srv "github.com/exbotanical/gouache/services"

	"github.com/exbotanical/corset"
	"github.com/exbotanical/turnpike"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	/* Config */
	origins := []string{"http://localhost:3000"}
	methods := []string{http.MethodOptions, http.MethodGet, http.MethodPut, http.MethodPatch}
	headers := []string{"*"}

	c := corset.NewCorset(corset.CorsetOptions{
		AllowedOrigins:   origins,
		AllowedMethods:   methods,
		AllowedHeaders:   headers,
		AllowCredentials: true,
	})

	bl := srv.NewLogger("cmd/serve")

	db, err := repositories.Connect()
	if err != nil {
		log.Fatal(err)
	}

	cache, err := cache.NewRedisStore()
	if err != nil {
		log.Fatal(err)
	}

	/* State */
	actx := controllers.NewSessionContext(cache, db)

	/* Routers */
	r := turnpike.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bl.Logf("NotFoundHandler - route %s%s\n", r.Host, r.URL.Path)
		models.FormatError(w, http.StatusNotFound, "invalid route", "", 0)
	})

	r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bl.Logf("MethodNotAllowedHandler - %s at route %s%s\n", r.Method, r.Host, r.URL.Path)
		models.FormatError(w, http.StatusMethodNotAllowed, "method not allowed", "", 0)
	})

	/* Health check */
	r.Handler("/", http.HandlerFunc(controllers.Health)).WithMethods(http.MethodGet).Register()

	/* Session @todo relocate to separate service */
	r.Handler("/session/login", http.HandlerFunc(actx.Login)).WithMethods(http.MethodPost).Register()
	r.Handler("/session/register", http.HandlerFunc(actx.Register)).WithMethods(http.MethodPost).Register()
	r.Handler("/session/renew", http.HandlerFunc(actx.RenewSession)).WithMethods(http.MethodPost).Use(actx.Authorize).Register()
	r.Handler("/session/logout", http.HandlerFunc(actx.Logout)).WithMethods(http.MethodPost).Use(actx.Authorize).Register()

	/* Init */
	fmt.Printf("Listening on port %s...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), c.Handler(r)); err != nil {
		log.Fatal(err)
	}
}