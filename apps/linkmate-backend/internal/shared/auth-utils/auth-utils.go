package authUtils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type TokenData struct {
	UserId int
	Email  string
}

func ParseJwtData(maybeToken interface{}) (*TokenData, error) {
	token, ok := maybeToken.(*jwt.Token)

	if !ok {
		return nil, errors.New("token parse error")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("claims parse error")
	}

	return &TokenData{
		UserId: int(claims["id"].(float64)),
		Email:  claims["email"].(string),
	}, nil
}
