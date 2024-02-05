package repositories

import (
	"errors"

	"github.com/jackyrusly/jrgo/dto"
	"github.com/jackyrusly/jrgo/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	RepositoryCreateUser(d dto.RegisterRequestBody) error
	RepositoryFindById(i int64) models.User
	RepositoryFindByUsername(Username string) models.User
	RepositoryFindByUsernameAndPassword(Username string, Password string) (models.User, error)
	RepositoryUpdateUserName(ID int64, Name string) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ar *UserRepository) RepositoryFindById(i int64) models.User {
	var user models.User
	ar.db.Where("ID = ?", i).First(&user)

	return user
}

func (ar *UserRepository) RepositoryFindByUsername(Username string) models.User {
	var user models.User
	ar.db.Where("Username = ?", Username).First(&user)

	return user
}

func (ar *UserRepository) RepositoryCreateUser(d dto.RegisterRequestBody) error {
	u := ar.RepositoryFindByUsername(d.Username)

	if u.ID != 0 {
		return errors.New("username already exists")
	}

	ar.db.Create(&models.User{
		Username: d.Username,
		Password: d.Password,
		Name:     d.Name,
	})

	return nil
}

func (ar *UserRepository) RepositoryFindByUsernameAndPassword(Username string, Password string) (models.User, error) {
	var user models.User
	ar.db.Where("Username = ? AND Password = ?", Username, Password).First(&user)

	if user.ID == 0 {
		return user, errors.New("invalid username or password")
	}

	return user, nil
}

func (ar *UserRepository) RepositoryUpdateUserName(ID int64, Name string) error {
	var user models.User
	ar.db.Where("ID = ?", ID).First(&user)

	if user.ID == 0 {
		return errors.New("invalid user")
	}

	user.Name = Name
	ar.db.Save(&user)

	return nil
}
