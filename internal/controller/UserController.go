package controller

import (
	"FoodOrder/internal/usecase/user"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userLogin *user.LoginUseCaseIntercafe
}

func NewUserController(login *user.LoginUseCaseIntercafe) *UserController {
	return &UserController{userLogin: login}
}

func Login(c echo.Context) (string, error) {

	return "", nil
}
