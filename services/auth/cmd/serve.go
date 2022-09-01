package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/exbotanical/gouache/controllers"
	"github.com/exbotanical/gouache/models"
	"github.com/exbotanical/gouache/repositories"
	"github.com/exbotanical/gouache/services"
	"github.com/exbotanical/gouache/utils"

	"github.com/exbotanical/corset"
	"github.com/exbotanical/turnpike"
	"github.com/joho/godotenv"
)

func main() {
	/* Environment */
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	clientPort := os.Getenv("CLIENT_PORT")
	gouacheHost := os.Getenv("CLIENT_HOST")

	/* Config */
	origins := []string{utils.ToEndpoint(gouacheHost, clientPort)}
	methods := []string{http.MethodOptions, http.MethodGet, http.MethodPut, http.MethodPatch}
	headers := []string{"*"}

	c := corset.NewCorset(corset.CorsetOptions{
		AllowedOrigins:   origins,
		AllowedMethods:   methods,
		AllowedHeaders:   headers,
		AllowCredentials: true,
	})

	bl := services.NewLogger("cmd/serve")

	t, err := repositories.InitUserTable()
	if err != nil {
		log.Fatal(err)
	}

	cache, err := repositories.NewRedisStore()
	if err != nil {
		log.Fatal(err)
	}

	q := repositories.NewReportRepository()

	/* State */
	ctx := controllers.NewSessionContext(cache, t, q)

	/* Routers */
	r := turnpike.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bl.Logf("NotFoundHandler - route %s%s\n", r.Host, r.URL.Path)
		models.SendGouacheException(w, http.StatusNotFound, "invalid route", "", 0)
	})

	r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bl.Logf("MethodNotAllowedHandler - %s at route %s%s\n", r.Method, r.Host, r.URL.Path)
		models.SendGouacheException(w, http.StatusMethodNotAllowed, "method not allowed", "", 0)
	})

	/* Health */
	r.Handler("/auth/health", http.HandlerFunc(controllers.Health)).WithMethods(http.MethodGet).Use(ctx.ReportRequest).Register()

	/* Auth */
	r.Handler("/auth/login", http.HandlerFunc(ctx.Login)).WithMethods(http.MethodPost).Register()
	r.Handler("/auth/register", http.HandlerFunc(ctx.Register)).WithMethods(http.MethodPost).Register()
	r.Handler("/auth/renew", http.HandlerFunc(ctx.RenewSession)).WithMethods(http.MethodPost).Use(ctx.Authorize).Register()
	r.Handler("/auth/logout", http.HandlerFunc(ctx.Logout)).WithMethods(http.MethodPost).Use(ctx.Authorize).Register()

	/* Init */
	fmt.Printf("Listening on port %s...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), c.Handler(r)); err != nil {
		log.Fatal(err)
	}
}
