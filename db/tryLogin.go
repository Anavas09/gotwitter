package db

import (
	"github.com/Anavas09/gotwitter/models"
	"golang.org/x/crypto/bcrypt"
)

//TryLogin : Try to login sending the data to the database
func TryLogin(email string, password string) (models.User, bool) {
	user, founded, _ := CheckIfUserExist(email)

	if founded == false {
		return user, false
	}

	//We need this type to compare with bcrypt.compareHashAndPassword
	passwordBytes := []byte(password)

	passwordDB := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}

	return user, true
}
