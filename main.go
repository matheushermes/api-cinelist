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
	allowedHeader := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})


	fmt.Printf("API cineList rodando corretamente na porta %d!", config.Port);
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), handlers.CORS(allowedOrigins, allowedHeader, allowedMethods)(r)))
}
