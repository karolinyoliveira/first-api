package controllers

import (
	"net/http"

	"github.com/first-api/utils"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	utils.ToJson(w, struct {
		Message string `json:"message"` // = "json:\"message\"" ?
	}{
		Message: "Welcome to first-api!",
	})
}
