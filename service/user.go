package service

import (
	"github.com/first-api/dao"
	"github.com/first-api/models"
)

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService() *UserService {
	return &UserService{
		userDao: dao.NewUserDao(),
	}
}

func (us UserService) GetUsers() ([]models.User, error) {
	con := models.Connect()
	defer con.Close()
	return us.userDao.GetUsers(con)
}

func (us UserService) GetById(id uint32) (models.User, error) {
	con := models.Connect()
	defer con.Close()
	return us.userDao.GetById(con, id)
}

func (us UserService) GetByEmail(email string) (models.User, error) {
	con := models.Connect()
	defer con.Close()
	return us.userDao.GetByEmail(con, email)
}

func (us UserService) NewUser(user models.User) (bool, error) {
	con := models.Connect()
	defer con.Close()
	return us.userDao.NewUser(con, user)
}

func (us UserService) UpdateUser(user models.User) (int64, error) {
	con := models.Connect()
	defer con.Close()
	return us.userDao.UpdateUser(con, user)
}

func (us UserService) DeleteUser(uid uint32) (int64, error) {
	con := models.Connect()
	defer con.Close()
	return us.userDao.DeleteUser(con, uid)
}
