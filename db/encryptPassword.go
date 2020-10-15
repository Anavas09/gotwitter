package db

import "golang.org/x/crypto/bcrypt"

//EncryptPassword : Encrypt the user password
func EncryptPassword(pass string) (string, error) {
	weigth := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), weigth)

	return string(bytes), err
}
