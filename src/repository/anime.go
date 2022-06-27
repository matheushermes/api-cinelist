package repository

import (
	"cinelist/src/models"
	"database/sql"
)

type animes struct {
	db *sql.DB
}

//NewRepositoryAnimes cria uma repositório de animes;
func NewRepositoryAnimes(db *sql.DB) *animes {
	return &animes{db}
}

//CreateAnime insere um novo anime no banco de dados;
func (a animes) CreateAnime(anime models.AnimeList) (uint64, error) {
	statement, err := a.db.Prepare("insert into animeList (name, genre, rating, user_id) values (?, ?, ?, ?)",)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(anime.Name, anime.Genre, anime.Rating, anime.UserID)
	if err != nil {
		return 0, err
	}

	lastIdInsert, err := result.LastInsertId()
	if err != nil {
		return 0,  err
	}

	return uint64(lastIdInsert), nil
}