package handler

import "github.com/kevin-luvian/goauth/server/usecases"

type Handler struct {
	authUC usecases.IAuthUseCase
}

type Dependencies struct {
	AuthUC usecases.IAuthUseCase
}

func New(dep Dependencies) *Handler {
	return &Handler{
		authUC: dep.AuthUC,
	}
}
