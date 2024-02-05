package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	AccessTokenSecret      []byte
	RefreshTokenSecret     []byte
	AccessTokenExpiration  int
	RefreshTokenExpiration int
}

var Config *AppConfig

func LoadConfig() {
	err := godotenv.Load()

	if err != nil {
		panic("failed to load env")
	}

	accessTokenExpiration, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRATION"))

	if err != nil {
		panic(err)
	}

	refreshTokenExpiration, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRATION"))

	if err != nil {
		panic(err)
	}

	Config = &AppConfig{
		AccessTokenSecret:      []byte(os.Getenv("ACCESS_TOKEN_SECRET")),
		RefreshTokenSecret:     []byte(os.Getenv("REFRESH_TOKEN_SECRET")),
		AccessTokenExpiration:  accessTokenExpiration,
		RefreshTokenExpiration: refreshTokenExpiration,
	}
}
