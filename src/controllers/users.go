package controllers

import (
	"cinelist/src/answers"
	"cinelist/src/database"
	"cinelist/src/models"
	"cinelist/src/repository"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//CreateUser vai registrar um usuário no banco de dados;
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body);
	if err != nil {
		answers.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("cadastro"); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConnectingDatabase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

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
	w.Write([]byte("Buscando todos os usuários!"))
}

//SearchUser vai buscar um único usuário registrado no banco de dados;
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um único usuário!"))
}

//UpdateUser vai atualizar os dados de usuário no banco de dados;
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário!"))
}

//DeleteUser vai deletar um usuário no banco de dados;
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário!"))
}