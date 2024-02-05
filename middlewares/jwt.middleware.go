package middlewares

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackyrusly/jrgo/config"
	"github.com/jackyrusly/jrgo/dto"
	"github.com/jackyrusly/jrgo/repositories"
	"github.com/jackyrusly/jrgo/utils"
	"github.com/labstack/echo/v4"
)

type JwtMiddleware struct {
	ur repositories.IUserRepository
}

func NewJwtMiddleware(ur repositories.IUserRepository) *JwtMiddleware {
	return &JwtMiddleware{
		ur: ur,
	}
}

func (j *JwtMiddleware) CheckAccessToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header.Get("Authorization")

		if authorization == "" {
			return c.JSON(http.StatusUnauthorized, utils.Response{
				Message: "Unauthorized",
			})
		}

		splitAuthorization := strings.Split(authorization, "Bearer")

		if len(splitAuthorization) != 2 {
			return c.JSON(http.StatusUnauthorized, utils.Response{
				Message: "Wrong Authorization",
			})
		}

		accessToken := strings.TrimSpace(splitAuthorization[1])

		claims := dto.AccessTokenClaims{}
		token, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
			return config.Config.AccessTokenSecret, nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, utils.Response{
				Message: err.Error(),
			})
		}

		if !token.Valid {
			return c.JSON(http.StatusUnauthorized, utils.Response{
				Message: "Invalid token",
			})
		}

		if claims.ID == 0 {
			return c.JSON(http.StatusUnauthorized, utils.Response{
				Message: "Invalid claims",
			})
		}

		user := j.ur.RepositoryFindById(claims.ID)

		if user.ID == 0 {
			return c.JSON(http.StatusUnauthorized, utils.Response{
				Message: "Invalid user",
			})
		}

		c.Set("User", user)

		return next(c)
	}
}
