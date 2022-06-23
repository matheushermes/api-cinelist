package repository

import (
	"cinelist/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

//NewRepositoryUsers cria um repositório de usuários
func NewRepositoryUsers(db *sql.DB) *users {
	return &users{db}
}

//Create insere um usuário no banco de dados
func (u users) Create(user models.User) (uint64, error) {

	statement, err := u.db.Prepare(
		"insert into users (name, email, username, password) values(?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Email, user.Username, user.Password)
	if err != nil {
		return 0, err
	}

	lastIdInsert, err := result.LastInsertId()

	return uint64(lastIdInsert), nil
}

//Update vai atualizar os dados de um usuário no banco de dados;
func (u users) Update(ID uint64, user models.User) error {

	statement, err := u.db.Prepare("update users set name = ?, email = ?, username = ? where id = ?",)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Email, user.Username, ID); err != nil {
		return err
	}

	return nil
}

//Delete vai deletar um usuário no banco de dados;
func (u users) Delete(ID uint64) error {

	statement, err := u.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}