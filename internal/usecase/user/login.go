package user

import domain "FoodOrder/internal/domain/entity"

type LoginUseCaseIntercafe interface {
	Login(cpf string)
}
type LoginUseCase struct {
	repository *domain.UserRepository
}

func NewLoginUseCase(repo *domain.UserRepository) *LoginUseCase {
	return &LoginUseCase{repository: repo}
}

func (uc *LoginUseCase) Login(cpf string) string {
	return ""
}
