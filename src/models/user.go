package models

import (
	"errors"
	"strings"
	"time"
)

//User representa um usuário utilizando a aplicação web
type User struct {
	ID			uint64		`json:"id,omitempty"`
	Name		string		`json:"name,omitempty"`
	Email 		string 		`json:"email,omitempty"`
	Username 	string 		`json:"username,omitempty"`
	Password 	string 		`json:"password,omitempty"`
	CreatedIn 	time.Time	`json:"createdIn,omitempty"`
}

//Prepare chama os métodos de validação para validar um usuário
func (user *User) Prepare(step string) error {
	if err := user.validateUser(step); err != nil {
		return err
	}

	user.format()
	return nil
}

//validateUser vai validar um usuário
func (user *User) validateUser(step string) error {

	if user.Name == "" {
		return errors.New("Campo nome está vazio!")
	}

	if user.Email == "" {
		return errors.New("Campo e-mail está vazio!")
	}

	if user.Username == "" {
		return errors.New("Campo username está vazio!")
	}

	if step == "register" && user.Password == "" {
		return errors.New("Campo senha está vazio!")
	}

	return nil
}

//format remove os espaços em branco;
func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Username = strings.TrimSpace(user.Username)
}