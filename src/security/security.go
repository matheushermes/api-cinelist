package security

import "golang.org/x/crypto/bcrypt"

//Hash recebe uma string e coloca um Hash nela;
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//CheckPassword compara a senha inserida pelo usuário no Login é a mesma armazenada com um hash no banco de dados;
func CheckPassword(passwordWithHash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordWithHash), []byte(password))
}