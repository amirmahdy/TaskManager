package db

import (
	"database/sql"
	"log"
	"os"
	"taskmanager/utils"
	"testing"

	_ "github.com/lib/pq"
)

var testStore Store

func TestMain(m *testing.M) {
	cfg, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load env", err)
	}

	conn, err := sql.Open(cfg.DBDriver, cfg.DBConn)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	defer conn.Close()

	testStore = SetupDB(conn)
	os.Exit(m.Run())
}
