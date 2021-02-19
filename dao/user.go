package dao

import (
	"database/sql"
	"errors"

	"github.com/first-api/models"
	"github.com/first-api/utils"
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

func (ud UserDao) GetUsers(con *sql.DB) ([]models.User, error) {
	sql := "SELECT * FROM users"
	rs, err := con.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	var users []models.User
	for rs.Next() {
		var user models.User
		err := rs.Scan(&user.UID, &user.Nickname, &user.Email, &user.Password, &user.Status, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (ud UserDao) GetByEmail(con *sql.DB, email string) (models.User, error) {
	sql := "SELECT * FROM users WHERE email= $1"
	rs, err := con.Query(sql, email)
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

func (ud UserDao) NewUser(con *sql.DB, user models.User) (bool, error) {
	tx, err := con.Begin()
	if err != nil {
		return false, err
	}

	// DB: users
	sql := "INSERT INTO users (nickname, email, password) VALUES ($1, $2, $3) RETURNING uid "
	{
		stmt, err := tx.Prepare(sql)
		if err != nil {
			tx.Rollback()
			return false, err
		}

		defer stmt.Close()
		hashedPassword, err := utils.Bcrypt(user.Password)
		if err != nil {
			tx.Rollback()
			return false, err
		}

		err = stmt.QueryRow(user.Nickname, user.Email, hashedPassword).Scan(&user.UID)
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}

	// DB: wallets
	sql = "INSERT INTO wallets (public_key,usr) VALUES ($1, $2)"
	wallet := models.Wallet{User: user}
	wallet.GeneratePublicKey()
	{
		stmt, err := tx.Prepare(sql)
		if err != nil {
			tx.Rollback()
			return false, err
		}
		_, err = stmt.Exec(wallet.PublicKey, wallet.User.UID)
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}
	return true, tx.Commit()
}

func (ud UserDao) UpdateUser(con *sql.DB, user models.User) (int64, error) {
	sql := "UPDATE USERS SET nickname= $1,email = $2, status = $3 WHERE uid = $4"

	stmt, err := con.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	rs, err := stmt.Exec(user.Nickname, user.Email, user.Status, user.UID)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}

func (ud UserDao) DeleteUser(con *sql.DB, uid uint32) (int64, error) {
	sql := "DELETE FROM users WHERE uid = $1"

	stmt, err := con.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	rs, err := stmt.Exec(uid)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}
