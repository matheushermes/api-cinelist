package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//String de conexão com o banco de dados;
	StringConnection = ""

	//Port onde a API vai estar rodando;
	Port = 0


	SecretKey []byte
)

//LoadingEnvironmentVariables vai carregar as váriaveis de ambiente;
func LoadingEnvironmentVariables() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Fatal("Não foi possível carregar a porta 5000")
	}

	StringConnection = fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("HOST"), os.Getenv("DB_NAME"))

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}