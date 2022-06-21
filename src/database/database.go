package database

import (
	"cinelist/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //driver
)

//ConnectingDatabase abre a conexão com o banco de dados;
func ConnectingDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}