package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // underline = Import for side effect
)

// pq - A pure Go postgres driver for Go's database/sql package
const (
	USER    = "postgres"
	PASS    = "postgres"
	DBNAME  = "fapidb"
	SSLMODE = "disable"
)

func Connect() *sql.DB {
	URL := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", USER, PASS, DBNAME, SSLMODE)
	db, err := sql.Open("postgres", URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}

func TestConnection() {
	con := Connect()
	defer con.Close() // adiar
	err := con.Ping()
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return
	}
	fmt.Println("Database Connected!")
	return
}
