package controllers

import (
	"cinelist/src/answers"
	"cinelist/src/auth"
	"cinelist/src/database"
	"cinelist/src/models"
	"cinelist/src/repository"
	"cinelist/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

//Login é responsável por autenticar um usuário;
func Login(w http.ResponseWriter, r *http.Request) {

	//Lendo o corpo da requisição
	bodyRequest, err := ioutil.ReadAll(r.Body)
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

	//Abrindo conexão com o banco de dados;
	db, err := database.ConnectingDatabase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	//Chamando repositório;
	repository := repository.NewRepositoryUsers(db)
	userSavedDataBase, err := repository.SearchUserByEmail(user.Email)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	//Comparando a senha inserida na requisição com a senha com Hash armazenada no banco de dados;
	if err = security.CheckPassword(userSavedDataBase.Password, user.Password); err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.CreateToken(userSavedDataBase.ID)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	userId := strconv.FormatUint(userSavedDataBase.ID, 10)

	answers.JSON(w, http.StatusOK, models.AuthData{ID: userId, Token: token})
}