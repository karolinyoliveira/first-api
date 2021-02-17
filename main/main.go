package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/firstAPI/models"
	"github.com/firstAPI/routes"
)

func main() {
	port := "8080"
	models.TestConnection()
	fmt.Printf("Welcome to FirstAPI! Running on port %s\n", port)
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
