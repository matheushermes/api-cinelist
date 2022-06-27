package controllers

import (
	"cinelist/src/answers"
	"cinelist/src/auth"
	"cinelist/src/database"
	"cinelist/src/models"
	"cinelist/src/repository"
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	w.Write([]byte("Buscando sua lista de animes assistidos"))
}

//SearchAnime traz o anime escolhido pelo usuário;
func SearchAnime(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando anime mencionado"))
}

//UpdateAnime altera os dados de um anime adicionado;
func UpdateAnime(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando anime"))
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