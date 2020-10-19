package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/Anavas09/gotwitter/db"
	"github.com/Anavas09/gotwitter/models"
)

//Email will be use this email on all endpoints
var Email string

//UserID is the ID from model. We will use this on all endpoints
var UserID string

//ProcessToken : Process the token to get the token data
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myKey := []byte("LaClaveSecretaQueUsoParaCrearUnJWT")

	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Invalid Token Format")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		_, founded, _ := db.CheckIfUserExist(claims.Email)

		if founded {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}

		return claims, founded, UserID, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}

	return claims, false, string(""), err
}
