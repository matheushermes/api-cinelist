package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Router representa todas as rotas da api;
type Router struct {
	URI    string
	Method string
	Func   func(http.ResponseWriter, *http.Request)
	RequiresAuth bool
}

//ConfigRouters coloca todas as rotas dentro do router;
func ConfigRouters(r *mux.Router) *mux.Router {
	routers := routerUser
	routers = append(routers, routerLogin)

	for _, router := range routers {
		r.HandleFunc(router.URI, router.Func).Methods(router.Method)
	}

	return r
}