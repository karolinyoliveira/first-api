package dao

import (
	"database/sql"
	"errors"

	"github.com/first-api/models"
)

var (
	ErrWalletNotFound = errors.New("Wallet Not Found")
)

type WalletDao struct {
}

func NewWalletDao() *WalletDao { return &WalletDao{} }

func (wd WalletDao) GetWallets(con *sql.DB) ([]models.Wallet, error) {

	sql := `SELECT u.uid, u.nickname, u.email, u.status, u.created_at, u.updated_at, w.public_key, w.balance, w.updated_at
	 FROM wallets AS w 
	 INNER JOIN users AS u
	 ON u.uid = w.usr;`

	rs, err := con.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	var wallets []models.Wallet
	for rs.Next() {
		var wallet models.Wallet
		err := rs.Scan(&wallet.User.UID, &wallet.User.Nickname, &wallet.User.Email, &wallet.User.Status, &wallet.User.CreatedAt, &wallet.User.UpdatedAt, &wallet.PublicKey, &wallet.Balance, &wallet.UpdatedAt)

		if err != nil {
			return nil, err

		}
		wallets = append(wallets, wallet)
	}
	return wallets, nil
}

func (wd WalletDao) GetWalletByPublicKey(con *sql.DB, publicKey string) (models.Wallet, error) {

	sql := `SELECT u.uid, u.nickname, u.email, u.password, u.status, u.created_at, u.updated_at, 
	 w.public_key, w.balance, w.updated_at
	 FROM wallets AS w 
	 INNER JOIN users AS u
	 ON u.uid = w.usr
	 WHERE w.public_key = $1`

	rs, err := con.Query(sql, publicKey)
	if err != nil {
		return models.Wallet{}, err
	}
	defer rs.Close()

	var wallet models.Wallet
	for rs.Next() {
		err := rs.Scan(&wallet.User.UID, &wallet.User.Nickname, &wallet.User.Email, &wallet.User.Password, &wallet.User.Status, &wallet.User.CreatedAt, &wallet.User.UpdatedAt, &wallet.PublicKey, &wallet.Balance, &wallet.UpdatedAt)
		if err != nil {
			return models.Wallet{}, err
		}
	}
	if wallet.PublicKey == "" {
		return models.Wallet{}, ErrWalletNotFound
	}
	return wallet, nil
}

func (wd WalletDao) UpdateWallet(con *sql.DB, w models.Wallet) (int64, error) {

	sql := "UPDATE wallets SET balance = $1 WHERE public_key = $2"
	stmt, err := con.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	rs, err := stmt.Exec(w.Balance, w.PublicKey)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}

func (wd WalletDao) AddBalance(con *sql.DB, w models.Wallet) (int64, error) {

	sql := "UPDATE wallets SET balance = (balance + $1) WHERE public_key = $2"
	stmt, err := con.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	rs, err := stmt.Exec(w.Balance, w.PublicKey)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}
