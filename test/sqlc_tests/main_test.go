package test

import (
	"database/sql"
	"os"
	db "simple-bank/db/sqlc"
	"testing"

	_ "github.com/lib/pq"
)

// recommend to set config values of testing database or delete all testing datas after tests  
const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testingQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		panic(err)
	}
	testingQueries = db.New(testDB)
	
	os.Exit(m.Run())
}