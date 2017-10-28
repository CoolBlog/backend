package databases

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var DB *sql.DB

func init() {
	os.Remove("./test.db")
	DB, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
