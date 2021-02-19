package routes

import (
	"github.com/first-api/controllers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", controllers.GetHome).Methods("GET")
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{uid}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/users", controllers.PostUsers).Methods("POST")
	r.HandleFunc("/users/{uid}", controllers.PutUser).Methods("PUT")
	r.HandleFunc("/users/{uid}", controllers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/wallets", controllers.GetWallets).Methods("GET")
	r.HandleFunc("/wallets/{public_key}", controllers.GetWallet).Methods("GET")
	r.HandleFunc("/wallets/{public_key}", controllers.PutWallet).Methods("PUT")

	return r
}
