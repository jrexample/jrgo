package routes

import (
	"github.com/jackyrusly/jrgo/controllers"
	"github.com/jackyrusly/jrgo/middlewares"
	"github.com/labstack/echo/v4"
)

type ProfileRoute struct {
	c controllers.IProfileController
	m *middlewares.JwtMiddleware
}

func NewProfileRoute(c controllers.IProfileController, m *middlewares.JwtMiddleware) *ProfileRoute {
	return &ProfileRoute{
		c: c,
		m: m,
	}
}

func (r *ProfileRoute) RegisterProfileRoutes(e *echo.Echo) {
	g := e.Group("/profile")
	g.Use(r.m.CheckAccessToken)
	g.GET("", r.c.ControllerGetProfile)
	g.PATCH("/name", r.c.ControllerUpdateProfileName)
}
