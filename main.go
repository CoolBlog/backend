package main

import (
	. "backend/databases"
)

func main() {
	defer DB.Close()
	r := initRouter()
	r.Run(":8000")
}
