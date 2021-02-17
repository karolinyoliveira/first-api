package dao

import (
	"database/sql"
	"errors"

	"github.com/firstAPI/models"
)

var (
	ErrUserNotFound = errors.New("User Not Found")
)

type UserDao struct {
}

func NewUserDao() *UserDao { return &UserDao{} }

func (ud UserDao) GetById(con *sql.DB, id uint32) (models.User, error) {
	sql := "SELECT * FROM users WHERE uid= $1"
	rs, err := con.Query(sql, id)
	if err != nil {
		return models.User{}, err
	}
	defer rs.Close()
	var user models.User
	for rs.Next() {
		err := rs.Scan(&user.UID, &user.Nickname, &user.Email, &user.Password, &user.Status, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return models.User{}, err
		}
	}
	if user.UID == 0 {
		return models.User{}, ErrUserNotFound
	}
	return user, nil
}
func CreateUser()
