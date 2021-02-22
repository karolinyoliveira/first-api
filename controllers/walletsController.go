package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/first-api/models"
	"github.com/first-api/service"
	"github.com/first-api/utils"

	"github.com/gorilla/mux"
)

func GetWallets(w http.ResponseWriter, r *http.Request) {
	wallets, err := service.NewWalletService().GetWallets()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, wallets)
}

func GetWallet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	wallet, err := service.NewWalletService().GetWalletByPublicKey(params["public_key"])
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, wallet)
}

func PutWallet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedWallet models.Wallet
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &updatedWallet)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	updatedWallet.PublicKey = params["public_key"]

	rows, err := service.NewWalletService().AddBalance(updatedWallet)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, rows)
}
