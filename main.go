package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/vantrung8794/simplebank/api"
	db "github.com/vantrung8794/simplebank/db/sqlc"
	"github.com/vantrung8794/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot init new server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
