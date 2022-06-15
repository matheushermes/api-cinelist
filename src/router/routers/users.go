package routers

import "net/http"

var routerUser = []Router{
	{
		URI:    "/register",
		Method: http.MethodPost,
		Func: func(w http.ResponseWriter, r *http.Request) {},
		RequiresAuth: false,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Func: func(w http.ResponseWriter, r *http.Request) {},
		RequiresAuth: false,
	},
	{
		URI:    "/user/{userId}",
		Method: http.MethodGet,
		Func: func(w http.ResponseWriter, r *http.Request) {},
		RequiresAuth: false,
	},
	{
		URI:    "/user/{userId}",
		Method: http.MethodPut,
		Func: func(w http.ResponseWriter, r *http.Request) {},
		RequiresAuth: false,
	},
	{
		URI:    "/user/{userId}",
		Method: http.MethodDelete,
		Func: func(w http.ResponseWriter, r *http.Request) {},
		RequiresAuth: false,
	},
	
}