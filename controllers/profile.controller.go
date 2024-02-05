package controllers

import (
	"net/http"

	"github.com/jackyrusly/jrgo/dto"
	"github.com/jackyrusly/jrgo/models"
	"github.com/jackyrusly/jrgo/services"
	"github.com/jackyrusly/jrgo/utils"
	"github.com/labstack/echo/v4"
)

type IProfileController interface {
	ControllerGetProfile(c echo.Context) error
	ControllerUpdateProfileName(c echo.Context) error
}

type ProfileController struct {
	ps services.IProfileService
}

func NewProfileController(ps services.IProfileService) *ProfileController {
	return &ProfileController{
		ps: ps,
	}
}

func (ac *ProfileController) ControllerGetProfile(c echo.Context) error {
	user := c.Get("User").(models.User)

	return c.JSON(http.StatusOK, utils.Response{
		Message: "Success",
		Data: dto.UserResponseBody{
			ID:       user.ID,
			Username: user.Username,
			Name:     user.Name,
		},
	})
}

func (ac *ProfileController) ControllerUpdateProfileName(c echo.Context) error {
	u := c.Get("User").(models.User)
	d := new(dto.UpdateUserNameRequestBody)

	if err := c.Bind(d); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(d); err != nil {
		return err
	}

	if err := ac.ps.ServiceUpdateProfileName(u.ID, d.Name); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
