package jwtToken

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	BirthDate string `json:"birthdate"`
	Role      int64  `json:"role"`
	jwt.StandardClaims
}

var jwtKey = []byte("iCoffeeProject")

func CreateJwtToken(claims *Claims) (string, error) {
	claims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

func AuthenticationJwtToken(token string) (Token string, ID string, sessionID string, err error) {
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}

	return "", "", "", err
}
