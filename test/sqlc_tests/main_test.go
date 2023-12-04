package test

import (
	"database/sql"
	"os"
	db "simple-bank/db/sqlc"
	"simple-bank/util"
	"testing"

	_ "github.com/lib/pq"
)

var testingQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		panic(err)
	}

	testDB, err = sql.Open(config.Dbdriver, config.Dbsource)
	if err != nil {
		panic(err)
	}
	testingQueries = db.New(testDB)
	
	os.Exit(m.Run())
}