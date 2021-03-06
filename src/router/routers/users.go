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
		RequiresAuth: 	true,
	},
	{
		URI:    		"/user/{userId}",
		Method: 		http.MethodGet,
		Func: 			controllers.SearchUser,
		RequiresAuth: 	true,
	},
	{
		URI:    		"/user/{userId}",
		Method: 		http.MethodPut,
		Func: 			controllers.UpdateUser,
		RequiresAuth: 	true,
	},
	{
		URI:    		"/user/{userId}",
		Method: 		http.MethodDelete,
		Func: 			controllers.DeleteUser,
		RequiresAuth: 	true,
	},
	{
		URI: 			"/user/{userId}/update-password",
		Method: 		http.MethodPost,
		Func: 			controllers.UpdatePassword,
		RequiresAuth: 	true,
	},
}