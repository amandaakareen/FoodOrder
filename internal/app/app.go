package app

import (
	"FoodOrder/internal/auth"
	"FoodOrder/internal/controller"
	"FoodOrder/internal/infra"
	"FoodOrder/internal/infra/repository"
	userCase "FoodOrder/internal/usecase/user"
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func NewEcho(lc fx.Lifecycle) *echo.Echo {
	e := echo.New()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
					fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	},
	)

	return e
}

func RegistreRoutes(e *echo.Echo, u *controller.UserController) {
	e.POST("/login", u.Login)
	e.POST("/cadastro", u.Login)
}

func Run() {

	app := fx.New(
		fx.Provide(NewEcho),
		fx.Provide(auth.ProvideJWTSecret),
		fx.Provide(auth.NewAuthService),
		fx.Provide(controller.NewUserController),
		fx.Provide(userCase.NewLoginUseCase),
		fx.Provide(repository.NewUserRepository),
		fx.Provide(infra.Connetc),
		fx.Invoke(RegistreRoutes),
	)

	app.Run()
}
