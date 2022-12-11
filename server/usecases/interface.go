package usecases

import (
	"context"

	"github.com/kevin-luvian/goauth/server/entity/user"
)

//go:generate mockgen -source=./interface.go -destination=../handler/mock_usecases.go -package=handler

type IAuthUseCase interface {
	GetGoogleLoginURL(c context.Context, state string) string
	GetGoogleProfileInfo(c context.Context, state string, code string) (user.User, error)
	SignJWT(c context.Context, usr user.User) (string, error)
}
