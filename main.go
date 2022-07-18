package main

import (
	"cinelist/src/config"
	"cinelist/src/router"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	config.LoadingEnvironmentVariables()
	r := router.Generate()

	//Cors
	allowedHeader := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"http://127.0.0.1:5500"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})


	fmt.Printf("API cineList rodando corretamente na porta %d!", config.Port);
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), handlers.CORS(allowedOrigins, allowedHeader, allowedMethods)(r)))
}
