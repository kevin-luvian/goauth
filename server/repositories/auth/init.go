package auth

//go:generate mockgen -source=./init.go -destination=mock_auth_repo.go -package=auth

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

type IOAuth interface {
	AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string
	Client(ctx context.Context, t *oauth2.Token) *http.Client
	Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
}

type Repo struct {
	oauth IOAuth
}

func New(oauth IOAuth) *Repo {
	return &Repo{
		oauth: oauth,
	}
}
