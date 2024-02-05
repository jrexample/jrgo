package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackyrusly/jrgo/dto"
	"github.com/jackyrusly/jrgo/repositories"
	"github.com/jackyrusly/jrgo/utils"
)

type IAuthService interface {
	ServiceGenerateToken(i int64) (*dto.TokenBody, error)
	ServiceRegister(d dto.RegisterRequestBody) error
	ServiceLogin(d dto.LoginRequestBody) (*dto.TokenBody, error)
	ServiceRefreshToken(d dto.TokenBody) (*dto.TokenBody, error)
}

type AuthService struct {
	ur repositories.IUserRepository
	c  *utils.Config
}

func NewAuthService(ur repositories.IUserRepository, c *utils.Config) *AuthService {
	return &AuthService{
		ur: ur,
		c:  c,
	}
}

func (as *AuthService) ServiceGenerateToken(i int64) (*dto.TokenBody, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": "jrgo-server",
			"id":  i,
			"exp": time.Now().Add(time.Hour * time.Duration(as.c.AccessTokenExpiration)).Unix(),
		})

	accessToken, err := at.SignedString(as.c.AccessTokenSecret)

	if err != nil {
		return nil, err
	}

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": "jrgo-server",
			"at":  accessToken,
			"exp": time.Now().Add(time.Hour * time.Duration(as.c.RefreshTokenExpiration)),
		})

	refreshToken, err := rt.SignedString(as.c.RefreshTokenSecret)

	if err != nil {
		return nil, err
	}

	return &dto.TokenBody{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (as *AuthService) ServiceRegister(d dto.RegisterRequestBody) error {
	return as.ur.RepositoryCreateUser(d)
}

func (as *AuthService) ServiceLogin(d dto.LoginRequestBody) (*dto.TokenBody, error) {
	user, err := as.ur.RepositoryFindByUsernameAndPassword(d.Username, d.Password)

	if err != nil {
		return nil, err
	}

	return as.ServiceGenerateToken(user.ID)
}

func (as *AuthService) ServiceRefreshToken(d dto.TokenBody) (*dto.TokenBody, error) {
	refreshTokenClaims := dto.RefreshTokenClaims{}
	refreshToken, err := jwt.ParseWithClaims(d.RefreshToken, &refreshTokenClaims, func(token *jwt.Token) (interface{}, error) {
		return as.c.RefreshTokenSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !refreshToken.Valid {
		return nil, errors.New("invalid refresh token")
	}

	if refreshTokenClaims.AccessToken != d.AccessToken {
		return nil, errors.New("invalid payload")
	}

	accessTokenClaims := dto.AccessTokenClaims{}
	accessToken, _, _ := new(jwt.Parser).ParseUnverified(d.AccessToken, &accessTokenClaims)
	_ = accessToken

	if err != nil {
		return nil, err
	}

	if accessTokenClaims.ID == 0 {
		return nil, errors.New("invalid access token")
	}

	return as.ServiceGenerateToken(accessTokenClaims.ID)
}
