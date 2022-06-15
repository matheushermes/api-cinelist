package router

import "github.com/gorilla/mux"

//Generate vai retornar um router com as rotas configuradas;
func Generate() *mux.Router {
	return mux.NewRouter()
}