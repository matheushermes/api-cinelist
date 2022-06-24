package models

import (
	"cinelist/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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

	if err := user.format(step); err != nil {
		return err
	}
	
	return nil
}

//validateUser vai validar um usuário
func (user *User) validateUser(step string) error {

	if user.Name == "" {
		return errors.New("Campo nome não pode ficar em branco!")
	}

	if user.Email == "" {
		return errors.New("Campo e-mail não pode ficar em branco!")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("E-mail inserido é inválido!")
	}

	if user.Username == "" {
		return errors.New("Campo username não pode ficar em branco!")
	}

	if step == "register" && user.Password == "" {
		return errors.New("Campo senha não pode ficar em branco!")
	}

	return nil
}

//format remove os espaços em branco;
func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Username = strings.TrimSpace(user.Username)
	user.Password = strings.TrimSpace(user.Password)

	if step == "register" {
		passwordWithHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(passwordWithHash)
	}

	return nil
}