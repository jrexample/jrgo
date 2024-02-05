package services

import (
	"github.com/jackyrusly/jrgo/repositories"
	"github.com/jackyrusly/jrgo/utils"
)

type IProfileService interface {
	ServiceUpdateProfileName(i int64, name string) error
}

type ProfileService struct {
	ur repositories.IUserRepository
	c  *utils.Config
}

func NewProfileService(ur repositories.IUserRepository, c *utils.Config) *ProfileService {
	return &ProfileService{
		ur: ur,
		c:  c,
	}
}

func (ps *ProfileService) ServiceUpdateProfileName(i int64, name string) error {
	return ps.ur.RepositoryUpdateUserName(i, name)
}
