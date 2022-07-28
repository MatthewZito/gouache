package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	controllers "github.com/MatthewZito/gouache/controllers"
	"github.com/MatthewZito/gouache/db"
	"github.com/MatthewZito/gouache/format"
	srv "github.com/MatthewZito/gouache/services"

	"github.com/MatthewZito/corset"
	"github.com/MatthewZito/turnpike"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	/* Config */
	origins := []string{"*"}
	methods := []string{http.MethodOptions, http.MethodGet, http.MethodPut, http.MethodPatch}
	headers := []string{"*"}

	c := corset.NewCorset(corset.CorsetOptions{
		AllowedOrigins: origins,
		AllowedMethods: methods,
		AllowedHeaders: headers,
	})

	bl := srv.NewLogger("cmd/serve")

	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	/* State */
	rctx := controllers.NewResourceContext(true, db)
	mctx := controllers.NewMetaContext(true)

	/* Routers */
	r := turnpike.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bl.Logf("NotFoundHandler - route %s%s\n", r.Host, r.URL.Path)
		format.FormatError(w, http.StatusNotFound, "invalid route")
	})

	r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bl.Logf("MethodNotAllowedHandler - %s at route %s%s\n", r.Method, r.Host, r.URL.Path)
		format.FormatError(w, http.StatusMethodNotAllowed, "method not allowed")
	})

	r.Handler("/", http.HandlerFunc(controllers.Health)).WithMethods(http.MethodGet).Register()

	r.Handler("/resource/:key[(.+)]", http.HandlerFunc(rctx.GetResource)).WithMethods(http.MethodGet).Register()
	r.Handler("/resource", http.HandlerFunc(rctx.CreateResource)).WithMethods(http.MethodPost).Register()
	r.Handler("/resource", http.HandlerFunc(rctx.GetAllResources)).WithMethods(http.MethodGet).Register()
	r.Handler("/resource/:key[(.+)]", http.HandlerFunc(rctx.UpdateResource)).WithMethods(http.MethodPatch).Register()
	// r.Handler("/resource/:key[(.+)]", http.HandlerFunc(rctx.DeleteResource)).WithMethods(http.MethodDelete).Register()

	r.Handler("/time", http.HandlerFunc(mctx.GetTime)).WithMethods(http.MethodGet).Register()

	/* Init */
	fmt.Printf("Listening on port %s...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), c.Handler(r)); err != nil {
		log.Fatal(err)
	}
}
