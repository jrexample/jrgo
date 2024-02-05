package dto

import "github.com/golang-jwt/jwt/v5"

type RegisterRequestBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type LoginRequestBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type TokenBody struct {
	AccessToken  string `json:"accessToken" validate:"required"`
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type AccessTokenClaims struct {
	jwt.MapClaims
	ID int64 `json:"id"`
}

type RefreshTokenClaims struct {
	jwt.MapClaims
	AccessToken string `json:"at"`
}
