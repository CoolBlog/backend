package main

import (
	. "backend/databases"
	. "backend/consts"
)

func main() {
	defer DB.Close()
	r := initRouter()
	r.Run(HOST + PORT)
}
