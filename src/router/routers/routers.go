package routers

import "net/http"

//Router representa todas as rotas da api;
type Router struct {
	URI    string
	Method string
	Func   func(http.ResponseWriter, *http.Request)
	RequiresAuth bool
}