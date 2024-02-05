package controllers

import (
	"net/http"

	"github.com/jackyrusly/jrgo/dto"
	"github.com/jackyrusly/jrgo/services"
	"github.com/jackyrusly/jrgo/utils"
	"github.com/labstack/echo/v4"
)

type IAuthController interface {
	ControllerRegister(c echo.Context) error
	ControllerLogin(c echo.Context) error
	ControllerRefreshToken(c echo.Context) error
}

type AuthController struct {
	as services.IAuthService
}

func NewAuthController(as services.IAuthService) *AuthController {
	return &AuthController{
		as: as,
	}
}

func (ac *AuthController) ControllerRegister(c echo.Context) error {
	d := new(dto.RegisterRequestBody)

	if err := c.Bind(d); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(d); err != nil {
		return err
	}

	if err := ac.as.ServiceRegister(*d); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}

func (ac *AuthController) ControllerLogin(c echo.Context) error {
	d := new(dto.LoginRequestBody)

	if err := c.Bind(d); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
		})
	}

	if err := c.Validate(d); err != nil {
		return err
	}

	tokenResponse, err := ac.as.ServiceLogin(*d)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &utils.Response{
		Message: "Success",
		Data:    tokenResponse,
	})
}

func (ac *AuthController) ControllerRefreshToken(c echo.Context) error {
	d := new(dto.TokenBody)

	if err := c.Bind(d); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
		})
	}

	if err := c.Validate(d); err != nil {
		return err
	}

	token, err := ac.as.ServiceRefreshToken(*d)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &utils.Response{
		Data: token,
	})
}
