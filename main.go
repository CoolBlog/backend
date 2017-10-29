package main

import (
	. "backend/databases"
	. "backend/commons"
)

func main() {
	defer DB.Close()
	r := initRouter()
	r.Run(HOST + PORT)
}
