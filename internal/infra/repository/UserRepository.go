package repository

import "github.com/jinzhu/gorm"

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(dbGorm *gorm.DB) *UserRepository {
	return &UserRepository{db: dbGorm}
}

func (u *UserRepository) FindByCpfPassword(cpf string, password string) {

}
