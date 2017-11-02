package main

import (
	. "backend/commons"
	"github.com/cihub/seelog"
	"backend/models"
)

func main() {
	db, err := models.InitDB()
	if err != nil {
		seelog.Critical("err open databases", err)
		return
	}
	defer db.Close()
	r := initRouter()
	r.Run(HOST + PORT)
}
