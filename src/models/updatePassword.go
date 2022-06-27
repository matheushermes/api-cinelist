package models

//UpdatePassword representa o formato da requisição de alteração de senha;
type UpdatePassword struct {
	CurrentPassword string 	`json:"currentPassword"`
	NewPassword 	string	`json:"newPassword"`
}