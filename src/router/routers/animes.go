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
		URI:    		"/animes/{animeId}",
		Method: 		http.MethodGet,
		Func: 			controllers.SearchAnime,
		RequiresAuth: 	true,
	},
	{
		URI:    		"/animes/{animeId}",
		Method: 		http.MethodPut,
		Func: 			controllers.UpdateAnime,
		RequiresAuth: 	true,
	},
	{
		URI:    		"/animes/{animeId}",
		Method: 		http.MethodDelete,
		Func: 			controllers.DeleteAnime,
		RequiresAuth: 	true,
	},
	{
		URI:    		"/animes/{animeId}/favorite",
		Method: 		http.MethodPost,
		Func: 			controllers.AddFavoriteAnime,
		RequiresAuth: 	true,
	},
	{
		URI:    		"/animes/{animeId}/remove-favorite",
		Method: 		http.MethodPost,
		Func: 			controllers.RemoveFavoriteAnime,
		RequiresAuth: 	true,
	},
}