package routers

import (
	"cinelist/src/controllers"
	"net/http"
)

var routerAnimeList = []Router{
	{
		URI:    		"/animes",
		Method: 		http.MethodPost,
		Func: 			controllers.CreateNewAnime,
		RequiresAuth: 	true,
	},
	{
		URI:    		"/animes",
		Method: 		http.MethodGet,
		Func: 			controllers.SearchAnimeList,
		RequiresAuth: 	true,
	},
	{
		URI:    		"/anime/{animeId}",
		Method: 		http.MethodGet,
		Func: 			controllers.SearchAnime,
		RequiresAuth: 	true,
	},
	{
		URI:    		"/anime/{animeId}",
		Method: 		http.MethodPut,
		Func: 			controllers.UpdateAnime,
		RequiresAuth: 	true,
	},
	{
		URI:    		"/anime/{animeId}",
		Method: 		http.MethodDelete,
		Func: 			controllers.DeleteAnime,
		RequiresAuth: 	true,
	},
}