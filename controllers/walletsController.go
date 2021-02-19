package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/first-api/models"
	"github.com/first-api/utils"

	"github.com/gorilla/mux"
)

func GetWallets(w http.ResponseWriter, r *http.Request) {
	wallets, err := models.GetWallets()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, wallets)
}

func GetWallet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	wallet, err := models.GetWalletByPublicKey(params["public_key"])
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, wallet)
}

func PutWallet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedWallet models.Wallet //por algum motivo isso tem que vir antes
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &updatedWallet)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	updatedWallet.PublicKey = params["public_key"]

	rows, err := models.AddBalance(updatedWallet)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, rows)
}
