package userCase

import (
	"FoodOrder/internal/infra/repository"
	"log"
)

type LoginUseCaseIntercafe interface {
	Login(email string, password string) (string, error)
}
type LoginUseCase struct {
	repository repository.UserRepositoryIntercafe
}

func NewLoginUseCase(repo repository.UserRepositoryIntercafe) LoginUseCaseIntercafe {
	return &LoginUseCase{repository: repo}
}

func (uc LoginUseCase) Login(email string, password string) (string, error) {

	err := uc.repository.FindByEmailPassword(email, password)

	if err == nil {
		return "token", nil
	}

	log.Println("Usuário não encontrado case", err)
	return "", err
}
