package helper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("clave02security")

type Claims struct {
	IdUser string `json:"id"`
	jwt.StandardClaims
}

func GenerateJWT(IdUser string) (string, error) {
	expirationTime := time.Now().Add(4 * time.Hour)
	claims := &Claims{
		IdUser: IdUser,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
