package usecases

//go:generate mockgen -source=./interface.go -destination=../handler/mock_usecases.go -package=handler

type IAuthUseCase interface {
}
