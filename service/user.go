package service

import (
	"github.com/firstAPI/dao"
	"github.com/firstAPI/models"
)

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService() *UserService {
	return &UserService{
		userDao: dao.NewUserDao(),
	}
}

func (us UserService) GetById(id uint32) (models.User, error) {
	con := Connect()
	defer con.Close()
	return us.userDao(con, id)
}
