package auth

//go:generate mockgen -source=./init.go -destination=mock_auth_usecase.go -package=auth

import (
	"context"

	"github.com/kevin-luvian/goauth/server/entity/google"
	"github.com/kevin-luvian/goauth/server/entity/user"
	"github.com/kevin-luvian/goauth/server/pkg/db"
	"github.com/kevin-luvian/goauth/server/pkg/setting"
)

type IAuthRepo interface {
	GoogleRedirectURL(ctx context.Context, state string) string
	GetGoogleUserInfo(ctx context.Context, code string) (google.GoogleUserInfo, error)
}

type IUserRepo interface {
	Get(ctx context.Context, param db.GetDBParam) ([]user.User, int, error)
	Create(ctx context.Context, usr user.User) (user.User, error)
}

type UseCase struct {
	authRepo IAuthRepo
	userRepo IUserRepo
	secrets  struct {
		access  string
		refresh string
	}
}

type Dependencies struct {
	AuthRepo IAuthRepo
	UserRepo IUserRepo
}

func New(dep Dependencies) *UseCase {
	return &UseCase{
		authRepo: dep.AuthRepo,
		userRepo: dep.UserRepo,
		secrets: struct {
			access  string
			refresh string
		}{
			access:  setting.App.JWTAccessSecret,
			refresh: setting.App.JWTRefreshSecret,
		},
	}
}
