package main

import (
	"cinelist/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("API cineList rodando corretamente!");

	r := router.Generate()

	log.Fatal(http.ListenAndServe(":5000", r))
}