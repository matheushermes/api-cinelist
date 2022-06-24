package routers

import (
	"cinelist/src/controllers"
	"net/http"
)

var routerUser = []Router{
	{
		URI:    		"/register",
		Method: 		http.MethodPost,
		Func: 			controllers.CreateUser,
		RequiresAuth: 	false,
	},
	{
		URI:    		"/users",
		Method: 		http.MethodGet,
		Func: 			controllers.SearchUsers,
		RequiresAuth: 	false,
	},
	{
		URI:    		"/user/{userId}",
		Method: 		http.MethodGet,
		Func: 			controllers.SearchUser,
		RequiresAuth: 	false,
	},
	{
		URI:    		"/user/{userId}",
		Method: 		http.MethodPut,
		Func: 			controllers.UpdateUser,
		RequiresAuth: 	false,
	},
	{
		URI:    		"/user/{userId}",
		Method: 		http.MethodDelete,
		Func: 			controllers.DeleteUser,
		RequiresAuth: 	false,
	},
	
}