package services

import (
	"github.com/jackyrusly/jrgo/repositories"
)

type IProfileService interface {
	ServiceUpdateProfileName(i int64, name string) error
}

type ProfileService struct {
	ur repositories.IUserRepository
}

func NewProfileService(ur repositories.IUserRepository) *ProfileService {
	return &ProfileService{
		ur: ur,
	}
}

func (ps *ProfileService) ServiceUpdateProfileName(i int64, name string) error {
	return ps.ur.RepositoryUpdateUserName(i, name)
}
