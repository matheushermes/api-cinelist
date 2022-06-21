package main

import (
	"cinelist/src/config"
	"cinelist/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadingEnvironmentVariables()
	r := router.Generate()

	fmt.Printf("API cineList rodando corretamente na porta %d!", config.Port);
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
