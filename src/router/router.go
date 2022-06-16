package router

import (
	"cinelist/src/router/routers"

	"github.com/gorilla/mux"
)

//Generate vai retornar um router com as rotas configuradas;
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routers.ConfigRouters(r)
}