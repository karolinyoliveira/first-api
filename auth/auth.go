package auth

import (
	"errors"

	"github.com/first-api/models"
	"github.com/first-api/service"
	"github.com/first-api/utils"
)

var (
	ErrInvalidPassword = errors.New("Invalid Password")
)

func SignIn(user models.User) (models.User, error) {
	password := user.Password
	user, err := service.NewUserService().GetByEmail(user.Email)
	if err != nil {
		return user, err
	}
	err = utils.IsPassword(user.Password, password)
	if err != nil {
		return models.User{}, ErrInvalidPassword
	}
	return user, nil
}
