package token

import (
	"github.com/golang-jwt/jwt/v4"
)

type AccessToken struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

type RefreshToken struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}
