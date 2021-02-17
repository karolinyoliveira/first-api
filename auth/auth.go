package auth

import (
	"errors"

	"github.com/firstAPI/models"
	"github.com/firstAPI/utils"
)

var (
	ErrInvalidPassword = errors.New("Invalid Password")
)

func SignIn(user models.User) (models.User, error) {
	password := user.Password
	user, err := models.GetUserByEmail(user.Email)
	if err != nil {
		return user, err
	}
	err = utils.IsPassword(user.Password, password)
	if err != nil {
		return models.User{}, ErrInvalidPassword
	}
	return user, nil
}
