package usecases

//go:generate mockgen -source=./interface.go -destination=../handler/mock_usecases.go -package=handler

import (
	"context"

	"github.com/kevin-luvian/goauth/server/entity/google"
	"github.com/kevin-luvian/goauth/server/entity/user"
)

type IAuthUseCase interface {
	GoogleRedirectURL(ctx context.Context, state string) string
	GetGoogleProfileInfo(ctx context.Context, code string) (google.GoogleUserInfo, error)

	SignJWT(ctx context.Context, usr user.User) (string, string, error)
	ParseJWTRefreshToken(c context.Context, refreshToken string) (user.User, error)

	Signup(ctx context.Context, tag, name, email string) (user.User, error)
	GetByEmail(ctx context.Context, email string) (user.User, error)
}
