package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kevin-luvian/goauth/server/entity/token"
	"github.com/kevin-luvian/goauth/server/entity/user"
	"github.com/kevin-luvian/goauth/server/pkg/db"
)

func (u *UseCase) SignJWT(c context.Context, usr user.User) (string, string, error) {
	// generate access token
	aTokenClaims := token.AccessToken{
		UserID:   usr.ID,
		UserName: usr.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, aTokenClaims)
	accessToken, err := jwtToken.SignedString([]byte(u.secrets.access))
	if err != nil {
		return "", "", err
	}

	// generate refresh token
	rTokenClaims := token.RefreshToken{
		UserID: usr.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Hour)),
		},
	}

	jwtToken = jwt.NewWithClaims(jwt.SigningMethodHS256, rTokenClaims)
	refreshToken, err := jwtToken.SignedString([]byte(u.secrets.refresh))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (u *UseCase) ParseJWTAccessToken(c context.Context, accessToken string) (user.User, error) {
	aTokenClaims := token.AccessToken{}

	tkn, err := jwt.ParseWithClaims(accessToken, &aTokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(u.secrets.access), nil
	})

	// is token valid
	if err != nil || !tkn.Valid {
		return user.User{}, errors.New("invalid token provided")
	}

	users, _, err := u.userRepo.Get(c, db.GetDBParam{
		Filters: []db.Filter{{
			Field: "id",
			Value: aTokenClaims.UserID,
		}},
		DisableCount: true,
	})
	if err != nil || len(users) <= 0 {
		return user.User{}, fmt.Errorf("can't get user with id %d", aTokenClaims.UserID)
	}

	return users[0], nil
}

func (u *UseCase) ParseJWTRefreshToken(c context.Context, refreshToken string) (user.User, error) {
	rTokenClaims := token.RefreshToken{}

	tkn, err := jwt.ParseWithClaims(refreshToken, &rTokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(u.secrets.refresh), nil
	})

	// is token valid
	if err != nil || !tkn.Valid {
		return user.User{}, errors.New("invalid token provided")
	}

	users, _, err := u.userRepo.Get(c, db.GetDBParam{
		Filters: []db.Filter{{
			Field: "id",
			Value: rTokenClaims.UserID,
		}},
		DisableCount: true,
	})
	if err != nil || len(users) <= 0 {
		return user.User{}, fmt.Errorf("can't get user with id %d", rTokenClaims.UserID)
	}

	return users[0], nil
}
