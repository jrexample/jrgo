package main

import (
	"github.com/jackyrusly/jrgo/controllers"
	"github.com/jackyrusly/jrgo/middlewares"
	"github.com/jackyrusly/jrgo/repositories"
	"github.com/jackyrusly/jrgo/routes"
	"github.com/jackyrusly/jrgo/services"
	"github.com/jackyrusly/jrgo/utils"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			utils.NewDatabase,
			utils.NewRouter,
			utils.NewConfig,
			routes.NewAuthRoute,
			routes.NewProfileRoute,
			middlewares.NewJwtMiddleware,
			fx.Annotate(
				controllers.NewAuthController,
				fx.As(new(controllers.IAuthController)),
			),
			fx.Annotate(
				services.NewAuthService,
				fx.As(new(services.IAuthService)),
			),
			fx.Annotate(
				repositories.NewUserRepository,
				fx.As(new(repositories.IUserRepository)),
			),
			fx.Annotate(
				controllers.NewProfileController,
				fx.As(new(controllers.IProfileController)),
			),
			fx.Annotate(
				services.NewProfileService,
				fx.As(new(services.IProfileService)),
			),
		),
		fx.Invoke(func(e *echo.Echo, ar *routes.AuthRoute, pr *routes.ProfileRoute) {
			ar.RegisterAuthRoutes(e)
			pr.RegisterProfileRoutes(e)

			e.Logger.Fatal(e.Start(":8080"))
		}),
	).Run()
}
