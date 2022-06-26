package controllers

import "net/http"

//CreateNewAnime vai adicionar um novo anime a lista de animes já assistidos;
func CreateNewAnime(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Adicionando um novo anime"))
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