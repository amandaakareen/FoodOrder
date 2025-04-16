package controller

import (
	userCase "FoodOrder/internal/usecase/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userLogin userCase.LoginUseCaseIntercafe
}

type LoginInput struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

func NewUserController(login userCase.LoginUseCaseIntercafe) *UserController {
	return &UserController{userLogin: login}
}

func (u *UserController) Login(c echo.Context) error {
	input := LoginInput{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"erro": "Erro ao ler o corpo da requisição: " + err.Error(),
		})
	}

	token, err := u.userLogin.Login(input.Email, input.Senha)

	if err == nil {
		return c.JSON(http.StatusOK, map[string]string{
			"mensagem": "Login OK",
			"email":    input.Email,
			"token":    token,
		})
	}

	return c.JSON(http.StatusUnauthorized, map[string]string{
		"mensagem": "Não autenticado",
	})

}
