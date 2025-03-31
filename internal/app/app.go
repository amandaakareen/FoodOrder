package app

import (
	"FoodOrder/internal/infra"
	"fmt"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func NewEcho() *echo.Echo {
	e := echo.New()

	e.GET("/", c)

	return e
}

func StartServer(e *echo.Echo) {
	e.Logger.Fatal(e.Start(":8080"))
}

func Run() {

	infra.Connetc()

	app := fx.New(
		fx.Provide(NewEcho),
		fx.Invoke(StartServer),
	)

	if err := app.Start(nil); err != nil {
		fmt.Println("Erro ao iniciar o app:", err)
	}
}
