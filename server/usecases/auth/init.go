package auth

import (
	"context"
)

type IAuthRepo interface {
	GoogleLoginURL(ctx context.Context, state string) string
	GetGoogleUserInfo(ctx context.Context, code string) error
}

type UseCase struct {
	authRepo    IAuthRepo
	stateLookup map[string]string
}

type Dependencies struct {
	AuthRepo IAuthRepo
}

// New will instantiate new user usecase
func New(dep Dependencies) *UseCase {
	return &UseCase{
		authRepo:    dep.AuthRepo,
		stateLookup: map[string]string{},
	}
}
