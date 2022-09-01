package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/exbotanical/gouache/controllers"
	"github.com/exbotanical/gouache/models"
	"github.com/exbotanical/gouache/services"
	"github.com/exbotanical/gouache/utils"

	"github.com/exbotanical/corset"
	"github.com/exbotanical/turnpike"
	"github.com/joho/godotenv"
)

func main() {
	/* Environment */
	envFile := ".env"

	if os.Getenv("LOCAL_MODE") != "" {
		envFile = ".env.local"
	}

	if err := godotenv.Load(envFile); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	port := os.Getenv("PORT")
	clientPort := os.Getenv("CLIENT_PORT")
	gouacheHost := os.Getenv("CLIENT_HOST")

	/* Config */
	origins := []string{utils.ToEndpoint(gouacheHost, clientPort)}
	methods := []string{http.MethodOptions, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch}
	headers := []string{"content-type"}

	c := corset.NewCorset(corset.CorsetOptions{
		AllowedOrigins:   origins,
		AllowedMethods:   methods,
		AllowedHeaders:   headers,
		ExposeHeaders:    []string{"X-Powered-By"},
		AllowCredentials: true,
	})

	bl := services.NewLogger("cmd/serve")

	us, err := services.NewUserService()
	if err != nil {
		log.Fatal(err)
	}

	ss, err := services.NewSessionService()
	if err != nil {
		log.Fatal(err)
	}

	rs, err := services.NewReportService()
	if err != nil {
		fmt.Printf("unable to initialize report service; see %v\n", err)
	}

	/* State */
	ctx := controllers.NewAuthProvider(ss, us, rs)

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
	r.Handler("/auth/login", http.HandlerFunc(ctx.Login)).
		WithMethods(http.MethodPost).
		Use(ctx.ReportRequest).
		Register()

	r.Handler("/auth/register", http.HandlerFunc(ctx.Register)).
		WithMethods(http.MethodPost).
		Use(ctx.ReportRequest).
		Register()

	r.Handler("/auth/renew", http.HandlerFunc(ctx.RenewSession)).
		WithMethods(http.MethodPost).
		Use(ctx.ReportRequest, ctx.Authorize).
		Register()

	r.Handler("/auth/logout", http.HandlerFunc(ctx.Logout)).
		WithMethods(http.MethodPost).
		Use(ctx.ReportRequest, ctx.Authorize).
		Register()

	/* Init */
	fmt.Printf("Listening on port %s...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), c.Handler(r)); err != nil {
		log.Fatal(err)
	}
}
