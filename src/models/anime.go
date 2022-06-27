package models

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

//AnimeList Representa ao um anime adicionado em sua lista pelo o usuário;
type AnimeList struct {
	ID        	uint64    	`json:"id,omitempty"`
	Name      	string    	`json:"name,omitempty"`
	Genre     	string    	`json:"genre,omitempty"`
	Rating    	string    	`json:"rating,omitempty"`
	Favorite	bool		`json:"favorite,omitempty"`
	CreatedIn 	time.Time 	`json:"createdIn,omitempty"`
}

//Prepare vai chamar validar um novo anime;
func (newAnime *AnimeList) Prepare() error {
	if err := newAnime.validate(); err != nil {
		return err
	}

	newAnime.format()
	return nil
}

//validate valida um novo anime que vai ser adicionado na lista do usuário;
func (newAnime *AnimeList) validate() error {
	if newAnime.Name == "" {
		return errors.New("Por favor, insira o nome do anime!")
	}

	if newAnime.Genre == "" {
		return errors.New("Por favor, insira o gênero do anime!")
	}

	if newAnime.Rating == ""  {
		return errors.New("Por favor, insira uma nota de avaliação para o anime que quer inserir!")
	}

	//Convertando a nota do anime para um uint;
	ratingUint, err := strconv.ParseUint(newAnime.Rating, 10, 64)
	if err != nil {
		return errors.New("Não foi possível converter a nota passada para um uint")
	}

	if ratingUint < 0 {
		return errors.New("A nota atribuida ao anime não pode ser menor do que 0")
	} else if ratingUint > 5 {
		return errors.New("A nota atribuida ao anime não pode ser maior do que 5")
	}

	return nil
}

//format formata o nome do anime passado na requisição pelo usuário;
func (newAnime *AnimeList) format() {
	//Removendo os espaços das estremidades;
	newAnime.Name = strings.TrimSpace(newAnime.Name)

	//Colocando a nome do anime em letra miniscula;
	newAnime.Name = strings.ToLower(newAnime.Name)
}

