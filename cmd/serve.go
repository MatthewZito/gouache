package main

import (
	"log"
	"net/http"

	cors "github.com/MatthewZito/gouache/cors"
	mux "github.com/MatthewZito/gouache/premux"

	controllers "github.com/MatthewZito/gouache/controllers"
)

func main() {
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

	r.Handler("/", http.HandlerFunc(controllers.Health)).WithMethods(http.MethodGet).Register()
	r.Handler("/resource/:key[(.+)]", http.HandlerFunc(rc.GetResource)).WithMethods(http.MethodGet).Register()
	r.Handler("/resource", http.HandlerFunc(rc.AddResource)).WithMethods(http.MethodPost).Register()

	/* Init */
	if err := http.ListenAndServe(":5000", c.Handler(r)); err != nil {
		log.Fatal(err)
	}
}
