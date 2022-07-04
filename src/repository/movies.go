package repository

import (
	"cinelist/src/models"
	"database/sql"
)

type movies struct {
	db *sql.DB
}

//NewRepositoryMovie cria uma repositório de filmes;
func NewRepositoryMovies(db *sql.DB) *movies {
	return &movies{db}
}

//CreateMovie insere um novo filme no banco de dados;
func (m movies) CreateMovie(movie models.Program) (uint64, error) {
	statement, err := m.db.Prepare("insert into movieList (name, genre, rating, favorite, user_id) values (?, ?, ?, ?, ?)",)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(movie.Name, movie.Genre, movie.Rating, movie.Favorite,movie.UserID)
	if err != nil {
		return 0, err
	}

	lastIdInsert, err := result.LastInsertId()
	if err != nil {
		return 0,  err
	}

	return uint64(lastIdInsert), nil
}

//GetMovie vai buscar um filme pelo seu ID no banco de dados;
func (m movies) GetMovie(movieID uint64) (models.Program, error) {
	line, err := m.db.Query("select m.id, m.name, m.genre, m.rating, m.favorite, m.user_id from movieList m inner join users u on u.id = m.user_id where m.id = ?", movieID)
	if err != nil {
		return models.Program{}, err
	}
	defer line.Close()

	var movie models.Program

	if line.Next() {
		if err = line.Scan(&movie.ID, &movie.Name, &movie.Genre, &movie.Rating, &movie.Favorite, &movie.UserID); err != nil {
			return models.Program{}, err
		}
	}

	return movie, nil
}

//GetMovieList Vai retornar todos os filmes inseridos pelo o usuário;
func (m movies) GetMovieList(userID uint64) ([]models.Program, error) {
	lines, err := m.db.Query("select m.id, m.name, m.genre, m.rating, m.favorite from animeList m inner join users u on u.id = m.user_id where m.user_id = ?", userID)
	if err != nil {
		return []models.Program{}, err
	}
	defer lines.Close()

	var movies []models.Program

	for lines.Next() {
		var movie models.Program

		if err = lines.Scan(&movie.ID, &movie.Name, &movie.Genre, &movie.Rating, &movie.Favorite); err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

//UpdateMovie atualiza um filme inserido pelo usuário no banco de dados;
func (m movies) UpdateMovie(movieID uint64, movie models.Program) error {
	statement, err := m.db.Prepare("update movieList set name = ?, genre = ?, rating = ?, favorite = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(movie.Name, movie.Genre, movie.Rating, movie.Favorite, movieID); err != nil {
		return err
	}

	return nil
}

//DeleteMovie deleta um anime da lista de filmes do usuário no banco de dados;
func (m movies) DeleteMovie(movieID uint64) error {
	statement, err := m.db.Prepare("delete from movieList where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(movieID); err != nil {
		return err
	}

	return nil
}