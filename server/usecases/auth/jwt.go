package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kevin-luvian/goauth/server/entity/user"
)

func (u *UseCase) SignJWT(c context.Context, usr user.User) (string, error) {
	expirationTime := time.Now().Add(180 * time.Minute)

	usrClaims := struct {
		user.User
		jwt.RegisteredClaims
	}{
		User: usr,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, usrClaims)
	tokenString, err := token.SignedString([]byte(u.secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (u *UseCase) ParseJWT(c context.Context, token string) (user.User, error) {
	usrClaims := struct {
		user.User
		jwt.RegisteredClaims
	}{}

	tkn, err := jwt.ParseWithClaims(token, &usrClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(u.secret), nil
	})

	// initial error
	if err != nil || !tkn.Valid {
		return user.User{}, errors.New("invalid token provided")
	}

	return usrClaims.User, nil
}
