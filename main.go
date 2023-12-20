package main

import (
	"database/sql"
	"log"
	"taskmanager/api"
	db "taskmanager/db/model"
	"taskmanager/utils"

	_ "github.com/lib/pq"
)

// @securityDefinitions.apikey	BearerAuth
// @name						Authorization
// @in							header
// @tokenurl					http://127.0.0.1:8080/api/user/login
func main() {
	cfg, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load env", err)
	}
	conn, err := sql.Open(cfg.DBDriver, cfg.DBConn)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	defer conn.Close()

	dbStore := db.SetupDB(conn)

	server, err := api.NewServer(cfg, dbStore)
	if err != nil {
		log.Fatal("cannot create server", err)
	}
	log.Println(server.Run(cfg.HTTPServerAddress))
}
