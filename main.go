package main

import (
	"database/sql"
	"fmt"
	"log"

	"kuann.me/simplebank/api"
	db "kuann.me/simplebank/db/sqlc"

	_ "github.com/lib/pq"
)

var conn *sql.DB

const (
	dbdriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5434/simple_bank?sslmode=disable"
	address  = "0.0.0.0:8089"
)

func main() {
	fmt.Println("Starting main")
	var err error
	conn, err = sql.Open(dbdriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	error := server.Start(address)

	if error != nil {
		log.Fatal("Faileeeeeddddd")
	}
}
