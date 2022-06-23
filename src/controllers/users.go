package controllers

import (
	"cinelist/src/answers"
	"cinelist/src/database"
	"cinelist/src/models"
	"cinelist/src/repository"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//CreateUser vai registrar um usuário no banco de dados;
func CreateUser(w http.ResponseWriter, r *http.Request) {

	//Lendo o corpo da requisição;
	bodyRequest, err := ioutil.ReadAll(r.Body);
	if err != nil {
		answers.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	//Colocando o corpo da requisição em um struct;
	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	//Validando informações do cadastro;
	if err = user.Prepare("register"); err != nil {
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

	//Chamando o repositório;
	repository := repository.NewRepositoryUsers(db)
	userId, err := repository.Create(user)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	user.ID = userId

	answers.JSON(w, http.StatusCreated, user)
}

//SearchUsers vai buscar todos os usuários registrados no banco de dados;
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ainda será implementado o método de buscar todos os usuários"))
}

//SearchUser vai buscar um único usuário registrado no banco de dados;
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ainda será implementado o método de buscar um único usuário"))
}

//UpdateUser vai atualizar os dados de usuário no banco de dados;
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	
	//Lendo o ID que vem por parâmetro;
	parameter := mux.Vars(r)
	userId, err := strconv.ParseUint(parameter["userId"], 10, 64)
	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	//Lendo o corpo da requisição;
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		answers.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	//Colocando o corpo os dados em um struct;
	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	//Validando campos de dados preenchidos;
	if err = user.Prepare("edit"); err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	//Abrindo conexão com o banco de dados;
	db, err := database.ConnectingDatabase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	//Chamando o repositório;
	repository := repository.NewRepositoryUsers(db)
	if err = repository.Update(userId, user); err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, err)
}

//DeleteUser vai deletar um usuário no banco de dados;
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	//Lendo o ID que vem por parâmetro;
	parameter := mux.Vars(r)
	userId, err := strconv.ParseUint(parameter["userId"], 10, 64)
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

	//Chamando o repositório;
	repository := repository.NewRepositoryUsers(db)
	if err = repository.Delete(userId); err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, err)
}