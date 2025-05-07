package controller

import (
	"FoodOrder/internal/auth"
	userCase "FoodOrder/internal/usecase/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userLogin   userCase.LoginUseCaseIntercafe
	authService *auth.AuthService
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserController(login userCase.LoginUseCaseIntercafe, auth *auth.AuthService) *UserController {
	return &UserController{userLogin: login, authService: auth}
}

func (u *UserController) Login(c echo.Context) error {
	input := LoginInput{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"erro": "Erro ao ler o corpo da requisição: " + err.Error(),
		})
	}

	userID, err := u.userLogin.Login(input.Email, input.Password)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"mensagem": "Não autenticado",
		})
	}

	token, err := u.authService.GenerateToken(userID)

	if err == nil {
		return c.JSON(http.StatusOK, map[string]string{
			"mensagem": "Login OK",
			"email":    input.Email,
			"token":    token,
		})
	}

	return c.JSON(http.StatusInternalServerError, map[string]string{
		"mensagem": "Ocorreu um erro inesperado. Por favor, tente novamente mais tarde.",
	})

}
