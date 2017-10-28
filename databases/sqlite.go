package databases

import (
	"os"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	. "backend/consts"
)

var DB *sql.DB

func init() {
	os.Remove(DB_NAME)
	DB, err := sql.Open(DB_TYPE, DB_NAME)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
