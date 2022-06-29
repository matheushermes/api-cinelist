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
	statement, err := a.db.Prepare("insert into animeList (name, genre, rating, favorite, user_id) values (?, ?, ?, ?, ?)",)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(anime.Name, anime.Genre, anime.Rating, anime.Favorite,anime.UserID)
	if err != nil {
		return 0, err
	}

	lastIdInsert, err := result.LastInsertId()
	if err != nil {
		return 0,  err
	}

	return uint64(lastIdInsert), nil
}

//GetAnime vai buscar um anime pelo seu ID no banco de dados;
func (a animes) GetAnime(animeID uint64) (models.AnimeList, error) {
	line, err := a.db.Query("select a.id, a.name, a.genre, a.rating, a.favorite from animeList a inner join users u on u.id = a.user_id where a.id = ?", animeID)
	if err != nil {
		return models.AnimeList{}, err
	}
	defer line.Close()

	var anime models.AnimeList

	if line.Next() {
		if err = line.Scan(&anime.ID, &anime.Name, &anime.Genre, &anime.Rating, &anime.Favorite); err != nil {
			return models.AnimeList{}, err
		}
	}

	return anime, nil
}

//GetAnimeList Vai retornar todos os animes inseridos pelo o usuário;
func (a animes) GetAnimeList(userID uint64) ([]models.AnimeList, error) {

}