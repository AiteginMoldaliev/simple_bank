package main

import (
	"database/sql"
	"simple-bank/api"
	db "simple-bank/db/sqlc"
	"simple-bank/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		panic(err)
	}		

	conn, err := sql.Open(config.Dbdriver, config.Dbsource)
	if err != nil {
		panic(err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err = server.Start(config.ServerAddress); err != nil {
		panic(err)
	}
}