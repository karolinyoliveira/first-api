package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/first-api/models"
	"github.com/first-api/routes"
)

func main() {
	port := "8080"
	models.TestConnection()
	fmt.Printf("Welcome to first-api! Running on port %s\n", port)
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
