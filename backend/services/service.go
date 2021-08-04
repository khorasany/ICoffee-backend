package services

import (
	"github.com/khorasany/coffee/api/backend/helpers/jwtToken"
	jwtTokenModel "github.com/khorasany/coffee/api/backend/models/authentication/permissions"
	"net/http"
	"time"
)

func SetToken(claims *jwtToken.Claims) (*http.Cookie, error) {
	token, err := jwtToken.CreateJwtToken(claims)
	if err != nil {
		return nil, err
	}

	err = jwtTokenModel.InsertToken(token)
	if err != nil {
		return nil, err
	}

	return &http.Cookie{
		Name:    "AuthenticationToken",
		Value:   token,
		Expires: time.Now().Add(1 * time.Hour),
		Path:    "localhost:9006/",
	}, nil
}
