package models

//User representa um usuário utilizando a aplicação web
type User struct {
	Email 		string 	`json:"email, omitempty"`
	Username 	string 	`json:"username, omitempty"`
	PhoneNumber string 	`json:"phoneNumber, omitempty"`
	Password 	string 	`json:"password, omitempty"`
}