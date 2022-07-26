package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	controllers "github.com/MatthewZito/gouache/controllers"
	"github.com/MatthewZito/gouache/cors"
	"github.com/MatthewZito/gouache/format"
	mux "github.com/MatthewZito/gouache/premux"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.example")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	/* Config */
	origins := []string{"*"}
	methods := []string{"PUT"}

	c := cors.New(cors.CorsOptions{
		AllowedOrigins: origins,
		AllowedMethods: methods,
	})

	/* State */
	rc := controllers.NewResourceCache()

	/* Routers */
	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		format.FormatError(w, http.StatusNotFound, "invalid route")
	})

	r.Handler("/", http.HandlerFunc(controllers.Health)).WithMethods(http.MethodGet).Register()

	r.Handler("/resource/:key[(.+)]", http.HandlerFunc(rc.GetResource)).WithMethods(http.MethodGet).Register()
	r.Handler("/resource", http.HandlerFunc(rc.AddResource)).WithMethods(http.MethodPost).Register()
	r.Handler("/resource/:key[(.+)]", http.HandlerFunc(rc.UpdateResource)).WithMethods(http.MethodPatch).Register()
	r.Handler("/resource/:key[(.+)]", http.HandlerFunc(rc.DeleteResource)).WithMethods(http.MethodDelete).Register()

	r.Handler("/time", http.HandlerFunc(controllers.GetTime)).WithMethods(http.MethodGet).Register()

	/* Init */
	fmt.Printf("Listening on port %s...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), c.Handler(r)); err != nil {
		log.Fatal(err)
	}
}
