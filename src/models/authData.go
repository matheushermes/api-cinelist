package models

//AuthData contém o token e o ID do usuário authenticado;
type AuthData struct {
	ID 		string `json:"id"`
	Token 	string `json:"token"`
}