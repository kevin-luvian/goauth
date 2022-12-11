package auth

import (
	"context"

	"github.com/kevin-luvian/goauth/server/entity/google"
	"github.com/kevin-luvian/goauth/server/pkg/setting"
)

type IAuthRepo interface {
	GoogleLoginURL(ctx context.Context, state string) string
	GetGoogleUserInfo(ctx context.Context, code string) (google.UserInfo, error)
}

type UseCase struct {
	authRepo IAuthRepo
	secret   string
}

type Dependencies struct {
	AuthRepo IAuthRepo
}

// New will instantiate new user usecase
func New(dep Dependencies) *UseCase {
	return &UseCase{
		authRepo: dep.AuthRepo,
		secret:   setting.App.JWTSecret,
	}
}
