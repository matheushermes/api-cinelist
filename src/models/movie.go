package models

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

//AnimeList Representa ao um anime adicionado em sua lista pelo o usuário;
type Program struct {
	ID        	uint64    	`json:"id,omitempty"`
	Name      	string    	`json:"name,omitempty"`
	Genre     	string    	`json:"genre,omitempty"`
	Rating    	string    	`json:"rating,omitempty"`
	UserID		uint64		`json:"userId,omitempty"`
	Favorite	int64		`json:"favorite,omitempty"`
	CreatedIn 	time.Time 	`json:"createdIn,omitempty"`
}

//Prepare vai chamar validar um novo anime;
func (newProgram *Program) Prepare() error {
	if err := newProgram.validate(); err != nil {
		return err
	}

	newProgram.format()
	return nil
}

//validate valida um novo anime que vai ser adicionado na lista do usuário;
func (newProgram *Program) validate() error {
	if newProgram.Name == "" {
		return errors.New("Por favor, insira o nome do filme!")
	}

	if newProgram.Genre == "" {
		return errors.New("Por favor, insira o gênero do filme!")
	}

	if newProgram.Rating == ""  {
		return errors.New("Por favor, insira uma nota de avaliação para o filme que quer inserir!")
	}

	//Convertando a nota do anime para um uint;
	ratingUint, err := strconv.ParseInt(newProgram.Rating, 10, 64)
	if err != nil {
		return errors.New("Não foi possível converter a nota passada para um uint")
	}

	if ratingUint < 0 {
		return errors.New("A nota atribuida ao filme não pode ser menor do que 0")
	} else if ratingUint > 5 {
		return errors.New("A nota atribuida ao filme não pode ser maior do que 5")
	}

	return nil
}

//format formata o nome do anime passado na requisição pelo usuário;
func (newProgram *Program) format() {
	//Removendo os espaços das estremidades;
	newProgram.Name = strings.TrimSpace(newProgram.Name)

	//Colocando a nome do anime em letra miniscula;
	newProgram.Name = strings.ToLower(newProgram.Name)
	newProgram.Genre = strings.ToLower(newProgram.Genre)
}

