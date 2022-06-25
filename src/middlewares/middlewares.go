package middlewares

import (
	"cinelist/src/answers"
	"cinelist/src/auth"
	"log"
	"net/http"
)

//Logger escreve informações da requisição no terminal;
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}
}

//Authenticate verifica se o usuário que está fazendo a requisição na aplicação está autenticado;
func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			answers.Erro(w, http.StatusUnauthorized, err)
			return
		}

		nextFunction(w, r)
	}
}