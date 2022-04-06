package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

const (
	dbdriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5434/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	fmt.Println("Starting main")
	var err error
	testDb, err = sql.Open(dbdriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
