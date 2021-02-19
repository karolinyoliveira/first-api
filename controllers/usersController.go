package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/first-api/models"
	"github.com/first-api/service"
	"github.com/first-api/utils"
	"github.com/first-api/validations"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := service.NewUserService().GetUsers()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	utils.ToJson(w, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["uid"], 10, 32)
	user, err := service.NewUserService().GetById(uint32(id))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	utils.ToJson(w, user)
}

func PostUsers(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var user models.User

	err := json.Unmarshal(body, &user)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}

	user, err = validations.ValidateNewUser(user)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}

	_, err = service.NewUserService().NewUser(user)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uid, _ := strconv.ParseUint(params["uid"], 10, 32)
	body, _ := ioutil.ReadAll(r.Body)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	user.UID = uint32(uid)

	rows, err := service.NewUserService().UpdateUser(user)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, rows)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uid, _ := strconv.ParseUint(params["uid"], 10, 32)

	_, err := service.NewUserService().DeleteUser(uint32(uid))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, http.StatusNoContent)
}
