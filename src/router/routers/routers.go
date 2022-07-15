package routers

import (
	"cinelist/src/middlewares"
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
	routers = append(routers, routerMovieList...)

	for _, router := range routers {

		if router.RequiresAuth == true {
			r.HandleFunc(router.URI, middlewares.Logger(middlewares.Authenticate(router.Func))).Methods(router.Method, http.MethodOptions)
		} else {
			r.HandleFunc(router.URI, middlewares.Logger(router.Func)).Methods(router.Method, http.MethodOptions)
		}
	}

	return r
}