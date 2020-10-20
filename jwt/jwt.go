package jwt

import (
	"time"

	"github.com/Anavas09/gotwitter/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//GenerateJWT : Generate a JWT
func GenerateJWT(u models.User) (string, error) {
	myKey := []byte("LaClaveSecretaQueUsoParaCrearUnJWT")

	payload := jwt.MapClaims{
		"email":     u.Email,
		"name":      u.Name,
		"lastname":  u.Lastname,
		"birthday":  u.Birthday,
		"location":  u.Location,
		"biography": u.Biography,
		"website":   u.Website,
		"_id":       u.ID.Hex(),
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(myKey)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
