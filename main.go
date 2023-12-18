package main

import (
	"database/sql"
	"log"
	"taskmanager/api"
	db "taskmanager/db/model"
	"taskmanager/token"
	"taskmanager/utils"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load env", err)
	}

	token, err := token.NewJWTMaker(cfg.TokenSymmetricKey)
	if err != nil {
		log.Fatal("cannot create token maker", err)
	}

	conn, err := sql.Open(cfg.DBDriver, cfg.DBConn)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	defer conn.Close()

	dbStore := db.SetupDB(conn)

	server, err := api.NewServer(cfg, dbStore, token)
	if err != nil {
		log.Fatal("cannot create server", err)
	}
	log.Println(server.Run(cfg.HTTPServerAddress))
}
