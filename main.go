package main

import (
	. "backend/databases"
	. "backend/global"
)

func main() {
	defer DB.Close()
	r := initRouter()
	r.Run(HOST + PORT)
}
