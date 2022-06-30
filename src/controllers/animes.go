package controllers

import (
	"cinelist/src/answers"
	"cinelist/src/auth"
	"cinelist/src/database"
	"cinelist/src/models"
	"cinelist/src/repository"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//CreateNewAnime vai adicionar um novo anime a lista de animes já assistidos;
func CreateNewAnime(w http.ResponseWriter, r *http.Request) {
	//Pegando o ID inserido no token do usuário;
	userId, err := auth.ExtractUserIdFromToken(r)
	if err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	//Lendo o corpo da requisição;
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	//Passando o corpo da requisição para um struct;
	var anime models.AnimeList
	if err = json.Unmarshal(bodyRequest, &anime); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	if anime.Favorite != 1 {
		anime.Favorite = 0
	}

	//Adicionando o ID extraido do token no struct;
	anime.UserID = userId

	//Validando anime inserido pelo usuário;
	if err = anime.Prepare(); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	//Abrindo conexão com o banco de dados;
	db, err := database.ConnectingDatabase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryAnimes(db)
	animeId, err := repository.CreateAnime(anime)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	anime.ID = animeId

	answers.JSON(w, http.StatusCreated, anime)
}

//SearchAnimeList traz todos os animes adicionados pelo usuário em sua lista de animes já assistidos;
func SearchAnimeList(w http.ResponseWriter, r *http.Request) {
	//Extrair ID do token do usuário que esta fazendo a requisição;
	userId, err := auth.ExtractUserIdFromToken(r)
	if err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	//Abrindo conexão com o banco de dados;
	db, err := database.ConnectingDatabase()
	if err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}
	defer db.Close()

	//chamando repositório;
	repository := repository.NewRepositoryAnimes(db);
	animes, err := repository.GetAnimeList(userId)
	if err != nil {
		answers.Erro(w , http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, animes)
}

//SearchAnime traz o anime escolhido pelo usuário;
func SearchAnime(w http.ResponseWriter, r *http.Request) {
		//Pegando ID vindo por parâmetro;
		parameters := mux.Vars(r)
		animeId, err:= strconv.ParseUint(parameters["animeId"], 10, 64)
		if err != nil {
			answers.Erro(w, http.StatusBadRequest, err)
			return
		}
	
		//Abrindo conexão com o banco de dados;
		db, err := database.ConnectingDatabase()
		if err != nil {
			answers.Erro(w, http.StatusInternalServerError, err)
			return
		}
		defer db.Close()
	
		//Chamando repositório
		repository := repository.NewRepositoryAnimes(db)
		anime, err := repository.GetAnime(animeId)
		if err != nil {
			answers.Erro(w, http.StatusInternalServerError, err)
			return
		}
	
		answers.JSON(w, http.StatusOK, anime)
}

//UpdateAnime altera os dados de um anime adicionado;
func UpdateAnime(w http.ResponseWriter, r *http.Request) {
	//ID inserido no token do usuário;
	userId, err := auth.ExtractUserIdFromToken(r)
	if err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	//Pegando o ID vindo por parâmetro;
	parameters:= mux.Vars(r)
	animeId, err := strconv.ParseUint(parameters["animeId"], 10, 64)
	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	//Abrindo conexão com o banco de dados;
	db, err := database.ConnectingDatabase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	//Chamando repositório;
	repository := repository.NewRepositoryAnimes(db)
	animeSavedDB, err := repository.GetAnime(animeId)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if animeSavedDB.UserID != userId {
		answers.Erro(w, http.StatusForbidden, errors.New("Você não pode atualizar um anime que não tenha sido inserido por você!"))
		return
	}

	//Lendo o corpo da requisição;
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	//Colocando o corpo da requisição dentro de um struct;
	var anime models.AnimeList
	if err = json.Unmarshal(bodyRequest, &anime); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	//Validar dados do anime;
	if err = anime.Prepare(); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	//Chamando novamente o repositório;
	if err = repository.UpdateAnime(animeId, anime); err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)
}

//DeleteAnime altera os dados de um anime adicionado;
func DeleteAnime(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando anime"))
}

//AddFavoriteAnime adiciona um anime a sua lista de animes favoritos;
func AddFavoriteAnime(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Adicionando anime a lista de favoritos"))
}

//RemoveFavoriteAnime remove um anime da lista de favoritos do usuário;
func RemoveFavoriteAnime(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Removendo anime da lista de favoritos"))
}