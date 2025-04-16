package repository

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
)

type UserRepositoryIntercafe interface {
	FindByEmailPassword(email string, password string) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(dbGorm *gorm.DB) UserRepositoryIntercafe {
	return &UserRepository{db: dbGorm}
}

type User struct {
	ID        int    `gorm:"column:id"`
	Name      string `gorm:"column:name"`
	Password  string `gorm:"column:password"`
	Cpf       string `gorm:"column:cpf"`
	Email     string `gorm:"column:email"`
	Telephone string `gorm:"column:telephone"`
}

func (u *UserRepository) FindByEmailPassword(email string, password string) error {
	var user User
	log.Println(email + password)
	err := u.db.Table("users").Where("email = ? AND password = ? ", email, password).First(&user).Error

	if err != nil {
		log.Println("Usuário não encontrado:", err)
		return err
	}
	if user.ID == 0 {
		log.Println("Usuário não encontrado 1 ")
		return errors.New("Usuário não encontrado")
	}

	return nil
}
