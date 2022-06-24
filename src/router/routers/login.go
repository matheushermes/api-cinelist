package routers

import (
	"cinelist/src/controllers"
	"net/http"
)

var routerLogin = Router{
	URI:    		"/login",
	Method: 		http.MethodPost,
	Func: 			controllers.Login,
	RequiresAuth: 	false,
}