package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/firstAPI/auth"
	"github.com/firstAPI/models"
	"github.com/firstAPI/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnauthorized)
		return
	}
	user, err = auth.SignIn(user)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnauthorized)
		return
	}
	utils.ToJson(w, user)
}
