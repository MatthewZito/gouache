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

	/* Routers */
	r := mux.NewRouter()

	r.Handler("/", http.HandlerFunc(controllers.Health)).WithMethods(http.MethodGet).Register()

	/* Init */
	if err := http.ListenAndServe(":5000", c.Handler(r)); err != nil {
		log.Fatal(err)
	}
}
