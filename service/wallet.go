package service

import (
	"github.com/first-api/dao"
	"github.com/first-api/models"
)

type WalletService struct {
	walletDao *dao.WalletDao
}

func NewWalletService() *WalletService {
	return &WalletService{
		walletDao: dao.NewWalletDao(),
	}
}

func (ws WalletService) GetWallets() ([]models.Wallet, error) {
	con := models.Connect()
	defer con.Close()
	return ws.walletDao.GetWallets(con)
}

func (ws WalletService) GetWalletByPublicKey(publicKey string) (models.Wallet, error) {
	con := models.Connect()
	defer con.Close()
	return ws.walletDao.GetWalletByPublicKey(con, publicKey)
}

func (ws WalletService) Updatewallet(wallet models.Wallet) (int64, error) {
	con := models.Connect()
	defer con.Close()
	return ws.walletDao.UpdateWallet(con, wallet)
}

func (ws WalletService) AddBalance(wallet models.Wallet) (int64, error) {
	con := models.Connect()
	defer con.Close()
	return ws.walletDao.AddBalance(con, wallet)
}
