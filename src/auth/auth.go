package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//CreateToken retorna um token assinado com as permissões do usuário;
func CreateToken(userId uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 8).Unix()
	permissions["userId"] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte("Secret"))
}
