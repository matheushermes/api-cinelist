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

//CreateNewMovie vai adicionar um novo anime a lista de filmes já assistidos;
func CreateNewMovie(w http.ResponseWriter, r *http.Request) {
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
	var movie models.Program
	if err = json.Unmarshal(bodyRequest, &movie); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	if movie.Favorite != 1 {
		movie.Favorite = 0
	}

	//Adicionando o ID extraido do token no struct;
	movie.UserID = userId

	//Validando anime inserido pelo usuário;
	if err = movie.Prepare(); err != nil {
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

	repository := repository.NewRepositoryMovies(db)
	movieId, err := repository.CreateMovie(movie)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	movie.ID = movieId

	answers.JSON(w, http.StatusCreated, movie)
}

//SearchMovieList traz todos os animes adicionados pelo usuário em sua lista de filmes já assistidos;
func SearchMovieList(w http.ResponseWriter, r *http.Request) {
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
	repository := repository.NewRepositoryMovies(db);
	movies, err := repository.GetMovieList(userId)
	if err != nil {
		answers.Erro(w , http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, movies)
}

//SearchMovie traz o filme escolhido pelo usuário;
func SearchMovie(w http.ResponseWriter, r *http.Request) {
		//Pegando ID vindo por parâmetro;
		parameters := mux.Vars(r)
		movieId, err:= strconv.ParseUint(parameters["movieId"], 10, 64)
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
		repository := repository.NewRepositoryMovies(db)
		movie, err := repository.GetMovie(movieId)
		if err != nil {
			answers.Erro(w, http.StatusInternalServerError, err)
			return
		}
	
		answers.JSON(w, http.StatusOK, movie)
}

//UpdateMovie altera os dados de um filme adicionado;
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	//ID inserido no token do usuário;
	userId, err := auth.ExtractUserIdFromToken(r)
	if err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	//Pegando o ID vindo por parâmetro;
	parameters:= mux.Vars(r)
	movieId, err := strconv.ParseUint(parameters["movieId"], 10, 64)
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
	repository := repository.NewRepositoryMovies(db)
	movieSavedDB, err := repository.GetMovie(movieId)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if movieSavedDB.UserID != userId {
		answers.Erro(w, http.StatusForbidden, errors.New("Você não pode atualizar um filme que não tenha sido inserido por você!"))
		return
	}

	//Lendo o corpo da requisição;
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	//Colocando o corpo da requisição dentro de um struct;
	var movie models.Program
	if err = json.Unmarshal(bodyRequest, &movie); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	//Validar dados do anime;
	if err = movie.Prepare(); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	//Chamando novamente o repositório;
	if err = repository.UpdateMovie(movieId, movie); err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)
}

//DeleteMovie altera os dados de um filme adicionado;
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	//ID inserido no token do usuário;
	userId, err := auth.ExtractUserIdFromToken(r)
	if err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	//Pegando o ID vindo por parâmetro;
	parameters:= mux.Vars(r)
	movieId, err := strconv.ParseUint(parameters["movieId"], 10, 64)
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
	repository := repository.NewRepositoryMovies(db)
	movieSavedDB, err := repository.GetMovie(movieId)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if movieSavedDB.UserID != userId {
		answers.Erro(w, http.StatusForbidden, errors.New("Você não pode deletar um filme que não tenha sido inserido por você!"))
		return
	}

	if err = repository.DeleteMovie(movieId); err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, err)
}