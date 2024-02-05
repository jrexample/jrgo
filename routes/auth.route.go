package routes

import (
	"github.com/jackyrusly/jrgo/controllers"
	"github.com/labstack/echo/v4"
)

type AuthRoute struct {
	c controllers.IAuthController
}

func NewAuthRoute(c controllers.IAuthController) *AuthRoute {
	return &AuthRoute{
		c: c,
	}
}

func (r *AuthRoute) RegisterAuthRoutes(e *echo.Echo) {
	g := e.Group("/auth")
	g.POST("/register", r.c.ControllerRegister)
	g.POST("/login", r.c.ControllerLogin)
	g.POST("/refresh-token", r.c.ControllerRefreshToken)
}
