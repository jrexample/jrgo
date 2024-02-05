package utils

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AccessTokenSecret      []byte
	RefreshTokenSecret     []byte
	AccessTokenExpiration  int
	RefreshTokenExpiration int
}

func NewConfig() *Config {
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

	config := &Config{
		AccessTokenSecret:      []byte(os.Getenv("ACCESS_TOKEN_SECRET")),
		RefreshTokenSecret:     []byte(os.Getenv("REFRESH_TOKEN_SECRET")),
		AccessTokenExpiration:  accessTokenExpiration,
		RefreshTokenExpiration: refreshTokenExpiration,
	}

	return config
}
