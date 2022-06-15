package controllers

import "net/http"

//CreateUser vai registrar um usuário no banco de dados;
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário!"))
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