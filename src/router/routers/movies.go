package routers

import (
	"cinelist/src/controllers"
	"net/http"
)

var routerMovieList = []Router{
	{
		URI:    		"/movies",
		Method: 		http.MethodPost,
		Func: 			controllers.CreateNewMovie,
		RequiresAuth: 	true,
	},
	{
		URI:    		"/movies",
		Method: 		http.MethodGet,
		Func: 			controllers.SearchMovieList,
		RequiresAuth: 	true,
	},
	{
		URI:    		"/movie/{movieId}",
		Method: 		http.MethodGet,
		Func: 			controllers.SearchMovie,
		RequiresAuth: 	true,
	},
	{
		URI:    		"/movie/{movieId}",
		Method: 		http.MethodPut,
		Func: 			controllers.UpdateMovie,
		RequiresAuth: 	true,
	},
	{
		URI:    		"/movie/{movieId}",
		Method: 		http.MethodDelete,
		Func: 			controllers.DeleteMovie,
		RequiresAuth: 	true,
	},
}